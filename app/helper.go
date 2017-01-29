package app

// must panics if error not nil
func must(err error) {
	if err != nil {
		panic(err)
	}
}
