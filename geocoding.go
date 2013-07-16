package gogeo

import (
	"io"
	"net/http"
)

type Location struct {
	FullAddress string
}

// Equals compare 2 locations. src and dst locations are equal if the content of every
// field in src matches the content of the same field name in the dst location.
func (src *Location) Equals(dst *Location) bool {
	if dst == nil {
		return false
	}
	return src.FullAddress == dst.FullAddress
}

type HttpMethod string

const (
	GET    HttpMethod = "GET"
	POST              = "POST"
	PUT               = "PUT"
	DELETE            = "DELETE"
)

// GpsPoint defines a location by latitude and longitude.
type GpsPoint struct {
	Lat, Long float64
}

// GeoCoder is an interface used to translate street addresses into
// Gps coordinates. This process is called geocoding.
type GeoCoder interface {
	Geocode(*Location) (*GpsPoint, error)
}

func sendRequest(m HttpMethod, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(string(m), path, body)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
