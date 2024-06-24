package main

type Card struct {
	name  string
	value int
}

type Game struct {
	cards []Card
}

func makeField() Game {
	cards := []Card{
		Card{"start", 0},
		Card{"brown", 15},
		Card{"chest", 0},
		Card{"brown", 57},
		Card{"tax", 100},
		Card{"neutral", 500},
		Card{"blue", 68},
		Card{"chance", 0},
		Card{"blue", 71},
		Card{"blue", 81},

		Card{"jail", 0},
		Card{"pink", 91},
		Card{"tax", 124},
		Card{"pink", 97},
		Card{"pink", 112},
		Card{"neutral", 700},
		Card{"orange", 125},
		Card{"chest", 0},
		Card{"orange", 148},
		Card{"orange", 208},

		Card{"park", 0},
		Card{"red", 211},
		Card{"chance", 0},
		Card{"red", 215},
		Card{"red", 228},
		Card{"neutral", 1000},
		Card{"yellow", 271},
		Card{"yellow", 320},
		Card{"tax", 8000},
		Card{"yellow", 370},

		Card{"go to jail", 0},
		Card{"green", 404},
		Card{"green", 440},
		Card{"chest", 0},
		Card{"green", 550},
		Card{"neutral", 1500},
		Card{"chance", 0},
		Card{"blue", 562},
		Card{"tax", 200},
		Card{"blue", 1800},
	}
	game := Game{cards: cards}
	return game
}
