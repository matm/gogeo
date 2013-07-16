package gogeo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	googleServiceUrl  = "https://maps.googleapis.com/maps/api/geocode/json?"
	googleServiceName = "Google Geocoding"
)

// Ad-hoc structs for easy unmarshaling Google's geocoding results
type googleResultSet struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Long float64 `json:"lng"`
				Lat  float64 `json:"lat"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
}

type GoogleGeoCoder struct {
	serviceName string
}

// NewGoogleGeoCoder creates a new GeoCoder for the MapQuest's
// geocoding service.
func NewGoogleGeoCoder() *GoogleGeoCoder {
	return &GoogleGeoCoder{serviceName: googleServiceName}
}

// Geocode uses the Google's geocoding service API to translate
// a location into GPS coordinates.
func (gc *GoogleGeoCoder) Geocode(loc *Location) (*GpsPoint, error) {
	if loc == nil {
		return nil, errors.New(fmt.Sprintf("%s: can't geocode a nil location", gc.serviceName))
	}

	resp, err := sendRequest(GET, googleServiceUrl+fmt.Sprintf("address=%s&sensor=false", loc.FullAddress), nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s: error sending request: %s", gc.serviceName, err.Error()))
	}

	dec := json.NewDecoder(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("%s: error %d - %s", gc.serviceName, resp.StatusCode, resp.Status))
	}
	var rset = new(googleResultSet)
	err = dec.Decode(rset)
	if err != nil {
		return nil, err
	}
	if len(rset.Results) > 0 {
		gps := rset.Results[0].Geometry.Location
		return &GpsPoint{Lat: gps.Lat, Long: gps.Long}, nil
	}
	return nil, nil
}
