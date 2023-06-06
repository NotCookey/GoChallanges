package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

const check="HelloWorld"

func main(){
	scanner:=bufio.NewReader(os.Stdin)
	fmt.Print(">> ")
	input, _ := scanner.ReadString('\n')

	fmt.Print(reverseCheck(strings.TrimSpace(input)))
}

func reverseCheck(str string) string{
	if (len(str)<=len(check)){
		if str[0]>='A' && str[0]<='Z'{
			if str==check{
				return "!!!"
			} else if str==check[0:len(str)]{
				return check[len(str):len(check)]
			}
		}
	}
	return "???"
}