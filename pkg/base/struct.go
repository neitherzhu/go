package base

import(
    . "fmt"
)
//可以声明新的类型
type person struct{
    name string
    age int
}

func StructTest(){

    var p person
    
    p.name = "neitherzhu"
    p.age = 25
    
    //p := person{"neitherzhu", 25} //必须按顺序
    //p := person{age: 25, name: "neither"}//可以任意顺序
    
    Printf("我的名字叫%v，我%v岁了 \n", p.name, p.age)
    
    p1 := person{"Tom", 30}
    p2 := person{"Bob", 31}
    
    olderperson, diff := older(p1, p2)
    Printf("Older person is %v, diff is %v \n", olderperson.name, diff)
    
    //struct匿名字段 student 用person的所有字段，当然也可覆盖person的字段
    type student struct{
        person
        number int
        age int
    }
    
    neither := student{person{"neither", 25}, 1, 24}
    
    Printf("The student's name is %v \n", neither.name)
    Printf("His age is %v \n", neither.person.age)
    Printf("His school number is %v \n", neither.number)
    Printf("Last year, his age is %v \n", neither.age)
}

func older(p1, p2 person)(person, int){
    if p1.age > p2.age{
        return p1, p1.age - p2.age
    }
    return p2, p2.age - p1.age
}
