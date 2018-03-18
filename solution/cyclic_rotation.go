package solution

func CyclicRotation(A []int, K int) []int {
    var L = len(A)
    B := make([]int, L)
    if L == 0 {
        return B
    }
    if K > L {
        K = K % L        
    }
    for i, v := range A {
        if i + K < L {
            B[i + K] = v
        } else {
            B[K - L + i] = v
        }
    }
    return B
}