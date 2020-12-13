package middleware

import (
	"github.com/gobuffalo/buffalo"
)

func DarkMode(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		mode, _ := c.Cookies().Get("halfmoon_preferredMode")
		switch mode {
		case "light-mode":
			c.Set("halfmoon_preferredMode", "light-mode")
		default:
			c.Set("halfmoon_preferredMode", "dark-mode")
		}
		err := next(c)
		return err
	}
}
