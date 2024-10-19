package swaggokratos

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/orzkratos/swaggokratos/swaggogin"
)

func RegisterSwaggoHTTPServer(srv *http.Server, prefix string, params []*swaggogin.Param) {
	//根据demo: https://github.com/go-kratos/examples/blob/main/http/gin/main.go 这里也用相同的方案
	engine := gin.New()
	//其实这里不恢复也行，毕竟还有洋葱卷，这里的日志也需要替换为kratos的日志，待做吧
	engine.Use(gin.Logger(), gin.Recovery())

	for idx := range params {
		swaggogin.SwaggerRoute(engine.Group(prefix), prefix, params[idx])
	}

	srv.HandlePrefix(prefix, engine)
}
