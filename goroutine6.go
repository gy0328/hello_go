package main
import("fmt"
	"runtime"
)
var quit chan int = make(chan int)
func loop(){
	for i:=0;i<10;i++{
	
		fmt.Printf("%d ",i)
	}
	quit <-0
}
func main(){
	runtime.GOMAXPROCS(1)
	for i:=0;i<4;i++{
		go loop()
	}

	for i:=0;i<4;i++{
		<- quit
	}
}