package solution

func FrogRiverOne(X int, A []int) int {
    // create a map where
    // key = position and value = earliest time
    M := make(map[int]int)
    for sec, pos := range A {
        if _, ok := M[pos]; !ok {
            M[pos] = sec
        }
    }
    
    var maxSec int
    for i := 1; i < X + 1; i++ {
        // check if leaves on all positions exists, if not => return -1
        if _, ok := M[i]; !ok {
            return -1
        }
        if M[i] > maxSec {
            maxSec = M[i]
        }
    }
    return maxSec
}