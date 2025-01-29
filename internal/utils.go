package internal

import (
	"fmt"
	"os"
	"strconv"
)

func parseID(idStr string) int {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID:", idStr)
		os.Exit(1)
	}
	return id
}
