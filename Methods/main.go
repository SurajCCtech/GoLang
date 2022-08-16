package main

import "fmt"

type cricket struct {
	playerName    string
	playerRole    string
	playerJnumber int
}

// Reciever type struct cricket and the name c
func (c cricket) details() {
	fmt.Printf("Player name: %v\n", c.playerName)
	fmt.Printf("Role of the player: %v\n", c.playerRole)
	fmt.Printf("Jersey number: %v\n", c.playerJnumber)
}

//Reciever type pointer *cricket and the name C
// func (C *cricket) setName(newName string) {
// 	C.playerName = newName
// }

func main() {
	cric := cricket{playerName: "Yuvraj Singh", playerRole: "Batsman", playerJnumber: 12}
	cric.details()
	//(*cricket).setName(&cric, "Virat Kohli")
	// (&cric).setName("MS Dhoni")
	// fmt.Printf("New player name: %v\n", cric.playerName)
}
