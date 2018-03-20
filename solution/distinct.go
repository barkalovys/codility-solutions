package solution

import "sort"

type Asc []int
func (A Asc) Len() int { return len(A) }
func (A Asc) Swap(i, j int)      { A[i], A[j] = A[j], A[i] }
func (A Asc) Less(i, j int) bool { return A[i] < A[j] }

func Distinct(A []int) int {
    if len(A) == 0 {
        return 0
    }
    sort.Sort(Asc(A))
    count := 1
    for i := 1; i < len(A); i++ {
        if A[i] != A[i-1] {
            count += 1
        }
    }
    return count
}