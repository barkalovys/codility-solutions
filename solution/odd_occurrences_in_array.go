package solution

func OddOccurencesInArray(A []int) int {
    m := make(map[int]int)
    for _, v := range A {
        m[v] += 1    
    }
    for v, count := range m {
        if count % 2 != 0 {
            return v
        }
    }
    return 0
}