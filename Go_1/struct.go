package main
import "fmt"

func main(){
	type person struct{
		name string
		age int	
	}
	p:=person{name:"hritik",age:22}
	fmt.Println(p)

}
