package goenv

import (
	"encoding/base64"
)

type stack []func(string) string

func (s stack) Play(o string) string {
	for _, m := range s {
		o = m(o)
	}

	return o
}

func defaultStack() stack {
	return stack{
		base64Middleware,
	}
}

func base64Middleware(in string) string {
	if len(in) < 6 || in[:7] != "base64:" {
		return in
	}

	o, err := base64.StdEncoding.DecodeString(in[7:])

	if err != nil {
		return ""
	}

	return string(o)
}
