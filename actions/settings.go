package actions

import (
    "net/http"

	"github.com/gobuffalo/buffalo"
)


// TODO Allow enabling and disabling webhooks to Slack

// SettingsIndex default implementation.
func SettingsIndex(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("settings/index.html"))
}

// SettingsEdit default implementation.
func SettingsEdit(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("settings/edit.html"))
}

