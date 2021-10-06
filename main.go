// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"

	"cloud.google.com/go/compute/metadata"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	regionUrl, err := metadata.Get("instance/region")
	if err != nil {
		log.Print(err.Error())
	} else {
		fullUrl, _ := url.Parse(regionUrl)
		computeRegion := path.Base(fullUrl.Path)
		log.Printf("REGION %s", computeRegion)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// Package tips contains tips for writing Cloud Functions in Go.
// package tips

// import (
// 	"log"
// 	"net/http"

// 	"cloud.google.com/go/compute/metadata"
// )

// // EnvVar is an example of getting an environment variable in a Go function.
// func F(w http.ResponseWriter, r *http.Request) {
// 	region, _ := metadata.Get("region")
// 	log.Printf("REGION %s", region)
// }
