package helper

// Loop runs through the numbers provided and prints them out
func Loop(n, m int) []int {
	arr := make([]int, m-n)
	v := n
	for i := 0; i < m-v; i++ {
		arr[i] = n
		n++
	}
	return arr
}
