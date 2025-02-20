package _uuid

import (
	"fmt"

	"github.com/google/uuid"
)

func New() {
	id := uuid.New().String()
	fmt.Println("UUID: ", id)
}
