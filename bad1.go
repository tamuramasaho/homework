package main

import "fmt"

type Instrument struct {
	manufacturer string
	price int
	kind string
}

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