package main

//每日一题：话说有十二个鸡蛋，有一个是坏的（重量与其余鸡蛋不同，不知道比正常蛋重还是轻），
//现要求用天平称三次，称出哪个鸡蛋是坏的。

func GetBadEgg(eggs []int) int {
	set1, set2, set3 := eggs[0:4], eggs[4:8], eggs[8:12]
	if GetListSum(set1) == GetListSum(set2) {

	}
}

func GetListSum(list []int) int {
	var ans int
	for _, val := range list {
		ans += val
	}
	return ans
}
