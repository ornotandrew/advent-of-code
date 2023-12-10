package day_05

type Edge [3]int

func (e Edge) MaxSrcValue() int {
	return e[1] + e[2] - 1
}
