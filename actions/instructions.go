package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// InstructionsIndex default implementation.
func InstructionsIndex(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("instructions/index.html"))
}
