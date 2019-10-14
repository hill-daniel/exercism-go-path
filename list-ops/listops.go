package listops

type binFunc func(x, y int) int

type IntList []int

func (list IntList) Foldr(fn binFunc, initial int) int {
	i2 := len(list)
	return i2
}

func (list IntList) Foldl(fn binFunc, initial int) int {
	return 0
}

func (list IntList) Length()int{
	var counter int
	for range list {
		counter++
	}
	return counter
}

