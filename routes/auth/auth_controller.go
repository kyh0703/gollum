
package auth

import (


	"github.com/gin-gonic/gin"

)

type authController struct {
	path string
}

func NewAuthController() *authController {
	return &authController {
		path: "auth",
	}
}

func (ctrl *authController) Path() string {
	return ctrl.path
}


func (ctrl *authController) Routes(router gin.Router) {
}

