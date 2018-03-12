package utils

import "os"

var (
	BuffetPath = os.Getenv("buffet_path")
)

const (
	DBServerFile      = "db.toml"
)