package main

import (
	"fmt"
	"time"
)
// Um channel em Go serve para permitir a comunicação segura entre goroutines (threads leves do Go). Ele funciona como uma "ponte" por onde dados podem ser enviados e recebidos, garantindo sincronização e evitando condições de corrida (race conditions).

// Principais usos de channels:

// Trocar dados entre goroutines.
// Sincronizar a execução (por exemplo, esperar uma goroutine terminar).
// Implementar padrões de concorrência, como pipelines.
// Exemplo simples:

// // Cria um channel de inteiros
// c := make(chan int)

// // Envia um valor para o channel em uma goroutine
// go func() {
//     c <- 42 // envia 42 para o channel
// }()

// // Recebe o valor do channel
// valor := <-c
// println(valor) // imprime 42


func simple_main() {
	println("Channels")
	
	c := make(chan int)
	println(c)
	go func() {
		c <- 42 // envia 42 para o channel
	}()
	go func() {
		c <- 43
	}()
	valor := <-c
	println(valor) // imprime 42
}


func main() {
	// Exemplo 1: Channel não bufferizado
	unbuffered := make(chan string)

	go func() {
		fmt.Println("Enviando mensagem no channel não bufferizado...")
		unbuffered <- "ping" // vai bloquear até alguém receber
		fmt.Println("Mensagem enviada no channel não bufferizado")
	}()

	time.Sleep(time.Second) // simula atraso
	fmt.Println("Recebendo do channel não bufferizado:", <-unbuffered)

	// Exemplo 2: Channel bufferizado (capacidade 2)
	buffered := make(chan string, 2)

	buffered <- "mensagem 1"
	fmt.Println("Enviada mensagem 1 no channel bufferizado")

	buffered <- "mensagem 2"
	fmt.Println("Enviada mensagem 2 no channel bufferizado")

	// Se tentar enviar mais uma, vai bloquear (buffer cheio)
	go func() {
		fmt.Println("Tentando enviar mensagem 3 no channel bufferizado...")
		buffered <- "mensagem 3"
		fmt.Println("Mensagem 3 enviada")
	}()

	time.Sleep(time.Second)
	fmt.Println("Recebendo do channel bufferizado:", <-buffered)
	fmt.Println("Recebendo do channel bufferizado:", <-buffered)
	fmt.Println("Recebendo do channel bufferizado:", <-buffered)
}