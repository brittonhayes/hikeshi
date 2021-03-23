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
	users := &models.Users{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	queryOpen := tx.Where("closed = ?", false)
	countOpen, err := queryOpen.Count(incidents)
	if err != nil {
		return err
	}

	queryClosed := tx.Where("closed = ?", true)
	countClosed, err := queryClosed.Count(incidents)
	if err != nil {
		return err
	}

	queryUsers := tx.Where("role = 'admin' OR role = 'analyst' OR role = 'guest'")
	countUsers, err := queryUsers.Count(users)
	if err != nil {
		return err
	}

	if err := tx.Limit(5).All(incidents); err != nil {
		return err
	}

	c.Set("OpenIncidents", countOpen)
	c.Set("ClosedIncidents", countClosed)
	c.Set("incidents", incidents)
	c.Set("ActiveUsers", countUsers)
	c.Set("pagination", q.Paginator)
	return c.Render(http.StatusOK, r.HTML("index.html"))
}
