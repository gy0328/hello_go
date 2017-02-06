package main
import "fmt"

type Books struct{
	title string
	author string
	subject string
	book_id int
}

func main(){
	var Book1 Books
	var Book2 Books

	Book1.title = "Go 语言"
	Book1.author ="www.test.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 123455

	Book2.title = "Python 语言"
	Book2.author = "www.python.com"
	Book2.subject = "python 语言教程"
	Book2.book_id = 234566

	fmt.Printf("Book 1 title : %s\n",Book1.title)
	fmt.Printf("Book 1 author : %s\n",Book1.author)
	fmt.Printf("Book 1 subject : %s\n",Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n",Book1.book_id)

	fmt.Printf("Book 2 title : %s\n",Book2.title)
	fmt.Printf("Book 2 author : %s\n",Book2.author)
	fmt.Printf("Book 2 subject : %s\n",Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n",Book2.book_id)
}