package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка чтения ввода:", err)
			continue
		}
		input = strings.TrimSpace(input)

		// Разбиение строки на команду и аргументы
		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "cd":
			if len(parts) < 2 {
				fmt.Fprintln(os.Stderr, "Не указан аргумент для команды 'cd'")
				continue
			}
			err := os.Chdir(parts[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при смене директории:", err)
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при получении текущей директории:", err)
				continue
			}
			fmt.Println(dir)
		case "echo":
			fmt.Println(strings.Join(parts[1:], " "))
		case "kill":
			if len(parts) < 2 {
				fmt.Fprintln(os.Stderr, "Не указан аргумент для команды 'kill'")
				continue
			}
			cmd := exec.Command("kill", parts[1])
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при выполнении команды 'kill':", err)
			}
		case "ps":
			cmd := exec.Command("ps")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при выполнении команды 'ps':", err)
			}
		default:
			fmt.Fprintf(os.Stderr, "Неизвестная команда: %s\n", parts[0])
		}
	}
}
