package grifts

import (
	"github.com/brittonhayes/hikeshi/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
