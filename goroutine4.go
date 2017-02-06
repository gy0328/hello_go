package main
import("fmt")
var ch chan int = make(chan int)
func foo(id int){
	ch <- id
}

func main(){
	for i:=0;i<5;i++{
		go foo(i)
	}
	for i:=0;i<5;i++{
		fmt.Println(<-ch)
	}
	test:=make(chan int,3)
	test <-1
	test <-2
	test <-3
	for v:=range test{
		fmt.Println(v)
		if len(test)<=0{
			break
		}
	}
}