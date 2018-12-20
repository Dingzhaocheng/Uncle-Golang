// struct 结构体
//一个结构体（struct）就是一个字段的集合。
package main
import "fmt"
type Person struct {
    name string
	age int
}
// 结构体文法通过直接列出字段的值来新分配一个结构体。
var(
		v1 = Person{"小红", 1}
		v2 = Person{age: 20}
		v3 = Person{}
		p = &Person{"22", 33} //返回一个指向结构体的指针
	)
func main(){
	/* fmt.Println(Person{"xiaoming", 20}) */
	/*v := Person{"xiaoming", 20} // 可用.来获取字段名的值
	 v.name = "xiaohong"
	fmt.Println(v.name) */
	/* p := &v  // 结构体指针绑定
	p.name = "xiaoli"
	fmt.Println(v.name)  */
	fmt.Println(v1,v2,v3,p)
	
}