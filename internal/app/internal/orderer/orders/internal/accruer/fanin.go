package accruals

import "sync"

func fanIn(chans ...<-chan string) <-chan string {
	outCh := make(chan string, len(chans))
	var wg sync.WaitGroup
	wg.Add(len(chans))

	output := func(inCh <-chan string) {
		defer wg.Done()
		for order := range inCh {
			outCh <- order
		}
	}

	for _, ch := range chans {
		go output(ch)
	}

	go func() {
		defer close(outCh)
		wg.Wait()
	}()

	return outCh
}
