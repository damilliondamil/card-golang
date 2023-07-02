package main

func main() {
	// card := "Royal Flush"
	// card = "Full House"
	// card := kartuBaru()
	// fmt.Println(card)
	// cards := deck{"Ace of Diamond", kartuBaru()}
	// cards = append(cards, "Straight")

	cards := kartuBaru()

	ditangan, sisaTumpukan := deal(cards, 2)

	ditangan.tampilkan()
	sisaTumpukan.tampilkan()
}
