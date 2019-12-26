package interfaces

func RunAssignExample() {
	var (
		speaker Speaker
		walker  Walker
	)

	man := Man{}           // variable man implements the `Walk` method
	manPointer := new(Man) // variable manPointer implements both the `Speak` and the `Walk` method

	walker = man        // valid
	walker = manPointer // valid

	speaker = manPointer // valid
	// speaker = man	 // invalid, compiler error, variable man does not implement the `Speak` method

	use(speaker, walker)
}
