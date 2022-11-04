package pointer

// GetIntP returns a pointer to the num argument.
func GetIntP(num int) *int {
	return &num
}
