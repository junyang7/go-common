package _toml

import (
	"fmt"
	"testing"
)

func TestReader_Get(t *testing.T) {

	text := `# Database configuration
[database]
host = "localhost"
port = 5432
username = "user"
password = "password"

# Server settings
[server]
host = "0.0.0.0"
port = 8080

# Logging options
[logging]
level = "debug"
file = "app.log"

# List of enabled features
enabled_features = ["feature1", "feature2", "feature3"]

# Nested table for a specific server
[servers]
  [servers.alpha]
  ip = "192.168.1.1"
  port = 8081

  [servers.beta]
  ip = "192.168.1.2"
  port = 8082


[[products]]
name = "product1"
price = 19.99

[[products]]
name = "product2"
price = 39.99
`

	reader := New().Text(text)
	fmt.Println(reader.Get("database"))
	fmt.Println(reader.Get("database.host"))
	fmt.Println(reader.Get("logging.enabled_features"))
	fmt.Println(reader.Get("logging.enabled_features.1"))
	fmt.Println(reader.Get("servers"))
	fmt.Println(reader.Get("servers.beta"))
	fmt.Println(reader.Get("servers.beta.ip"))
	fmt.Println(reader.Get("products"))
	fmt.Println(reader.Get("products.0.price"))

}
