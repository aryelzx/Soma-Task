package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	total float64
	totalHours int
	totalMinutes int
) 

func main() {
	fmt.Printf("Bem vindo ao Manager Tasks! Digite quantas tasks você quer criar:\n")

	scanner := bufio.NewScanner(os.Stdin) 
	scanner.Scan()
	countInput := scanner.Text()

	count, err := strconv.Atoi(countInput)
	if err != nil {
		panic(err)
	}

	for i := 0; i < count; i++ {
		fmt.Printf("Digite o nome da %dª task : ", i+1)
		scanner.Scan()
		taskName := scanner.Text()
		
		fmt.Println("Digite o tempo parcial de duração da task:\n(digite o total de horas, seguindo por um espaço em branco, digite os minutos) ")
		scanner.Scan()
		hour := scanner.Text()
		
		splitHourMinutes := strings.Split(hour, " ")

		if len(splitHourMinutes) != 2 {
			fmt.Println("Por favor, insira exatamente dois valores separados por espaço. ")
			return
		}

		var splitedHour, splitedMinute int
		_, err1 := fmt.Sscanf(splitHourMinutes[0], "%d", &splitedHour)
		_, err2 := fmt.Sscanf(splitHourMinutes[1], "%d", &splitedMinute)

		if err1 != nil || err2 != nil {
			fmt.Println("Por favor, insira dois valores válidos. ")
			return
		}
		
		if splitedHour < 0 || splitedMinute >= 60 || splitedMinute < 0 {
			fmt.Println("Por favor, insira um valor válido para hora. ")
			return
		}

		totalHours += splitedHour
		totalMinutes += splitedMinute

		fmt.Printf("A %vª task: %v, foi computada.\n", i+1, taskName)

	}

	extraHours := totalMinutes / 60
	totalHours += extraHours
	totalMinutes %= 60

	fmt.Printf("O tempo total é de: %d horas e %d minutos\n", totalHours, totalMinutes)
}
