package main

import "fmt"

const a = 121
//包级的声明在整个包中都能使用
//反之局部变量则作用域很小*

func main() {

  // himeko1()
   //himeko()
 //  himeko2()
   himeko3()
   var c = himeko6(3.5)
   fmt.Println(c)
}

func himeko1() {
   // a:= "himeko"
   //局部变量不可不使用
    for n:= 0; n < 5; n++{
       fmt.Println(n)
    }
 
    himeko()
 }

func himeko()  {
   fmt.Println(5)
}
//函数未声明 且顺序不对但还是正常调用?
//important 
//令人惊讶的是在go中与c中在函数的声明上有很大的差异,在包一级的声明中没有声明顺序的要求，只要声明了在包中都可以随意调用
//而对于函数本身更是没有定义的说法，我们常用的func实际上已经是在声明函数

//函数声明格式 函数名+参数列表+返回值列表+函数体

//example
func himeko6(f float64) float64{
   fmt.Println(f)
   return 2.5
}

func himeko2() {
   // a:= "himeko"
   //局部变量不可不使用
    for n:= 0; n < 5; n++{
       fmt.Println(n)
       break
    }
}

func himeko3() {
   // a:= "himeko"
   //局部变量不可不使用
    for n:= 0; n < 5; n++{
      if n-3 == 0{
         continue
      }
       fmt.Println(n)
    }
}
