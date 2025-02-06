package _datetimeSE

import (
	"fmt"
	"testing"
)

func TestFormateByY(t *testing.T) {

	datetimeS := "2020-09-01 00:59:00"
	datetimeE := "2024-01-01 12:00:00"

	datetimeSEList := FormatByY(datetimeS, datetimeE)
	fmt.Println(datetimeSEList)

}
func TestFormatByYm(t *testing.T) {

	datetimeS := "2020-09-01 00:59:00"
	datetimeE := "2024-01-01 12:00:00"

	datetimeSEList := FormatByYm(datetimeS, datetimeE)
	fmt.Println(datetimeSEList)

}
