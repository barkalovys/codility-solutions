package solution

func PermMissingElem(A []int) int {
    M := make(map[int]bool)
    //put all values from slice to map
    for _, v := range A {
        M[v] = true
    }
    N := len(A)
    for i := 1; i < N + 2; i++ {
        if _,ok := M[i]; !ok {
            return i
        }
    }
    return -1
}

