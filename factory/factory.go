package factory

import "fmt"

// Monster is the interface that all monsters must implement.
type Monster interface {
    Name() string
    SetHealth(health int)
    AttackPower() int
    DefensePower() int
    Health() int
    SpecialAbility()
}

// Dragon represents a Dragon monster.
type Dragon struct {
    health int
}

func (d *Dragon) Name() string           { return "Dragon" }
func (d *Dragon) AttackPower() int        { return 50 }
func (d *Dragon) DefensePower() int       { return 30 }
func (d *Dragon) Health() int             { return d.health }
func (d *Dragon) SpecialAbility()         { fmt.Println("Dragon uses Fire Breath!") }
func (d *Dragon) SetHealth(h int)         { d.health = h }

// Zombie represents a Zombie monster.
type Zombie struct {
    health int
}

func (z *Zombie) Name() string            { return "Zombie" }
func (z *Zombie) AttackPower() int         { return 20 }
func (z *Zombie) DefensePower() int        { return 10 }
func (z *Zombie) Health() int              { return z.health }
func (z *Zombie) SpecialAbility()          { fmt.Println("Zombie uses Regenerate!") }
func (z *Zombie) SetHealth(h int)          { z.health = h }

// MonsterFactory is the Factory for creating monsters.
func MonsterFactory(monsterType string) (Monster, error) {
    switch monsterType {
    case "Dragon":
        return &Dragon{health: 100}, nil
    case "Zombie":
        return &Zombie{health: 150}, nil
    default:
        return nil, fmt.Errorf("Invalid monster type: %s", monsterType)
    }
}