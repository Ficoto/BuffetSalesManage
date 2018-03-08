package router

import (
	"flag"
	"fmt"
	"log"
)

// FlagHostPort - format host and port
func FlagHostPort(defaultPot int) string {
	var (
		host = flag.String("host", "0.0.0.0", "http web host.")
		port = flag.Int("port", defaultPot, "http web port.")
	)
	log.Printf("http://127.0.0.1:%d", defaultPot)
	return fmt.Sprintf("%s:%d", *host, *port)
}
