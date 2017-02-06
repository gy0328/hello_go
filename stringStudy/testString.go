package main
import(
	"fmt"
	"strings"
)
func main(){
	fmt.Println("查找子串是否在指定的字符串中")
	fmt.Println("Contains 函数用法")
	fmt.Println(strings.Contains("seafood","foo"))
	fmt.Println(strings.Contains("seafood","bar"))
	fmt.Println(strings.Contains("seafood",""))
	fmt.Println(strings.Contains("",""))
	fmt.Println(strings.Contains("我是中国人","我"))

	fmt.Println("")
    fmt.Println("ContainsAny 函数的用法")
    fmt.Println(strings.ContainsAny("team", "i"))        // false
    fmt.Println(strings.ContainsAny("failure", "u & i")) // true
    fmt.Println(strings.ContainsAny("foo", ""))          // false
    fmt.Println(strings.ContainsAny("", ""))             // false

	fmt.Println("")
	fmt.Println(" ContainsRune 函数的用法")
	fmt.Println(strings.ContainsRune("我是中国", '我')) // true 注意第二个参数，用的是字符
	fmt.Println("")
	fmt.Println(" Count 函数的用法")
	fmt.Println(strings.Count("cheese", "e")) // 3
	fmt.Println(strings.Count("five", "")) // before & after each rune result: 5 , 源码中有实现
	fmt.Println("")
	fmt.Println(" EqualFold 函数的用法")
	fmt.Println(strings.EqualFold("Go", "go")) //大小写忽略 fmt.Println("")
	fmt.Println(" Fields 函数的用法")
	fmt.Println("Fields are: %q", strings.Fields(" foo bar baz ")) //["foo" "bar" "baz"] 返回一个列表

	//相当于用函数做为参数，支持匿名函数
	for _, record := range []string{" aaa*1892*122", "aaa\taa\t", "124|939|22"}{
		fmt.Println(strings.FieldsFunc(record, func(ch rune) bool {
			switch {
			case ch > '5':
				return true
			}
			return false
		}))
	}
}
