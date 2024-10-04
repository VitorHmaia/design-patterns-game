package observer

import (
    "fmt"
    "game/factory"
)

// Observer is the interface for objects that observe changes in the game.
type Observer interface {
    Update(monster factory.Monster)
}

// GameObserver observes changes in monsters' health.
type GameObserver struct{}

func (g *GameObserver) Update(monster factory.Monster) {
    fmt.Printf("%s now has %d health left.\n", monster.Name(), monster.Health())
}
