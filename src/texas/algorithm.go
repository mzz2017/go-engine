package texas

import (
	"sort"
	"strconv"
	"strings"
)

func StrToByteValue(str string) int8 {

	if str == "A" {
		return PokeValue_A
	} else if str == "K" {
		return PokeValue_K
	} else if str == "Q" {
		return PokeValue_Q
	} else if str == "J" {
		return PokeValue_J
	} else {
		ret, _ := strconv.ParseInt(str, 10, 8)
		return int8(ret)
	}
}

func StrToBytes(str string) []int8 {
	var ret []int8

	if len(str) == 0 {
		return ret
	}

	strs := strings.Split(str, ",")
	for _, s := range strs {
		ret = append(ret, StrToByte(s))
	}

	return ret
}

func StrToByte(str string) int8 {
	if strings.HasPrefix(str, "方") {
		p := Poke{PokeColor_FANG, StrToByteValue(string([]rune(str)[1:]))}
		return p.ToByte()
	} else if strings.HasPrefix(str, "梅") {
		p := Poke{PokeColor_MEI, StrToByteValue(string([]rune(str)[1:]))}
		return p.ToByte()
	} else if strings.HasPrefix(str, "红") {
		p := Poke{PokeColor_HONG, StrToByteValue(string([]rune(str)[1:]))}
		return p.ToByte()
	} else if strings.HasPrefix(str, "黑") {
		p := Poke{PokeColor_HEI, StrToByteValue(string([]rune(str)[1:]))}
		return p.ToByte()
	} else if strings.HasPrefix(str, "鬼") {
		p := Poke{PokeColor_GUI, PokeValue_GUI}
		return p.ToByte()
	}

	return 0
}

func BytesToStr(bytes []int8) string {
	return KeyToStr(GenCardBind(bytes))
}

func GenCardBind(bytes []int8) int64 {
	var ret int64

	for _, b := range bytes {
		ret = ret*100 + int64(b)
	}

	return ret
}

func KeyToPoke(k int64) []Poke {

	var cs []Poke

	if k > 1000000000000 {
		cs = append(cs, NewPoke((int8)(k%100000000000000/1000000000000)))
	}
	if k > 10000000000 {
		cs = append(cs, NewPoke((int8)(k%1000000000000/10000000000)))
	}
	if k > 100000000 {
		cs = append(cs, NewPoke((int8)(k%10000000000/100000000)))
	}
	if k > 1000000 {
		cs = append(cs, NewPoke((int8)(k%100000000/1000000)))
	}
	if k > 10000 {
		cs = append(cs, NewPoke((int8)(k%1000000/10000)))
	}
	if k > 100 {
		cs = append(cs, NewPoke((int8)(k%10000/100)))
	}
	if k > 1 {
		cs = append(cs, NewPoke((int8)(k%100/1)))
	}
	return cs
}

func KeyToBytes(k int64) []int8 {

	var cs []int8

	if k > 1000000000000 {
		cs = append(cs, (int8)(k%100000000000000/1000000000000))
	}
	if k > 10000000000 {
		cs = append(cs, (int8)(k%1000000000000/10000000000))
	}
	if k > 100000000 {
		cs = append(cs, (int8)(k%10000000000/100000000))
	}
	if k > 1000000 {
		cs = append(cs, (int8)(k%100000000/1000000))
	}
	if k > 10000 {
		cs = append(cs, (int8)(k%1000000/10000))
	}
	if k > 100 {
		cs = append(cs, (int8)(k%10000/100))
	}
	if k > 1 {
		cs = append(cs, (int8)(k%100/1))
	}
	return cs
}

func KeyToStr(key int64) string {
	var ret string

	cs := KeyToPoke(key)
	for _, cs := range cs {
		ret += cs.ToString()
	}

	return ret
}

func IsGui(i int8) bool {

	color := (int8)(i >> 4)
	value := (int8)(i % 16)

	return value == PokeValue_GUI && color == PokeColor_GUI
}

func GetKeyData(bytes []int8) *KeyData {
	key := GenCardBind(bytes)

	return GetKeyDataByKey(key)
}

func GetKeyDataByKey(key int64) *KeyData {
	colorKey := ChangeColor(key)
	color, okcolor := colorMap[colorKey]

	normalKey := RemoveColor(key)
	normal, oknormal := normalMap[normalKey]
	if !okcolor {
		return &normal
	}
	if !oknormal {
		return &color
	}
	if color.index > normal.index {
		return &color
	} else {
		return &normal
	}
}

func RemoveColor(k int64) int64 {
	var cs []int8
	if k > 1000000000000 {
		cs = append(cs, (int8)(k%100000000000000/1000000000000))
	}
	if k > 10000000000 {
		cs = append(cs, (int8)(k%1000000000000/10000000000))
	}
	if k > 100000000 {
		cs = append(cs, (int8)(k%10000000000/100000000))
	}
	if k > 1000000 {
		cs = append(cs, (int8)(k%100000000/1000000))
	}
	if k > 10000 {
		cs = append(cs, (int8)(k%1000000/10000))
	}
	if k > 100 {
		cs = append(cs, (int8)(k%10000/100))
	}
	if k > 1 {
		cs = append(cs, (int8)(k%100/1))
	}

	for i := 0; i < len(cs); i++ {
		c := cs[i]
		if !IsGui(c) {
			cs[i] = PokeColor_FANG<<4 | (c % 16)
		}
	}

	sort.Slice(cs, func(i, j int) bool { return cs[i] < cs[j] })

	return GenCardBind(cs)
}

func ChangeColor(k int64) int64 {
	var cs []int8
	if k > 1000000000000 {
		cs = append(cs, (int8)(k%100000000000000/1000000000000))
	}
	if k > 10000000000 {
		cs = append(cs, (int8)(k%1000000000000/10000000000))
	}
	if k > 100000000 {
		cs = append(cs, (int8)(k%10000000000/100000000))
	}
	if k > 1000000 {
		cs = append(cs, (int8)(k%100000000/1000000))
	}
	if k > 10000 {
		cs = append(cs, (int8)(k%1000000/10000))
	}
	if k > 100 {
		cs = append(cs, (int8)(k%10000/100))
	}
	if k > 1 {
		cs = append(cs, (int8)(k%100/1))
	}

	var color [4]int

	for _, i := range cs {
		if !IsGui((int8)(i)) {
			color[i>>4]++
		}
	}

	var maxColor int8
	maxColorNum := 0
	for i := 0; i < len(color); i++ {
		if color[i] > maxColorNum {
			maxColor = int8(i)
			maxColorNum = color[i]
		}
	}

	for i := 0; i < len(cs); i++ {
		c := cs[i]
		if !IsGui(int8(c)) {
			if (c >> 4) == maxColor {
				cs[i] = PokeColor_HEI<<4 | (c % 16)
			} else {
				cs[i] = PokeColor_FANG<<4 | (c % 16)
			}
		}
	}

	sort.Slice(cs, func(i, j int) bool { return cs[i] < cs[j] })

	return GenCardBind(cs)
}

func GetMax(str string) (string, []int8) {
	max, trans := GetMaxBytes(StrToBytes(str))
	return KeyToStr(GenCardBind(max)), trans
}

func GetMaxBytes(bytes []int8) ([]int8, []int8) {

	var max []int8
	var trans []int8

	if len(bytes) < 5 || len(bytes) > 7 {
		return max, trans
	}

	var hand []int8
	hand = append(hand, bytes[0])
	hand = append(hand, bytes[1])

	var pub []int8
	for i := 2; i < len(bytes); i++ {
		pub = append(pub, bytes[i])
	}

	return GetMaxBytesHandPub(hand, pub)
}

func GetMaxBytesHandPub(hand []int8, pub []int8) ([]int8, []int8) {

	var ret []int8
	var trans []int8

	if len(hand) != 2 {
		return ret, trans
	}

	if len(pub) < 3 || len(pub) > 5 {
		return ret, trans
	}

	var tmp []int8

	tmp = append(tmp, hand...)
	tmp = append(tmp, pub...)

	keyData := GetKeyData(tmp)
	if keyData == nil {
		return ret, trans
	}

	max := KeyToBytes(keyData.max)

	var pubtmp []int8
	pubtmp = append(pubtmp, pub...)

	var handtmp []int8
	handtmp = append(handtmp, hand...)

	if keyData.ty == TEXAS_CARD_TYPE_TONGHUA || keyData.ty == TEXAS_CARD_TYPE_TONGHUASHUN || keyData.ty == TEXAS_CARD_TYPE_KINGTONGHUASHUN {
		var srccolor [4]int
		for i := 0; i < len(tmp); i++ {
			if !IsGui(tmp[i]) {
				srccolor[tmp[i]>>4]++
			}
		}

		var srcmaxColor int8
		srcmaxColorNum := 0
		for i := 0; i < len(srccolor); i++ {
			if srccolor[i] >= srcmaxColorNum {
				srcmaxColor = int8(i)
				srcmaxColorNum = srccolor[i]
			}
		}

		for i := 0; i < len(max); i++ {
			for j := 0; j < len(pubtmp); j++ {
				if (pubtmp[j]%16) == (max[i]%16) && (pubtmp[j]>>4) == srcmaxColor && max[i] != 0 && pubtmp[j] != 0 {
					ret = append(ret, pubtmp[j])
					max[i] = 0
					pubtmp[j] = 0
					break
				}
			}
		}

		if len(ret) < 5 {
			for i := 0; i < len(max); i++ {
				for j := 0; j < len(handtmp); j++ {
					if (handtmp[j]%16) == (max[i]%16) && (handtmp[j]>>4) == srcmaxColor && max[i] != 0 && handtmp[j] != 0 {
						ret = append(ret, handtmp[j])
						max[i] = 0
						handtmp[j] = 0
						break
					}
				}
			}
		}

		for i := 0; i < len(max); i++ {
			if max[i] != 0 {
				max[i] = (int8)(srcmaxColor<<4 | (max[i] % 16))
			}
		}
	} else {
		for j := 0; j < len(pubtmp); j++ {
			for i := 0; i < len(max); i++ {
				if pubtmp[j] == max[i] && max[i] != 0 && pubtmp[j] != 0 && !IsGui(pubtmp[j]) {
					ret = append(ret, pubtmp[j])
					max[i] = 0
					pubtmp[j] = 0
					break
				}
			}
		}

		for j := 0; j < len(handtmp); j++ {
			for i := 0; i < len(max); i++ {
				if handtmp[j] == max[i] && max[i] != 0 && handtmp[j] != 0 && !IsGui(handtmp[j]) {
					ret = append(ret, handtmp[j])
					max[i] = 0
					handtmp[j] = 0
					break
				}
			}
		}

		for i := 0; i < len(max); i++ {
			for j := 0; j < len(pubtmp); j++ {
				if (pubtmp[j]%16) == (max[i]%16) && max[i] != 0 && pubtmp[j] != 0 && !IsGui(pubtmp[j]) {
					ret = append(ret, pubtmp[j])
					max[i] = 0
					pubtmp[j] = 0
					break
				}
			}
		}

		if len(ret) < 5 {
			for i := 0; i < len(max); i++ {
				for j := 0; j < len(handtmp); j++ {
					if (handtmp[j]%16) == (max[i]%16) && max[i] != 0 && handtmp[j] != 0 && !IsGui(handtmp[j]) {
						ret = append(ret, handtmp[j])
						max[i] = 0
						handtmp[j] = 0
						break
					}
				}
			}
		}
	}

	for len(ret) < 5 {
		ret = append(ret, GUI.ToByte())
	}

	var guiTrans []int8
	for i := 0; i < len(max); i++ {
		if max[i] != 0 {
			guiTrans = append(guiTrans, max[i])
		}
	}

	sort.Slice(ret, func(i, j int) bool { return ret[i] < ret[j] })

	return ret, guiTrans
}
