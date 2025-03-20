package freeGameTimes

type Win struct {
	FreeAddTimes [5]int
}

func (w *Win) Get(count int) (bool, int) {
	if count < 1 || count > 5 {
		return false, 0
	}
	times := w.FreeAddTimes[count-1]
	if times == 0 {
		return false, 0
	}
	return true, times
}
