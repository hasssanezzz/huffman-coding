package huffman

import (
	"fmt"
	"io"
	"sync"
)

func countBytes(data []byte, ch chan map[byte]int, wg *sync.WaitGroup) {
	defer wg.Done()
	freq := make(map[byte]int)
	for _, b := range data {
		freq[b]++
	}
	ch <- freq
}

func mergeMaps(ch chan map[byte]int, result map[byte]int, done chan struct{}) {
	for freq := range ch {
		for b, count := range freq {
			result[b] += count
		}
	}
	done <- struct{}{}
}

func ConcurrentFrequencyRead(r io.Reader) (map[byte]int, error) {
	var wg sync.WaitGroup
	ch := make(chan map[byte]int)
	done := make(chan struct{})
	chunkSize := 1024 * 1024
	result := map[byte]int{}

	go mergeMaps(ch, result, done)

	for {
		chunk := make([]byte, chunkSize)
		n, err := r.Read(chunk)
		if err != nil && err != io.EOF {
			return nil, fmt.Errorf("can not read chunk of sise %d: %v", chunkSize, err)
		}

		if n == 0 {
			break
		}

		wg.Add(1)
		go countBytes(chunk[:n], ch, &wg)
	}

	wg.Wait()
	close(ch)

	<-done // Wait for the mergeMaps goroutine to finish

	return result, nil
}
