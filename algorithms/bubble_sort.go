package algorithms

func BubbleSort(input []int) []int {
	var temp int
	swapped := false

	for i := 0; i < len(input)-1; i++ {
		swapped = false
		for j := 0; j < len(input)-i-1; j++ {
			if input[j] > input[j+1] {
				temp = input[j]
				input[j] = input[j+1]
				input[j+1] = temp
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	return input
}
