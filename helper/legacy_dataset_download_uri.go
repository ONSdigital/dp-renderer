package helper

import "fmt"

// LegacyDataSetDownloadURI builds a URI string for a legacy dataset download URI.
func LegacyDatasetDownloadURI(pageURI, filename string) string {
	legacyDatasetURIFormat := "/file?uri=%s/%s"
	// Concatenation of strings inside a Href tag causes the URI value to be HTML escaped.
	// The preference is for our links not to be escaped to maintain readability. To remedy this we build
	// the link inside this func which is then inserted into template.
	return fmt.Sprintf(legacyDatasetURIFormat, pageURI, filename)
}
