// Boa. Em Go não existe herança clássica como em Java ou C#. O modelo é baseado em:

// Composição (um struct contém outro)

// Interfaces (definem comportamentos, e qualquer struct que implemente os métodos é do tipo da interface)

// Isso dá o efeito prático de polimorfismo.


package main

import "fmt"

// Interface comum (comportamento)
type Animal interface {
	Falar() string
}

// Struct "base"
type AnimalBase struct {
	Nome string
}

// Structs específicos
type Cachorro struct {
	AnimalBase
}

type Gato struct {
	AnimalBase
}

// Métodos específicos de cada tipo
func (c Cachorro) Falar() string {
	return c.Nome + " diz: Au au!"
}

func (g Gato) Falar() string {
	return g.Nome + " diz: Miau!"
}

func main() {
	animals := []Animal{
		Cachorro{AnimalBase{"Rex"}},
		Gato{AnimalBase{"Mimi"}},
	}

	for _, a := range animals {
		fmt.Println(a.Falar())
	}
}