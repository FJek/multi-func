package ctrl

import "fzw/proto/dao"

import "fmt"

func AddNumber() {
	err := dao.AddNumber()
	if err != nil {
		fmt.Println("写入失败")
	}
	fmt.Println("succeed")
}

func GetNumbers() {
	numbers := dao.GetNumbers()
	if len(numbers) <= 0 {
		fmt.Println("no record")
		return
	}
	dao.PrintNumber(numbers)
}
