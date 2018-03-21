package solution

import (
    "sort"
    "math"
)

func MaxProductOfThree(A []int) int {
	//for Asc type description see types.go
	sort.Sort(Asc(A))
    L := len(A)
    return int(math.Max(float64(A[0] * A[1] * A[L-1]), float64(A[L-1] * A[L-2] * A[L-3])))
}