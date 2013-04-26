package base

import(
    f "fmt"
)

func MapVars(){
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
