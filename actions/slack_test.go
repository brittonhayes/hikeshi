package actions_test

import (
	"github.com/brittonhayes/hikeshi/actions"
	"github.com/brittonhayes/hikeshi/models"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/uuid"
	"testing"
	"time"
)

func TestSendSlackNotification(t *testing.T) {
	type args struct {
		baseURL    string
		webhookURL string
		incident   *models.Incident
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Base test", args{
			"http://localhost:3000",
			envy.Get("SLACK_URL", ""),
			&models.Incident{
				ID:                uuid.UUID{},
				Date:              time.Time{},
				DateClosed:        time.Time{},
				Severity:          "High",
				Title:             "Example incident",
				Summary:           "summary",
				Scope:             "example scope",
				ResponsibleParty:  "Example party",
				Result:            "example result",
				Mitigation:        "example mitigation",
				AffectedCustomers: "example affected",
				RootCause:         "example root cause",
				SlackChannel:      "#slack-channel-example",
				CreatedAt:         time.Time{},
				UpdatedAt:         time.Time{},
			},
		},
			false,
		},
		{"Blocks", args{
			"http://localhost:3000",
			envy.Get("SLACK_URL", ""),
			&models.Incident{
				ID:                uuid.UUID{},
				Date:              time.Time{},
				DateClosed:        time.Time{},
				Severity:          "High",
				Title:             "Example incident",
				Summary:           "summary",
				Scope:             "example scope",
				ResponsibleParty:  "Example party",
				Result:            "example result",
				Mitigation:        "example mitigation",
				AffectedCustomers: "example affected",
				RootCause:         "example root cause",
				SlackChannel:      "#slack-channel-example",
				CreatedAt:         time.Time{},
				UpdatedAt:         time.Time{},
			},
		},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := actions.SendSlackNotification(tt.args.baseURL, tt.args.webhookURL, *tt.args.incident); (err != nil) != tt.wantErr {
				t.Errorf("SendSlackNotification() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
