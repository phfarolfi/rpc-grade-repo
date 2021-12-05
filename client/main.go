package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type ItemNota struct {
	Matricula string
	CodDisciplina string
	Nota float64
}

type ConsultaNota struct {
	Matricula string
	CodDisciplina string
}

func showDatabase(database []ItemNota) {
	for _, data := range database {
		fmt.Printf("{ Matrícula: %s, CodDisciplina: %s, Nota: %.2f }\n", data.Matricula, data.CodDisciplina, data.Nota)
	}
}

func main() {
	var reply string
	var reply_notas []float64
	var database []ItemNota

	var option int
	var end = true

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	for end {
		fmt.Println("Digite o número correspondente à operação que você deseja executar e aperte 'Enter':")
		fmt.Println("")
		fmt.Println("1 - Cadastrar uma nota de um aluno")
		fmt.Println("2 - Consultar uma nota de um aluno")
		fmt.Println("3 - Consultar todas as notas de um aluno")
		fmt.Println("4 - Consultar CR de um aluno")
		fmt.Println("5 - Consultar todas as notas de todos os alunos")
		fmt.Println("0 - Sair")
		fmt.Println("")
		fmt.Print("Opção escolhida: ")
		fmt.Scan(&option)
		fmt.Println("")

		switch option {
			case 1:
				var matricula, cod_disciplina string
				var nota float64

				fmt.Println("Digite a matrícula do aluno: ")
				fmt.Scan(&matricula)

				fmt.Println("Digite o código da disciplina: ")
				fmt.Scan(&cod_disciplina)

				fmt.Println("Digite a nota do aluno: ")
				fmt.Scan(&nota)

				var item = ItemNota{matricula, cod_disciplina, nota}

				client.Call("API.CadastrarNota", item, &reply)
				fmt.Println("\n", reply)
			case 2:
				var matricula, cod_disciplina string

				fmt.Print("Digite a matrícula do aluno que deseja consultar a nota: ")
				fmt.Scan(&matricula)

				fmt.Print("Digite o código da displina do aluno que deseja consultar a nota: ")
				fmt.Scan(&cod_disciplina)

				var item_nota = ConsultaNota{matricula, cod_disciplina}

				client.Call("API.ConsultarNota", item_nota, &reply)
				fmt.Println("\n", reply)
			case 3:
				var matricula string

				fmt.Print("Digite a matrícula do aluno que deseja consultar as notas: ")
				fmt.Scan(&matricula)

				client.Call("API.ConsultarNotas", matricula, &reply_notas)
				fmt.Println("Notas:", reply_notas)
				fmt.Println("")
			case 4:
				var matricula string

				fmt.Print("Digite a matrícula do aluno que deseja consultar o CR: ")
				fmt.Scan(&matricula)

				client.Call("API.ConsultarCR", matricula, &reply)
				fmt.Println("\n", reply)
			case 5:
				client.Call("API.GetDatabase", "", &database)
				showDatabase(database)
			case 0:
				client.Call("API.GetDatabase", "", &database)
				showDatabase(database)

				end = false
		}
		fmt.Println("")
	}
}
