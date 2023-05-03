package test

func Fibu(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	return Fibu(num-1) + Fibu(num-2)
}
