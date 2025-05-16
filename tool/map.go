package tool

// MapMerge map合并
func MapMerge(array1 map[string]interface{}, array2 ...map[string]interface{}) map[string]interface{} {
	if len(array2) == 0 {
		return array1
	}
	for _, v := range array2 {
		for k, v2 := range v {
			array1[k] = v2
		}
	}
	return array1
}
