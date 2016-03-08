package jobs

//Reachable
type Reachable interface {
	GetLat() float64
	GetLng() float64
	SaveGeo(string, string)
}
