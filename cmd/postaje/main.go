package main

func main() {
	var postaje map[string]Postaja
	postaje = make(map[string]Postaja)

	postaje["MZG"] = Postaja{"OŠ Marezige", "Marezige", []float64{23.1, 26.2, 27.3, 21.4}}
	postaje["KPN"] = Postaja{"PGD Koprivnik", "Koprivnik pri Kočevju", []float64{19.1, 10.2, 12.3, 18.4}}
	postaje["APA"] = Postaja{"ŠRC Apače", "Apače", []float64{13.1, 12.2}}
	postaje["SOL"] = Postaja{"Solčava, kapela", "Solčava", []float64{8.1, 7.2, 13.3, 12.4}}
	postaje["GOR"] = Postaja{"Spodnje Gorje, cerkev", "Gorje", []float64{5.1, 2.2, 3.3, 0.4}}

	// 1.
	// DODAJANJE MERITEV
	println("\nDodajanje meritev:")
	dodajMeritev(postaje, "BLABLA", 10.0)
	dodajMeritev(postaje, "APA", 100.0)
	dodajMeritev(postaje, "GOR", -2.5)

	// 3.
	// IZPIS MERITEV
	println("\n\nInformacije o postajah:")
	izpisMeritev(postaje)

	// 2. in 4.
	// POROČILO
	println("\nPoročilo:")
	izpisPorocila(postaje)

}
