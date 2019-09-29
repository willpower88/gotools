package tools

//error检查
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
