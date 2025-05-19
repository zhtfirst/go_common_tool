package tool

// Ternary 三元运算
func Ternary[T comparable](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}
