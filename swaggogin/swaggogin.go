package swaggogin

import (
	"strings"

	"github.com/gin-gonic/gin"
	swaggofiles "github.com/swaggo/files"
	ginswaggers "github.com/swaggo/gin-swagger"
	"github.com/yyle88/done"
)

type Param struct {
	SwaggerPath string //相当于设置自路由名称 只能设置为 "/swagger/a/*any" 或者 "/swagger/b/*any" 这样的内容
	ExplorePath string //请求 openapi.yaml 资源的路径，就是网页如何获取到接口文档，因此要符合路径规则，比如"/abc"或者"/abc/openapi-xyz.yaml"或者"/abc/openapi-uvw.yaml"这样的(带"/"前缀的)
	ContentData []byte //就是 openapi.yaml 内容的数据，就是整个 openapi.yaml 里面的内容
}

func SwaggerRoute(routes gin.IRoutes, prefix string, param *Param) {
	//被访问时，会请求一堆的资源，这是接口文档的资源完整路径
	resourcePath := strings.TrimSuffix(prefix, "/") + "/" + strings.TrimPrefix(param.ExplorePath, "/")

	//被访问时，会请求一堆的资源，这是swagger的静态资源内容
	routes.GET(param.SwaggerPath, ginswaggers.WrapHandler(
		swaggofiles.Handler,
		ginswaggers.URL(resourcePath), //在swagger静态页面中，需要根据接口文档渲染页面，因此需要有资源路径
	))

	//其中一个资源就是文档的内容，因此这个资源是必须存在的，否则页面就请求不到啦
	routes.GET(param.ExplorePath, func(c *gin.Context) {
		c.Header("Content-Type", "text/yaml")
		done.VNE(c.Writer.Write(param.ContentData)).Nice()
	})
}
