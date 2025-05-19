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

// GetValueWithDefault GetValueWithDefault[T comparable]
//
//	@Description: 获取map中的值，如果不存在则返回默认值
//	@param m
//	@param value
//	@param defaultValue
//	@return T
func GetValueWithDefault[T comparable](m map[string]T, value string, defaultValue T) T {
	if v, ok := m[value]; ok {
		return v
	}

	return defaultValue
}
