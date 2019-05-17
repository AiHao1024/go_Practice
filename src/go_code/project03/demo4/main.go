package main
// import "fmt"
// import "unsafe"
//当要引入多个多个包时
import(
	"fmt"
	"unsafe"
)
//golang的数据类型
func main(){
	//整数类型  i=-128  int8的范围是-128--127   ---有符号  
	var i int8
	i=-128
	//无符号的整数类型 uint8  范围0---255
	var a uint8
	a=255
	var b=100;
	fmt.Println("i=",i)
	fmt.Println("a=",a)
	//Printf可以输出数据的类型 输出类型必须带 %T   查看一个变量占用的字节大小  必须带%d 后面跟上函数unsafe.Sizeof()
	fmt.Printf("b的类型是%T  b占用的字节数是%d",b,unsafe.Sizeof(b))

	//在整型变量使用时，尽量使用保小不保大的原则，即尽量使用占用空间小的数据类型  var age byte =50
	


 
 
}