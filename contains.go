package sutils

func StringInSlice(str string, slice []string) bool {
	for _, el := range slice {
		if str == el {
			return true
		}
	}
	return false
}

func IntInSlice(i int, slice []int) bool {
	for _, el := range slice {
		if i == el {
			return true
		}
	}
	return false
}
