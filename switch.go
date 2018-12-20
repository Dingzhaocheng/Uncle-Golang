package main
import (
	"fmt"
	"runtime"
)
func main(){
	/* fmt.Print() */
	switch os := runtime.GOOS; os{
		case "darwin" :
		fmt.Print("OS X")
		case "linux" :
		fmt.Print("Linux");
		default:
		fmt.Printf("%s.\n", os);
	}
}
// Print一行输出,Println多行输出,Printf可以加入%s这样的C语言的变量
