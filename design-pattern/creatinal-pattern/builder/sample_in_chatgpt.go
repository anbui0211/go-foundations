package builder

import "fmt"

// Product là đối tượng cuối cùng cần được xây dựng.
type Product struct {
	partA string
	partB string
	partC string
}

// Builder là interface mô tả các bước để xây dựng đối tượng.
type Builder interface {
	BuildPartA()
	BuildPartB()
	BuildPartC()
	GetResult() Product
}

// ConcreteBuilder xây dựng và lưu trữ đối tượng sản phẩm.
type ConcreteBuilder struct {
	product Product
}

func (b *ConcreteBuilder) BuildPartA() {
	b.product.partA = "PartA"
}

func (b *ConcreteBuilder) BuildPartB() {
	b.product.partB = "PartB"
}

func (b *ConcreteBuilder) BuildPartC() {
	b.product.partC = "PartC"
}

func (b *ConcreteBuilder) GetResult() Product {
	return b.product
}

// Director là lớp quản lý quy trình xây dựng với một Builder.
type Directorr struct {
	builder Builder
}

func (d *Directorr) Construct() {
	d.builder.BuildPartA()
	d.builder.BuildPartB()
	d.builder.BuildPartC()
}

func main() {
	builder := &ConcreteBuilder{}
	director := Directorr{builder: builder}
	director.Construct()
	product := builder.GetResult()

	fmt.Printf("Product: %+v\n", product)
}
