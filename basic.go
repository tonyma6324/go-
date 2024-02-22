package main

import "fmt"

/*
内建常量: true false iota nil

内建类型: int int8 int16 int32 int64
          uint uint8 uint16 uint32 uint64 uintptr
          float32 float64 complex128 complex64
          bool byte rune string error

内建函数: make len cap new append copy close delete
          complex real imag
          panic recover

*/

/*
主要声明手段
var, const ,type,func
*/
var s1 string 
const s2 float32 = 2.41

//type用法比较多元 example
//1,自定义类型
type mystr string
//2,定义结构体
type mystruct struct{
	name string
    bound int
}
//3,定义接口
type getsomething interface{
    start()
	end()
}
//4,定义函数类型 -----这样可以将函数比较自由的通过函数传递，而不需要用类似c语言中的函数指针

type my_func func (int,int) string

func main(){

//如果num1或num2被声明过，则转为赋值，但不可都被声明过，无法编译通过
//转为赋值的情况必须变量为同一作用域，不然会变成重新声明一个新变量

	num1,num2 := basic_func(1)
	point()
	fmt.Println(num1,num2)
}
 
//指针和c语言的大差不差

func point(){
	x := 1
	p := &x         // p, of type *int, points to x
	fmt.Println(*p) // "1"
	*p = 2          // equivalent to x = 2
	fmt.Println(x)  // "2"
	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"
}
/*
在Go语言中，返回函数中局部变量的地址也是安全的。例如下面的代码，调用f函数时创建局部变量v，在局部变量地址被返回之后依然有效，因为指针p依然引用这个变量。
*/

func mynew(){
	//可以用以下的方式创建变量，跟普通的没啥区别，可以看做语法糖
	p := new(int) 
}

//变量的生命周期
//显然包级全局变量的声明周期和整个程序的生命周期一样
//局部变量的生命周期则是动态的
//c语言中，函数返回之后函数内的变量的生命周期就结束了，特别在c中没有垃圾回收的机制,内存都由自己来管理
//堆栈上的内存被回收了之后，非常容易导致引用的错误以及内存的泄露
//而在go中,大多数情况下我们不太需要考虑这些,内存如何分配，在堆上还是栈上由编译器来决定，所以局部变量的生命周期是动态的
//在函数返回之后，这块内存也不一定会被回收。需要等待它的引用结束之后，才会被回收，生命周期结束

/*
关于go的垃圾回收的机制
可以简易理解为
从每个包级的变量以及函数内的局部变量开始遍历,如果变量不存在引用，不可达。则可以将这块内存给回收了
所以局部变量在函数返回之后不一定就生命周期结束了，如果还有引用，变量的生命周期仍将继续
这也就是为何在上面的指针一节中，指向局部变量的指针在函数返回后仍然有用，仍然指向这个变量

也就是说:局部变量的生命周期很有可能会超出它的作用域，这在c语言中是很难想象的事情

*/

var global *int

func stack(){
	a := 1
	p := &a
}

func heap(){
	a := 1
	global := &a
}
//显然即使不需要返回，stack函数结束之后a的生命周期就结束了，而在heap函数中，仍然能通过global指针来访问a，a仍在被引用

//所以在heap函数中，a的内存必须在堆上来分配，来保证仍能引用
//而在stack函数中，a的内存可以由栈来分配（当然也可能是堆）

//go的gc机制显然能帮助我们来管理内存，但也不是就完全不关注
//比如我们最好对变量的生命周期更好的管理，不然将一个短生命周期的指针保存在长生命周期的函数中，会阻止内存的回收


func basic_func(a int)(int,int){
    b:= 1
	fmt.Println(a+b)
	return 2,3
}
