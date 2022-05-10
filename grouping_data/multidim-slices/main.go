package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Chocolote", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Ferengana", "Strawberry", "Jenkings"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
