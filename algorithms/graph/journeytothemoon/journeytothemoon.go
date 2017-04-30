package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	a int
	b int
}

type vertex struct {
	visited    bool
	neighbours []int
}

type input struct {
	numberOfAstronauts int
	numberOfPairs      int
}

func dfs(n, size int, adjMat map[int]vertex) int {
	size++
	v := adjMat[n]
	v.visited = true
	adjMat[n] = v

	for _, x := range v.neighbours {
		xv := adjMat[x]
		if !xv.visited {
			//if vertex is not visited
			dfs(x, size, adjMat)
		}
	}
	return size
}

func initAdjMat(pairs []pair) map[int]vertex {
	adjMat := make(map[int]vertex)
	for _, p := range pairs {
		var nv1, nv2 vertex
		u := adjMat[p.a]
		nv1.neighbours = append(u.neighbours, p.b)
		adjMat[p.a] = nv1

		v := adjMat[p.b]
		nv2.neighbours = append(v.neighbours, p.a)
		adjMat[p.b] = nv2
	}
	return adjMat
}

func (p *input) calcCombinations(pairs []pair) int {
	var disjointSets []int
	adjMat := initAdjMat(pairs)
	for j := 0; j < p.numberOfAstronauts; j++ {
		v := adjMat[j]
		if !v.visited {
			size := dfs(j, 0, adjMat)
			disjointSets = append(disjointSets, size)
		}
	}
	var sum, total int

	for k := 0; k < len(disjointSets); k++ {
		total += sum * disjointSets[k]
		sum += disjointSets[k]
	}

	return total
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var p input
	var pairs []pair

	fmt.Fscanf(reader, "%d %d\n", &p.numberOfAstronauts, &p.numberOfPairs)
	for i := 0; i < p.numberOfPairs; i++ {
		var p pair

		fmt.Fscanf(reader, "%d %d\n", &p.a, &p.b)
		pairs = append(pairs, p)
	}

	fmt.Println(p.calcCombinations(pairs))
}
