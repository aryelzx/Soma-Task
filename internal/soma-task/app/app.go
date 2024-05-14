package app

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aryelzx/Soma-Task/internal/soma-task/task"
	"github.com/gen2brain/beeep"
)

type App struct {
	TasksRunning []*task.Task
	quitChan     chan any
	TaskChan     chan task.Task
}

func NewApp() *App {
	return &App{
		quitChan: make(chan any),
		TaskChan: make(chan task.Task, 10),
	}
}

func (a *App) SetNotifys() {
	for {
		select {
		case task := <-a.TaskChan:
			fmt.Println(task.Name)
			err := beeep.Notify(task.Name, string(task.Message.Payload), "assets/information.png")
			if err != nil {
				panic(err)
			}
		default:
		}
	}
}

func (a *App) settedMessage(task *task.Task) {
	time.Sleep(time.Duration(*&task.TimeSeconds) * time.Second)
	a.TaskChan <- *task

	if task.Repeat {
		a.settedMessage(task)
	}
}

func (a *App) Start() {
	go a.SetNotifys()

	fmt.Printf("Bem vindo ao Manager Tasks! Digite quantas tasks você quer criar:\n")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	countTarefas, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	countTarefasTotal := countTarefas
	var tarefas = "Tarefas:\n(hour:minute:second, date_now)\n\n"

	for countTarefas > 0 {
		countTarefasPrinter := countTarefasTotal - countTarefas + 1
		fmt.Printf("Digite o nome da task %d: ", countTarefasPrinter)
		scanner.Scan()
		tarefaName := scanner.Text()

		fmt.Println("Digite o tempo parcial de duração da task, numero em segundos:")
		scanner.Scan()
		tarefaTime := scanner.Text()

		fmt.Println("Digite a mensagem agendar:")
		scanner.Scan()
		message := scanner.Text()

		fmt.Println("Deseja que essa mensagem se repita? (y/N)")
		scanner.Scan()
		repeatAnswer := scanner.Text()
		repeat := false
		if repeatAnswer == "y" {
			repeat = true
		}

		tarefaTimeSeconds, err := strconv.Atoi(tarefaTime)
		if err != nil {
			fmt.Println(err)
			continue
		}

		var tarefaTimeHours int
		var tarefaTimeMinutes int
		countTarefas--

		go a.settedMessage(task.NewTask(uint64(countTarefas), tarefaName, bytes.NewReader([]byte(message)), uint16(tarefaTimeSeconds), repeat))

		if tarefaTimeSeconds >= 60 {
			tarefaTimeMinutes := tarefaTimeSeconds / 60
			tarefaTimeSeconds = tarefaTimeSeconds % 60

			if tarefaTimeMinutes >= 60 {
				tarefaTimeHours := tarefaTimeMinutes / 60
				tarefaTimeMinutes = tarefaTimeMinutes % 60

				tarefas += fmt.Sprintf("%s: %d:%d:%d, %v\n", tarefaName, tarefaTimeHours, tarefaTimeMinutes, tarefaTimeSeconds, time.Now())
				continue
			}

			tarefas += fmt.Sprintf("%s: %d:%d:%d, %v\n", tarefaName, tarefaTimeHours, tarefaTimeMinutes, tarefaTimeSeconds, time.Now())
			continue
		}

		tarefas += fmt.Sprintf("%s: %d:%d:%d, %v\n", tarefaName, tarefaTimeHours, tarefaTimeMinutes, tarefaTimeSeconds, time.Now())
	}

	fmt.Println("\n", tarefas)

	<-a.quitChan
}
