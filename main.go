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

// Function to select action with validation
func selectAction() string {
    var action string
    for {
        fmt.Println("Escolha sua ação (atacar, defender, habilidade):")
        fmt.Scan(&action)
        
        if action == "atacar" || action == "defender" || action == "habilidade" {
            return action
        }
        fmt.Println("Ação inválida. Por favor, escolha uma ação válida.")
    }
}

// Function to select an opponent in PvP with validation
func selectOpponent(players []Player, currentPlayerIndex int) *Player {
    var opponentIndex int
    for {
        fmt.Println("Escolha um oponente para atacar:")
        for i, player := range players {
            if i != currentPlayerIndex && player.IsAlive {
                fmt.Printf("%d - %s (Vida: %d)\n", i, player.Name, player.Monster.Health())
            }
        }

        fmt.Scan(&opponentIndex)
        if opponentIndex >= 0 && opponentIndex < len(players) && opponentIndex != currentPlayerIndex && players[opponentIndex].IsAlive {
            return &players[opponentIndex]
        }
        fmt.Println("Escolha inválida, tente novamente.")
    }
}

// Function to execute turn
func executeTurn(player *Player, opponent *Player, action string) {
    gameObserver := &observer.GameObserver{}

    attackStrategy := &strategy.AttackStrategy{}
    defendStrategy := &strategy.DefendStrategy{}
    specialStrategy := &strategy.SpecialAbilityStrategy{}

    // Separador visual
    fmt.Println("\n-------------------------------")
    fmt.Printf("Ação de %s: %s\n", player.Name, action)

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

    // Exibe o status após a ação
    fmt.Printf("%s agora tem %d de vida.\n", opponent.Name, opponent.Monster.Health())

    // Verifica se o oponente perdeu
    if opponent.Monster.Health() <= 0 {
        fmt.Printf("%s foi derrotado!\n", opponent.Name)
        opponent.IsAlive = false
    }

    // Mais um separador visual
    fmt.Println("-------------------------------\n")
    
    // Pause para permitir a leitura
    time.Sleep(500 * time.Millisecond) // Espera meio segundo
}

// Bot action function with improved intelligence
func botAction(bot *Player, player *Player) string {
    // If the bot's health is below 30, it has a higher chance of defending or using abilities
    if bot.Monster.Health() < 30 {
        if rand.Float32() < 0.5 {
            return "defender"
        } else {
            return "habilidade"
        }
    }

    // If player's health is low, bot aggressively attacks
    if player.Monster.Health() < 50 {
        return "atacar"
    }

    // Randomized but with a more intelligent approach based on health levels
    actions := []string{"atacar", "defender", "habilidade"}
    return actions[rand.Intn(len(actions))]
}

func main() {
    rand.Seed(time.Now().UnixNano())

    var mode string
    // Loop de validação para modo de jogo
    for {
        fmt.Println("Escolha o modo de jogo: PvP ou PvE")
        fmt.Scan(&mode)

        if mode == "PvP" || mode == "PvE" {
            break
        } else {
            fmt.Println("Modo inválido, escolha PvP ou PvE.")
        }
    }

    if mode == "PvP" {
        // Configura jogadores
        var numPlayers int
        fmt.Println("Quantos jogadores?")
        fmt.Scan(&numPlayers)

        players := make([]Player, numPlayers)

        for i := 0; i < numPlayers; i++ {
            var name, monsterType string
            fmt.Printf("Nome do jogador %d: ", i+1)
            fmt.Scan(&name)
            fmt.Printf("Escolha seu monstro (Dragon, Zombie): ")
            fmt.Scan(&monsterType)

            // Tratamento de erro ao criar o monstro
            monster, err := factory.MonsterFactory(monsterType)
            if err != nil {
                fmt.Println("Tipo de monstro inválido. Tente novamente.")
                i--
                continue
            }

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
                        fmt.Printf("%s venceu a batalha!\n", player.Name)
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
    } else if mode == "PvE" {
        // Configura PvE
        var playerName, monsterType string
        fmt.Println("Digite seu nome:")
        fmt.Scan(&playerName)
        fmt.Println("Escolha seu monstro (Dragon, Zombie):")
        fmt.Scan(&monsterType)

        // Tratamento de erro ao criar o monstro
        playerMonster, err := factory.MonsterFactory(monsterType)
        if err != nil {
            fmt.Println("Tipo de monstro inválido. Encerrando o jogo.")
            return
        }

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
                fmt.Printf("%s venceu a batalha contra o Bot!\n", player.Name)
                gameManager.AddScore(100)
                return
            }

            // Turno do bot
            fmt.Printf("Turno do %s!\n", bot.Name)
            botAct := botAction(&bot, &player)
            executeTurn(&bot, &player, botAct)

            if player.Monster.Health() <= 0 {
                fmt.Println("O Bot venceu a batalha!")
                return
            }

            round++
        }
    }
}