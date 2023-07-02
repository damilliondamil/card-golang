package main

import "fmt"

type deck []string

func (d deck) tampilkan() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func kartuBaru() deck {
	cards := deck{}

	jenisKartu := []string{"Sekop", "Wajik", "Hati", "Keriting"}
	nilaiKartu := []string{"As", "Dua", "Tiga", "Empat"}

	for _, jenis := range jenisKartu {
		for _, nilai := range nilaiKartu {
			cards = append(cards, nilai+" "+jenis)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
