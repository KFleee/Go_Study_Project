package main

var sortList [6]int

func init() {
	sortList[0] = 1
	sortList[1] = 2
	sortList[2] = 5
	sortList[3] = 4
	sortList[4] = 3
	sortList[5] = 10
}

func merge(i, j int, sign chan interface{}) {
	if i == j {
		sign <- 1
		return
	}
	med := (i + j) / 2
	if (j - i + 1) > 2 {
		first_sign := make(chan interface{})
		second_sign := make(chan interface{})
		go merge(i, med, first_sign)
		go merge(med+1, j, second_sign)
		<-first_sign
		<-second_sign
	}
	len := j - i + 1
	tempList := make([]int, len)
	p := i
	q := med + 1
	for index := 0; index < len; index++ {
		if p == med+1 && q <= j {
			tempList[index] = sortList[q]
			q++
		} else if q == j+1 {
			tempList[index] = sortList[p]
			p++
		} else {
			if sortList[p] > sortList[q] {
				tempList[index] = sortList[p]
				p++
			} else {
				tempList[index] = sortList[q]
				q++
			}
		}
	}
	for index, tempIndex := i, 0; index <= j; index, tempIndex = index+1, tempIndex+1 {
		sortList[index] = tempList[tempIndex]
	}
	sign <- 1
}
