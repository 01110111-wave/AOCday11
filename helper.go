package main

type Monkeys struct {
	M0    []int64
	M1    []int64
	M2    []int64
	M3    []int64
	M4    []int64
	M5    []int64
	M6    []int64
	M7    []int64
	count []int64
}

var constant int64 = 17 * 3 * 5 * 7 * 11 * 19 * 2 * 13

func (Monkey *Monkeys) M0throw() {
	for _, worryness := range Monkey.M0 {
		Monkey.count[0]++
		worryness = worryness * 7
		worryness = worryness % constant
		if worryness%17 == 0 {
			Monkey.M5 = append(Monkey.M5, worryness)
		} else {
			Monkey.M3 = append(Monkey.M3, worryness)
		}
	}
	Monkey.M0 = []int64{}
}

func (Monkey *Monkeys) M1throw() {
	for _, worryness := range Monkey.M1 {
		Monkey.count[1]++
		worryness = worryness + 4
		worryness = worryness % constant
		if worryness%3 == 0 {
			Monkey.M0 = append(Monkey.M0, worryness)
		} else {
			Monkey.M3 = append(Monkey.M3, worryness)
		}
	}
	Monkey.M1 = []int64{}
}

func (Monkey *Monkeys) M2throw() {
	for _, worryness := range Monkey.M2 {
		Monkey.count[2]++
		worryness = worryness + 2
		worryness = worryness % constant
		if worryness%5 == 0 {
			Monkey.M7 = append(Monkey.M7, worryness)
		} else {
			Monkey.M4 = append(Monkey.M4, worryness)
		}
	}
	Monkey.M2 = []int64{}
}
func (Monkey *Monkeys) M3throw() {
	for _, worryness := range Monkey.M3 {
		Monkey.count[3]++
		worryness = worryness + 7
		worryness = worryness % constant
		if worryness%7 == 0 {
			Monkey.M5 = append(Monkey.M5, worryness)
		} else {
			Monkey.M2 = append(Monkey.M2, worryness)
		}
	}
	Monkey.M3 = []int64{}
}
func (Monkey *Monkeys) M4throw() {
	for _, worryness := range Monkey.M4 {
		Monkey.count[4]++
		worryness = worryness * 17
		worryness = worryness % constant
		if worryness%11 == 0 {
			Monkey.M1 = append(Monkey.M1, worryness)
		} else {
			Monkey.M6 = append(Monkey.M6, worryness)
		}
	}
	Monkey.M4 = []int64{}
}
func (Monkey *Monkeys) M5throw() {
	for _, worryness := range Monkey.M5 {
		Monkey.count[5]++
		worryness = worryness + 8
		worryness = worryness % constant
		if worryness%19 == 0 {
			Monkey.M2 = append(Monkey.M2, worryness)
		} else {
			Monkey.M7 = append(Monkey.M7, worryness)
		}
	}
	Monkey.M5 = []int64{}
}
func (Monkey *Monkeys) M6throw() {
	for _, worryness := range Monkey.M6 {
		Monkey.count[6]++
		worryness = worryness + 6
		worryness = worryness % constant
		if worryness%2 == 0 {
			Monkey.M0 = append(Monkey.M0, worryness)
		} else {
			Monkey.M1 = append(Monkey.M1, worryness)
		}
	}
	Monkey.M6 = []int64{}
}
func (Monkey *Monkeys) M7throw() {
	for _, worryness := range Monkey.M7 {
		Monkey.count[7]++
		worryness = worryness * worryness
		worryness = worryness % constant
		if worryness%13 == 0 {
			Monkey.M6 = append(Monkey.M6, worryness)
		} else {
			Monkey.M4 = append(Monkey.M4, worryness)
		}
	}
	Monkey.M7 = []int64{}
}
