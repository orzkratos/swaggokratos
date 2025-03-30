package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/orzkratos/swaggokratos"
	"github.com/orzkratos/swaggokratos/swaggogin"
	"github.com/yyle88/done"
)

func main() {
	srv := http.NewServer(http.Address(":8080"))

	// see the swag go gin test case in the project
	param := &swaggogin.Param{
		SwaggerPath: "/swagger/b/*any",
		ExplorePath: "/abc/openapi-uvw.yaml",
		ContentData: []byte("openapi: 3.0.3\ninfo:\n    title: DEMO-2-TITLE\n    version: 0.0.1"),
	}

	swaggokratos.RegisterSwaggoHTTPServer(srv, "/doc/", []*swaggogin.Param{param})

	app := kratos.New(
		kratos.Name("demo"),
		kratos.Server(
			srv,
		),
	)

	// please explore in chrome: http://127.0.0.1:8080/doc/swagger/b/index.html
	done.Done(app.Run())
}
