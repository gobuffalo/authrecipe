package grifts

import (
	"github.com/gobuffalo/authrecipe/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
