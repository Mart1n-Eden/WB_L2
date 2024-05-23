package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func main() {
	// ntpTime, err := ntp.Time("pool.ntp.org")
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при получении времени: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Точное время: %s\n", ntpTime)
}
