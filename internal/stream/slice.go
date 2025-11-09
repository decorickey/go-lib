package stream

func FromSlice[T any](values []T) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for _, v := range values {
			ch <- v
		}
	}()
	return ch
}

func Filter[T any](values <-chan T, fn func(v T) bool) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		if values == nil {
			return
		}
		for v := range values {
			if fn(v) {
				ch <- v
			}
		}
	}()
	return ch
}
