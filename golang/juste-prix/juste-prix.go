package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func generate_random_number(maxValue int)int{
	randomNumber := rand.IntN(maxValue+1)
	if randomNumber == 0 {
		randomNumber++
	}
	return randomNumber
}

var lecteur = bufio.NewScanner(os.Stdin)

func demanderEntier(message string, mini, maxi int) int {
	for {
		fmt.Print(message)
		if !lecteur.Scan() {
			os.Exit(1) // entrée fermée (Ctrl+D)
		}
		n, err := strconv.Atoi(strings.TrimSpace(lecteur.Text()))
		if err != nil {
			fmt.Println("Ce n'est pas un nombre.")
			continue
		}
		if n >= mini && n <= maxi {
			return n
		}
		fmt.Printf("Entre %d et %d stp.\n", mini, maxi)
	}
}

func main(){
	maxTryNumber := 10
	tryNumber := 0
	minValue := 1
	maxValue := 100
	randomNumber := generate_random_number(maxValue) 

	fmt.Printf("Devine le nombre entre %d et %d\n", minValue, maxValue)

	for tryNumber<maxTryNumber {
		userNumber := demanderEntier("> ", minValue, maxValue)

		tryNumber++
			
		if userNumber==randomNumber {
			fmt.Printf("Bravo, trouvé en %d coups !\n", tryNumber)
			return
		} else if userNumber>randomNumber {
			fmt.Println("C'est moins !")
		} else {
			fmt.Println("C'est plus !")
		}
	}


}