package letter

import "sync"

// ConcurrentFrequency3 calculates count of occurrence of a letter in given string, but concurrently with a sync.Map.
func ConcurrentFrequency3(ss []string) FreqMap {
	var wg sync.WaitGroup
	var sm sync.Map
	wg.Add(len(ss))

	for _, s := range ss {
		go func(value string) {
			defer wg.Done()
			for _, r := range value {
				value, ok := sm.LoadOrStore(r, 1)
				if ok {
					sm.Store(r, value.(int)+1)
				}
			}
		}(s)
	}
	wg.Wait()
	result := FreqMap{}
	sm.Range(func(key, value interface{}) bool {
		result[key.(rune)] = value.(int)
		return true
	})
	return result
}
