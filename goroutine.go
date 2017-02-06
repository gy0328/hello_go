package main
import(
	"fmt"
	"time"
)
var c chan int
func ready(w string,sec int64){
	time.Sleep(time.Duration(sec * 1e9))
	fmt.Println(w," is ready")
	c <- 1
}
func main(){
    c = make(chan int)
	go ready("Tee ",2)
	go ready("Coffee",1)
fmt.Println("I'm waiting, but not too long")
//time.Sleep(5*1e9)
<-c
<-c
}

