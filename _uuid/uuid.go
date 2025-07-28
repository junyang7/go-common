package _uuid

import (
	"github.com/google/uuid"
	"strings"
)

func V4() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
