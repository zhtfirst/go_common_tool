package tool

// Reverse 数组逆序输出
func SeliceReverse(array []string) []string {
	for i := 0; i < len(array)/2; i++ {
		array[i], array[len(array)-1-i] = array[len(array)-1-i], array[i]
	}
	return array
}

// Unique 数组去重
func SeliceUnique(array []string) []string {
	m := make(map[string]bool)
	for _, v := range array {
		m[v] = true
	}
	var result []string
	for k := range m {
		result = append(result, k)
	}
	return result
}
