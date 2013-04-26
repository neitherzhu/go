package vars

import(
    f "fmt"
)

func ArrayVars(){
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
