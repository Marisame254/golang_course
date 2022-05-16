// package acdc asks if you are ready to rock!!!
package acdc

// Sum add an unlimited number of values
func Sum(x1 ...int) int {
	s := 0
	for _, v := range x1 {
		s += v
	}
	return s
}
