package solution

import "strconv"

func BinaryGap(N int) int {
	bin := strconv.FormatUint(uint64(N), 2)
	curGap, maxGap := 0, 0
	for _, c := range bin {
		if string(c) == "1" {
			if curGap > maxGap {
				maxGap = curGap
			}
			curGap = 0
			continue
		}
		curGap += 1
	}
	return maxGap
}
