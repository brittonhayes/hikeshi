package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// Incident is used by pop to map your incidents database table to your go code.
type Incident struct {
	ID                uuid.UUID `json:"id" db:"id"`
	Date              time.Time `json:"date" db:"date"`
	DateClosed        time.Time `json:"date_closed" db:"date_closed"`
	Severity          string    `json:"severity" db:"severity"`
	Title             string    `json:"title" db:"title"`
	Summary           string    `json:"summary" db:"summary"`
	Scope             string    `json:"scope" db:"scope"`
	ResponsibleParty  string    `json:"responsible_party" db:"responsible_party"`
	Result            string    `json:"result" db:"result"`
	Mitigation        string    `json:"mitigation" db:"mitigation"`
	AffectedCustomers string    `json:"affected_customers" db:"affected_customers"`
	RootCause         string    `json:"root_cause" db:"root_cause"`
	SlackChannel      string    `json:"slack_channel" db:"slack_channel"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (i Incident) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// Incidents is not required by pop and may be deleted
type Incidents []Incident

// String is not required by pop and may be deleted
func (i Incidents) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (i *Incident) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (i *Incident) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (i *Incident) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
