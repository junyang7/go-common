package _json

import (
	"fmt"
	"testing"
)

func TestReader_Get(t *testing.T) {

	text := `{
		"name": "Example Project",
		"version": 1.0,
		"features": {
			"auth": true,
			"logging": false
		},
		"databases": [
			{ "type": "mysql", "host": "localhost", "port": 3306 },
			{ "type": "postgres", "host": "remotehost", "port": 5432 }
		],
		"enabled": true
	}`

	reader := New().Text(text)
	fmt.Println(reader.Get("databases.1.port").Int64().Value())

}
