//Go is a statically type language
//Basic data type - int,string,bool,float64
package main

func main() {

	cards := newDeckFromFile("MyCards")
	cards.shuffle()
	cards.print()

}
