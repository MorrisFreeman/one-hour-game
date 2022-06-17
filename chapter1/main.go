package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/mattn/go-tty"
)

type Character struct {
	hp      int
	maxHp   int
	mp      int
	maxMp   int
	name    string
	aa      string
	command int
	target  int
	attack  int
}

const (
	MONSTER_PLAYER int = iota
	MONSTER_SLIME
	MONSTER_BOSS
	MONSTER_MAX
)

var monsters [MONSTER_MAX]Character

const (
	CHARACTER_PLAYER int = iota
	CHARACTER_MONSTER
	CHARACTER_MAX
)

var characters [CHARACTER_MAX]Character

const (
	COMMAND_FIGHT int = iota
	COMMAND_SPEL
	COMMAND_RUN
	COMMAND_MAX
)

var commandNames [COMMAND_MAX]string = [COMMAND_MAX]string{
	"たたかう",
	"じゅもん",
	"にげる",
}

const SPELL_COST = 3

func main() {
	Init()
	Battle(MONSTER_BOSS)
}

func Init() {
	rand.Seed(time.Now().UnixNano())
	monsters[MONSTER_PLAYER] = Character{
		hp:     100,
		maxHp:  100,
		mp:     15,
		maxMp:  15,
		name:   "ゆうしゃ",
		attack: 30,
	}
	monsters[MONSTER_SLIME] = Character{
		hp:    3,
		maxHp: 3,
		mp:    0,
		maxMp: 0,
		name:  "スライム",
		aa: `／・Д・＼
〜〜〜〜〜`,
		attack: 2,
	}
	monsters[MONSTER_BOSS] = Character{
		hp:     255,
		maxHp:  255,
		mp:     0,
		maxMp:  0,
		attack: 50,
		name:   "まおう",
		aa: ` 　A＠A
ψ（▼皿▼）ψ
`,
	}
	characters[CHARACTER_PLAYER] = monsters[MONSTER_PLAYER]
	fmt.Print()
}

func Battle(monster int) {
	characters[CHARACTER_MONSTER] = monsters[monster]
	characters[CHARACTER_PLAYER].target = CHARACTER_MONSTER
	characters[CHARACTER_MONSTER].target = CHARACTER_PLAYER
	DrawBattleScreen()
	fmt.Printf("%sが 現れた\n", characters[CHARACTER_MONSTER].name)
	WaitInput()
	for {
		SelectCommand()
		for i := 0; i < CHARACTER_MAX; i++ {
			DrawBattleScreen()
			switch characters[i].command {
			case COMMAND_FIGHT:
				fmt.Printf("%sの こうげき！\n", characters[i].name)
				WaitInput()
				DrawBattleScreen()
				damage := 1 + rand.Int()%characters[i].attack
				characters[characters[i].target].hp -= damage
				if characters[characters[i].target].hp < 0 {
					characters[characters[i].target].hp = 0
				}
				fmt.Printf("%sに %dの ダメージ\n", characters[characters[i].target].name, damage)
				WaitInput()
				DrawBattleScreen()
			case COMMAND_SPEL:
				if characters[i].mp < SPELL_COST {
					fmt.Printf("MPが たりない！\n")
					WaitInput()
					DrawBattleScreen()
					break
				}
				fmt.Printf("%sは ヒールを となえた！\n", characters[i].name)
				WaitInput()
				DrawBattleScreen()
				fmt.Printf("%sのきずが かいふくした！\n", characters[i].name)
				characters[i].hp = characters[i].maxHp
				characters[i].mp -= SPELL_COST
				WaitInput()
				DrawBattleScreen()
			case COMMAND_RUN:
				fmt.Printf("%sは にげだした！\n", characters[i].name)
				return
			}
			if characters[characters[i].target].hp <= 0 {
				switch characters[i].target {
				case CHARACTER_PLAYER:
					fmt.Printf("%sは たおされてしまった...\n", characters[characters[i].target].name)
				case CHARACTER_MONSTER:
					characters[characters[i].target].aa = "\n"
					DrawBattleScreen()
					fmt.Printf("%sを たおした！\n", characters[characters[i].target].name)
				}
				WaitInput()
				return
			}
		}
	}
}

func DrawBattleScreen() {
	ClearScreen()
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

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func SelectCommand() {
	characters[CHARACTER_PLAYER].command = COMMAND_FIGHT
	for {
		DrawBattleScreen()
		for i := 0; i < COMMAND_MAX; i++ {
			if i == characters[CHARACTER_PLAYER].command {
				fmt.Print(">")
			} else {
				fmt.Print(" ")
			}
			fmt.Printf("%s\n", commandNames[i])
		}
		r := WaitInput()
		switch r {
		case "w":
			characters[CHARACTER_PLAYER].command--
		case "s":
			characters[CHARACTER_PLAYER].command++
		default:
			return
		}
		characters[CHARACTER_PLAYER].command = (COMMAND_MAX + characters[CHARACTER_PLAYER].command) % COMMAND_MAX
	}
}

func WaitInput() string {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()
	r, err := tty.ReadRune()
	if err != nil {
		log.Fatal(err)
	}
	return string(r)
}
