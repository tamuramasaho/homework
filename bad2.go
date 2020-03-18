package main

import "fmt"

type Instrument interface {	
	sound() string
	findMore() string
}

// 委譲せずに構造体を複数定義、新しい楽器が増えた場合に、同じフィールドを再び書くことになる。
type Guitar struct {
	manufacturer string
	price int
}

type Drums struct {
	manufacturer string
	price int
}

/*　Instrumentインターフェースを実装するために、Guitar構造体とDrums構造体にそれぞれメソッド
	を実装している。ここでも新しく楽器が増えた場合は同じようなコードを再び書かなければいけない　*/
func (g Guitar) sound() string {
	return "ギター は鳴るよ JUGJUG\n"
}
func (g Guitar) findMore() string {
	return fmt.Sprintf("メーカー:%v\n値段: %v", g.manufacturer, g.price)
}


func (d Drums) sound() string {
	return "ドラム は鳴るよ DAMDAM\n"
}
func (d Drums) findMore() string {
	return fmt.Sprintf("メーカー:%v\n値段: %v", d.manufacturer, d.price)
}

func main() {

	instruments := []Instrument{
		Guitar{
			manufacturer: "gibson",
			price: 10,
		},
		Drums{
			manufacturer: "gretch",
			price: 20,
		},
	}

	for _, v := range instruments{
		fmt.Println("------------------------------")
		fmt.Println(v.sound())
		fmt.Println(v.findMore())
		fmt.Println("------------------------------")
	}
}

