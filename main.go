package main

import (
	
	"fmt"
	"time"
)

type No struct {
	valor   int
	proximo *No
}

type ListaCircular struct {
	cabeca *No
}

func (lista *ListaCircular) Adicionar(valor int) {
	no := &No{valor: valor}
	if lista.cabeca == nil {
		lista.cabeca = no
		no.proximo = lista.cabeca
	} else {
		atual := lista.cabeca
		for atual.proximo != lista.cabeca {
			atual = atual.proximo
		}
		atual.proximo = no
		no.proximo = lista.cabeca
	}
}

func (lista *ListaCircular) Remover(valor int) {
	if lista.cabeca == nil {
		return
	}
	atual := lista.cabeca
	anterior := lista.cabeca

	for atual.valor != valor {
		anterior = atual
		atual = atual.proximo
	}

	if atual == lista.cabeca {
		proximoNo := lista.cabeca.proximo
		for anterior.proximo != lista.cabeca {
			anterior = anterior.proximo
		}
		lista.cabeca = proximoNo
		anterior.proximo = lista.cabeca
	} else {
		anterior.proximo = atual.proximo
	}
}
func pausar(duracao time.Duration) {
	time.Sleep(duracao)
}


func main() {
	var numCrianca, numPasses int

	fmt.Println("Bem-vindo ao jogo da batata quente!")
	fmt.Print("Digite o número de crianças: ")
	fmt.Scanln(&numCrianca)
	fmt.Print("Digite o número de passes antes da eliminação: ")
	fmt.Scanln(&numPasses)

	crianca := &ListaCircular{}
	for i := 1; i <= numCrianca; i++ {
		crianca.Adicionar(i)
	}

	fmt.Println("Começando o jogo!")

	currentCrianca := crianca.cabeca

	for numCrianca > 1 {
		for i := 0; i < numPasses-1; i++ {
			currentCrianca = currentCrianca.proximo
		}

		fmt.Println("Eliminando a criança", currentCrianca.valor)
		crianca.Remover(currentCrianca.valor)
		numCrianca--
		currentCrianca = currentCrianca.proximo
		pausar(1 * time.Second)
	}

	fmt.Println("A criança vencedora é:", crianca.cabeca.valor)
}
