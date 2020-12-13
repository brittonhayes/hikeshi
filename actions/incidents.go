package actions

import (
	"fmt"
	"net/http"

	"github.com/brittonhayes/hikeshi/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/x/responder"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Incident)
// DB Table: Plural (incidents)
// Resource: Plural (Incidents)
// Path: Plural (/incidents)
// View Template Folder: Plural (/templates/incidents/)

// IncidentsResource is the resource for the Incident model
type IncidentsResource struct {
	buffalo.Resource
}

// List gets all Incidents. This function is mapped to the path
// GET /incidents
func (v IncidentsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	incidents := &models.Incidents{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Incidents from the DB
	if err := q.All(incidents); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// Add the paginator to the context so it can be used in the template.
		c.Set("pagination", q.Paginator)
		c.Set("incidents", incidents)
		return c.Render(http.StatusOK, r.HTML("/incidents/index.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(incidents))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(incidents))
	}).Respond(c)
}

// Show gets the data for one Incident. This function is mapped to
// the path GET /incidents/{incident_id}
func (v IncidentsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Incident
	incident := &models.Incident{}

	// To find the Incident the parameter incident_id is used.
	if err := tx.Find(incident, c.Param("incident_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		c.Set("incident", incident)
		c.Set("maxlength", false)
		if len(incident.Summary) > MaxStringLen {
			c.Set("maxlength", true)
		}
		return c.Render(http.StatusOK, r.HTML("/incidents/show.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(incident))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(incident))
	}).Respond(c)
}

// New renders the form for creating a new Incident.
// This function is mapped to the path GET /incidents/new
func (v IncidentsResource) New(c buffalo.Context) error {
	c.Set("incident", &models.Incident{})

	return c.Render(http.StatusOK, r.HTML("/incidents/new.plush.html"))
}

// Create adds a Incident to the DB. This function is mapped to the
// path POST /incidents
func (v IncidentsResource) Create(c buffalo.Context) error {
	// Allocate an empty Incident
	incident := &models.Incident{}

	// Bind incident to the html form elements
	if err := c.Bind(incident); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(incident)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the new.html template that the user can
			// correct the input.
			c.Set("incident", incident)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("/incidents/new.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "incident.created.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/incidents/%v", incident.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.JSON(incident))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.XML(incident))
	}).Respond(c)
}

// Edit renders a edit form for a Incident. This function is
// mapped to the path GET /incidents/{incident_id}/edit
func (v IncidentsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Incident
	incident := &models.Incident{}

	if err := tx.Find(incident, c.Param("incident_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("incident", incident)
	return c.Render(http.StatusOK, r.HTML("/incidents/edit.plush.html"))
}

// Update changes a Incident in the DB. This function is mapped to
// the path PUT /incidents/{incident_id}
func (v IncidentsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Incident
	incident := &models.Incident{}

	if err := tx.Find(incident, c.Param("incident_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Incident to the html form elements
	if err := c.Bind(incident); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(incident)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the edit.html template that the user can
			// correct the input.
			c.Set("incident", incident)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("/incidents/edit.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "incident.updated.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/incidents/%v", incident.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(incident))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(incident))
	}).Respond(c)
}

// Destroy deletes a Incident from the DB. This function is mapped
// to the path DELETE /incidents/{incident_id}
func (v IncidentsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Incident
	incident := &models.Incident{}

	// To find the Incident the parameter incident_id is used.
	if err := tx.Find(incident, c.Param("incident_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(incident); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a flash message
		c.Flash().Add("success", T.Translate(c, "incident.destroyed.success"))

		// Redirect to the index page
		return c.Redirect(http.StatusSeeOther, "/incidents")
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(incident))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(incident))
	}).Respond(c)
}
