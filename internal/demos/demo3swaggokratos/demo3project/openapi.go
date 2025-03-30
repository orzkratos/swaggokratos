package demo3project

import (
	"embed"

	"github.com/go-xlan/yaml-go-edit/yamlv3edit"
	"github.com/yyle88/rese"
)

//go:embed openapi.yaml
var files embed.FS

func GetOpenapiContent(docTitle string) []byte {
	// 读取文档的内容
	content := rese.A1(files.ReadFile("openapi.yaml"))
	// 设置文档的标题
	content = yamlv3edit.ModifyYamlFieldValue(content, "info.title", docTitle)
	return content
}
