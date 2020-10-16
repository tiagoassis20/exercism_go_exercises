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

// calc the frequency and put it in an channel
func chFrequency(ch chan FreqMap, s string) {
	ch <- Frequency(s)
}

// Frequency counts the frequency of each rune in a given texts slice concurrently and returns this
// data as a FreqMap.
func ConcurrentFrequency(texts []string) FreqMap {
	m := FreqMap{}
	ch := make(chan FreqMap)
	for i := 0; i < len(texts); i++ {
		go chFrequency(ch, texts[i])
	}
	for i := 0; i < len(texts); i++ {
		a := <-ch
		for k, v := range a {
			m[k] += v
		}
	}
	return m
}
