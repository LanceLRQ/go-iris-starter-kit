package views

import "github.com/kataras/iris"

func IndexView (ctx iris.Context) {
	ctx.HTML("Congratulation! Your server is running.")
}