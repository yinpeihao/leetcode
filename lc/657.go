package main

func main() {

}

func judgeCircle(moves string) bool {
	var cntL, cntU int
	var i int
	for i = 0; i < len(moves); i++ {
		switch moves[i] {
		case 'L':
			cntL++
		case 'R':
			cntL--
		case 'U':
			cntU++
		case 'D':
			cntU--
		}
	}
	return cntL == 0 && cntU == 0
}
