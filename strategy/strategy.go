package strategy

import (
    "fmt"
    "game/factory"
)

// ActionStrategy is the interface for different combat actions.
type ActionStrategy interface {
    Execute(attacker, defender factory.Monster)
}

// AttackStrategy handles the attack action.
type AttackStrategy struct{}

func (a *AttackStrategy) Execute(attacker, defender factory.Monster) {
    fmt.Printf("%s attacks %s!\n", attacker.Name(), defender.Name())
    defender.SetHealth(defender.Health() - attacker.AttackPower())
}

// DefendStrategy handles the defense action.
type DefendStrategy struct{}

func (d *DefendStrategy) Execute(attacker, defender factory.Monster) {
    fmt.Printf("%s defends!\n", attacker.Name())
    // Implement defense logic (e.g., reduce damage taken next turn).
}

// SpecialAbilityStrategy handles using the monster's special ability.
type SpecialAbilityStrategy struct{}

func (s *SpecialAbilityStrategy) Execute(attacker, defender factory.Monster) {
    attacker.SpecialAbility()
}
