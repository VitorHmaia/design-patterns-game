# 🐲 Batalha de Monstros

Este projeto implementa um jogo de batalha por turnos entre monstros, jogado no terminal. O jogador pode escolher entre diferentes tipos de monstros, cada um com habilidades específicas, e pode jogar contra outro jogador (PvP) ou contra um bot (PvE). O objetivo é derrotar os oponentes por meio de ataques, defesas e habilidades especiais.

## ⚙️ Requisitos

- **Go**: Certifique-se de ter o [Go](https://golang.org/dl/) instalado em sua máquina (versão 1.20+ recomendada).
- **Terminal**: Este jogo roda diretamente no terminal, sem interface gráfica.
  
## Funcionalidades Principais

1. **Modos de Jogo**:
   - 🆚 **PvP (Jogador vs Jogador)**: Dois ou mais jogadores competem em uma batalha por turnos, podendo atacar uns aos outros. O jogador escolhe quem atacar durante o turno.
   - 🤖 **PvE (Jogador vs Computador)**: Um jogador enfrenta um bot controlado por IA em uma batalha por turnos.

2. **Criação de Monstros**:
   - O jogador pode escolher entre diferentes tipos de monstros (ex.: dragões, zumbis), cada um com atributos como pontos de vida, ataque e defesa, além de habilidades especiais.

3. **Sistema de Combate por Turnos**:
   - Cada jogador, durante seu turno, pode escolher atacar, defender ou usar uma habilidade especial. O turno alterna entre os jogadores.

4. **Habilidades Especiais**:
   - Cada tipo de monstro tem uma habilidade especial exclusiva que pode ser usada no combate.

5. **Sistema de Pontuação**:
   - Pontuação é atribuída aos vencedores, e o progresso do jogo pode ser salvo e carregado futuramente.

## 📐 Como Cada Arquivo se Comunica

### `main.go`

Este é o ponto de entrada do jogo. Ele gerencia a lógica principal do jogo, incluindo a seleção de modos (PvP ou PvE), gerenciamento de turnos, e interação do jogador com o jogo.

1. O jogador pode escolher entre jogar no modo PvP ou PvE.
2. Baseado na escolha do jogador, monstros são criados utilizando a **Factory** (`factory.MonsterFactory`).
3. Durante cada turno, o jogador escolhe sua ação (atacar, defender ou usar habilidade especial). Dependendo da ação:
   - A estratégia correspondente é selecionada da pasta **strategy** e aplicada ao monstro.
   - As mudanças no estado do jogo (como redução de vida) são notificadas usando o padrão **Observer**.
4. O **Singleton** `GameManager` é utilizado para gerenciar a pontuação global do jogo e fornecer outras funcionalidades globais, como salvar progresso.

### `factory/factory.go`

A pasta `factory` implementa o padrão **Factory**, responsável pela criação de diferentes tipos de monstros. Ela contém:
- A interface `Monster`, que define os métodos que cada monstro deve implementar (vida, ataque, defesa, etc.).
- Implementações específicas de monstros, como `Dragon` e `Zombie`, que definem atributos e comportamentos diferentes para cada monstro.

### `strategy/*`

A pasta `strategy` define o comportamento de cada ação possível no jogo, utilizando o padrão **Strategy**. Os arquivos são:
- `attack_strategy.go`: Define a lógica de ataque.
- `defend_strategy.go`: Define a lógica de defesa.
- `special_ability.go`: Define a lógica para habilidades especiais dos monstros.

Cada uma dessas estratégias pode ser escolhida durante o turno de um jogador e aplicada ao monstro inimigo.

### `observer/observer.go`

O padrão **Observer** é utilizado para notificar quando há mudanças no estado de um monstro (por exemplo, quando sua vida é reduzida). O arquivo `observer.go` contém a lógica de notificação, que pode ser utilizada para exibir mensagens para o jogador.

### `singleton/game_manager.go`

Este arquivo implementa o padrão **Singleton**, garantindo que exista apenas uma instância do `GameManager` em todo o jogo. O `GameManager` é responsável por:
- Gerenciar a pontuação dos jogadores.
- Armazenar o progresso do jogo.
- Fornecer acesso global a funções de gerenciamento do jogo.

## 🚀 Como Executar o Jogo

1. Certifique-se de ter o [Go](https://golang.org/dl/) instalado em sua máquina.
2. Clone este repositório e navegue até o diretório do projeto.
3. No terminal, execute o seguinte comando para rodar o jogo:

```bash
go run main.go
```
Escolha o modo de jogo: PvP ou PvE
PvP
Quantos jogadores?
2
Nome do jogador 1: Alice
Escolha seu monstro (Dragon, Zombie): Dragon
Nome do jogador 2: Bob
Escolha seu monstro (Dragon, Zombie): Zombie

Iniciando batalha PvP!

--- Turno 1 ---
Alice, é o seu turno!
Escolha sua ação (atacar, defender, habilidade): atacar
Escolha um oponente para atacar:
1 - Bob (Vida: 100)
Bob foi atacado! Vida restante: 70

Bob, é o seu turno!
Escolha sua ação (atacar, defender, habilidade): habilidade
Bob usou sua habilidade especial! Vida restaurada: 90

## Conclusão
Este jogo utiliza vários padrões de design para implementar uma lógica modular e extensível. Os padrões Factory, Strategy, Observer, Singleton e Memento são usados para criar e gerenciar monstros, implementar estratégias de combate, notificar mudanças no jogo e gerenciar o estado global. A estrutura modular facilita a adição de novos monstros, habilidades e modos de jogo.