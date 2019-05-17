package main
import "fmt"

//定义全局变量，在函数外声明，输出在函数内输出
// var a1=100
// var a2=200
// var a3=300
// //上面的声明方式,也可以改成
// var (
// 	a1=100
// 	a2=200
// 	a3=300
// )

func main(){
   //golang变量的使用方式
   //第一种：指定变量类型，声明后若不赋值，使用默认值 输出效果：i=0
   var i int
   fmt.Println("i=",i) 
   //int的默认值是0


   //第二种：根据值自行判断变量类型（类型推导）
   var num = 10.11
   fmt.Println("num=",num)

   //第三种：省略关键字var 直接使用:=  注意： :=左侧的变量不应该是已经声明过的，不然会导致编译出错
   sum:=10
   fmt.Println("sum=",sum)

   //多变量声明
   //一次性声明多个变量  注意：变量声明了就必须使用，只输出n1  编译出错
   var n1,n2,n3 int
   fmt.Println("n1=",n1,"n2=",n2,"n3=",n3)

   //一次性声明多个变量2
   // var a,name,ss=100,"tom",18.33
   // fmt.Println("a=",a,"name=",name,"ss=",ss)

    //一次性声明多个变量3，省略关键字
	 a,name,ss:=100,"tom",18.33
	fmt.Println("a=",a,"name=",name,"ss=",ss)

}