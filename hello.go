package main
import "fmt"

func main(){
	fmt.Printf("hello,world\n")
	const(
		a = iota
		b
		c
		d="haha"
		e
		f = 100
		g
		h = iota
		i
	)
	fmt.Println(a,b,c,d,e,f,g,h,i)
}