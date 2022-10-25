package templates

var Service = `
package {{ .Name | ToLower }}

type {{ .Name | ToLower }}Service struct{}

func New{{ .Name | ToTitle }}Service() *authService {
	return &{{ .Name }}Service{}
}
`
