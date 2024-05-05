package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	name   string
	hora   int
	minuto int
	count  int
	total  int
)

func main() {

	scanner := bufio.NewScanner(os.Stdin) //cria um scanner

	fmt.Printf("Bem vindo ao Manager Tasks! Digite quantas tasks você quer criar:\n")
	if scanner.Scan() {
		countTotal := scanner.Text()
		count, _ = strconv.Atoi(countTotal)

		for i := 0; i <= count-1; i++ {
			fmt.Printf("Digite o nome da %d task: ", i+1)
			if scanner.Scan() {
				name = scanner.Text()
			}

			fmt.Println("Digite o tempo parcial de duração da task: digite o total de horas, seguindo por um espaço em branco, digite os minutos. ")
			if scanner.Scan() {
				totalHora := scanner.Text()
				//divide em 2 atraves dos espacos em branco
				valores := strings.Split(totalHora, " ")

				if len(valores) != 2 {
					fmt.Println("Por favor, insira exatamente dois valores separados por espaço. ")
					return
				}

				//valida e atribui os valores
				_, err1 := fmt.Sscanf(valores[0], "%d", &hora)
				_, err2 := fmt.Sscanf(valores[1], "%d", &minuto)

				if hora < 0 || minuto > 60 || minuto < 0 {
					fmt.Println("Por favor, insira um valor válido para hora. ")
					return
				}

				//converter string para int
				horaemInt, _ := strconv.Atoi(valores[0])
				minemInt, _ := strconv.Atoi(valores[1])

				// Concatenar os dois números como uma string
				strHour := strconv.Itoa(horaemInt) + "." + strconv.Itoa(minemInt)
				// Converter string concatenada em float
				horaMaisMin, _ := strconv.ParseFloat(strHour, 64)

				//soma o total
				total += int(horaMaisMin)

				//valida o erro
				if err1 != nil || err2 != nil {
					fmt.Println("Por favor, insira dois valores válidos. ")
					return
				}

				fmt.Printf("A %vª task: %v, demorará em torno de %v hora's e %v minuto's\n", i+1, name, hora, minuto)

				if count-1 == i && count > 1 {
					//total
					println("O tempo total é de: ", total, " horas")
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Erro ao ler a entrada:", err)
	}

}
