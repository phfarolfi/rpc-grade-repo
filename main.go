package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
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

type API int

var database[]ItemNota

func (a *API) GetDatabase(empty string, reply *[]ItemNota) error {
	*reply = database

	return nil
}

func (a *API) CadastrarNota(item ItemNota, reply *string) error {
	var alreadyExist = false
	var indexFound int

	for idx, data := range database {
		if data.Matricula == item.Matricula && data.CodDisciplina == item.CodDisciplina {
			alreadyExist = true
			indexFound = idx
		}
	}

	if alreadyExist {
		database[indexFound].Nota = item.Nota
		str := "A nota foi sobreescrita."
		*reply = str
	} else {
		database = append(database, item)
		str := "A nota foi adicionada."
		*reply = str
	}
	
	return nil
}

func (a *API) ConsultarNota(item ConsultaNota, reply *string) error {
	var getItem ItemNota
	var found = false

	for _, data := range database {
		if data.Matricula == item.Matricula && data.CodDisciplina == item.CodDisciplina {
			found = true
			getItem = data
		}
	}

	if found {
		nota := "Nota: "+fmt.Sprintf("%.2f", getItem.Nota)
		*reply = nota
	} else {
		str := "ERRO: A nota não foi encontrada"
		*reply = str
	}

	return nil
}

func (a *API) ConsultarNotas(matricula string, reply *[]float64) error {
	var itemList []float64

	for _, data := range database {
		if data.Matricula == matricula {
			itemList = append(itemList, data.Nota)
		}
	}

	*reply = itemList // Ele deve retornar uma lista com todos as notas do aluno, mas como fazer isso em formato string?

	return nil
}

func (a *API) ConsultarCR(matricula string, reply *string) error {
	var quantidade_notas = 0.0
	var soma_notas = 0.0
	var found = false

	for _, data := range database {
		if matricula == data.Matricula {
			found = true
			soma_notas += data.Nota
			quantidade_notas += 1
		}
	}

	if found {
		cr := fmt.Sprintf("%.2f", (soma_notas/quantidade_notas))
		*reply = "CR: "+cr
	} else {
		str := "ERRO: Não há notas cadastradas para essa matrícula"
		*reply = str
	}

	return nil
}

func main() {
	var api = new(API)
	err := rpc.Register(api)

	if err != nil {
		log.Fatal("Error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("Listener error", err)
	}

	log.Printf("serving rpc on port %d", 4040)

	err = http.Serve(listener, nil)
	
	if err != nil {
		log.Fatal("error serving:", err)
	}
}