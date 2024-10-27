package calculator

const limitItems = 10

func Average(sensorCh <-chan int, averageCh chan<- float64) {
	defer close(averageCh)
	var counter int
	var acm float64

	for v := range sensorCh {
		counter++
		acm += float64(v)
		if counter == limitItems {
			averageCh <- acm / limitItems
			acm = 0
			counter = 0
		}
	}
}
