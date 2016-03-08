package controllers

import (
	"fmt"
	"strconv"

	appJobs "github.com/mrkaspa/socialdonor/app/jobs"
	"github.com/mrkaspa/socialdonor/app/models"
	"github.com/revel/modules/jobs/app/jobs"

	"github.com/revel/revel"
)

// Profile is the main controller
type Profile struct {
	SecuredController
}

// Index for complete info
func (c Profile) Index() revel.Result {
	c.LoadBloodTypes()
	if _, ok := c.Flash.Data["user.Email"]; !ok {
		user := (&c).Connected()
		c.Flash.Data["user.Email"] = user.Email
		c.Flash.Data["user.BloodType"] = user.BloodType
		c.Flash.Data["user.Name"] = user.Name
		c.Flash.Data["user.PhoneNumber"] = user.PhoneNumber
		c.Flash.Data["user.Available"] = strconv.FormatBool(user.Available)
		c.Flash.Data["user.Lat"] = fmt.Sprintf("%.8f", user.Lat)
		c.Flash.Data["user.Lng"] = fmt.Sprintf("%.8f", user.Lng)
	}
	return c.Render()
}

// Save the profile
func (c Profile) Save(user models.User) revel.Result {
	userSession := (&c).Connected()

	// Handle errors
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		c.Flash.Error("Corrige los errores por favor")
		return c.Redirect(Profile.Index)
	}

	if available, err := strconv.ParseBool(c.Params.Values["user.Available"][0]); err == nil {
		userSession.Available = available
	}

	// user.ID = userSession.ID
	userSession.Email = user.Email
	userSession.BloodType = user.BloodType
	userSession.Name = user.Name
	userSession.PhoneNumber = user.PhoneNumber
	userSession.Lng = user.Lng
	userSession.Lat = user.Lat

	if c.Txn.Save(userSession).Error != nil {
		c.FlashParams()
		c.Flash.Error("No se pudo guardar la solicitud")
		return c.Redirect(Profile.Index)
	}

	c.Flash.Success("Se ha actualizado tu información")
	jobs.Now(appJobs.GetGeo{userSession})
	jobs.Now(appJobs.Cartoo{userSession})

	return c.Redirect(Profile.Index)
}
