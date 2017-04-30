package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Query struct
type Query struct {
	n           int //number of cities
	m           int //number of roads
	clib        int //cost to build a library
	croad       int //cost to build a road
	connections []*Connect
}

//Connect struct
type Connect struct {
	u int //city number
	v int //city number
}

func readData(reader *bufio.Reader) string {
	data, _ := reader.ReadString('\n')
	return data
}

func toInt(input string) int {
	s := strings.Split(input, "\n")
	data, _ := strconv.Atoi(s[0])
	return data
}

func readQuery(reader *bufio.Reader) Query {
	var q Query
	line := readData(reader)
	s := strings.Split(line, " ")
	q.n = toInt(s[0])
	q.m = toInt(s[1])
	q.clib = toInt(s[2])
	q.croad = toInt(s[3])
	for i := 0; i < q.m; i++ {
		var c Connect
		line := readData(reader)
		s := strings.Split(line, " ")
		c.u = toInt(s[0])
		c.v = toInt(s[1])
		q.connections = append(q.connections, &c)
	}
	return q
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	data := readData(reader)
	var input []Query
	for i := 0; i < toInt(data); i++ {
		input = append(input, readQuery(reader))
	}
	fmt.Println(input)
}
