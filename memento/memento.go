package memento

import "game/factory"

// GameState is the structure that stores the current state of the game.
type GameState struct {
    PlayerScore int
    Monsters    []factory.Monster
}

// Memento stores and restores the game state.
type Memento struct {
    savedState GameState
}

// SaveState saves the current game state.
func (m *Memento) SaveState(playerScore int, monsters []factory.Monster) {
    m.savedState = GameState{
        PlayerScore: playerScore,
        Monsters:    monsters,
    }
}

// RestoreState restores the saved game state.
func (m *Memento) RestoreState() GameState {
    return m.savedState
}
