package utils

func HandleError(e error) {
	if e != nil {
		panic(e)
	}
	return
}
