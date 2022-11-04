package pointer

// GetStringP returns a pointer to the str argument.
func GetStringP(str string) *string {
	return &str
}
