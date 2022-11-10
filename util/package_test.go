package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetModinfo(t *testing.T) {
	var pkg Package
	modFile, err := pkg.GetGoModFromExec()
	assert.NoError(t, err)
	t.Log(modFile)
}

func TestGoGetExec(t *testing.T) {
	var (
		pkg Package
		mod = "github.com/spf13/viper"
	)
	err := pkg.ExecGoGet(mod)
	assert.NoError(t, err)
}
