package containerwithmostwater

func maxArea(height []int) int {
	var i, maxArea int
	j := len(height) - 1

	for j > i {
		size := j - i
		area := calculateArea(height[i], height[j], size)

		if area > maxArea {
			maxArea = area
		}

		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}

	return maxArea
}

func calculateArea(height1, height2, width int) int {
	h := height1
	if h > height2 {
		h = height2
	}

	return h * width
}
