package main

import "fmt"

type Instrument interface {	
	sound() string
	findMore() string
}

type Guitar struct {
	manufacturer string
	price int
}

type Drums struct {
	manufacturer string
	price int
}

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

