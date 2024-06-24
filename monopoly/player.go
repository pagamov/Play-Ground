package main

type Player struct {
	name  string
	money int
	turn  int
}

func makePlayers(number_of_players int, start_money int) []Player {
	res := []Player{}
	for i := 0; i < number_of_players; i++ {
		res = append(res, Player{name: "Player " + string(i+1), money: start_money, turn: i})
	}
	return res
}
