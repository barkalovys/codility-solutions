package solution

func PermCheck(A []int) int {
    M := make(map[int]bool)
    for _, v := range A {
        M[v] = true
    }
    N := len(A)
    //if some element from 1 to N is missing - A is not a permutation
    for i := 1; i < N + 1; i++ {
        if _, ok := M[i]; !ok {
            return 0
        }
    }
    return 1
}

