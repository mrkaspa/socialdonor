package jobs

import (
	"fmt"

	"github.com/mrkaspa/socialdonor/app/models"
)

const EARTH_RADIUS = 6371

type Alerts struct {
	Request *models.Request
}

var bloodCompatibility = map[string][]string{
	"A+":  []string{"O+", "O-", "A+", "A-"},
	"A-":  []string{"O-", "A-"},
	"B+":  []string{"O+", "O-", "B+", "B-"},
	"B-":  []string{"O-", "B-"},
	"AB+": []string{"O+", "O-", "A+", "A-", "B+", "B-", "AB+", "AB-"},
	"AB-": []string{"O-", "AB-", "A-", "B-"},
	"O+":  []string{"O+"},
	"O-":  []string{"O-"},
}

// Looks for the nearest donors and set
func (a Alerts) Run() {
	users, _ := a.PointsWithinRadius(1000)
	for _, user := range users {
		user.NotifyFB(a.Request.UUID, "Nueva solicitud de sangre")
	}
}

func (a Alerts) PointsWithinRadius(radius float64) (users []models.User, err error) {
	in := ""
	if value, ok := bloodCompatibility[a.Request.BloodType]; ok {
		for i, iter := range value {
			in += "'" + iter + "'"
			if i+1 < len(value) {
				in += ","
			}
		}
	}
	lat1 := fmt.Sprintf("sin(radians(%f)) * sin(radians(lat))", a.Request.Lat)
	lng1 := fmt.Sprintf("cos(radians(%f)) * cos(radians(lat)) * cos(radians(lng) - radians(%f))", a.Request.Lat, a.Request.Lng)
	whereGeo := fmt.Sprintf("acos(%s + %s) * %f <= %f", lat1, lng1, float64(EARTH_RADIUS), radius)
	whereUser := fmt.Sprintf("id != %d and blood_type in (%s) and available = 1", a.Request.UserID, in)
	whereAll := fmt.Sprintf("%s and %s", whereUser, whereGeo)
	err = models.Gdb.Find(&users, whereAll).Error
	return
}
