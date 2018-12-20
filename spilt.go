// spilt function
package main

import "fmt"
func spilt(sum int) (x, y int){
	x = sum * 4 / 9
	y = sum - x
	return
}
func main(){
	fmt.Println(spilt(20))
}