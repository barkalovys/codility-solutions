package solution

func MissingInteger(A []int) int {
    M := make(map[int]bool)
    for _, v := range A {
        M[v] = true
    }
    for i := 1; ;i++ {
        if _, ok := M[i]; !ok {
            return i
        }
    }
}

