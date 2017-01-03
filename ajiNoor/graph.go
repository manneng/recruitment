package main

import( "fmt"
	"strconv"
)

func main()  {

	arrays := []int{1, 4, 5, 6, 8, 2} /*change arrays value to try another graph*/

	fmt.Print("\n\nASCENDING")
	ascending(arrays)
	fmt.Print("\n\nREVERSE")
	descending(arrays)

}

func ascending(allData []int){
	i:= 0
	j:=0
	input := allData

	graph(input)
	for i = len(input)-1; i > 0; i-- {
		for j = i ; j <= len(input)-1 ; j++{
			if(input[j] <= input[j-1]){
				input[j], input[j-1] = input[j-1], input[j];
				graph(input)
			}
		}
	}
}

func descending(allData []int){
	i:= 0
	j:=0
	input := allData

	graph(input)
	for i = len(input)-1; i > 0; i-- {
		for j = i ; j <= len(input)-1 ; j++{
			if(input[j] > input[j-1]){
				input[j], input[j-1] = input[j-1], input[j];
				graph(input)
			}
		}
	}
}

func graph(allData []int) {
	j:= 0
	k:=0
	var n, biggest int

	data := allData

	for _,v:=range data {
		if v>n {
			n = v
			biggest = n
		}
	}
	for i := 0; i <= biggest; i++ {
		fmt.Println("")
		for j = 0; j < len(data); j++ {
			if biggest-data[j] >= i{
				fmt.Print("   ")
			}else {
				fmt.Print(" | ")
			}
		}
	}
	fmt.Println("")
	for k = 0; k < len(data); k++ {
		t := strconv.Itoa(data[k])
		fmt.Print(" "+t+" ")
	}
}
