package src

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/lancelrq/go-iris-starter-kit/src/routes"
	"github.com/lancelrq/go-iris-starter-kit/src/utils"
)

type MasterServer struct {
	app *iris.Application
}

func NewServer() *MasterServer {
	tas := &MasterServer{}
	tas.app = iris.New()
	return tas
}

func (server *MasterServer) RunServer() (err error) {

	if utils.GlobalConfig.DebugMode {
		server.app.Logger().SetLevel("debug")
	} else {
		server.app.Logger().SetLevel("info")
	}
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	server.app.Use(recover.New())
	server.app.Use(logger.New())

	routes.InitRoutes(server.app)

	err = server.app.Run(iris.Addr(fmt.Sprintf("%s:%d", utils.GlobalConfig.Server.Listen, utils.GlobalConfig.Server.Port)))
	return err
}
