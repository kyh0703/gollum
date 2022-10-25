package templates

var Controller = `
package {{ .Name | ToLower }}

import (

{{ if eq .Framework "fiber" }}
	"github.com/gofiber/fiber/v2"
{{ else if eq .Framework "gin" }}
	"github.com/gin-gonic/gin"
{{ end }}
)

type {{ .Name | ToLower }}Controller struct {
	path string
}

func New{{ .Name | ToTitle }}Controller() *{{ .Name | ToLower }}Controller {
	return &{{ .Name | ToLower }}Controller {
		path: "{{ .Name }}",
	}
}

func (ctrl *{{ .Name | ToLower }}Controller) Path() string {
	return ctrl.path
}

{{ if eq .Framework "fiber" }}
func (ctrl *{{ .Name | ToLower }}Controller) Routes(router fiber.Router) {
}
{{ else if eq .Framework "gin" }}
func (ctrl *{{ .Name | ToLower }}Controller) Routes(router gin.Router) {
}
{{ end }}
`
