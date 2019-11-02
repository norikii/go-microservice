package utils

// not really efficient Bubble sort function
func BubbleSort(elements []int) []int {
	keepRunning := true
	for keepRunning {
		keepRunning = false

		for i := 0; i< len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				// we are swapping elements
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRunning = true
			}
		}
	}

	return elements
}
