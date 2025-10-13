package _coordinate

import (
	"fmt"
	"github.com/junyang7/go-common/_file"
	"github.com/junyang7/go-common/_json"
	"testing"
)

func TestWriter_ConvertLngLatToXY(t *testing.T) {
	{
		p := [2]float64{116.671264, 39.892798}
		fmt.Println(p)
		xy := ConvertLngLatToXY(p)
		fmt.Println(xy)
	}
}

func TestWriter_ConvertXYToLngLat(t *testing.T) {
	{
		xy := [2]float64{12987785.698687589, 4850376.183964485}
		fmt.Println(xy)
		p := ConvertXYToLngLat(xy)
		fmt.Println(p)
	}
}

func TestWriter_PToP1P2(t *testing.T) {
	{
		p := [2]float64{116.671264, 39.892798}
		fmt.Println(p)
		p1 := [2]float64{116.671586, 39.891744}
		fmt.Println(p1)
		p2 := [2]float64{116.560079, 39.787293}
		fmt.Println(p2)
		px := PToP1P2(p, p1, p2)
		fmt.Println(px)
	}
}
func TestWriter_Distance(t *testing.T) {
	{
		p1 := [2]float64{116.671586, 39.891744}
		fmt.Println(p1)
		p2 := [2]float64{116.560079, 39.787293}
		fmt.Println(p2)
		s := Distance(p1, p2)
		fmt.Println(s)
	}
}
func TestWriter_DistancePToP1P2(t *testing.T) {
	{
		p := [2]float64{116.671264, 39.892798}
		fmt.Println(p)
		p1 := [2]float64{116.671586, 39.891744}
		fmt.Println(p1)
		p2 := [2]float64{116.560079, 39.787293}
		fmt.Println(p2)
		s := DistancePToP1P2(p, p1, p2)
		fmt.Println(s)
	}
}
func TestWriter_DistancePToPath(t *testing.T) {
	{
		p := [2]float64{116.671264, 39.892798}
		fmt.Println(p)
		path := [][2]float64{
			[2]float64{116.671586, 39.891744},
			[2]float64{116.560079, 39.787293},
			[2]float64{116.539019, 39.806957},
			[2]float64{116.532574, 39.80433},
		}
		fmt.Println(path)
		s := DistancePToPath(p, path)
		fmt.Println(s)
	}
}
func TestWriter_DistancePToPath2(t *testing.T) {
	{
		p := [2]float64{116.671264, 39.892798}
		fmt.Println(p)
		var path [][2]float64
		_json.Decode(_file.ReadAll("camera_list.json"), &path)
		fmt.Println(path)
		s := DistancePToPath(p, path)
		fmt.Println(s)
	}
}
