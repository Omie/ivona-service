package main // import "github.com/omie/ivona-service"

import (
	"log"
	"os"
)

func main() {

	accessKey := os.Getenv("IVONA_ACCESSKEY")
	secretKey := os.Getenv("IVONA_SECRETKEY")
	if accessKey == "" || secretKey == "" {
		log.Println("main: ivona credentials not set")
		return
	}

	host := os.Getenv("IVONA_SERVICE_HOST")
	port := os.Getenv("IVONA_SERVICE_PORT")
	if host == "" || port == "" {
		log.Println("main: host or port not set")
		return
	}
	initIvona(accessKey, secretKey)

	err := loadVoices()
	if err != nil {
		log.Println("Error fetching voices", err)
	}

	err = StartHTTPServer(host, port)
	if err != nil {
		log.Println("Error starting server", err)
	}

}
