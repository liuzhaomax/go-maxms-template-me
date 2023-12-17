package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/liuzhaomax/go-maxms-template-me/internal/api"
	"github.com/liuzhaomax/go-maxms-template-me/internal/core"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Engine   *gin.Engine
	Handler  *api.Handler
	Response *core.Response
}
