package templates

var Controller = `
package {{ .Name | ToLower }}

import (

	"github.com/gofiber/fiber/v2"
)

type {{ .Name | ToLower }}Controller struct {
	path string
}

func (ctrl *{{ .Name | ToLower }}Controller) Path() string {
	return ctrl.path
}

func (ctrl *{{ .Name | ToLower }}Controller) Routes(router fiber.Router) {
}
`
