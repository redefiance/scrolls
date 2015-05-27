package scrolls

func deny(err error) {
	if err != nil {
		panic(err)
	}
}
