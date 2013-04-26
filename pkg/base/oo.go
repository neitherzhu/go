package base

import(
    . "fmt"
    "math"
)

type rectangle struct{
    width, height float32
}

type circle struct{
    radius float32
}

type square struct{
    rectangle
}

func (r rectangle) area() float32{
    return r.width*r.height
}

func (c circle) area() float32{
    return c.radius*c.radius*math.Pi
}
//通过struct前面加*来改变指针
func (c *circle) changeRadius(r float32){
    //此处不需要写成*c.radius，go自动识别是要修改指针
    c.radius = r;
}

//method 重写
func (s square) area() float32{
    return s.width*s.width
}

func OO(){
    r := rectangle{12, 10}
    c := circle{10}
    s := square{rectangle{10, 12}}
    
    Printf("长方形的面积：%v \n", r.area());
    Printf("圆形的面积：%v \n", c.area());
    Printf("正方形的面积：%v \n", s.area());
    
    c.changeRadius(20)
    Printf("圆形的面积：%v \n", c.area());
}

