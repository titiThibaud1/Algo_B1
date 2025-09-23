package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var Classe [100000000]int
	for i := 0; i < len(Classe); i++ {
		fmt.Println(Classe[i])
	}
	fin := time.Since(start)
	fmt.Println(fin)
}

/*
PB 1 :

var Classe []string{"prenom1", ..., "prenom26"}
Pour i de 0 à len(Classe) exclu par pas de 1
	Print Classe[i]

2
+3*n
= n linéaire
*/
