package slowimage

import (
	"fmt"
	"time"
)

// Image represents any image
type Image struct {
	Filename string
}

// Download will download the image to disk. Slowly.
// Upon completion, the name of the file is returned.
func (i Image) Download() string {
	fmt.Println("downloading image . . .")
	time.Sleep(60 * time.Minute)
	return i.Filename
}
