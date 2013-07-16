// Copyright 2013 Mathias Monnerville. All rights reserved.
// Use of this source code is governed by a GPL
// license that can be found in the LICENSE file.

package gogeo

import (
	"testing"
)

func TestGeocode(t *testing.T) {
	gps, err := gc.Geocode(gpsFixes["frontignan"].loc)
	if err != nil {
		t.Errorf("error geocoding: %s", err.Error())
	}
	if gps == nil {
		t.Fatal("no GPS coordinates returned")
	}
	if gps.Lat != 43.448762 || gps.Long != 3.753064 {
		t.Errorf("wrong GPS coords for Frontignan, got (%v, %v)", gps.Lat, gps.Long)
	}
}
