package main
import("fmt")
func rand() chan int{
	ch :=make(chan int)
	go func(){
		for{
			select{
			case ch<-0:
			case ch<-1:
			case ch<-6:
			}
		}
	}()
	return ch
}

func main(){
	generator := rand()
	for i:=0;i<10;i++{
		fmt.Println(<-generator)
	}
}