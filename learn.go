package main

import(
	f "fmt"
)

func baseVars(){
    //定义多个变量
	var v1, v2, v3 int
	//定义变量并初始化
	//var v4 int = 1
	//同时初始化多个变量
	var v5, v6, v7 int = 5, 6, 7
    //对于已经申明但是没用使用的变量，在编译时会报错
    
    v1, v2 ,v3 = v5, v6, v7
    f.Printf("v1 = %v; v2 = %v; v3 = %v \n", v1, v2, v3)
    //简短申明（只能用在内部函数）
	v8, v9, v10 := 8, 9, 10
	
	f.Printf("v8 = %v; v8 = %v; v10 = %v \n", v8, v9, v10)
	//_(下划线)是一个特殊的变量名，任何赋予它的值都会被丢弃
	_, v0 := 2, 0
	
	f.Printf("v0 = %v; and _ is gone! \n", v0)
	
	//定义常量
    const a string = "string"
    const pi float64 = 3.1415926
    
    f.Printf("const a = %v; pi = %v \n", a, pi)
    /****************************************************/
    
    var s string = "hello"
    //s[0] = "c"不能通过这种方式修改字符串的值
    //第一种修改字符串的方式
    c := []byte(s) //将字符串转成byte类型
    c[0] = 'c'
    s2 := string(c) //将byte类型再转换成字符串类型
    
    f.Printf("s is changed to s2: %v \n", s2)
    //第二种修改字符串的方式
    s = "c" + s[1:] //切片操作
    
    f.Printf("s is changed to s: %v \n", s)
    /****************************************************/
    
    //申明一个多行的字符串
    sl := `hello
          world`
          
    f.Printf("A muli string \"sl\": %v \n", sl)
    
    //iota关键字，默认从0开始，每次调用+1
    const(
        x = iota
        y
        z
    )
    f.Printf("x = %v; y = %v; z = %v \n", x, y, z)
    
    //每次遇到const关键字，iota会被重置为0
    const zz = iota
    f.Printf("zz = %v \n", zz)
}

func arrayVars(){
    //定长数组定义
    var arr [2]int
    arr[0] = 1
    arr[1] = 2
    
    f.Printf("arr[0] = %v; arr[1] = %v \n", arr[0], arr[1])
    //另一种定长数组定义
    a1 := [3]string{"a1", "a2", "a3"}
    
    f.Printf("a1[2] = %v \n", a1[2])
    //可以省略长度采用...，Go会根据元素个数来计算长度
    b1 := [...]int{1, 2, 3, 4, 5}
    f.Printf("b1数组的长度为： %v \n", len(b1))
    //多维数组定义
    c1 := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
    //如果数组内部的类型与外部一样，可以省略类型
    //c1 := [2][4]int{{1 ,2 ,3 ,4}, {5, 6, 7, 8}}
    
    
    f.Printf("c1[1][2] = %v \n", c1[1][2])
    
    //数组内部没有定义元素值时，默认为0/空／false
    d1 := [2][4]int{{1, 2}, {3, 4}}
    
    f.Printf("d1[1][3] = %v \n", d1[1][3])
    
    e1 := [2][3]string{{"a", "b"}}
    
    f.Printf("e1[1][0] = %v \n", e1[1][0])
    
    f1 := [2][2]bool{{true}, {false}}
    
    f.Printf("f1[1][1] = %v \n", f1[1][1])
}

func mapVars(){
    //声明map map[keytype]valuetype 必须要make初始化
    var numbers map[string]int
    numbers = make(map[string]int)
    numbers["one"] = 1
    f.Printf("numbers的第一个数是：%v \n", numbers["one"])
    //或者另外一种声明方式
    anotherNumbers := make(map[string]int)
    anotherNumbers["one"] = 1
    anotherNumbers["two"] = 2
    f.Printf("anotherNumbers的二个数是：%v \n", anotherNumbers["two"])
    //还有一种声明方式
    an := map[string]float32{"one":0.1, "two": 1.2, "three": 2.3}
    f.Printf("an的第三个数是： %v \n", an["three"])
    
    //map[key]有2个返回值， 第一个是value 第二个是是否存在这个key, 存在则为true，不存在为false
    four, hasFour := an["four"]
    if hasFour{
        f.Printf("an存在第四个元素： %v \n", four)
    }else{
        f.Printf("an不存在第四个元素 \n")
    }
}

func process(){
    //if 不需要括号
    var x int = 11
    f.Printf("x = %v \n", x)
    if x > 10{
        f.Printf("x > 10 \n")
    }else{
        f.Printf("x <= 10 \n")
    }
    
    //可以在if条件语句中声明一个变量，作用域只在这个if条件逻辑中
    if y := 10; y > 10{
        f.Printf("y > 10 \n")
    }else {
        f.Printf("y <= 10 \n")
    }
    
    //goto
    gotoFunc()
    
    //for
    sum := 0
    for index := 0; index < 10; index++{
       sum += index
    }
    f.Printf("从0加到9的和为： %v \n", sum)
    
    myMap := map[int]string{1: "one", 2: "two", 3: "three", 4: "four"}
    for k, v := range myMap{
        f.Printf("key = %v; value = %v \n", k, v)
    }
}

func gotoFunc(){
    i := 0
Here:
    f.Printf("%v \n", i)
    i++
    if i <= 5{
        goto Here
    }
}

func funcTest(){
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
func main(){
    //baseVars()
    //arrayVars()
    //mapVars()
    //process()
    funcTest()
}
