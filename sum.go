package main
import(
	"fmt"
	"time"
	"runtime"
)
var c = make(chan int)
func main(){
	cores:=runtime.NumCPU()
	start := time.Now().UnixNano()
	fmt.Println(cores)
	runtime.GOMAXPROCS(cores)
	//he :=0
	max,num:=10000000000,10
	for i:=0;i<num;i++{
		go sum((max/num) * i +1,(max/num)*(i+1),i)
		
	}
	/*for i:=0;i<num; i++{
		he = he + <-c
	}
	fmt.Println("he:",he)*/
	fmt.Println("take times is :",(time.Now().UnixNano() - start))
}

func sum(min,max,number int){
	s:=0
	fmt.Println("min:",min,"max:",max,"|",number)
	for i:=min;i<=max;i++{
		s= s+i
	}
}