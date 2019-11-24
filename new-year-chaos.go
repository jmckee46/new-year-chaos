package main

import "fmt"

func minimumBribes(q []int32) {
	qLength := int32(len(q))
	var cost int32
	var pos int32
	var chaotic bool
	var temp int32

	for i := qLength - 1; i >= 0; i-- {
		pos = i + 1
		if q[i] != pos {
			if q[i] > pos+2 {
				chaotic = true
				break
			} else if i-1 >= 0 && q[i-1] == pos {
				temp = q[i-1]
				q[i-1] = q[i]
				q[i] = temp
				cost++
			} else if i-2 >= 0 && q[i-2] == pos {
				temp = q[i-2]
				q[i-2] = q[i-1]
				q[i-1] = q[i]
				q[i] = temp
				cost += 2
			}
		}

	}
	if chaotic {
		fmt.Println("Too chaotic")
	} else {
		fmt.Println(cost)
	}
}

func main() {}
