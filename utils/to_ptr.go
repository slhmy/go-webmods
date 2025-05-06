package gwm_utils

func ToPtr[T any](v T) *T {
	result := new(T)
	*result = v
	return result
}
