package solution

import "sort"

func Triangle(A []int) int {
    //for Desc type description see types.go
    sort.Sort(Desc(A))
    for i := 0; i < len(A) - 2; i++ {
        if A[i] < A[i+1] + A[i+2] {
            return 1
        } 
    }
    return 0
}