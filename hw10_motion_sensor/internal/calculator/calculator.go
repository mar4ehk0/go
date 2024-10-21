package calculator

const limitItems = 10

func ProcessData(dataCh <-chan int, resultCh chan<- float64) {
	var counter int

	data := make([]int, 0, limitItems)

	for v := range dataCh {
		data = append(data, v)
		counter++
		if counter == limitItems {
			result := average(data)

			resultCh <- result
			data = nil
			counter = 0
		}
	}
}

func average(data []int) float64 {
	var result float64
	var acm int

	for i := 0; i < len(data); i++ {
		acm += data[i]
	}

	result = float64(acm) / float64(len(data))

	return result
}
