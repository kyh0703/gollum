package templates

var Service = `
package {{ .Name }}

type {{ .Name }}Service struct{}

func New{{ .Name | ToTitle }}Service() *{{ .Name }}Service {
	return &{{ .Name }}Service{}
}
`
