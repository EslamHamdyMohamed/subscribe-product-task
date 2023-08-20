package helpers

func CheckError(err error) bool {
	if err != nil {
		return true
	}
	return false
}
