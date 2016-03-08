package models

import (
	"fmt"
	"os"
	"strconv"
	"time"

	fb "github.com/huandu/facebook"
	"github.com/mrkaspa/socialdonor/app/helpers"
	"github.com/revel/revel"
)

// User is the donor
type User struct {
	ID              int `gorm:"primary_key"`
	Name            string
	Email           string `sql:"default:null"`
	BloodType       string
	PhoneNumber     string
	City            string
	Country         string
	Lat             float64
	Lng             float64
	Available       bool `sql:"default:1"`
	FbId            string
	FbToken         string
	Requests        []Request
	FbExirationDate time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Validate the user in the complete request
func (u *User) Validate(validation *revel.Validation) {
	validation.Required(u.Email)
	validation.Email(u.Email)
	validation.Required(u.BloodType)
	validation.Required(u.Name)
	validation.Required(u.PhoneNumber)
	validation.Required(u.Lat)
	validation.Required(u.Lng)
}

func (u *User) AfterCreate() {
	properties := map[string]interface{}{
		"$first_name": u.Name,
	}
	go u.WelcomeFB()
	go helpers.GetMixClient().CreateProfile(strconv.Itoa(u.ID), properties)
}

func (u *User) AfterUpdate() (err error) {
	properties := map[string]interface{}{
		"$email":     u.Email,
		"$phone":     u.PhoneNumber,
		"blood_type": u.BloodType,
		"name":       u.Name,
		"city":       u.City,
		"country":    u.Country,
		"lat":        u.Lat,
		"lng":        u.Lng,
	}
	go helpers.GetMixClient().SetPropertiesOnProfileOnce(strconv.Itoa(u.ID), properties)
	return
}

func (u *User) GetLat() float64 {
	return u.Lat
}
func (u *User) GetLng() float64 {
	return u.Lng
}

func (u *User) SaveGeo(city, country string) {
	u.City = city
	u.Country = country
	Gdb.Save(&u)
}

func (u User) NotifyFB(requestId string, text string) {
	var globalApp = fb.New(os.Getenv("FACEBOOK_APP_ID"), os.Getenv("FACEBOOK_APP_SECRET"))
	var appAccess = globalApp.Session(globalApp.AppAccessToken())
	appAccess.Post(u.FbId+"/notifications", fb.Params{
		"href":     fmt.Sprintf("requests/%s/show", requestId),
		"template": text,
	})
}

func (u User) WelcomeFB() {
	var globalApp = fb.New(os.Getenv("FACEBOOK_APP_ID"), os.Getenv("FACEBOOK_APP_SECRET"))
	var appAccess = globalApp.Session(globalApp.AppAccessToken())
	appAccess.Post(u.FbId+"/notifications", fb.Params{
		"href":     "tour",
		"template": "Gracias por unirte a Gota de vida", //TODO mejorar texto
	})
}
