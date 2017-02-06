package main
import(
	"fmt"
	"time"
)
var quit chan int
func foo(id int){
	fmt.Println(id)
	quit <-0
}
func main(){
	start:= time.Now().UnixNano()
	count:=1000
	quit = make(chan int)
	for i:=0;i< count;i++{
		go foo(i)
	}
	for i:=0;i<count;i++{
		<-quit
	}
	fmt.Println("take times is :",time.Now().UnixNano()-start)
}