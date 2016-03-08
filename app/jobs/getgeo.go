package jobs

import (
	"sort"

	"github.com/nf/geocode"
)

type GetGeo struct {
	Reachable Reachable
}

func (g GetGeo) Run() {
	req := &geocode.Request{
		Region:   "us",
		Provider: geocode.GOOGLE,
		Location: &geocode.Point{g.Reachable.GetLat(), g.Reachable.GetLng()},
	}
	if resp, err := req.Lookup(nil); err == nil {
		result := resp.GoogleResponse.Results[0]
		city := ""
		country := ""
		for _, addressPart := range result.AddressParts {
			if idx := sort.SearchStrings(addressPart.Types, "administrative_area_level_1"); idx < len(addressPart.Types) && addressPart.Types[idx] == "administrative_area_level_1" {
				city = addressPart.Name
			}
			if idx := sort.SearchStrings(addressPart.Types, "country"); idx < len(addressPart.Types) && addressPart.Types[idx] == "country" {
				country = addressPart.Name
			}
		}
		g.Reachable.SaveGeo(city, country)
	}
}
