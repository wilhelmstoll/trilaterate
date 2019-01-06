// Copyright 2019 Wilhelm Stoll. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package trilaterate

import (
	"testing"

	"github.com/wilhelmstoll/testx"
)

func TestSolve(t *testing.T) {
	b1 := Beacon{Lat: 35.000000, Lon: -120.000000, Dist: 189.419265289145}
	b2 := Beacon{Lat: 35.000005, Lon: -120.000010, Dist: 189.420325082156}
	b3 := Beacon{Lat: 35.000000, Lon: -120.000020, Dist: 189.420689733286}

	lat, lon := Solve(b1, b2, b3)

	testx.EqualFloat(t, 33.997467251559286, lat, "Latitude")
	testx.EqualFloat(t, -118.39745961414125, lon, "Longitude")
}
