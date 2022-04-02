package main

// automaticShit 自动排班
// people 人员
// perUserNum 每天排班的人数
// num 从第几个开始排
func automaticShit(people []string, perUserNum, num int, days int) (shit [][]string) {
	totalNum := days * perUserNum
	tmp := make([]string, totalNum)
	shit = make([][]string, 0, days)
	for i := 0; i < totalNum; i++ {
		if num < len(people) {
			tmp[i] = people[num]
		} else {
			num = 0
			tmp[i] = people[num]
		}
		num++
	}
	for i := 0; i < totalNum; i += perUserNum {
		shitTmp := make([]string, 0, perUserNum)
		for j := 0; j < perUserNum; j++ {
			shitTmp = append(shitTmp, tmp[i+j])
		}
		shit = append(shit, shitTmp)
	}
	return
}
