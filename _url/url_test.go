package _url

import (
	"fmt"
	"testing"
)

func TestGetOrigin(t *testing.T) {

	{
		give := "https://www.example.com:8080/path/to/page?search=query#section"
		f := GetOrigin(give)
		fmt.Println(f)
	}

	{
		give := "https://www.example.com/path/to/page?search=query#section"
		f := GetOrigin(give)
		fmt.Println(f)
	}

}
