package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Virus - структура, яка представляє вірус
type Virus struct {
	Genome        string
	InfectedCells int
}

// People - структура, яка представляє людину, яку вірус може заражати
type People struct {
	ID       int
	Infected bool
}

// InfectionInterface - інтерфейс для представлення властивостей віруса
type Cell interface {
	BecomeInfected() bool
	IsInfected() bool
}

// перевіряємо чи заражена людина
func (c *People) IsInfected() bool {
	return c.Infected
}

// Рандомно заражаємо людину
func (c *People) BecomeInfected() bool {
	infectionAttempt := RandBool()
	c.Infected = infectionAttempt
	return infectionAttempt
}

// NewVirus - функція для створення нового вірусу з заданим геномом
func NewVirus(genome string) *Virus {
	return &Virus{
		Genome: genome,
	}
}

// Infect - метод намагається інфикувати людину вірусом
func (v *Virus) Infect(cell Cell) {
	if !cell.IsInfected() {
		// Додав перевірку чи відбулось зараження
		if cell.BecomeInfected() {
			v.InfectedCells++
		}
	}
}

// Replicate - метод для реплікації вірусу
func (v *Virus) Replicate() *Virus {
	newVirus := &Virus{
		Genome: v.Genome,
	}
	fmt.Println("Virus is replicating...")
	return newVirus
}

// Replicate - метод для реплікації вірусу
func (v *Virus) ReplicateWithMutation(newGenome string) *Virus {
	v.mutate(newGenome)
	newVirus := &Virus{
		Genome: v.Genome,
	}
	fmt.Println("Virus is replicating...")
	return newVirus
}

// Mutate - метод для мутації вірусу
func (v *Virus) mutate(newGenome string) {
	v.Genome = newGenome

	fmt.Println("Virus is mutating...")
}

func RandBool() bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(2) == 1
}

func main() {

	virus := NewVirus("AGCTAGCT")
	numberOfPeople := 50
	// Створення клітин і вірус заражає їх
	for i := 0; i < numberOfPeople; i++ {
		pers := &People{ID: i + 1}
		virus.Infect(pers)
		// fmt.Println(i, pers.ID)
	}

	// Вірус реплікується та мутує
	virus2 := virus.Replicate() // копіюємо вірус
	virusMutable := virus.ReplicateWithMutation("AGCAAGCT")

	fmt.Println(virus, virus2, virusMutable)
	fmt.Printf("Number of infected people is %d out of %d", virus.InfectedCells, numberOfPeople)
}
