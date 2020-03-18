package main

import "fmt"

//構造体の定義

type Instrument struct {
	manufacturer string
	price int
	kind string
}

/* Instrument型で定義した構造体のそれぞれの音を鳴らすためのメソッド。
	楽器によって音が変わるので、switch文で分岐させている。
	楽器が増えるたびにcaseを増やす修正が必要なので、SOLID原則のOCPに適合していない*/
func sound(v Instrument) {
	switch v.kind {
	case "guitar":
		fmt.Printf("%vで作られた%v円のギターはJUG JUG鳴るよ\n",
						v.manufacturer, v.price)
	case "drums":
		fmt.Printf("%vで作られた%v円のドラムはDAM DAM鳴るよ\n",
						v.manufacturer, v.price)
	default:
		return
	}
}

func main() {
	guitar := Instrument{
		manufacturer: "gibson",
		price: 10,
		kind: "guitar",
	}
	drum := Instrument{
		manufacturer: "gretch",
		price: 20,
		kind: "drums",
	}

	instruments := []Instrument{guitar, drum}
	for _, v := range instruments{
		fmt.Println("------------------------------")
		sound(v)
		fmt.Println("------------------------------")
	}
}