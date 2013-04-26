package process

import(
    f "fmt"
)

func Process(){
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
