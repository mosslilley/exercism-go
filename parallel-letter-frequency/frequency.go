package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}

	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency runs the frequency function in parallel
// for different stories and adds the result through a channel
func ConcurrentFrequency(strings []string) FreqMap {
	ch := make(chan FreqMap)

	result := FreqMap{}

	for _, story := range strings {
		go func(story string) {
			ch <- Frequency(story)
		}(story)
	}

	for range strings {
		for rune, count := range <-ch {
			result[rune] += count
		}
	}
	return result
}
