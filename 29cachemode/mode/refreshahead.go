package mode

import (
	"fmt"
	"log"
	"time"

	"github.com/withlin/canal-go/client"
	pbe "github.com/withlin/canal-go/protocol/entry"
	"google.golang.org/protobuf/proto"
	"marcho.life/cachemode/cache"
	"marcho.life/cachemode/db"
)

// 借助 Canal 中间件
// 监听binlog日志，客户端订阅binlog增量日志
// 从RedisRedis读取
// 直接写入到mysql数据库
type RefreshAhead struct {
	cache    cache.Cache
	database db.Database
}

func NewRefreshAhead() *RefreshAhead {

	// 开一个client 读取 canal. binlog 日志，并且同步到 redis.
	r := &RefreshAhead{
		cache:    cache.NewRedisClient(),
		database: db.NewMysql(),
	}

	go r.listen()

	return r
}

func (r *RefreshAhead) listen() {

	connector := client.NewSimpleCanalConnector("127.0.0.1", 11111, "", "", "example", 60000, 60*60*1000)

	err := connector.Connect()

	if err != nil {
		panic(err)
	}

	err = connector.Subscribe(".*\\..*")

	if err != nil {
		panic(err)
	}

	for {

		message, err := connector.Get(100, nil, nil)

		if err != nil {
			log.Println(err)
			panic(err)
		}

		batchId := message.Id

		if batchId == -1 || len(message.Entries) <= 0 {
			time.Sleep(300 * time.Millisecond)
			continue
		}

		log.Println("Receive a message.")

		for _, entity := range message.Entries {
			go r.process(entity)
		}
	}
}

func (r *RefreshAhead) process(item pbe.Entry) {

	if item.GetEntryType() == pbe.EntryType_TRANSACTIONBEGIN || item.GetEntryType() == pbe.EntryType_TRANSACTIONEND {
		return
	}

	rowChange := new(pbe.RowChange)
	err := proto.Unmarshal(item.GetStoreValue(), rowChange)

	if err != nil {
		log.Printf("error message = %v\n", err)
		return
	}

	if rowChange == nil {
		log.Printf("error message = %v\n", err)
		return
	}

	eventType := rowChange.GetEventType()
	header := item.GetHeader()

	log.Printf("binlog[%s : %d],name[%s,%s], eventType: %s", header.GetLogfileName(), header.GetLogfileOffset(), header.GetSchemaName(), header.GetTableName(), header.GetEventType())

	for _, rowData := range rowChange.GetRowDatas() {

		if eventType == pbe.EventType_INSERT || eventType == pbe.EventType_UPDATE {

			m := rowData2Map(rowData.GetAfterColumns())

			k := m["k"]
			val := m["val"]

			r.cache.Set(k, val)
		}
	}

}

func rowData2Map(columns []*pbe.Column) map[string]string {

	m := make(map[string]string, len(columns))

	for _, col := range columns {
		m[col.GetName()] = col.GetValue()
	}

	return m
}

func printColumn(columns []*pbe.Column) {
	for _, col := range columns {
		fmt.Println(fmt.Sprintf("%s : %s  update= %t", col.GetName(), col.GetValue(), col.GetUpdated()))
	}
}

func (r *RefreshAhead) Read(key string) string {

	v := r.cache.Get(key)

	if v != "" {
		return v
	}

	// 不存在，需要从数据库加载
	v = r.database.Select(key)

	// 写入cache.
	r.cache.Set(key, v)

	return v
}

func (r *RefreshAhead) Write(key, val string) {
	r.database.Update(key, val)
}
