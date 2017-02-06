package main

import(
	"fmt"
)
func main(){
	var a int = 101
	var b int = 1
	var result ,result2 int
	result = max(a,b)
	fmt.Printf("最大值是 %d\n",result)
	array:=[]int{1,23,32,44,33,23,100,300,232,0}
	result2 = max2(array)
	fmt.Printf("数组中最大值是 %d\n",result2)
}
func max(a,b int) int{
	var result int
	if(a > b){
		result=a
	}else{
		result=b
	}
	return result
}

func max2(a []int) int{
    num:= len(a)
    maxNum := 0
	for i:=0;i<num;i++{
		if a[i] > maxNum{
			maxNum =a[i]
		}
	}
	return maxNum
}