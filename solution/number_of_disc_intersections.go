package solution

import (
     "sort"
)

func NumberOfDiscIntersections(A []int) int {
    N := len(A)
    // R1 + R2 >= i1 - i2 - condition of disks intersection
    // or
    // A[i] + A[j] >= j - i
    // A[i] + i >= -A[j] + j:
    C := make([]int, N)
    D := make([]int, N)
    for i := range A {
        C[i] = A[i] + i
        D[i] = - A[i] + i
    }
    //for Asc type description see types.go
    sort.Sort(Asc(C))
    sort.Sort(Asc(D))
    //Find number of values of D less than C[i]
    counter := 0
    for i := range C {
        j := sort.Search(N, func(j int) bool { return D[j] > C[i] })
        //Substract duplicate intersections from result:
        counter += j - (N - i)
        if counter > 10000000 {
            return -1
        }
    }
    return counter
}