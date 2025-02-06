package _toml

import (
	"fmt"
	"testing"
)

func TestReader_Get(t *testing.T) {

	text := `[general]
name = "MyApp"
version = 1.0
description = "This is a sample app"

[database]
host = "localhost"
port = 3306
username = "user"
password = "password123"
test = [1,3,4,5,6]
`

	reader := New().Text(text)
	fmt.Println(reader.Get("database"))
	fmt.Println(reader.Get("database.host"))
	fmt.Println(reader.Get("database.test"))
	fmt.Println(reader.Get("database.test.2"))
	fmt.Println(reader.Get("database.host").String().Value())

}
