package main

import (
	"fmt"
	"log"

	"github.com/mattn/go-tty"
)

type Character struct {
	hp    int
	maxHp int
	mp    int
	maxMp int
	name  string
	aa    string
}

const (
	MONSTER_PLAYER int = iota
	MONSTER_SLIME
	MONSTER_MAX
)

var monsters [MONSTER_MAX]Character

const (
	CHARACTER_PLAYER int = iota
	CHARACTER_MONSTER
	CHARACTER_MAX
)

var characters [CHARACTER_MAX]Character

func main() {
	Init()
	Battle(MONSTER_SLIME)
}

func Init() {
	monsters[MONSTER_PLAYER] = Character{
		hp:    15,
		maxHp: 15,
		mp:    15,
		maxMp: 15,
		name:  "ゆうしゃ",
	}
	monsters[MONSTER_SLIME] = Character{
		hp:    3,
		maxHp: 3,
		mp:    0,
		maxMp: 0,
		name:  "スライム",
		aa: `／・Д・＼
〜〜〜〜〜`,
	}
	characters[CHARACTER_PLAYER] = monsters[MONSTER_PLAYER]
	fmt.Print()
}

func Battle(monster int) {
	characters[CHARACTER_MONSTER] = monsters[monster]
	DrawBattleScreen()

	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {
		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		// handle key event
		fmt.Println(r)
	}
}

func DrawBattleScreen() {
	fmt.Printf("%s\n", characters[CHARACTER_PLAYER].name)
	fmt.Printf("HP:%d/%d MP:%d/%d\n",
		characters[CHARACTER_PLAYER].hp,
		characters[CHARACTER_PLAYER].maxHp,
		characters[CHARACTER_PLAYER].mp,
		characters[CHARACTER_PLAYER].maxMp)
	fmt.Printf("\n")

	fmt.Printf("%s", characters[CHARACTER_MONSTER].aa)
	fmt.Printf(" (HP:%d/%d) \n",
		characters[CHARACTER_MONSTER].hp,
		characters[CHARACTER_MONSTER].maxHp)
	fmt.Printf("\n")
}
