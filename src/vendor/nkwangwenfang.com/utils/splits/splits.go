package splits

// SplitUint64s 将[]uint64分割为长度最多为length的二级[][]uint64,length不能为0，否则会panic
func SplitUint64s(items []uint64, length uint) [][]uint64 {
	var (
		results [][]uint64
		result  []uint64
	)
	for i, item := range items {
		if i != 0 && uint(i)%length == 0 {
			results = append(results, result)
			result = nil
		}
		result = append(result, item)
	}
	if len(result) != 0 {
		results = append(results, result)
	}
	return results
}
