package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("Hangman.txt")//sert à lire le fichier Hangman.txt
	lignes := strings.Split(string(file), "\r")
	fmt.Print(lignes[0],//print le premier element
		lignes[len(lignes)-1])//print le dernier element
	for i, e := range lignes {
		if strings.Contains(e, "before") {//si jamais trouve le mot string "before" enchaine avec le reste du code
			e_pos, _ := strconv.Atoi(lignes[i+1][1:3]) // Mot après before convertit en int.
			fmt.Print(lignes[e_pos-1])                 // Mot à la ligne de la valeur de e_pos.
		} else if strings.Contains(e, "now") {//si dans la ligne il y a "now"
			mot_ascii := []byte(lignes[i-1]) // Mot précédent à now en ascii.
			e_pos := int(mot_ascii[2]) / len(lignes)//divise la valeur ascii du deuxieme charactere de "are" et le divise par la lenght de lignes
			fmt.Println(lignes[(e_pos)-1]) // Mot à la ligne de la valeur e_pos.
		}
	}
}