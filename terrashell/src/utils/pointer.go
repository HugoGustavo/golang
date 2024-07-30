package utils

func PointerIsNil(input *interface{}) bool {
	return input == nil
}

func PointerGetDefaultIfNil[T any](input *T, def *T) *T {
	if input == nil {
		return def
	}
	return input
}
