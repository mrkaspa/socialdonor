package models

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/mrkaspa/socialdonor/app/helpers"

	"github.com/revel/revel"
)

//Request for blood
type Request struct {
	ID          int `gorm:"primary_key"`
	UUID        string
	BloodType   string
	Description string
	PatientName string
	Lat         float64
	Lng         float64
	UserID      int  `sql:"index"`
	Approved    bool `sql:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Validate the user in the complete request
func (r *Request) Validate(validation *revel.Validation) {
	validation.Required(r.BloodType)
	validation.Required(r.Description)
	validation.Required(r.PatientName)
	validation.Required(r.Lat)
	validation.Required(r.Lng)
}

// BeforeCreate is callback
func (r *Request) BeforeCreate() {
	r.UUID, _ = newUUID()
}

// AfterCreate is callback
func (r *Request) AfterCreate() {
	properties := map[string]interface{}{
		"id":         r.ID,
		"blood_type": r.BloodType,
		"lat":        r.Lat,
		"lng":        r.Lng,
	}
	go r.sendSlackMessage()
	go helpers.TrackEvent("new_request", properties)
}

func (r *Request) sendSlackMessage() {
	messageTemplate := `
  {
    "text": "El paciente %s necesita sangre %s",
    "icon_url": "https://gotadevida.co/public/img/logo.png",
    "username": "BloodBot",
    "attachments": [{
      "mrkdwn": true,
      "fields": [
        {
          "title": "Descripción",
          "value": "%s"
        },
        {
          "title": "Aprobar",
          "value": "<%s|:+1:>",
          "short": true
        }
      ]
    }]
  }
  `
	approvalURL := fmt.Sprintf("https://"+"%s/requests/kr3atti3w3/%s/approve", os.Getenv("DOMAIN"), r.UUID)
	message := fmt.Sprintf(messageTemplate, r.PatientName, r.BloodType, r.Description, approvalURL)
	postSlack(message)
}

func postSlack(message string) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", os.Getenv("SLACK_PATH"), bytes.NewBuffer([]byte(message)))
	req.Header.Set("Content-Type", "application/json")
	return client.Do(req)
}

func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
