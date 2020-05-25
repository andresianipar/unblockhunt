package builtin

// Remove function
func Remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])

	return slice[:len(slice)-1]
}
