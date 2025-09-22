package main

import "fmt"

func main() {
	var Salle int
	fmt.Println("Numéro de la salle de l'exam :")
	fmt.Scan(&Salle)
	fmt.Println("L'élève a obtenu un : ", QualiteNote(), ", à l'exam en salle", Salle)

}

func QualiteNote() string {
	var Note float32
	fmt.Println("Note de l'élève :")
	fmt.Scan(&Note)
	if Note < 0 || Note > 20 {
		fmt.Println("\033[034mNote invalide\033[0m")
		QualiteNote()
	}
	if Note < 7 {
		return "\033[031mEchec\033[0m"
	}
	if Note > 13 {
		return "\033[032mExcellent\033[0m"
	} else {
		return "\033[033mPassable\033[0m"
	}
}

/*
Exo 6
Note : passable, echec, excellent en float
Forcement entre 0 et 20 sinon erreur
Salle de classe


Pseudo code :

input note
(input classe)
(lore)

Si note < 0 ou > 20
	retourner une erreur
	redemander la note
Si note < 7
	Note echec
Si note > 13
	Note Excellente
Sinon
	Note passable

Afficher le résultat + le numero de la salle + lore
*/
