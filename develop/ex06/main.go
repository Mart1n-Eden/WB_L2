package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Определение флагов командной строки
	fields := flag.String("f", "", "fields to select")
	delimiter := flag.String("d", "\t", "delimiter")
	separated := flag.Bool("s", false, "only lines with delimiter")
	flag.Parse()

	// Разбор запрашиваемых полей
	fieldSet := make(map[int]bool)
	if *fields != "" {
		for _, fieldStr := range strings.Split(*fields, ",") {
			fieldNum := atoi(fieldStr)
			if fieldNum != 0 {
				fieldSet[fieldNum] = true
			}
		}
	}

	// Чтение стандартного ввода
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// Проверка, содержит ли строка разделитель, если это требуется
		if !*separated || strings.Contains(line, *delimiter) {
			// Разделение строки на поля
			fields := strings.Split(line, *delimiter)

			// Выбор запрошенных полей и их объединение с разделителем
			var selectedFields []string
			for i, field := range fields {
				if len(fieldSet) == 0 || fieldSet[i+1] {
					selectedFields = append(selectedFields, field)
				}
			}
			fmt.Println(strings.Join(selectedFields, *delimiter))
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения стандартного ввода:", err)
		os.Exit(1)
	}
}

// atoi преобразует строку в целое число. Если строка пуста, возвращает 0.
func atoi(s string) int {
	if s == "" {
		return 0
	}
	n, _ := strconv.Atoi(s)
	return n
}
