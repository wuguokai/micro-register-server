package util

func CheckErr(errMasg error) {
	if errMasg != nil {
		panic(errMasg)
	}
}
