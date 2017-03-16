package comp

// StringsEqual 比较string slice是否相等
func StringsEqual(ss1, ss2 []string) bool {
	if len(ss1) != len(ss2) {
		return false
	}

	for i := range ss1 {
		if ss1[i] != ss2[i] {
			return false
		}
	}
	return true
}

// IntsEqual 比较int slice是否相等
func IntsEqual(is1, is2 []int) bool {
	if len(is1) != len(is2) {
		return false
	}

	for i := range is1 {
		if is1[i] != is2[i] {
			return false
		}
	}
	return true
}
