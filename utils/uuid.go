package utils

import (
	"time"

	"github.com/google/uuid"
)

func GenUuid() string {
	id := uuid.New()
	return id.String()
}
func GetCurrentTimestamp() int64 {
	currentTime := time.Now()
	timestamp := currentTime.Unix()
	return timestamp
}
