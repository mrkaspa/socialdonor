package helpers

import (
	"os"

	"github.com/nitrous-io/go-mixpanel"
)

func GetMixClient() *mixpanel.Mixpanel {
	return mixpanel.NewMixpanelClient(os.Getenv("MIXPANEL_KEY"))
}

func TrackEvent(evt string, data map[string]interface{}) {
	GetMixClient().Track(evt, data)
}
