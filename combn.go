package main

import "fmt"

func Combn(n int) {
	//création du tableau modulé par n
	nbr := make([]int, n)

	//initialisation des valeurs par défaut
	for a := 0; a < n-1; a++ {
		nbr[a] = n - a - 1
	}

	//WHILE tant que mon premier chiffre(quand affiché) est différent de sa valeur finale on ajoute 1
	for nbr[n-1] != 10-n {

		nbr[0]++

		//WHILE tant qu'un chiffre est égal à 10 dans mon tableau
		for In(10, nbr) {
			//Pour chaque élément du tableau
			for i := 0; i < len(nbr); i++ {

				//si cet élément est égal à 10
				if nbr[i] == 10 {
					//le prochain augmente
					nbr[i+1]++

					//Tous les nombres à gauche de l'élément actuel (lui y compris) prennent la valeur de l'élément de droite +1
					for j := 0; j <= i; j++ {
						nbr[i-j] = nbr[i-j+1] + 1
					}
				}
			}
		}

		//affiche nombres en inversant l'ordre
		for z := n - 1; z >= 0; z-- {
			fmt.Print(nbr[z])
		}
		if nbr[n-1] < 10-n {
			fmt.Print(", ")
		}

	}

}

func In(nb int, TBL []int) bool {
	for k := 0; k < len(TBL); k++ {
		if TBL[k] == nb {
			return true
		}
	}
	return false
}
