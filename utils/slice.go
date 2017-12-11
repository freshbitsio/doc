package utils

func ContainsStr(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}