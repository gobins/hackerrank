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
	cost        int
}

//Connect struct
type Connect struct {
	u int //city number
	v int //city number
}

//City struct
type City struct {
	neighbours []int
	visited    bool
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

func (q *Query) initAdjMat() map[int]City {
	adjCities := make(map[int]City)
	for _, c := range q.connections {
		var nc1 City
		c1 := adjCities[c.u]
		nc1.neighbours = append(c1.neighbours, c.v)
		adjCities[c.u] = nc1

		var nc2 City
		c2 := adjCities[c.v]
		nc2.neighbours = append(c2.neighbours, c.u)
		adjCities[c.v] = nc2
	}
	return adjCities
}

func dfs(city int, adjC map[int]City) {
	c := adjC[city]
	c.visited = true
	adjC[city] = c
	for _, n := range c.neighbours {
		nc := adjC[n]
		if !nc.visited {
			//if a city has not been visited
			dfs(n, adjC)
		}
	}
}

func (q *Query) processQuery() int {
	adjCities := q.initAdjMat()
	var connectedCom int
	var cost int
	if q.croad >= q.clib || q.m == 0 {
		cost = q.clib * q.n
	} else {
		for j := 1; j <= q.n; j++ {
			v := adjCities[j]
			if !v.visited {
				dfs(j, adjCities)
				connectedCom++
			}
		}
		cost = q.croad*(q.n-connectedCom) + (q.clib * connectedCom)
	}
	return cost
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	queries := readData(reader)

	for i := 0; i < toInt(queries); i++ {
		q := readQuery(reader)
		fmt.Println(q.processQuery())
	}
}
