package letter

import "sync"

type concurrentFreqMap struct {
	fm  FreqMap
	mux sync.RWMutex
}

func (cfm *concurrentFreqMap) Inc(r rune) {
	cfm.mux.Lock()
	defer cfm.mux.Unlock()
	cfm.fm[r]++
}

// ConcurrentFrequency2 calculates count of occurrence of a letter in given string, but concurrently with a mutex protected map.
func ConcurrentFrequency2(ss []string) FreqMap {
	cfm := concurrentFreqMap{
		fm: FreqMap{},
	}
	var wg sync.WaitGroup
	wg.Add(len(ss))

	for _, s := range ss {
		go func(value string) {
			defer wg.Done()
			for _, r := range value {
				cfm.Inc(r)
			}
		}(s)
	}
	wg.Wait()
	return cfm.fm
}

