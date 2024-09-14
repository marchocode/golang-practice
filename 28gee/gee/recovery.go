package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func trace(message string) string {

	var pcs [32]uintptr
	n := runtime.Callers(5, pcs[:])

	var str strings.Builder
	str.WriteString(message + "\n Tranceback:")

	for _, pc := range pcs[:n] {

		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)

		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}

	return str.String()
}

func Recovery() HandlerFunc {

	return func(ctx *Context) {

		defer func() {

			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)

				str := trace(message)
				log.Printf("%s \n\n", str)

				ctx.String(http.StatusInternalServerError, "Internal Server Error")
			}

		}()

		ctx.Next()

	}

}
