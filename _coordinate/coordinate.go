package _coordinate

import (
	"math"
)

const A = 6378137 // 6378245
const EE = 0.00669342162296594323
const VR = math.Pi / 180
const MR = 180 / math.Pi
const YR = 6378137
const EPSG3857MACLAT = 85.0511287798

func rad(d float64) float64 {
	return d * math.Pi / 180.0
}
func isOutOfChina(lngLat [2]float64) bool {
	lon, lat := lngLat[0], lngLat[1]
	if lon < 72.004 || lon > 137.8347 {
		return true
	}
	if lat < 0.8293 || lat > 55.8271 {
		return true
	}
	return false
}
func transformLat(x float64, y float64) float64 {
	ret := -100.0 + 2.0*x + 3.0*y + 0.2*y*y + 0.1*x*y + 0.2*math.Sqrt(math.Abs(x))
	ret += (20.0*math.Sin(6.0*x*math.Pi) + 20.0*math.Sin(2.0*x*math.Pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(y*math.Pi) + 40.0*math.Sin(y/3.0*math.Pi)) * 2.0 / 3.0
	ret += (160.0*math.Sin(y/12.0*math.Pi) + 320.0*math.Sin(y*math.Pi/30.0)) * 2.0 / 3.0
	return ret
}
func transformLon(x float64, y float64) float64 {
	ret := 300.0 + x + 2.0*y + 0.1*x*x + 0.1*x*y + 0.1*math.Sqrt(math.Abs(x))
	ret += (20.0*math.Sin(6.0*x*math.Pi) + 20.0*math.Sin(2.0*x*math.Pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(x*math.Pi) + 40.0*math.Sin(x/3.0*math.Pi)) * 2.0 / 3.0
	ret += (150.0*math.Sin(x/12.0*math.Pi) + 300.0*math.Sin(x/30.0*math.Pi)) * 2.0 / 3.0
	return ret
}
func transform(lngLat [2]float64) [2]float64 {
	lon, lat := lngLat[0], lngLat[1]
	if isOutOfChina(lngLat) {
		return [2]float64{lon, lat}
	}
	dLat := transformLat(lon-105.0, lat-35.0)
	dLon := transformLon(lon-105.0, lat-35.0)
	radLat := lat / 180.0 * math.Pi
	magic := math.Sin(radLat)
	magic = 1 - EE*magic*magic
	sqrtMagic := math.Sqrt(magic)
	dLat = (dLat * 180.0) / ((A * (1 - EE)) / (magic * sqrtMagic) * math.Pi)
	dLon = (dLon * 180.0) / (A / sqrtMagic * math.Cos(radLat) * math.Pi)
	return [2]float64{lon + dLon, lat + dLat}
}
func convertWGS84ToGCJ02(lngLat [2]float64) [2]float64 {
	lon, lat := lngLat[0], lngLat[1]
	if isOutOfChina(lngLat) {
		return [2]float64{lon, lat}
	}
	dLat := transformLat(lon-105.0, lat-35.0)
	dLon := transformLon(lon-105.0, lat-35.0)
	radLat := lat / 180.0 * math.Pi
	magic := math.Sin(radLat)
	magic = 1 - EE*magic*magic
	sqrtMagic := math.Sqrt(magic)
	dLat = (dLat * 180.0) / ((A * (1 - EE)) / (magic * sqrtMagic) * math.Pi)
	dLon = (dLon * 180.0) / (A / sqrtMagic * math.Cos(radLat) * math.Pi)
	return [2]float64{lon + dLon, lat + dLat}
}
func convertGCJ02ToBD09(lngLat [2]float64) [2]float64 {
	lon, lat := lngLat[0], lngLat[1]
	x := lon
	y := lat
	z := math.Sqrt(x*x+y*y) + 0.00002*math.Sin(y*math.Pi)
	theta := math.Atan2(y, x) + 0.000003*math.Cos(x*math.Pi)
	bdLon := z*math.Cos(theta) + 0.0065
	bdLat := z*math.Sin(theta) + 0.006
	return [2]float64{bdLon, bdLat}
}
func convertBD09ToGCJ02(lngLat [2]float64) [2]float64 {
	bdLon, bdLat := lngLat[0], lngLat[1]
	x := bdLon - 0.0065
	y := bdLat - 0.006
	z := math.Sqrt(x*x+y*y) - 0.00002*math.Sin(y*math.Pi)
	theta := math.Atan2(y, x) - 0.000003*math.Cos(x*math.Pi)
	gcjLon := z * math.Cos(theta)
	gcjLat := z * math.Sin(theta)
	return [2]float64{gcjLon, gcjLat}
}
func convertGCJ02ToWGS84(lngLat [2]float64) [2]float64 {
	t := transform(lngLat)
	lon := t[0]
	lat := t[1]
	wgsLon := lngLat[0]*2 - lon
	wgsLat := lngLat[1]*2 - lat
	return [2]float64{wgsLon, wgsLat}
}
func getWGS84Distance(lngLatFrom [2]float64, lngLatTo [2]float64) float64 {
	lon1 := lngLatFrom[0]
	lat1 := lngLatFrom[1]
	lon2 := lngLatTo[0]
	lat2 := lngLatTo[1]
	radLat1 := rad(lat1)
	radLat2 := rad(lat2)
	a := radLat1 - radLat2
	b := rad(lon1) - rad(lon2)
	s := 2 * math.Asin(math.Sqrt(math.Pow(math.Sin(a/2), 2)+math.Cos(radLat1)*math.Cos(radLat2)*math.Pow(math.Sin(b/2), 2)))
	s = s * A
	return s
}
func getGCJ02Distance(lngLatFrom [2]float64, lngLatTo [2]float64) float64 {
	return getWGS84Distance(convertGCJ02ToWGS84(lngLatFrom), convertGCJ02ToWGS84(lngLatTo))
}
func getBD09Distance(lngLatFrom [2]float64, lngLatTo [2]float64) float64 {
	return getGCJ02Distance(convertBD09ToGCJ02(lngLatFrom), convertBD09ToGCJ02(lngLatTo))
}
func getAngle(lngLatFrom [2]float64, lngLatTo [2]float64) float64 {
	lonA := lngLatFrom[0] * math.Pi / 180
	latA := lngLatFrom[1] * math.Pi / 180
	lonB := lngLatTo[0] * math.Pi / 180
	latB := lngLatTo[1] * math.Pi / 180
	dLon := lonB - lonA
	x := math.Cos(latA)*math.Sin(latB) - math.Sin(latA)*math.Cos(latB)*math.Cos(dLon)
	y := math.Sin(dLon) * math.Cos(latB)
	return math.Atan2(y, x) * 180 / math.Pi
}
func convertLngLatToXY(lngLat [2]float64) [2]float64 {
	lng := lngLat[0]
	lat := lngLat[1]
	lat = math.Max(math.Min(EPSG3857MACLAT, lat), -EPSG3857MACLAT)
	return [2]float64{
		YR * lng * VR,
		YR * math.Log(math.Tan(math.Pi/4+lat*VR/2)),
	}
}
func convertXYToLngLat(xy [2]float64) [2]float64 {
	x := xy[0]
	y := xy[1]
	return [2]float64{
		x / YR * MR,
		(2*math.Atan(math.Exp(y/YR)) - math.Pi/2) * MR,
	}
}
func distance(p1 [2]float64, p2 [2]float64) float64 {
	rad := VR
	lat1 := p1[1] * rad
	lng1 := p1[0] * rad
	lat2 := p2[1] * rad
	lng2 := p2[0] * rad
	dLng := lng2 - lng1
	a := (1 - math.Cos(lat2-lat1) + (1-math.Cos(dLng))*math.Cos(lat1)*math.Cos(lat2)) / 2
	return 2 * YR * math.Asin(math.Sqrt(a))
}
func distancePToP1P2(p [2]float64, p1 [2]float64, p2 [2]float64) float64 {
	return distance(p, pToP1P2(p, p1, p2))
}
func pToP1P2(p [2]float64, p1 [2]float64, p2 [2]float64) [2]float64 {
	xy := convertLngLatToXY(p)
	xy1 := convertLngLatToXY(p1)
	xy2 := convertLngLatToXY(p2)
	x := xy[0]
	y := xy[1]
	x1 := xy1[0]
	y1 := xy1[1]
	x2 := xy2[0]
	y2 := xy2[1]
	dx := x2 - x1
	dy := y2 - y1
	t := (dx*(x-x1) + dy*(y-y1)) / (dx*dx + dy*dy)
	t = math.Max(0, math.Min(1, t))
	return convertXYToLngLat([2]float64{x1 + t*dx, y1 + t*dy})
}
func distancePToPath(p [2]float64, path [][2]float64) float64 {
	s := math.Inf(1)
	for i := 0; i < len(path)-1; i++ {
		s = math.Min(s, distancePToP1P2(p, path[i], path[i+1]))
	}
	return s
}
