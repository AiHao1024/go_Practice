package main
import "fmt"
func main(){
	//变量使用的注意事项
	//1:该区域的数据值可以在同一类型范围内不断变化
	//2：变量在同一个作用域（函数）不能重名
	//3：变量三要素  变量=变量名+值+数据类型
	var i = 10
	i=20
	i=30
	fmt.Println("i=",i)
	//i=1.2  出错  不能改变数据类型
	


	//程序中+的使用  如果两边是数值 则是加法运算   字符串则是拼接
	var a=1
	var b=2
	var c=a+b
	fmt.Println("c=",c)

	var d="hello"
	var e="world"
	var f=d+e
	fmt.Println("f=",f)
}