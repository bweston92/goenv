package goenv

import (
	"testing"

	"github.com/spf13/afero"

	"github.com/stretchr/testify/assert"
)

func stubResolver(m map[string]string) EnvResolver {
	return func(k string) string {
		if v, ok := m[k]; ok {
			return v
		} else {
			return ""
		}
	}
}

func Test_EnvLocator_Get_UsesValueFromResolver(t *testing.T) {
	esr := stubResolver(map[string]string{
		"ABC": "123",
	})
	el := New(esr, afero.NewMemMapFs())

	assert.Equal(t, "123", el.Get("ABC"))
}

func Test_EnvLocator_Get_GetsValueFromFileIfExists(t *testing.T) {
	esr := stubResolver(map[string]string{
		"ABC_FILE": "/run/secrets/abc",
	})
	mfs := afero.NewMemMapFs()
	f, _ := mfs.Create("/run/secrets/abc")
	f.Write([]byte("123"))
	f.Close()
	el := New(esr, mfs)

	assert.Equal(t, "123", el.Get("ABC"))
}

func Test_EnvLocator_Get_ReturnsNothingIfNotFound(t *testing.T) {
	esr := stubResolver(map[string]string{})
	el := New(esr, afero.NewMemMapFs())

	assert.Equal(t, "", el.Get("unknown"))
}
