package main

import (
	"github.com/gin-gonic/gin"
	"github.com/orzkratos/swaggokratos/swaggogin"
	"github.com/yyle88/done"
)

func main() {
	engine := gin.Default()

	// this is the prefix of your doc route path
	const prefix = "/doc/"

	// read openapi.yaml content data
	// example:
	// contentData := done.VAE(os.ReadFile("/tmp/openapi.yaml")).Nice()
	// this is just demo data:
	contentData := []byte("openapi: 3.0.3\ninfo:\n    title: DEMO-TITLE\n    version: 0.0.1")

	swaggogin.SwaggerRoute(engine.Group(prefix), prefix, &swaggogin.Param{
		SwaggerPath: "/swagger/a/*any",
		ExplorePath: "/abc/openapi-xyz.yaml",
		ContentData: contentData,
	})

	// please explore in chrome: http://127.0.0.1:8080/doc/swagger/a/index.html
	done.Done(engine.Run(":8080"))
}
