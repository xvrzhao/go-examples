package channel

// NoDeadlock illustrates that only when all goroutines are blocked can a deadlock occur.
func NoDeadlock() {
	go func() {
		// always do something
		for {
			func() {}()
		}
	}()

	ch := make(chan int)
	ch <- 1 // block but not deadlock
}
