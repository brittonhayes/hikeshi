package actions

import (
	"html/template"
	"strings"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush/v4"
)

var r *render.Engine
var assetsBox = packr.New("app:assets", "../public")

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.plush.html",

		// Box containing all of the templates:
		TemplatesBox: packr.New("app:templates", "../templates"),
		AssetsBox:    assetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{
			"activeClass": activeClass,
			"toTitle":     toTitle,
			"prettyDate":  prettyDate,
			"csrf": func() template.HTML {
				return `<input name="authenticity_token" value="<%= authenticity_token %>" type="hidden">`
			},
		},
	})
}

func activeClass(n string, help plush.HelperContext) string {
	if p, ok := help.Value("current_route").(buffalo.RouteInfo); ok {
		if p.PathName == n {
			return "active primary"
		}
	}
	return ""
}

func toTitle(help plush.HelperContext) (string, error) {
	s, err := help.Block()
	if err != nil {
		return "", err
	}
	return strings.Title(s), nil
}

func prettyDate(help plush.HelperContext) (string, error) {
	s, err := help.Block()
	if err != nil {
		return "", err
	}
	layoutRFC := "02 Jan 06 15:04 MST"
	t, _ := time.Parse(layoutRFC, s)
	return t.Format(layoutRFC), nil
}
