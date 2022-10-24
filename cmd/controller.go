package templates

var ControllerTemplate = `
package {{ .Name }}

import (

{{ if .Framework.Fiber }}
	"github.com/gofiber/fiber/v2"
{{ end }}
)

type {{ .Name }}Controller struct {
	path string
}

func New{{ .Name }}Controller() *{{ .Name }}Controller {
	return &{{ .Name }}Controller {
		path: "{{ .Name }}"
	}
}

func (ctrl *{{ .Name }}Controller) Path() string {
	return ctrl.path
}

{{ if .Framework.Fiber }}
func (ctrl *{{ .Name }}Controller) Routes(router fiber.Router) {
}
{{ end }}
`
