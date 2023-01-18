package uid

import (
	uuid "github.com/google/uuid"
)

func GenerateUUID() uint32 {
	return uuid.New().ID()
}
