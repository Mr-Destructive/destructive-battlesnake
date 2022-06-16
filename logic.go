package main

import (
	"log"
	"math/rand"
    "reflect"
)

func exists(list []string, element string) bool {
    for _, item := range list {
        if element == item{
            return true
        }
    }
    return false
}

func map_exists(mapList []Coord, m Coord) bool {
    for _, pos := range mapList {
        if reflect.DeepEqual(pos, m){
            return true
        }
    }
    return false
}

func info() BattlesnakeInfoResponse {
	log.Println("INFO")
	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "mr_destructive",       
		Color:      "#000030", 
		Head:       "pixel",
		Tail:       "pixel", 
    }
}

func start(state GameState) {
	log.Printf("%s START\n", state.Game.ID)
}

func end(state GameState) {
	log.Printf("%s END\n\n", state.Game.ID)
}

func move(state GameState) BattlesnakeMoveResponse {
	possibleMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}


	myHead := state.You.Body[0] 
	myNeck := state.You.Body[1] 
	if myNeck.X < myHead.X {
		possibleMoves["left"] = false
	} else if myNeck.X > myHead.X {
		possibleMoves["right"] = false
	} else if myNeck.Y < myHead.Y {
		possibleMoves["down"] = false
	} else if myNeck.Y > myHead.Y {
		possibleMoves["up"] = false
	}

	// TODO: Step 1 - Don't hit walls.
	// Use information in GameState to prevent your Battlesnake from moving beyond the boundaries of the board.
	boardWidth := state.Board.Width
	boardHeight := state.Board.Height

    if myHead.X == boardWidth - 1{
        possibleMoves["right"] = false
    }
    if myHead.X == 0 {
        possibleMoves["left"] = false
    }
    if myHead.Y == boardHeight - 1 {
        possibleMoves["up"] = false
    }
    if myHead.Y == 0 {
        possibleMoves["down"] = false
    }
    

	// TODO: Step 2 - Don't hit yourself.
	// Use information in GameState to prevent your Battlesnake from colliding with itself.
    myBody := state.You.Body



	// TODO: Step 3 - Don't collide with others.
	// Use information in GameState to prevent your Battlesnake from colliding with others.

	// TODO: Step 4 - Find food.
	// Use information in GameState to seek out and find food.

	// Finally, choose a move from the available safe moves.
	// TODO: Step 5 - Select a move to make based on strategy, rather than random.
	var nextMove string

	safeMoves := []string{}
	for move, isSafe := range possibleMoves {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}

	if len(safeMoves) == 0 {
		nextMove = "down"
		log.Printf("%s MOVE %d: No safe moves detected! Moving %s\n", state.Game.ID, state.Turn, nextMove)
	} else {
		nextMove = safeMoves[rand.Intn(len(safeMoves))]
		for isSelfColliding(myBody, nextMove) == false{
		    nextMove = safeMoves[rand.Intn(len(safeMoves))]
        }
        log.Printf("%s MOVE %d: %s\n", state.Game.ID, state.Turn, nextMove)
	}
	return BattlesnakeMoveResponse{
		Move: nextMove,
	}
}

func isSelfColliding(body []Coord, move string) bool {
    currentHead := body[0]
    futureHead := currentHead
    move_direction := map[string]int{"left": -1, "right": 1, "up": 1, "down": -1}
    
    if exists([]string{"left", "right"}, move){
        futureHead.X = currentHead.X + move_direction[move]
    }else if exists([]string{"up", "down"}, move){
        futureHead.Y = currentHead.Y + move_direction[move]
    }
    //return future_head
    if map_exists(body, futureHead){
        return false
    }
    return true
}

