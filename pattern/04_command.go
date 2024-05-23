package main

import "fmt"

// Receiver - Получатель, который выполняет конкретные операции
type Receiver struct{}

func (r *Receiver) Action() {
	fmt.Println("Выполнение операции")
}

// Command - Интерфейс команды
type Command interface {
	Execute()
}

// ConcreteCommand - Конкретная команда
type ConcreteCommand struct {
	receiver *Receiver
}

func NewConcreteCommand(receiver *Receiver) *ConcreteCommand {
	return &ConcreteCommand{receiver: receiver}
}

func (c *ConcreteCommand) Execute() {
	c.receiver.Action()
}

// Invoker - Инициатор, который запускает команды
type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func main() {
	receiver := &Receiver{}
	command := NewConcreteCommand(receiver)
	invoker := &Invoker{}

	invoker.SetCommand(command)
	invoker.ExecuteCommand()
}

// Паттерн "Команда" относится к поведенческим паттернам проектирования и используется для инкапсуляции запроса
// в виде объекта, позволяя тем самым параметризовать клиентов с другими запросами, организовать
// отложенное выполнение запросов, а также поддерживать отмену операций. Он состоит из следующих основных компонентов:
// - Receiver: объект, который выполняет конкретные операции;
// - Command: интерфейс, определяющий общий метод для всех команд;
// - ConcreteCommand: конкретная реализация команды, связывающая команду с получателем;
// - Invoker: инициатор, который запускает команды.
// В данном примере Receiver представляет объект, выполняющий операцию Action. ConcreteCommand реализует интерфейс Command,
// связывая команду с получателем (Receiver), а Invoker запускает команду. Паттерн "Команда" позволяет обеспечить
// отделение отправителя команды от получателя, а также параметризовать объекты клиентов с различными запросами.
