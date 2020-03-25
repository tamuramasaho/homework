package main

import (
	"fmt"
	"strconv"
)

// ポリモーフィズム

// ポイントは、関数を使う側はGuitarやDrumの細かな差異や実装の詳細を把握しなくて良いこと（カプセル化、抽象化）。
// sound()が使えることさえわかっていれば、細かい差異はGuitarやDrumの側で吸収してくれるため。

type IInstrument interface {
	sound() string
}

type Guitar struct {
	numOfStrings int
}

func (g Guitar) sound() string {
	return "ギターは" + strconv.Itoa(g.numOfStrings) + "本の弦で鳴るよ JUGJUG"
}

type Drums struct {
	material string
}

func (d Drums) sound() string {
	return d.material + "でできているドラムは鳴るよ DAMDAM"
}

func main() {

	instruments := []IInstrument{
		Guitar{
			numOfStrings: 6,
		},
		Drums{
			material: "steel",
		},
	}

	for _, v := range instruments {
		fmt.Println("------------------------------------------")
		fmt.Println(v.sound())
		fmt.Println("------------------------------------------")
	}
}
