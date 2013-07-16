// Copyright 2013 Mathias Monnerville. All rights reserved.
// Use of this source code is governed by a GPL
// license that can be found in the LICENSE file.

package gogeo

import (
	"errors"
	"fmt"
)

// Fake geocoder that do not issue any outgoing HTTP requests.
type testGeoCoder struct {
	serviceName string
}

// Describes a fake position
type testGps struct {
	loc *Location
	gps *GpsPoint
}

var gc GeoCoder

func init() {
	gc = newTestGeoCoder()
}

// Some GPS fixes for the tests
var gpsFixes = map[string]*testGps{
	"montgallet": {
		loc: &Location{
			FullAddress: "rue Montgallet, Paris",
		},
		gps: &GpsPoint{48.843517, 2.387779},
	},
	"frontignan": {
		loc: &Location{
			FullAddress: "Frontignan",
		},
		gps: &GpsPoint{43.448762, 3.753064},
	},
}

func newTestGeoCoder() *testGeoCoder {
	return &testGeoCoder{serviceName: "Fake geocoder"}
}

func (gc *testGeoCoder) Geocode(loc *Location) (*GpsPoint, error) {
	if loc == nil {
		return nil, errors.New(fmt.Sprintf("%s: can't geocode a nil location", gc.serviceName))
	}
	for _, f := range gpsFixes {
		if f.loc.Equals(loc) {
			return f.gps, nil
		}
	}
	return nil, errors.New("gps fix for testing not found")
}
