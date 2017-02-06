package main

import "fmt"

func fabonaci(n int) int{
	if(n<2){
		return n
	}else{
		return fabonaci(n-2)+fabonaci(n-1)
	}
}

func main(){
	var i int
	for i=0;i<10;i++{
		fmt.Printf("%d\t",fabonaci(i))
	}
}