package grifts

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"log"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/brittonhayes/hikeshi/models"
	"github.com/gofrs/uuid"
	. "github.com/markbates/grift/grift"
)

func randomDate() time.Time {
	min := time.Date(2017, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2021, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
func randomIncident() string {
	profile := randomdata.GenerateProfile(randomdata.Male | randomdata.Female | randomdata.RandomGender)
	return fmt.Sprintf("The Security incident was initially caused by %s %s having their password compromised. This severity event was performed by a malicious actor who appeared under the username, %s when authenticating to our single sign-on service. The employee who had their credentials compromised is based in our %s office. The attacker seems to be targeting employees of %s origin.", profile.Name.First, profile.Name.Last, profile.Login.Username, profile.Location.City, profile.Nat)
}

var _ = Namespace("db", func() {
	_ = Namespace("seed", func() {
		_ = Desc("incidents", "Seed incidents into the database")
		_ = Add("incidents", func(c *Context) error {
			for i := 0; i < 15; i++ {
				incident := &models.Incident{
					ID:                uuid.UUID{},
					Date:              randomDate(),
					DateClosed:        randomDate(),
					Closed:            randomdata.Boolean(),
					Severity:          randomdata.StringSample("Low", "Moderate", "High", "Critical"),
					Title:             randomdata.StringSample("Compromised password", "Leaked API token", "Github repo public", "Insider threat"),
					Summary:           randomIncident(),
					Scope:             fmt.Sprintf("%v internal stakeholders", randomdata.Number(10, 50)),
					ResponsibleParty:  randomdata.FullName(randomdata.RandomGender),
					Result:            randomdata.StringSample("The situation was resolved by the user", "The incident is ongoing", "We are in communication with authorities"),
					Mitigation:        randomdata.StringSample("All credentials rotated", "Affected customers notified", "Employee enrolled in thorough security training"),
					AffectedCustomers: fmt.Sprintf("%v customers were affected", randomdata.Number(10, 10000)),
					RootCause:         randomdata.StringSample("Compromised password", "Leaked API token", "Github repo public", "Insider threat"),
					SlackChannel:      fmt.Sprintf("#security-incident-%v", randomdata.Number(12, 200)),
					CreatedAt:         randomDate(),
					UpdatedAt:         randomDate(),
				}
				err := models.DB.Create(incident)
				if err != nil {
					log.Fatal(err)
				}
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
