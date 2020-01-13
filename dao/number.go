package dao

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

type Number struct {
	Name   string
	Number string
}

// 读文件
func GetNumbers() []Number {
	numbers := make([]Number, 0)
	// =========
	inputFile, err := os.Open("book/number.txt")
	if err != nil {
		log.Error().Msg(err.Error())
		return nil
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	var number Number
	for {
		inputString, readerError := inputReader.ReadString('\n')
		info := strings.Split(inputString, " ")
		number.Name = info[0]
		number.Number = info[1]
		numbers = append(numbers, number)
		if readerError == io.EOF {
			return numbers
		}
	}
}

func AddNumber() error {
	var (
		name, number string
	)
	// 控制台输入
	fmt.Println("请输入名称：")
	fmt.Scanln(&name)
	fmt.Println("请输入手机：")
	fmt.Scanln(&number)
	// 写入文件
	record := name + " " + number
	outputFile, outputError := os.OpenFile("book/number.txt", os.O_WRONLY|os.O_CREATE, 0666)

	if outputError != nil {
		fmt.Printf("open fileError")
		return outputError
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	_, err := outputWriter.WriteString(record)
	if err != nil {
		fmt.Println("写入文件错误")
		return err
	}
	// 刷新文件流
	// outputWriter.Flush()

	return nil
}

// 打印信息
func PrintNumber(numbers []Number) {
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("名称：%s , 手机号：%s\n", numbers[i].Name, numbers[i].Number)
	}
}
