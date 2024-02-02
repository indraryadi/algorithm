package main

import (
	"fmt"
)

// dfs using recursion
func dfs(start int, listVisited []bool, label []int, graph [][]int) {
	// check if out start index is already visited
	if listVisited[start] { //base case to exit recursion
		return
	} else {
		listVisited[start] = true
		fmt.Printf("vertice %d has been visited\n", start)
	}

	for i := range label {
		if graph[start][i] == 1 {
			dfs(i, listVisited, label, graph)
		}
	}
}
func main() {
	// bfs on matrix graph
	row := 5
	label := []int{0, 1, 2, 3, 4}

	graph := make([][]int, row)
	for i := range graph {
		graph[i] = make([]int, row)
	}

	graph[0][1] = 1
	graph[0][2] = 1
	graph[0][3] = 1
	graph[1][0] = 1
	graph[2][0] = 1
	graph[2][1] = 1
	graph[3][0] = 1
	graph[1][2] = 1
	graph[2][4] = 1
	graph[4][2] = 1

	fmt.Printf("  |")
	for x := range label {
		fmt.Printf("%d ", label[x])
	}
	fmt.Println()
	for i := 0; i < row; i++ {
		fmt.Printf("%d |", label[i])
		for j := 0; j < row; j++ {
			fmt.Printf("%d ", graph[i][j])
		}
		fmt.Println()
	}
	// end create the graph matrix
	listVisited := make([]bool, row)

	dfs(3, listVisited, label, graph)

}

//note

// on dfs( depth first search algo) we will go through all grap, base on starting point

// if we have graph like this

// 0----3
// |\
// | \
// |  \
// |---2---4
// |  /
// | /
// |/
// 1

// 0 connected to 1,2,and 3
// 1 connected to 0 and 2
// 2 connected to 0,1,and 4
// 3 connected to 0
// 4 connected to 2

// and if it represented to matrix will be like this
//   |0 1 2 3 4
// 0 |0 1 1 1 0
// 1 |1 0 1 0 0
// 2 |1 1 0 0 1
// 3 |1 0 0 0 0
// 4 |0 0 1 0 0

// on dfs wee need helper variable to store value of visited vertice, lets call it listVisitied, it will be an array of boolean that has max length same like our data(5 data)

// next we will create a function called dfs which require some parameter
// 1. start -> index of vertice
// 2. listVisited -> our helper variable that store visited vertice
// 3. label -> since we use labeling for the vertice we need to use it
// 4. graph -> our graph which is matrix 2d

// let say we start from index 0 of label

// (recursion 1)
// start = 0
// listVisited = {false,false,false,false,false}
// label = [0,1,2,3,4]

// first we check if listVisited[start] (if true it indicate the vertice of list at index start already visited) this checker also use as base case for our recursion
// if true
//     return
// else
//     it will change the value of listVisited[start] into true, so our listVisited will be like this:

//     listVisited = {true,false,false,false,false}

//     we print on terminal we already visited vertice 0

// and then we loop through our label
// remember now we are at label 0
// //   |0 1 2 3 4
// // 0 |0 1 1 1 0 => we are on this row

// wee loop through that row to check if we connected to other vertice or not

// i=0 => 0,0 = 0 not connected
// i=1 => 0,1 = 1 connected
// if we are connected we call again dfs function but now with start = 1

// (recursion 2)
// start = 1
// listVisited = {true,false,false,false,false}
// label = [0,1,2,3,4]

// same like recursion no 1 we check again if listVisited[start] is true or not, now listVisited[1] is false, so
// change the value of listVisited[start] into true, so our listVisited will be like this:

//     listVisited = {true,true,false,false,false}

//     we print on terminal we already visited vertice 1

// and then we loop through our label
// remember now we are at label 1
// //   |0 1 2 3 4
// // 1 |1 0 1 0 0 => we are on this row

// wee loop through that row to check if we connected to other vertice or not

// i=0 => 1,0 = 1 connected
// if we are connected we call again dfs function but now with start = 0

// (recursion 3)
// start = 0
// listVisited = {true,true,false,false,false}
// label = [0,1,2,3,4]

// same again we check if listVisited[start] is true
// now the value is true because we already visited it at recursion 1
// so we hit our base case, an we finish our recursion 3 and jump again to recursion 2

// re enter (recursion 2)
// start = 1
// listVisited = {true,false,false,false,false}
// label = [0,1,2,3,4]

// we continue our loop

// before  i=0 => 1,0 = 1 connected
// after   i=1 => 1,1 = 0 not connected
// after   i=2 => 1,2 = 1 connected
// if we are connected we call again dfs function but now with start = 2

// we do recursion again and a lot of recursion again untill we check all vertice

// because we use undirected graph so it will alway visited all vertice
