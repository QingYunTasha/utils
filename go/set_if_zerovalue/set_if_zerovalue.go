package set_if_zerovalue

func SetStringIfZerovalue(old string, new string) string {
	if old == "" {
		return new
	}
	return old
}

func SetIntIfZerovalue(old int, new int) int {
	if old == 0 {
		return new
	}
	return old
}
