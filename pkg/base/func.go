package base

import(
    f "fmt"
)

func gotoFunc(){
    i := 0
Here:
    f.Printf("%v \n", i)
    i++
    if i <= 5{
        goto Here
    }
}

func FuncTest(){
    a, b := 10, 11
    c := getMax(a, b)
    f.Printf("%v \n", c)
    
    //变参
    uncertainArgs("a", 1, 2, 3, 4, 5, 8)
    
    //一般的函数修改的是传入参数的副本，不会真正修改参数
    //要修改指针，只需将传入参数的类型前加*， 传入参数是，在前面加&表示参数所在的地址
    
    z := 1
    f.Printf("z改变前：%v \n", z)
    changePointer(&z)
    f.Printf("z改变后：%v \n", z)
    
    //defer延迟执行 
    //defer是先进后出的顺序执行的
    deferTest()
    
}

func getMax(a, b int)int{
    if a > b {
        return a
    }
    
    return b
}

func uncertainArgs(a string, arg ...int){
    f.Printf("传入的第1个确定的参数为： %v \n", a)
    for k, v := range arg{
        f.Printf("传入的第%v个参数为：%v \n", k+1, v)
    }
}

func changePointer(a *int){
    *a = *a + 1
}

func deferTest(){
    for i := 0; i < 5; i++{
        defer f.Printf("%v \n", i)
    }
}
