package actions

import (
	"net/http"
	"time"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	c.Cookies().Set("halfmoon_preferredMode", "dark-mode", 30*24*time.Hour)
	return c.Render(http.StatusOK, r.HTML("index.html"))
}
