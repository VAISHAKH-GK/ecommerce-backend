package helpers

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
