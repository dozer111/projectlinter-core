package utilTest

func ToPointer[T any](val T) *T {
	return &val
}
