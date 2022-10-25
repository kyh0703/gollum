package templates

var Service = `
package {{ .Name }}

type {{ .Name }}Service struct{}

func NewAuthService() *authService {
	return &authService{}
}
`
