package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	//i1, i2 := 0, 1
	//return func() int {
	
		//fmt.Printf("i1:%d\n" , i1)
//		fmt.Printf("i2:%d\n" , i2)
	//	ret := i1+i2
	//	i1 = i2
	//	i2 = ret
	//	return ret
	//}
		f, g := 1, 0
	return func() int {
	    //f =g
		//g = f+g
		fmt.Printf("before\nf:%d\n" , f)
		fmt.Printf("g:%d\n" , g)		
		tmp := f
		f = g
		g = tmp + g
		//f, g = g, f+g
		fmt.Printf("after\nf:%d\n" , f)
		fmt.Printf("g:%d\n" , g)
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
