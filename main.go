package main

import (
    "game/factory"
    "game/strategy"
    "game/observer"
    "game/singleton"
    "fmt"
    "math/rand"
    "time"
)

// Player structure
type Player struct {
    Name    string
    Monster factory.Monster
    IsAlive bool
}

// Function to select action
func selectAction() string {
    var actionNumber int
    fmt.Println("Escolha sua ação:")
    fmt.Println("1 - Atacar")
    fmt.Println("2 - Defender")
    fmt.Println("3 - Habilidade")

    fmt.Scan(&actionNumber)
    switch actionNumber {
    case 1:
        return "atacar"
    case 2:
        return "defender"
    case 3:
        return "habilidade"
    default:
        fmt.Println("Ação inválida, escolha novamente.")
        return selectAction()
    }
}

// Function to select an opponent in PvP
func selectOpponent(players []Player, currentPlayerIndex int) *Player {
    fmt.Println("Escolha um oponente para atacar:")
    for i, player := range players {
        if i != currentPlayerIndex && player.IsAlive {
            fmt.Printf("%d - %s (Vida: %d)\n", i, player.Name, player.Monster.Health())
        }
    }

    var opponentIndex int
    fmt.Scan(&opponentIndex)

    if opponentIndex == currentPlayerIndex || !players[opponentIndex].IsAlive {
        fmt.Println("Escolha inválida, tente novamente.")
        return selectOpponent(players, currentPlayerIndex)
    }

    return &players[opponentIndex]
}

// Function to execute turn
func executeTurn(player *Player, opponent *Player, action string) {
    gameObserver := &observer.GameObserver{}

    attackStrategy := &strategy.AttackStrategy{}
    defendStrategy := &strategy.DefendStrategy{}
    specialStrategy := &strategy.SpecialAbilityStrategy{}

    switch action {
    case "atacar":
        attackStrategy.Execute(player.Monster, opponent.Monster)
    case "defender":
        defendStrategy.Execute(player.Monster, opponent.Monster)
    case "habilidade":
        specialStrategy.Execute(player.Monster, opponent.Monster)
    default:
        fmt.Println("Ação inválida.")
    }

    // Notificar mudança de estado
    gameObserver.Update(opponent.Monster)

    // Verifica se o oponente perdeu
    if opponent.Monster.Health() <= 0 {
        fmt.Printf("%s foi derrotado!\n", opponent.Name)
        opponent.IsAlive = false
    }
}

// Function for bot action in PvE
func botAction() string {
    actions := []string{"atacar", "defender", "habilidade"}
    return actions[rand.Intn(len(actions))]
}

// Function to display victory message
func displayVictoryMessage(playerName string) {
    fmt.Printf("Parabéns, %s! Você demonstrou uma força impressionante e dominou a batalha!\n", playerName)
    fmt.Println("Você é o grande campeão!")
}

func main() {
    rand.Seed(time.Now().UnixNano())

    var mode int
    fmt.Println("Escolha o modo de jogo:")
    fmt.Println("1 - PvP")
    fmt.Println("2 - PvE")
    fmt.Scan(&mode)

    if mode == 1 {
        // Configura jogadores
        var numPlayers int
        fmt.Println("Quantos jogadores?")
        fmt.Scan(&numPlayers)

        players := make([]Player, numPlayers)

        for i := 0; i < numPlayers; i++ {
            var name, monsterType string
            fmt.Printf("Nome do jogador %d: ", i+1)
            fmt.Scan(&name)
            fmt.Println("Escolha seu monstro:")
            fmt.Println("1 - Dragon")
            fmt.Println("2 - Zombie")
            
            var monsterChoice int
            fmt.Scan(&monsterChoice)

            if monsterChoice == 1 {
                monsterType = "Dragon"
            } else {
                monsterType = "Zombie"
            }

            monster, _ := factory.MonsterFactory(monsterType)
            players[i] = Player{Name: name, Monster: monster, IsAlive: true}
        }

        fmt.Println("Iniciando batalha PvP!")

        // Sistema de turnos
        gameManager := singleton.GetInstance()
        round := 1

        for {
            fmt.Printf("\n--- Turno %d ---\n", round)

            // Verifica se ainda existem jogadores vivos
            alivePlayers := 0
            for _, player := range players {
                if player.IsAlive {
                    alivePlayers++
                }
            }
            if alivePlayers <= 1 {
                for _, player := range players {
                    if player.IsAlive {
                        displayVictoryMessage(player.Name)
                        gameManager.AddScore(100)
                        return
                    }
                }
            }

            // Turnos dos jogadores vivos
            for i := 0; i < numPlayers; i++ {
                if players[i].IsAlive {
                    fmt.Printf("%s, é o seu turno!\n", players[i].Name)

                    action := selectAction()
                    opponent := selectOpponent(players, i)

                    executeTurn(&players[i], opponent, action)

                    // Se o oponente foi derrotado, ele sai da lista de turnos
                    if opponent.Monster.Health() <= 0 {
                        fmt.Printf("%s foi derrotado!\n", opponent.Name)
                    }
                }
            }

            round++
        }
    } else if mode == 2 {
        // Configura PvE
        var playerName, monsterType string
        fmt.Println("Digite seu nome:")
        fmt.Scan(&playerName)
        
        fmt.Println("Escolha seu monstro:")
        fmt.Println("1 - Dragon")
        fmt.Println("2 - Zombie")
        
        var monsterChoice int
        fmt.Scan(&monsterChoice)
        
        if monsterChoice == 1 {
            monsterType = "Dragon"
        } else {
            monsterType = "Zombie"
        }

        playerMonster, _ := factory.MonsterFactory(monsterType)
        player := Player{Name: playerName, Monster: playerMonster, IsAlive: true}

        // Configura o bot
        botMonster, _ := factory.MonsterFactory("Zombie") // O bot sempre usa um zumbi neste exemplo
        bot := Player{Name: "Bot", Monster: botMonster, IsAlive: true}

        fmt.Println("Iniciando batalha PvE!")

        // Sistema de turnos
        gameManager := singleton.GetInstance()
        round := 1

        for {
            fmt.Printf("\n--- Turno %d ---\n", round)

            // Exibe a vida de ambos
            fmt.Printf("Jogador: %s (Vida: %d)\n", player.Name, player.Monster.Health())
            fmt.Printf("Bot: %s (Vida: %d)\n", bot.Name, bot.Monster.Health())

            // Turno do jogador
            fmt.Printf("%s, é o seu turno!\n", player.Name)
            action := selectAction()
            executeTurn(&player, &bot, action)

            if bot.Monster.Health() <= 0 {
                displayVictoryMessage(player.Name)
                gameManager.AddScore(100)
                return
            }

            // Turno do bot
            fmt.Printf("Turno do %s!\n", bot.Name)
            botAct := botAction()
            executeTurn(&bot, &player, botAct)

            if player.Monster.Health() <= 0 {
                fmt.Println("O Bot venceu a batalha! Tente novamente para melhorar sua estratégia.")
                return
            }

            round++
        }
    } else {
        fmt.Println("Modo inválido, escolha 1 para PvP ou 2 para PvE.")
    }
}
