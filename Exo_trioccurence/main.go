package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var ListeEntrée = []int{}
	for nb := 0; nb <= 5000000; nb++ {
		intrdm := rand.Intn(15)
		ListeEntrée = append(ListeEntrée, intrdm)
	}
	for random := len(ListeEntrée) - 1; random > 0; random-- {
		j := rand.Intn(random + 1)
		ListeEntrée[random], ListeEntrée[j] = ListeEntrée[j], ListeEntrée[random]
	}

	start := time.Now()

	ValeurMax := MaxValue(ListeEntrée)
	ListeSortie := Occurence(ListeEntrée, ValeurMax)
	fmt.Println(ListeSortie)

	fin := time.Since(start)
	fmt.Println(fin)
}

func MaxValue(ListeEntrée []int) int {
	var ValeurMax int
	for i := 0; i < len(ListeEntrée); i++ {
		if ValeurMax < ListeEntrée[i] {
			ValeurMax = ListeEntrée[i]
		}
	}
	return ValeurMax
}

func Occurence(ListeEntrée []int, ValeurMax int) []int {
	var ListeSortie []int
	for CounterValue := 0; CounterValue <= ValeurMax; CounterValue++ {
		var ConteurNb int
		for CounterListe := 0; CounterListe < len(ListeEntrée); CounterListe++ {
			if ListeEntrée[CounterListe] == CounterValue {
				ConteurNb++
			}
		}
		ListeSortie = append(ListeSortie, ConteurNb)
	}
	return ListeSortie
}
