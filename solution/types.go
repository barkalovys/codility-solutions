package solution

type Asc []int
func (A Asc) Len() int { return len(A) }
func (A Asc) Swap(i, j int)      { A[i], A[j] = A[j], A[i] }
func (A Asc) Less(i, j int) bool { return A[i] < A[j] }

type Desc []int
func (A Desc) Len() int { return len(A) }
func (A Desc) Swap(i, j int)      { A[i], A[j] = A[j], A[i] }
func (A Desc) Less(i, j int) bool { return A[i] > A[j] }