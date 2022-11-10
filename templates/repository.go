package templates

var Repository = `
package {{ .Name | ToLower }}

import (

	"github.com/gofiber/fiber/v2"
)

type {{ .Name | ToLower }}Repository struct {}
`
