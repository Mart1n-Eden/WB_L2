package main

import "fmt"

// Product - продукт, который мы строим
type Product struct {
    Part1 string
    Part2 string
    Part3 string
}

// Builder - интерфейс для строителя
type Builder interface {
    BuildPart1()
    BuildPart2()
    BuildPart3()
    GetResult() *Product
}

// ConcreteBuilder - конкретный строитель, реализующий интерфейс Builder
type ConcreteBuilder struct {
    product *Product
}

func NewConcreteBuilder() *ConcreteBuilder {
    return &ConcreteBuilder{product: &Product{}}
}

func (b *ConcreteBuilder) BuildPart1() {
    b.product.Part1 = "Part 1 built"
}

func (b *ConcreteBuilder) BuildPart2() {
    b.product.Part2 = "Part 2 built"
}

func (b *ConcreteBuilder) BuildPart3() {
    b.product.Part3 = "Part 3 built"
}

func (b *ConcreteBuilder) GetResult() *Product {
    return b.product
}

// Director - директор, который управляет строителем
type Director struct {
    builder Builder
}

func NewDirector(builder Builder) *Director {
    return &Director{builder: builder}
}

func (d *Director) Construct() {
    d.builder.BuildPart1()
    d.builder.BuildPart2()
    d.builder.BuildPart3()
}

func main() {
    builder := NewConcreteBuilder()
    director := NewDirector(builder)

    director.Construct()
    product := builder.GetResult()

    fmt.Println("Product Parts:")
    fmt.Println(product.Part1)
    fmt.Println(product.Part2)
    fmt.Println(product.Part3)
}

// Строитель (Builder) - это порождающий паттерн проектирования, который используется для создания сложных объектов
// пошагово. Он позволяет создавать различные представления объекта, используя один и тот же процесс построения.
// В этом паттерне создается отдельный класс (Строитель), ответственный за создание объекта, а также класс
// Директор, который контролирует процесс сборки. Это позволяет разделить процесс создания объекта на отдельные этапы
// и концентрироваться на их реализации пошагово.