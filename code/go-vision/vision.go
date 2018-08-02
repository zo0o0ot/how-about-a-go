// Sample vision-quickstart uses the Google Cloud Vision API to label an image.
package main

import (
	"fmt"
	"log"
	"os"

	// Imports the Google Cloud Vision API client package.
	vision "cloud.google.com/go/vision/apiv1"
	"golang.org/x/net/context"
)

func sample() {
	// This is a sample of code.
}

func main() {
	// Note: if running this command from the console, you need to be in the
	// how-about-a-go/code/go-vision/ directory, since this uses relative pathing.
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the name of the image file to annotate.
	filename := "testdata/clark.png"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	defer file.Close()
	image, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Fatalf("Failed to create image: %v", err)
	}

	labels, err := client.DetectLabels(ctx, image, nil, 10)
	if err != nil {
		log.Fatalf("Failed to detect labels: %v", err)
	}

	fmt.Println("Labels:")
	for _, label := range labels {
		fmt.Println(label.Description)
	}
}
