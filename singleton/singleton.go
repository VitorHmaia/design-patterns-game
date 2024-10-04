package singleton

// GameManager is a singleton that manages the game state.
type GameManager struct {
    playerScore int
}

var instance *GameManager

// GetInstance returns the singleton instance of GameManager.
func GetInstance() *GameManager {
    if instance == nil {
        instance = &GameManager{}
    }
    return instance
}

func (g *GameManager) AddScore(points int) {
    g.playerScore += points
}

func (g *GameManager) PlayerScore() int {
    return g.playerScore
}