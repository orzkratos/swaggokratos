package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/orzkratos/swaggokratos"
	"github.com/orzkratos/swaggokratos/internal/demos/demo3swaggokratos/demo3project"
	"github.com/orzkratos/swaggokratos/swaggogin"
	"github.com/yyle88/done"
)

func main() {
	srv := http.NewServer(http.Address(":8080"))

	// see the swag go gin test case in the project
	param := &swaggogin.Param{
		SwaggerPath: "/swagger/c/*any",
		ExplorePath: "/abc/openapi-abc.yaml",
		ContentData: demo3project.GetOpenapiContent("demo-3-title"),
	}

	swaggokratos.RegisterSwaggoHTTPServer(srv, "/doc/", []*swaggogin.Param{param})

	app := kratos.New(
		kratos.Name("demo"),
		kratos.Server(
			srv,
		),
	)

	// please explore in chrome: http://127.0.0.1:8080/doc/swagger/c/index.html
	done.Done(app.Run())
}
