package main
import(
	"fmt"
)
type SortInterface interface{
	sort()
}
type Sortor struct{
	name string
}

func main(){
	arry:=[]int{6,1,3,4,8,4,2,0,9,7}
	learnsort:= Sortor{name:"选择排序--从小到大--不稳定--n*n---"}
	learnsort.sort(arry)
	fmt.Println(learnsort.name,arry)
	arry2:=[]int{6,1,3,4,8,4,2,0,9,7}
	learnsort2:=Sortor{name:"冒泡排序--从小到大--稳定--n*n---"}
	learnsort2.sort2(arry2)
	fmt.Println(learnsort2.name,arry2)
	arry3:=[]int{6,1,3,4,8,4,2,0,9,7}
	learnsort3:=Sortor{name:"快速排序--从小到大--不稳定--n*n---"}
	learnsort3.sort3(arry3)
	fmt.Println(learnsort3.name,arry3)
	arry4:=[]int{6,1,3,4,8,4,2,0,9,7}
	learnsort4:=Sortor{name:"插入排序--从小到大--稳定--n*n---"}
	learnsort4.sort4(arry4)
	fmt.Println(learnsort4.name,arry4)
}
func(sorter Sortor)sort(arry[]int){
	arrylength:=len(arry)
	for i:=0;i<arrylength;i++{
		min:=i
		for j:=i+1;j<arrylength;j++{
			if arry[j]< arry[min]{
				min =j
			}
		}
		t:=arry[i]
		arry[i] =arry[min]
		arry[min] = t
	}
}
func (sorter Sortor) sort2(arry []int){
	done:=true
	arrylength:=len(arry)
	for i:=0;i<arrylength && done;i++{
		done = false
		for j:=0;j<arrylength-i-1;j++{
			done = true
			if arry[j] > arry[j+1]{
				t:=arry[j]
				arry[j]=arry[j+1]
				arry[j+1]=t
			}
		}
	}
}
func (sorter Sortor) sort3(arry []int){
	if len(arry) <= 1{
		return
	}
	mid:=arry[0]
	i:=1
	head,tail:=0,len(arry)-1
	for head <tail{
		if arry[i] > mid {
			arry[i],arry[tail]=arry[tail],arry[i]
			tail--
		}else{
			arry[i],arry[head]=arry[head],arry[i]
			head++
			i++
		}
	}
	arry[head]= mid
	sorter.sort3(arry[:head])
	sorter.sort3(arry[head+1:])
}

func (sorter Sortor) sort4(arry []int){
	arrylength:=len(arry)
	for i,j:=1,0;i<arrylength;i++{
		temp:=arry[i]
		for j=i;j>0 && arry[j-1] > temp;j--{
			arry[j] =arry[j-1]
		}
		arry[j]=temp
	}
}