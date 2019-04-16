package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)
func main(){
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, notFound)
	app.OnErrorCode(iris.StatusInternalServerError, internalServerError)
	// 注册一个处理函数处理HTTP错误codes >=400:
	// app.OnAnyErrorCode(handler)
	app.Get("/", index)
	app.Run(iris.Addr(":8080"))
}
func notFound(ctx iris.Context) {
	// 当http.status=400 时向客户端渲染模板$views_dir/errors/404.html
	ctx.View("errors/404.html")
}
//当出现错误的时候，再试一次
func internalServerError(ctx iris.Context) {
	ctx.WriteString("Oups something went wrong, try again")
}
func index(ctx context.Context) {
	ctx.View("index.html")
}
