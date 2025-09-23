package main

import "fmt"

func main() {
	var DispoPlaces = [25]bool{true, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}
	for i := 1; i <= len(DispoPlaces); i++ {
		if DispoPlaces[i-1] {
			fmt.Println("La place numéro", i, "est \033[032mdisponible\033[0m")
		} else {
			fmt.Println("La place numéro", i, "est \033[031mindisponible\033[0m")
		}
	}
}

/*
Enoncé 2:
Liste bool de 25 places

Pseudo code :

boucle i
creer liste 25 bool
si liste[i] est true
	écrire libre
sinon écrire false

*/
