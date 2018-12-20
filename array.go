// array 数组
package main
import "fmt"
var c[2] string
func main(){
    c[0] = "h"
	c[1] = "a"
	fmt.Println(c[0],c[1])
	 a := [3] int {1,2,3}
	fmt.Println(a[0],a[1])
}
// 数组的长度是其类型的一部分，因此数组不能改变大小