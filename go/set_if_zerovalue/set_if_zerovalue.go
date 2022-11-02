package set_if_zerovalue

import "reflect"

func SetStringIfZerovalue(old string, new string) string {
	if reflect.ValueOf(old).IsZero() {
		return new
	}
	return old
}

func SetIntIfZerovalue(old int, new int) int {
	if reflect.ValueOf(old).IsZero() {
		return new
	}
	return old
}
