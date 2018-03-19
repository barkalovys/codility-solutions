package solution

import "math"

func MaxCounters(N int, A []int) []int {
    res := make([]int, N)
    var maxCounter, lastMaxCounter float64 = 0, 0
    for _, i := range A {
        if i == N + 1 {
            lastMaxCounter = maxCounter
        } else {
            res[i-1] = int(math.Max(float64(res[i-1]), lastMaxCounter)) + 1
            maxCounter = math.Max(float64(res[i-1]), maxCounter)
        }
    }
    for i := range res {
        res[i] = int(math.Max(float64(res[i]), lastMaxCounter))
    }
    return res
}