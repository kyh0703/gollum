package templates

var ServiceTemplate = `
package {{ .Name }}

type {{ .Name }}Service struct{}

func NewAuthService() *authService {
	return &authService{}
}
`
