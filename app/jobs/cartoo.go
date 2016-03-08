package jobs

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/mrkaspa/socialdonor/app/models"
	"github.com/revel/revel"
)

type Cartoo struct {
	User *models.User
}

func (c Cartoo) Run() {
	if err := doDelete(c.User); err != nil {
		revel.INFO.Println(err)
	}
	if err := doUpdate(c.User); err != nil {
		revel.INFO.Println(err)
	}
}

func doUpdate(user *models.User) error {
	query := fmt.Sprintf("INSERT INTO untitled_table (fb_id,email,blood_type,city,country,name,the_geom) "+
		"VALUES ('"+user.FbId+"', '"+user.Email+"','"+user.BloodType+"','"+user.City+"','"+user.Country+"','"+user.Name+"',ST_SetSRID(ST_MakePoint(%f, %f),4326))", user.Lng, user.Lat)
	_, err := makeRequest("POST", query)
	return err
}

func doDelete(user *models.User) error {
	query := "DELETE FROM untitled_table WHERE fb_id='" + user.FbId + "'"
	_, err := makeRequest("DELETE", query)
	return err
}

func baseURI() string {
	return "https://" + os.Getenv("CARTOO_DOMAIN") + ".cartodb.com/api/v2/sql"
}

func makeRequest(method, query string) (*http.Response, error) {
	client := &http.Client{}
	values := url.Values{"api_key": {os.Getenv("CARTOO_TOKEN")}, "q": {query}}
	path := fmt.Sprintf("%s?%s", baseURI(), values.Encode())
	req, _ := http.NewRequest(method, path, nil)
	return client.Do(req)
}
