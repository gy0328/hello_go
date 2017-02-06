package main
import(
	"fmt"
	"sync"
	"runtime"
)
var count int = 0
func counter(lock * sync.Mutex){
	lock.Lock()
	count++
	fmt.Println(count)
	lock.Unlock()
}
func main(){
	lock:=&sync.Mutex{}
	for i:=0;i<10;i++{
		go counter(lock)
	}
	for {
		lock.Lock()
		c:=count
		lock.Unlock()
		runtime.Gosched()
		if c >= 10{
			fmt.Println("goroutine end")
			break
		}
	}
}