package solution

import "sort"

func Distinct(A []int) int {
    if len(A) == 0 {
        return 0
	}
	//for Asc type description see types.go
    sort.Sort(Asc(A))
    count := 1
    for i := 1; i < len(A); i++ {
        if A[i] != A[i-1] {
            count += 1
        }
    }
    return count
}