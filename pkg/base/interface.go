package base

import(
    . "fmt"
)
type human struct{
    name string
    age string
    phone string
}

func (h human) String() string{
    return h.name + " - " + h.age + " - " + h.phone
}
func InterfaceTest(){
    neither := human{"neither", "25", "13122222222"}
    Printf("I'm %v \n", neither)
}
