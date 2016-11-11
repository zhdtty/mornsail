package logic

import (
	"fmt"
	"sort"
)

type PlayType int

const (
	PLAYTYPE_SINGLE                               PlayType = 1  //单张
	PLAYTYPE_PAIR                                 PlayType = 2  //一对
	PLAYTYPE_TRIPLE                               PlayType = 3  //三条
	PLAYTYPE_TRIPLE_WITH_PAIR                     PlayType = 4  //条带对
	PLAYTYPE_SEQUENCE                             PlayType = 5  //顺子
	PLAYTYPE_SEQUENCE_PAIRS                       PlayType = 6  //连对
	PLAYTYPE_SEQUENCE_TRIPLES                     PlayType = 7  //连条
	PLAYTYPE_SEQUENCE_TRIPLES_WITH_PAIRS          PlayType = 8  //连条带对
	PLAYTYPE_SEQUENCE_TRIPLES_WITH_SEQUENCE_PAIRS PlayType = 9  //连条带连对
	PLAYTYPE_BOMB                                 PlayType = 10 //炸弹
	PLAYTYPE_ROCKET                               PlayType = 11 //王炸
)

type IdentityType int

const (
	IDENTITY_LANDLORD IdentityType = 1 //地主
	IDENTITY_PEASANT  IdentityType = 2 //农民
)

type CardPlayer struct {
	RoleId   int64        //角色ID
	identity IdentityType //身份
	Cards    []int        //当前手牌
}
type CardLog struct {
	Cards        []int //本轮出的牌
	Seq          int   //出牌轮次
	Ts           int   //出牌时间（非时间戳，牌局时间线上的节点）
	BelongRoleId int64 //出牌者
}
type LandLord struct {
	Players    map[int64]*CardPlayer //组队玩家
	CardDeck   *CardDeck             //牌组
	CardStates map[int]int           //0:未打出 1:已打出
	CardLogs   []CardLog             //整个牌局纪录
}

func NewLandLord() *LandLord {
	rule := &LandLord{
		Players:    make(map[int64]*CardPlayer),
		CardDeck:   NewCardDeck(2),
		CardStates: make(map[int]int),
	}
	return rule
}
func (rule *LandLord) Init() {
}

func (rule *LandLord) Deal() { //分牌
	rule.CardDeck.Shuffle()
}

func (rule *LandLord) AddPlayer(roleId int64) { //加入

}

func (rule *LandLord) Start(roleId int64) { //开始

}

func (rule *LandLord) CallPoint(roleId int64, point int) { //叫分

}

func (rule *LandLord) Disband() { //解散

}

func (rule *LandLord) Play(roleId int64, cards []int) { //出牌

}

func (rule *LandLord) Pass(roleId int64) { //不出

}

func (rule *LandLord) CheckResult() { //检查结果

}

func (rule *LandLord) checkPlayType(cards []int) (int, int, error) { //检查出牌类型，返回（playType，首牌号，错误）
	len := len(cards)

	switch len {
	case 0:
		return 0, 0, fmt.Errorf("No any cards")
	case 1: //单张
		return PLAYTYPE_SINGLE, cards[1], nil
	case 2: //一对
		{
			cardId1 := GetCardSeq(cards[1])
			cardId2 := GetCardSeq(cards[2])
			if cardId1 != cardId2 {
				return 0, 0, fmt.Errorf("Pair not same")
			}
			return PLAYTYPE_PAIR, cards[1], nil
		}
	case 3: //三条
		cardId1 := GetCardSeq(cards[1])
		cardId2 := GetCardSeq(cards[2])
		cardId3 := GetCardSeq(cards[3])
		if cardId1 != cardId2 || cardId1 != cardId3 {
			return 0, 0, fmt.Errorf("Triple not same")
		}
		return PLAYTYPE_TRIPLE, cards[1], nil
	}
	if len >= 4 {
		groups := make(map[int]int)
		for _, v := range cards {
			cardId := GetCardSeq(v)
			if num, ok := groups[cardId]; ok {
				num++
				groups[cardId] = num
			} else {
				groups[cardId] = 1
			}
		}

		groupLen := len(groups)
		if groupLen == 1 {
			return PLAYTYPE_BOMB, cards[1], nil
		}
		//获取牌号并排序
		MASK := 1000 //创建一个系数，保证主对在后，副对在前('条带对'好检测)
		cardSeqs := make([]int, groupLen)
		index := 0
		hasSingle := false //是否存在单张
		hasPair := false   //是否存在一对
		for k, v := range groups {
			if v == 1 {
				hasSingle = true
			} else if v == 2 {
				hasPair = true
			}
			cardSeqs[index] = v*MASK + k
			index++
		}
		s := sort.IntSlice(cardSeqs)
		s.Sort()

		rate := len / groupLen
		mod := len % groupLen
		if mod > 0 {
			if (len*groupLen)%10 != 0 {
				return 0, 0, fmt.Errorf("Invalid card group")
			}
			if groupLen%2 != 0 {
				return 0, 0, fmt.Errorf("Invalid card group, triples nums not same as pairs nums")
			}
			if groupLen == 2 { //条带对
				return PLAYTYPE_TRIPLE_WITH_PAIR, cardSeqs[groupLen-1] % MASK, nil
			}
			pairFirst := cardSeqs[0] % MASK
			pairLast := cardSeqs[groupLen/2-1] % MASK
			tripleFirst := cardSeqs[groupLen/2] % MASK
			tripleLast := cardSeqs[groupLen-1] % MASK
			if tripleFirst < CARD_3 || tripleFirst > CARD_A ||
				tripleLast < CARD_3 || tripleLast > CARD_A ||
				tripleLast-tripleFirst+1 > groupLen/2 {
				return 0, 0, fmt.Errorf("Invalid sequence triples with pairs, not triples sequence")
			}
			//若发过来的牌是333444555668899，变成333444555666899过滤掉
			if hasSingle == true {
				return 0, 0, fmt.Errorf("Invalid sequence triples with pairs, has single card")
			}
			if pairFirst < CARD_3 || pairFirst > CARD_A ||
				pairLast < CARD_3 || pairLast > CARD_A ||
				pairLast-pairFirst+1 > groupLen/2 { //连条带对
				return PLAYTYPE_SEQUENCE_TRIPLES_WITH_PAIRS, tripleFirst, nil
			}
			return PLAYTYPE_SEQUENCE_TRIPLES_WITH_SEQUENCE_PAIRS, tripleFirst, nil //连条带连对（飞机）
		}
		switch rate { //如果整除
		case 1: //顺子
			if groupLen < 5 {
				return 0, 0, fmt.Errorf("Invalid sequence, must least 5")
			}
			first := cardSeqs[0] % MASK
			last := cardSeqs[groupLen-1] % MASK
			if first < CARD_3 || first > CARD_A ||
				last < CARD_3 || last > CARD_A ||
				last-first+1 > groupLen {
				return 0, 0, fmt.Errorf("Invalid sequence, not sequence")
			}
			return PLAYTYPE_SEQUENCE, first, nil
		case 2: //连对或天炸
			if hasSingle == true { //若发过来的牌是333455，过滤掉
				return 0, 0, fmt.Errorf("Invalid sequence pairs, has single card")
			}
			if groupLen < 3 {
				return 0, 0, fmt.Errorf("Invalid sequence pairs, must least 3")
			}
			first := cardSeqs[0] % MASK
			last := cardSeqs[groupLen-1] % MASK
			if first == CARD_KING && last == CARD_SUPER_KING {
				return PLAYTYPE_ROCKET, 0, nil
			}
			if first < CARD_3 || first > CARD_A ||
				last < CARD_3 || last > CARD_A ||
				last-first+1 > groupLen {
				return 0, 0, fmt.Errorf("Invalid sequence, not sequence pairs")
			}
			return PLAYTYPE_SEQUENCE_PAIRS, first, nil
		case 3: //连条 333444555
			//若发过来的牌是333445555，过滤掉
			//若发过来的牌是333455555，过滤掉
			if hasSingle == true || hasPair == true {
				return 0, 0, fmt.Errorf("Invalid sequence pairs, has single card or pair cards")
			}
			if groupLen < 2 {
				return 0, 0, fmt.Errorf("Invalid sequence triples, must least 2")
			}
			first := cardSeqs[0] % MASK
			last := cardSeqs[groupLen-1] % MASK
			if first < CARD_3 || first > CARD_A ||
				last < CARD_3 || last > CARD_A ||
				last-first+1 > groupLen {
				return 0, 0, fmt.Errorf("Invalid sequence, not sequence triples")
			}
			return PLAYTYPE_SEQUENCE_TRIPLES, first, nil
		}
	}
}
