// Declare it in the `main` package, as we're creating an executable
package main

// Import dependencies
import (
	// Formatting and outputting text to the console
	"fmt"

	// A libraty for reading and interpreting the EXIF picture data
	"github.com/rwcarlsen/goexif/exif"

	// A librarty for cross-platform opening of file in default apps
	"github.com/skratchdot/open-golang/open"

	// Operating-system provided features (command arguments,
	// opening files, exiting the program
	"os"
)

// The URL template we need to show a set of coordinates on Google Maps
const googleMapsURLTemplate = "https://maps.google.com/?q=%f,%f"

// A helper function that checks for an error condition and exits
// the process with a helpful message in case of error
func exitOnError(err error, message string) {
	if err != nil {
		fmt.Printf("%v (%v)\n", message, err)
		os.Exit(1)
	}
}

// Our programs starts executing here
func main() {
	// We need exactly 2 parameters: the first one is the program name,
	// the second one should be the photo we want to operate on
	if len(os.Args) != 2 {
		fmt.Println("Please give a single file name as an argument!")
		os.Exit(1)
	}

	// Retrieve the photo file name from the arguments array
	fileName := os.Args[1]

	// Try to opern the given file, error out on failure
	file, err := os.Open(fileName)
	exitOnError(err, "Couldn't open file")

	// Try to extract the EXIF data, error out on failure
	exifData, err := exif.Decode(file)
	exitOnError(err, "Couldn't find EXIF data")

	// Try to find a GPS coordinates entry in the EXIF data structure.
	// Error out on failure
	lat, long, err := exifData.LatLong()
	exitOnError(err, "Couldn't read GPS coordinates")

	// Create the final URL by using the Google Maps URL template
	url := fmt.Sprintf(googleMapsURLTemplate, lat, long)

	// Try to start the default browser for the current OS.
	// Show the computer URL on error, so that the user can still
	// access it manually
	err = open.Start(url)
	exitOnError(err, fmt.Sprintf(
		"Couldn't start the default browser, please visit %v manually", url))
}
