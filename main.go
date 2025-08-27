package main

import (
	"os"

	"github.com/Itros97/MokApp/internal/api"
)

func main() {
	api.Start()
	os.Exit(0)
}
