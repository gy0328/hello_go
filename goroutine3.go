package main
import(
	"fmt"
)
var ch1 chan int = make(chan int)
var ch2 chan int = make(chan int)
func say(s string){
	fmt.Println(s)
	ch1  <-1
	ch2 <-1
}
func main(){
	go say("hello")
	<- ch1
	<- ch2
}