package routes

import (
	"github.com/kataras/iris"
	"github.com/lancelrq/go-iris-starter-kit/src/views"
)

func InitRoutes(app *iris.Application) {
	app.Get("/", views.IndexView)
}