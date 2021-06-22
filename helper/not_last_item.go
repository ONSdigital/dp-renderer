package helper

// NotLastItem returns true/false based on if the index equals the length
// Example of use is in JSON-LD partials, where we must determine whether or not a comma should be rendered in a range
func NotLastItem(length, index int) bool {
	if index < length-1 {
		return true
	}
	return false
}
