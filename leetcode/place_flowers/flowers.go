package flowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	canPlant := 0

	emptyThree := 1
	for i, plot := range flowerbed {
		if plot == 1 {
			emptyThree = 0
			continue
		}

		emptyThree++
		if emptyThree >= 3 {
			canPlant++
			emptyThree = 1
		}

		// last one
		if i == len(flowerbed)-1 && emptyThree >= 2 {
			canPlant++
		}
	}

	return canPlant >= n
}
