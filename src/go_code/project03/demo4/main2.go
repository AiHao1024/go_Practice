package main
import(
	"fmt"
)
func main(){
	//浮点类型数据 用于存放小数  一般推荐使用float64  更加准确
	//浮点数据类型分类  单精度float32（4字节）  双精度float64（8字节）
	//十进制
	var n1 float32
	n1=12.66
    //默认的小数的类型是float64
	var n2=1.2  

	fmt.Println("n1=",n1)
	fmt.Printf("n2的数据类型是%T \n",n2 )

	//科学计数法
	n3:=5.1234e2 //等价于5.1234*10的2次方
	n4:=5.1234E2 //等价于5.1234*10的2次方
	n5:=5.1234e-2 //等价于5.1234c除以10的2次方
	fmt.Println("n3=",n3,"n4=",n4,"n5=",n5)
}