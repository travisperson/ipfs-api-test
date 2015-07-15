package config

import (
	"fmt"
	"os"
)

func U (uri string) string {
	return fmt.Sprintf("http://%s:%s%s", os.Getenv("IPFS_HOST"), os.Getenv("IPFS_PORT"), uri)
}
