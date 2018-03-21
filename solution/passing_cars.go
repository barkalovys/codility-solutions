package solution

func PassingCars(A []int) int {
    suffSum := make([]int, len(A) + 1)
    suffSum[len(A)] = 0
    for i := len(A)-1; i >= 0; i-- {
        suffSum[i] = suffSum[i+1] + A[i]
    }
    count := 0
    for i := range A {
        if A[i] != 0 {
            continue
        }
        count += suffSum[i+1]
        if count > 1000000000 {
            return -1
        }
    }
    return count
}