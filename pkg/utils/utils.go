package utils

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

func DeepCopyMap(o map[string]string) map[string]string {
	copyMap := make(map[string]string)
	for key, value := range o {
		copyMap[key] = value
	}
	return copyMap
}
func CreateUniqueFileName(ext string) string {
	timestamp := time.Now().UnixNano()
	uuid := uuid.New().String()
	return fmt.Sprintf("%d_%s.%s", timestamp, uuid, ext)
}
