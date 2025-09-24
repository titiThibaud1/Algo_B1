package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var ListeEntrée = []int{}          //Initialisation de la Liste donnée
	for nb := 0; nb <= 5000000; nb++ { //Génère un tableau de 10 valeur aléatoire entre 0 et 5
		intrdm := rand.Intn(15)
		ListeEntrée = append(ListeEntrée, intrdm)
	}
	for random := len(ListeEntrée) - 1; random > 0; random-- { //Mélange aléatoirement la liste
		j := rand.Intn(random + 1)
		ListeEntrée[random], ListeEntrée[j] = ListeEntrée[j], ListeEntrée[random]
	}

	start := time.Now() //Début du timer

	ValeurMax := MaxValue(ListeEntrée)         //Appel de la fonction MaxValue qui sélectionne la valeur max du tableau
	ListeSortie := Tri(ListeEntrée, ValeurMax) //Appel de la fonction Tri qui tri le tableau
	fmt.Println(ListeSortie)                   //Affichage de ListeSortie

	fin := time.Since(start) //Fin du timer
	fmt.Println(fin)         //Affichage du temps de traitement
}

func MaxValue(ListeEntrée []int) int { //Fonction MaxValue pour déterminer le nombre de fois que la 1ère boucle de Tri va devoir boucler
	var ValeurMax int                       //Initialisation de ValeurMax
	for i := 0; i < len(ListeEntrée); i++ { //Boucle qui va parcourir toute la liste
		if ValeurMax < ListeEntrée[i] { // Condition qui sélectione le plus grand nombre possible
			ValeurMax = ListeEntrée[i]
		}
	}
	return ValeurMax //Retour du nombre maximum
}

func Tri(ListeEntrée []int, ValeurMax int) []int { //Fonction Tri avec un tableau de 8 valeur en entrée et un tableau de taille x en sortie
	var ListeSortie []int                                              //Initialisation de ValeurMax, tableau qui affichera le résultat
	for CounterValue := 0; CounterValue <= ValeurMax; CounterValue++ { //Boucle qui va parcourir chaque nombre jusqu'à la valeur max
		var ConteurNb int                                                        //Initialisation de CounterNb qui se fera écraser à chaque lecture de la boucle prédécente
		for CounterListe := 0; CounterListe < len(ListeEntrée); CounterListe++ { //Boucle qui va parcourire toute ma Liste
			if ListeEntrée[CounterListe] == CounterValue { //Condition que va incrémenter CounterNb pour chaque nombre trouvé égale à CounterValue (de la boucle précédente)
				ConteurNb++
			}
		}
		ListeSortie = append(ListeSortie, ConteurNb) //Ajout à un nouveau tableau, les valeurs sortie de CounterNb
	}
	return ListeSortie //retourne le la liste du résultat
}
