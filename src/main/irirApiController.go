package main

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)
type IT struct {
	Company  string
	Subjects []string
	IsOk     bool
	Price    float64
}

func getJson() IT {

	t1 := IT{"itcast", []string{"Go", "C++", "Python", "Test"}, true, 666.666}
	fmt.Println(t1)
	b, err := json.MarshalIndent(t1, "", "    ")
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
	return t1
}
func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"a":"b"})

	})
	app.Get("/hello1", func(ctx iris.Context) {
		ctx.JSON(getJson())

	})
	app.Run(iris.Addr(":8080"))//8080 监听端口
}
