package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	algoritmo que recebe dois valores:
	1 string declarando o nome
	1 float64 declarando o tempo;

	armazenar o nome e o tempo em um Map,
	com chave sendo o nome e valor sendo o tempo;

	sempre que um novo nome for inserido,
	verificar se o nome já existe no Map.

	se o nome já existir, retorna mensagem de erro;
	se não existir, armazenar o nome e o tempo no Map.

	ao final, imprimir o Map.
*/

var name string
var hora int
var minuto int

func main() {
	Tasks := make(map[string]float64)

	scanner := bufio.NewScanner(os.Stdin) //cria um scanner

	fmt.Println("Bem vindo ao soma Task!\nDigite o nome da task:")
	if scanner.Scan() {
		name = scanner.Text()
	}

	fmt.Println("Agora Digite o tempo parcial de duração da task:\ndigite o total de horas, seguindo por um espaço em branco, digite os minutos.")
	if scanner.Scan() {
		totalHora := scanner.Text()
		//divide em 2 atraves dos espacos em branco
		valores := strings.Split(totalHora, " ")

		if len(valores) != 2 {
			fmt.Println("Por favor, insira exatamente dois valores separados por espaço.")
			return
		}

		//pega o valor e erro e adiciona as variaveis
		_, err1 := fmt.Sscanf(valores[0], "%d", &hora)
		_, err2 := fmt.Sscanf(valores[1], "%d", &minuto)

		//converter string para int
		horaemInt, _ := strconv.Atoi(valores[0])
		minemInt, _ := strconv.Atoi(valores[1])

		// Concatenar os dois números como uma string
		strHour := strconv.Itoa(horaemInt) + "." + strconv.Itoa(minemInt)
		// Converter string concatenada em float
		horaMaisMin, _ := strconv.ParseFloat(strHour, 64)

		Tasks[name] = (horaMaisMin)

		for chave, valor := range Tasks {
			fmt.Println("chave:", chave, "valor:", valor)
		}

		//valida o erro
		if err1 != nil || err2 != nil {
			fmt.Println("Por favor, insira dois valores inteiros válidos.")
			return
		}
	}

	fmt.Printf("A task %v demorará em torno de %v hora's e %v minuto's", name, hora, minuto)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Erro ao ler a entrada:", err)
	}

}
