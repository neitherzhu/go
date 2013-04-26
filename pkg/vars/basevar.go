package vars

import(
    f "fmt"
)

func BaseVars(){
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
