package logic

import (
//	"fmt"
)

type ColorMode int

const (
	COLOR_DIAMOND ColorMode = 1 //方块
	COLOR_HEART   ColorMode = 2 //红桃
	COLOR_CLUB    ColorMode = 3 //梅花
	COLOR_SPADE   ColorMode = 4 //黑桃
)

type CardSeq int

const (
	CARD_3          CardSeq = 3
	CARD_4          CardSeq = 4
	CARD_5          CardSeq = 5
	CARD_6          CardSeq = 6
	CARD_7          CardSeq = 7
	CARD_8          CardSeq = 8
	CARD_9          CardSeq = 9
	CARD_10         CardSeq = 10
	CARD_J          CardSeq = 11
	CARD_Q          CardSeq = 12
	CARD_K          CardSeq = 13
	CARD_A          CardSeq = 14
	CARD_2          CardSeq = 15
	CARD_KING       CardSeq = 21
	CARD_SUPER_KING CardSeq = 22
)

const CARD_MASK int = 100

func GetCardSeq(card int) int {
	return card / CARD_MASK
}

func GetColorMode(card int) int {
	return card % CARD_MASK
}

const MAX_DECK_CARDS int = 54 //一副牌最大数量

type CardDeck struct {
	Nums  int //几副牌
	Cards []int
}

func NewCardDeck(nums int) *CardDeck {
	deckNums := 1
	if nums == 2 {
		deckNums = 2
	}
	deck := &CardDeck{
		Nums:  deckNums,
		Cards: make([]int, MAX_DECK_CARDS*deckNums),
	}
	deck.Init()
	return deck
}

func (ck *CardDeck) Init() {
	index := 0
	for d := 1; d <= ck.Nums; d++ {
		for i := CARD_3; i <= CARD_2; i++ {
			for j := COLOR_DIAMOND; j <= COLOR_SPADE; j++ {
				cardId := i*CARD_MASK + j
				ck.Cards[index] = cardId
				index++
			}
		}
	}
	for d := 1; d < ck.Nums; d++ {
		ck.Cards[index] = CARD_KING * CARD_MASK
		index++
		ck.Cards[index] = CARD_SUPER_KING * CARD_MASK
		index++
	}
}

func (ck *CardDeck) Shuffle() {
	RandShuffle(ck.Cards, 0, len(ck.Cards))
}
