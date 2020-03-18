package main

import (
	"fmt"
)

//抽象クラスのような位置付けのInstrument構造体を定義
type Instrument struct {
	name string
	manufacturer string
	price int
	tone string
}

// ダックタイピング（のようなこと）を行うためのインターフェースを定義
type InstrumentInterface interface {
	sound() string
	findMore() string
}

// Instrument構造体にインターフェースを実装
func (i Instrument) sound() string {
	return fmt.Sprintf("%v は鳴るよ %v\n", i.name, i.tone)
}
func (i Instrument) findMore() string {
	return fmt.Sprintf("メーカー:%v\n値段: %v円", i.manufacturer, i.price)
}


/*　Guitar構造体にInstrument構造体を埋め込む（deligate)。
	これにより、InstrumentInterfaceも
	実装していることになりメソッドが使える。これがダックタイピング。*/

type Guitar struct {	
	Instrument
}

type Drums struct {
	Instrument
}

func main() {
	instruments := []InstrumentInterface{
		Guitar{
			Instrument: Instrument{
				name: "ギター",
				manufacturer: "gibson",
				price: 10,
				tone: "JUGJUG",
			},
		},
		Drums{
			Instrument: Instrument{
				name: "ドラム",
				manufacturer: "gretch",
				price: 20,
				tone: "DAMDAM",
			},
		},
	}

	for _, v := range instruments{
		fmt.Println("------------------------------")
		fmt.Println(v.sound())
		fmt.Println(v.findMore())
		fmt.Println("------------------------------")
	}
}

