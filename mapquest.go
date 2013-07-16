package gogeo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	mapquestServiceUrl  = "http://open.mapquestapi.com/geocoding/v1/address?"
	mapquestServiceName = "MapQuest Geocoding"
)

// Ad-hoc structs for easy unmarshaling MapQuest's geocoding results
type mapquestResultSet struct {
	Results []struct {
		Locations []struct {
			LatLng struct {
				Long float64 `json:"lng"`
				Lat  float64 `json:"lat"`
			} `json:"latLng"`
		} `json:"locations"`
	} `json:"results"`
}

type MapQuestGeoCoder struct {
	serviceName string
}

// NewMapQuestGeoCoder creates a new GeoCoder for the MapQuest's
// geocoding service.
func NewMapQuestGeoCoder() *MapQuestGeoCoder {
	return &MapQuestGeoCoder{serviceName: mapquestServiceName}
}

// Geocode uses the MapQuest's geocoding service to translate
// a location into GPS coordinates.
func (gc *MapQuestGeoCoder) Geocode(loc *Location) (*GpsPoint, error) {
	if loc == nil {
		return nil, errors.New(fmt.Sprintf("%s: can't geocode a nil location", gc.serviceName))
	}
	body := strings.NewReader(fmt.Sprintf(`
{
    location: {
        street: "%s"
    }
}`, fmt.Sprintf("%s", loc.FullAddress)))
	resp, err := sendRequest(POST, mapquestServiceUrl, body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s: error sending request: %s", gc.serviceName, err.Error()))
	}

	dec := json.NewDecoder(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("%s: error %d - %s", gc.serviceName, resp.StatusCode, resp.Status))
	}
	var rset = new(mapquestResultSet)
	err = dec.Decode(rset)
	if err != nil {
		return nil, err
	}
	if len(rset.Results) > 0 {
		if len(rset.Results[0].Locations) > 0 {
			gps := rset.Results[0].Locations[0].LatLng
			return &GpsPoint{Lat: gps.Lat, Long: gps.Long}, nil
		}
	}
	return nil, nil
}
