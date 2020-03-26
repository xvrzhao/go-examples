package channel

import "fmt"

// FactoryPattern demonstrates the factory pattern of channel.
// The factory function returns a channel that produces data.
func FactoryPattern() {
	ch := factory()
	for n := range ch {
		fmt.Println(n)
	}
}

func factory() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}
