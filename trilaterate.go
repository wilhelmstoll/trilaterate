// Copyright 2019 Wilhelm Stoll. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Pure go implementation of trilaterate function,
// transcribed in part from Python originals by StackExchange user
// wwnick from <http://gis.stackexchange.com/a/415/41129>.

package trilaterate

import (
	"math"

	"github.com/wilhelmstoll/mathx"
)

// Earth radius (in km)
const earthR = 6371.0

// Beacon represents a coordinate with a distance.
type Beacon struct {
	Lat  float64
	Lon  float64
	Dist float64
}

// Convert geodetic Lat/Long to ECEF xyz
//    1. Convert Lat/Long to radians
//    2. Convert Lat/Long(radians) to ECEF
func convertLatLonToECEF(lat float64, lon float64) [3]float64 {
	x := earthR * (math.Cos(mathx.Rad(lat)) * math.Cos(mathx.Rad(lon)))
	y := earthR * (math.Cos(mathx.Rad(lat)) * math.Sin(mathx.Rad(lon)))
	z := earthR * (math.Sin(mathx.Rad(lat)))

	return [3]float64{x, y, z}
}

// Solve performs a trilateration calculation to determine a location
// based on 3 beacons and their respective distances (in km) to the desired location.
// It returns the calculated location (in lat/lon).
func Solve(b1 Beacon, b2 Beacon, b3 Beacon) (float64, float64) {
	// assuming elevation = 0

	// using authalic sphere
	// if using an ellipsoid this step is slightly different
	// Convert geodetic Lat/Long to ECEF xyz
	p1 := convertLatLonToECEF(b1.Lat, b1.Lon)
	p2 := convertLatLonToECEF(b2.Lat, b2.Lon)
	p3 := convertLatLonToECEF(b3.Lat, b3.Lon)

	// from wikipedia
	// transform to get circle 1 at origin
	// transform to get circle 2 on x axis
	ex := mathx.Divide(mathx.Subtract(p2, p1), mathx.Norm(mathx.Subtract(p2, p1)))
	i := mathx.Dot(ex, mathx.Subtract(p3, p1))

	ey := mathx.Divide(
		mathx.Subtract(
			mathx.Subtract(p3, p1),
			mathx.Multiply(ex, i)),
		mathx.Norm(
			mathx.Subtract(
				mathx.Subtract(p3, p1),
				mathx.Multiply(ex, i))))

	ez := mathx.Cross(ex, ey)
	d := mathx.Norm(mathx.Subtract(p2, p1))
	j := mathx.Dot(ey, mathx.Subtract(p3, p1))

	// from wikipedia
	// plug and chug using above values
	x := (math.Pow(b1.Dist, 2) - math.Pow(b2.Dist, 2) + math.Pow(d, 2)) / (2 * d)
	y := ((math.Pow(b1.Dist, 2) - math.Pow(b3.Dist, 2) + math.Pow(i, 2) + math.Pow(j, 2)) / (2 * j)) - ((i / j) * x)

	// only one case shown here
	//
	// This line was editing from the original post. If the radical value was negative, it did not work.
	// I solved the problem by calc with the absolute value.
	z := math.Sqrt(math.Abs(math.Pow(b1.Dist, 2) - math.Pow(x, 2) - math.Pow(y, 2)))

	// triPt is an array with ECEF x,y,z of trilateration point
	triPt := mathx.Add(
		mathx.Add(
			mathx.Add(p1,
				mathx.Multiply(ex, x)),
			mathx.Multiply(ey, y)),
		mathx.Multiply(ez, z))

	// convert back to lat/long from ECEF
	// convert to degrees
	lat := mathx.Deg(math.Asin(triPt[2] / earthR))
	lon := mathx.Deg(math.Atan2(triPt[1], triPt[0]))

	return lat, lon
}
