package actions

import (
	"fmt"
	"github.com/brittonhayes/hikeshi/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"net/http"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	incidents := &models.Incidents{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Incidents from the DB
	count, err := q.Count(incidents)
	if err != nil {
		return err
	}

	c.Set("numIncidents", count)

	return c.Render(http.StatusOK, r.HTML("index.html"))
}
