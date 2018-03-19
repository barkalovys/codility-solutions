package solution

func GenomicRangeQuery(S string, P []int, Q []int) []int {
	res := make([]int, len(P))
	//Array of prefix sums for each nucleotides:
    p := make([][4]uint16, len(S) + 1)
    p[0] = [4]uint16{0,0,0,0} 
    for i := 1; i < len(p); i++ {
        p[i] = p[i-1]
        switch S[i-1] {
	    //A
            case 65:
                p[i][0] += 1
	    //C
            case 67:
                p[i][1] += 1
	    //G
	    case  71:
                p[i][2] += 1
	    //T
	    default:
                p[i][3] += 1
        }
    }
    for i := range P {
        switch {
            case p[Q[i]+1][0] - p[P[i]][0] != 0:
                res[i] = 1
            case p[Q[i]+1][1] - p[P[i]][1] != 0:
                res[i] = 2
            case p[Q[i]+1][2] - p[P[i]][2] != 0:
                res[i] = 3
            default:
                res[i] = 4
        }
    }
    return res
}
