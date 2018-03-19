package solution

import "math"

func TapeEquilibrium(A []int) int {
    var totalSum int
    //find sum of all elements
    for _, v := range A {
        totalSum += v 
    }
    leftSum := A[0]
    rightSum := totalSum - leftSum
    minDiff := math.Abs(float64(leftSum - rightSum))
    for i := 2; i < len(A); i++ {
        leftSum += A[i-1]
        rightSum = totalSum - leftSum
        curDiff := math.Abs(float64(leftSum - rightSum))
        if curDiff < minDiff {
            minDiff = curDiff
        } 
    }
    return int(minDiff)
}

