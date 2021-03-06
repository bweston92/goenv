package goenv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_stack_Play(t *testing.T) {
	b := stack{
		func(in string) string {
			return in + "abc"
		},
		func(in string) string {
			return in + "def"
		},
	}

	assert.Equal(t, "123abcdef", b.Play("123"))
}

func Test_base64Middleware_DoesNotModifyValueWithoutPrefix(t *testing.T) {
	assert.Equal(t, "abc", base64Middleware("abc"))
	assert.Equal(t, "abcdefghijklmnop", base64Middleware("abcdefghijklmnop"))
}

func Test_base64Middleware_DoesModifyValueWithPrefix(t *testing.T) {
	assert.Equal(t, "abc", base64Middleware("base64:YWJj"))
}

func Test_base64Middleware_ReturnsEmptyStringIfDecodeFails(t *testing.T) {
	assert.Equal(t, "", base64Middleware("base64:WH@T"))
}

func Test_trimSpaceMiddleware(t *testing.T) {
	assert.Equal(t, "abc", trimSpaceMiddleware("abc\n"))
	assert.Equal(t, "abc", trimSpaceMiddleware("abc\t"))
	assert.Equal(t, "abc", trimSpaceMiddleware("abc "))
	assert.Equal(t, "abc", trimSpaceMiddleware(" abc "))
}
