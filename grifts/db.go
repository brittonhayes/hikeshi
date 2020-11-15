package grifts

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"

	"github.com/brittonhayes/hikeshi/models"
	"github.com/gofrs/uuid"
	. "github.com/markbates/grift/grift"
)

var _ = Namespace("db", func() {
	_ = Namespace("seed", func() {
		summary := `Lorem ipsum dolor sit amet, consectetur adipisicing elit. Blanditiis consequatur doloribus eius et,
                itaque molestias neque officia quaerat! Ab asperiores expedita fugit ipsa itaque, iure iusto nesciunt
                odio quaerat quidem rerum, sint sit, temporibus vel veniam. Aperiam dolor, dolore dolorum error facilis
                hic molestiae odit repellendus similique temporibus! Aliquam consequatur deserunt earum, fugiat, illum
                ipsam libero molestias, quasi quo repellat sed soluta velit veniam! A numquam quaerat quia tempora
                veniam!`

		_ = Desc("incidents", "Seed incidents into the database")
		_ = Add("incidents", func(c *Context) error {
			i := &models.Incident{
				ID:                uuid.UUID{},
				Date:              time.Time{},
				DateClosed:        time.Time{},
				Severity:          "High",
				Title:             "Example incident",
				Summary:           summary,
				Scope:             "example scope",
				ResponsibleParty:  "Example party",
				Result:            "example result",
				Mitigation:        "example mitigation",
				AffectedCustomers: "example affected",
				RootCause:         "example root cause",
				SlackChannel:      "#slack-channel-example",
				CreatedAt:         time.Time{},
				UpdatedAt:         time.Time{},
			}

			err := models.DB.Create(i)
			if err != nil {
				log.Fatal(err)
			}
			return nil
		})

		_ = Add("users", func(c *Context) error {

			passwd := "Password"
			passwdHash, _ := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)

			u := []models.User{
				{
					ID:           uuid.UUID{},
					FirstName:    "John",
					LastName:     "Doe",
					Email:        "jdoe@example.com",
					Role:         "admin",
					Password:     "Password",
					PasswordHash: string(passwdHash),
					CreatedAt:    time.Time{},
					UpdatedAt:    time.Time{},
				},
				{
					ID:           uuid.UUID{},
					FirstName:    "Jane",
					LastName:     "Doseph",
					Email:        "jdoseph@example.com",
					Role:         "analyst",
					Password:     "Password",
					PasswordHash: string(passwdHash),
					CreatedAt:    time.Time{},
					UpdatedAt:    time.Time{},
				},
				{
					ID:           uuid.UUID{},
					FirstName:    "Jeremy",
					LastName:     "Does",
					Email:        "jdoes@example.com",
					Role:         "guest",
					Password:     "Password",
					PasswordHash: string(passwdHash),
					CreatedAt:    time.Time{},
					UpdatedAt:    time.Time{},
				},
			}

			err := models.DB.Create(u)
			if err != nil {
				log.Fatal(err)
			}
			return nil
		})

	})

})
