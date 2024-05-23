package main

import "fmt"

// Product - интерфейс продукта
type Product interface {
    Name() string
}

// ConcreteProductA - конкретный продукт A
type ConcreteProductA struct{}

func (p *ConcreteProductA) Name() string {
    return "Product A"
}

// ConcreteProductB - конкретный продукт B
type ConcreteProductB struct{}

func (p *ConcreteProductB) Name() string {
    return "Product B"
}

// Creator - интерфейс создателя
type Creator interface {
    FactoryMethod() Product
}

// ConcreteCreatorA - конкретный создатель A
type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) FactoryMethod() Product {
    return &ConcreteProductA{}
}

// ConcreteCreatorB - конкретный создатель B
type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) FactoryMethod() Product {
    return &ConcreteProductB{}
}

func main() {
    // Используем конкретных создателей напрямую
    creatorA := &ConcreteCreatorA{}
    productA := creatorA.FactoryMethod()
    fmt.Println(productA.Name())

    creatorB := &ConcreteCreatorB{}
    productB := creatorB.FactoryMethod()
    fmt.Println(productB.Name())
}

// Паттерн "Фабричный метод" относится к порождающим паттернам проектирования и используется для создания объектов без
// необходимости указания конкретного класса создаваемого объекта. Он определяет интерфейс для создания объекта, но оставляет
// выбор конкретного класса создаваемого объекта на усмотрение подклассов. Это позволяет инкапсулировать создание объектов,
// делая систему более гибкой и расширяемой. В этом примере Product представляет интерфейс для продукта, а ConcreteProductA
// и ConcreteProductB - конкретные реализации продуктов. Creator представляет интерфейс для создателя, а ConcreteCreatorA и
// ConcreteCreatorB - конкретные реализации создателей, каждый из которых реализует свой фабричный метод для создания
// определенного типа продукта. Клиентский код может использовать интерфейс Creator для создания продуктов, не заботясь
// о конкретных классах продуктов.
