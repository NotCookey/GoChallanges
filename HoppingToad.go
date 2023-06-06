package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main(){
	scanner:=bufio.NewReader(os.Stdin)
	fmt.Print(">> ")
	input, _ := scanner.ReadString('\n')
	input=strings.TrimSpace(input)

	Max, Repeat := Max(strings.Split(input,"_"))
	
	if Repeat>1{
		fmt.Printf("It took %d seconds to cross the river, and the longest jump, which occurred %d times, was %d blocks long.\n", CountTime(input), Repeat, Max)
	} else {
		fmt.Printf("It took %d seconds to cross the river, and the longest jump was %d blocks long.\n", CountTime(input), Max)
	}
}

func CountTime(input string) int{
	cactus:=strings.Count(input,"#")*2
	pad:=strings.Count(input,"_")

	return cactus+pad
}

func Max(list []string) (int,int){
	chains:=make([]int,len(list))
	repeat:=0
	
	for index,_:=range list{
		chains[index]=len(list[index])
	}

	max := chains[0]
	for i := 1; i < len(chains); i++ {
		if chains[i] > max {
			max = chains[i]
		}
	}

	for x:=0;x<len(chains);x++{
		if max==chains[x]{
			repeat++
		}
	}

	return max, repeat
}