package datastruct

// 前缀树
// 单词树
type Node struct {
	Char     string
	Word     string
	Children []*Node
}

func NewNode() *Node {
	return &Node{
		Children: make([]*Node, 0),
	}
}

func (n *Node) Insert(c string) {
	n.insert(c, 0)
}

// 查询指定单词开始的长度
// 并且返回指定个数的前缀
func (n *Node) SearchPrefix(c string, l int) []string {

	child := n.search(c, 0)

	if child == nil {
		return make([]string, 0)
	}

	// 从当前结点开始搜索
	res := make([]string, l)
	index := 0
	child.searchWord(res, &index)

	return res
}

// 在指定结点中搜索单词
func (n *Node) searchWord(re []string, index *int) {

	if len(re) == *index {
		return
	}

	if n.Word != "" {
		re[*index] = n.Word
		*index = *index + 1
	}

	for _, item := range n.Children {
		item.searchWord(re,index)
	}

}

func (n *Node) Search(c string) *Node {
	return n.search(c, 0)
}

// 找到指定节点
func (n *Node) search(c string, height int) *Node {

	if len(c) == height {
		return n
	}

	// 在子节点中搜索
	child := n.searchNode(string(c[height]))

	if child == nil {
		return nil
	}

	return child.search(c, height+1)
}

func (n *Node) insert(c string, height int) {

	if len(c) == height {
		n.Word = c
		return
	}

	node := n.searchNode(string(c[height]))

	if node == nil {
		node = &Node{
			Char:     string(c[height]),
			Children: make([]*Node, 0),
		}
		n.Children = append(n.Children, node)
	}

	node.insert(c, height+1)
}

// 在某个节点的所有子节点中，搜索需要的节点
// 找到则返回节点
// 未找到则返回nil
func (n *Node) searchNode(c string) *Node {

	for _, item := range n.Children {

		if item.Char == c {
			return item
		}
	}

	return nil
}
