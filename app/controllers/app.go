package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"code.google.com/p/goauth2/oauth"
	appJobs "github.com/mrkaspa/socialdonor/app/jobs"
	"github.com/mrkaspa/socialdonor/app/models"
	"github.com/revel/modules/jobs/app/jobs"
	"github.com/revel/revel"
)

// App is the mai	n controller
type App struct {
	GormController
}

var facebook *oauth.Config

func InitFacebook() {
	facebook = &oauth.Config{
		ClientId:     os.Getenv("FACEBOOK_APP_ID"),
		ClientSecret: os.Getenv("FACEBOOK_APP_SECRET"),
		AuthURL:      "https://graph.facebook.com/oauth/authorize",
		TokenURL:     "https://graph.facebook.com/oauth/access_token",
		RedirectURL:  fmt.Sprintf("http://%s/app/auth", os.Getenv("DOMAIN")),
		Scope:        "email public_profile user_friends",
	}
}

func (c *App) SetHeaders() revel.Result {
	c.Response.Out.Header().Add("X-Frame-Options", "ALLOW-FROM facebook.com")
	return nil
}

//Index is the main entry point of the app
func (c App) Index() revel.Result {
	user := (&c).Connected()
	if _, ok := c.Session["token"]; ok {
		resp, _ := http.Get("https://graph.facebook.com/me?fields=id,name,email,location&access_token=" +
			url.QueryEscape(c.Session["token"]))
		defer resp.Body.Close()
		me := map[string]interface{}{}
		if err := json.NewDecoder(resp.Body).Decode(&me); err != nil {
			revel.ERROR.Println(err)
		}
		if _, ok := me["id"]; ok {
			user.FbId = me["id"].(string)
			var userFound models.User
			if c.Txn.Where("fb_id = ?", user.FbId).First(&userFound); userFound.ID != 0 {
				user = &userFound
			}
			user.Name = me["name"].(string)
			revel.INFO.Println("fb info")
			revel.INFO.Println(me)
			if user.Email == "" && me["email"] != nil {
				user.Email = me["email"].(string)
			}
			user.FbToken = c.Session["token"]
			t1, _ := time.Parse(time.RFC3339, c.Session["token_expiry"])
			user.FbExirationDate = t1
			c.Txn.Save(user)
			c.Session["uid"] = strconv.Itoa(user.ID)
			if user.ID != 0 {
				if user.BloodType == "" {
					return c.Redirect(Profile.Index)
				}
				return c.Redirect(App.Map)
			}
		} else {
			delete(c.Session, "token")
		}
	}
	authURL := facebook.AuthCodeURL("")
	return c.Render(authURL)
}

//Auth is the callback for facebook
func (c App) Auth(code string) revel.Result {
	revel.INFO.Println("Entro a Auth")
	t := &oauth.Transport{Config: facebook}
	tok, err := t.Exchange(code)
	if err != nil {
		revel.ERROR.Println(err)
		return c.Redirect(App.Index)
	}

	revel.INFO.Println(tok.AccessToken)
	c.Session["token"] = tok.AccessToken
	c.Session["token_expiry"] = tok.Expiry.Format(time.RFC3339)
	return c.Redirect(App.Index)
}

//Logout closes the session
func (controller App) Logout() revel.Result {
	delete(controller.Session, "token")
	delete(controller.Session, "token_expiry")
	delete(controller.Session, "uid")
	return controller.Redirect(App.Index)
}

func (c App) Map() revel.Result {
	return c.Render()
}

func (c App) Terms() revel.Result {
	return c.Render()
}

func (c App) Privacy() revel.Result {
	return c.Render()
}

func (c App) Tour() revel.Result {
	return c.Render()
}

// Show the request for Blood
func (c App) Show(uuid string) revel.Result {
	var request models.Request
	if c.Txn.Where("uuid = ?", uuid).First(&request); request.ID != 0 {
		var userRequest models.User
		if c.Txn.Model(&request).Related(&userRequest); userRequest.ID != 0 {
			return c.Render(request, userRequest)
		}
	}
	return c.NotFound("No se encontro la solicitud")
}

// Approve the request for Blood
func (c App) Approve(uuid string) revel.Result {
	var request models.Request
	if c.Txn.Where("uuid = ? and approved = 0", uuid).First(&request); request.ID != 0 {
		request.Approved = false
		if c.Txn.Save(&request).Error == nil {
			jobs.Now(appJobs.Alerts{&request})
		}
	}
	return c.Redirect(App.Index)
}
