package main

func QuadA(x, y int) {
	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			//Coins
			if (col == 0 || col == x-1) && (row == 0 || row == y-1) {
				print(0)

				//lignes Haut et Bas
			} else if (col != x-1 || col != 0) && (row == 0 || row == y-1) {
				print("-")

				//Bords
			} else if (col == 0 || col == x-1) && (row != 0 || row != y-1) {
				print("|")

				//Centre
			} else {
				print(" ")
			}
		}
		print("\n")
	}
}
