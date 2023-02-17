package sutils

func StringSlicesIntersect(slice1, slice2 []string) bool {
	for i := range slice1 {
		for j := range slice2 {
			if slice1[i] == slice2[j] {
				return true
			}
		}
	}
	return false
}

func IntSlicesIntersect(slice1, slice2 []int) bool {
	for i := range slice1 {
		for j := range slice2 {
			if slice1[i] == slice2[j] {
				return true
			}
		}
	}
	return false
}
