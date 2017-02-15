package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readData(reader *bufio.Reader) string {
	data, _ := reader.ReadString('\n')
	return data
}

func checkContains(s string) bool {
	stringArray := strings.Split(s, "")
	var flag bool
	for i := 0; i < len(s); i++ {
		if strings.Count(s, stringArray[i]) > 1 {
			flag = true
			break
		}
	}
	return flag
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	inputCount, _ := strconv.Atoi(readData(reader))
	// for i := 0; i < inputCount; i++ {
	// 	a := readData(reader)
	//
	// }
	fmt.Println(inputCount)

}
