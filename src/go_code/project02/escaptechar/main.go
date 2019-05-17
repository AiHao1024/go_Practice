package main
import "fmt" //fmt包中提供格式化，输出，输出的函数
func main(){
	//演示转义字符的使用
	//\t 制表符 用于格式输出   输出效果 tom jack
	fmt.Println("tom\tjack")

	// \n: 换行符 输出效果 hello 
	//                 jack
	fmt.Println("hello\njack")
	
	// \\: 一个\ 输出效果 a\b
	fmt.Println("a\\b")
	
	// \": 一个"  输出效果 tom"jack
	fmt.Println("tom\"jack")
	
	// \r:一个回车   输出效果  李四有钱   相当于覆盖了李四长度字段的以前的内容
	fmt.Println("张三有钱\r李四")

	fmt.Println("姓名\t年龄\t籍贯\t住址  \njoin\t12\t河北\t 北京")
}