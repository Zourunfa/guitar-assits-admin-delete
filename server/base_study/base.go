package base_study

import (
	"fmt"
	"math"
)

// 定义 Point 结构体
type Point struct {
	x float64
	y float64
}

// Point 结构体的构造函数
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

/*
当你使用指针类型作为方法的接收者时，
这意味着你可以在方法内部修改接收者的字段。
如果你使用值类型作为接收者，那么在方法内部修改接收者的字段将只影响方法内部的副本，
而不会影响原始对象。因此，如果你希望在方法内部修改接收者的字段，并且希望这些修改影响到原始对象，
那么你应该选择指针类型作为方法的接收者。
*/
// 计算两点之间的距离
func (p *Point) Distance() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func main() {
	// 使用构造函数创建 Point 对象
	point := NewPoint(3, 4)
	// 调用方法计算距离
	distance := point.Distance()
	fmt.Println("Distance:", distance)
}
