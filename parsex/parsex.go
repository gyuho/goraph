package parsex

// UniqElemStr deletes the duplicate elements
// and returns the new string slice of unique elements.
func UniqElemStr(slice []string) []string {
	// var model map[string]bool
	// model := map[string]bool{}
	model := make(map[string]bool)

	// var result []string
	// result := make([]string, 5)
	result := []string{}

	// traverse the input array and map each to boolean value
	for _, v := range slice {
		if _, checked := model[v]; !checked {
			result = append(result, v)
			model[v] = true
		}
	}
	return result
}

// EqualSliceElem returns true if the two string slices
// contains equal elements regardless of its order.
func EqualSliceElem(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for _, v1 := range s1 {
		if !CheckStr(v1, s2) {
			return false
		}
	}
	return true
}

// CheckStr returns true if the string exist
// in the string slice.
func CheckStr(str string, slice []string) bool {
	exist := false
	for _, v := range slice {
		if v == str {
			exist = true
		}
	}
	return exist
}
