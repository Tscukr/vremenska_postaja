// simulacija vremenske postaje iz 1. naloge VPAS
package vreme

import (
	"fmt"
	"strconv"
	"strings"
)

type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

func DodajPostajo(postaje map[string]Postaja, oznaka string, ime string, kraj string, meritve []float64) {
	postaje[oznaka] = Postaja{ime: ime, kraj: kraj}
}

// postaji z dano oznako doda novo meritev
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64) {
	if temperatura > 50 || temperatura < -50 {
		fmt.Printf("[!] %s: Temperatura %.1f je izven dovoljenega območja [-50, 50]\n", oznaka, temperatura)
		return
	}

	p, ok := postaje[oznaka]

	if !ok {
		fmt.Printf("[!] Postaja z oznako %s ne obstaja\n", oznaka)
		return
	}

	fmt.Printf("    %s: dodana meritev %.1f", oznaka, temperatura)

	p.meritve = append(p.meritve, temperatura)
	postaje[oznaka] = p

}

// vrne najnižjo, najvišjo in povprečno temperaturo postaje
func Statistika(postaje map[string]Postaja, oznaka string) (najnizja, najvisja, povprecna float64, ok bool) {
	p, obstaja := postaje[oznaka]
	ok = true

	if !obstaja {
		ok = false
		return
	}

	if len(p.meritve) < 3 {
		ok = false
		return
	}

	najnizja = 51.0
	najvisja = -51.0
	var vsota float64 = 0

	for _, vrednost := range p.meritve {
		if vrednost > najvisja {
			najvisja = vrednost
		}
		if vrednost < najnizja {
			najnizja = vrednost
		}

		vsota += vrednost
	}

	povprecna = vsota / float64(len(p.meritve))
	return
}

// izpiše oznake, imena in kraje vseh postaj ter njihove meritve
func IzpisMeritev(postaje map[string]Postaja) {
	for kr, po := range postaje {
		fmt.Printf("%s - %s (%s): [%s]\n", kr, po.ime, po.kraj, izpisSeznama(po.meritve))
	}
}

// pomožna funkcija, vrne string, enak zapisu fmt.Printf(seznam)
func izpisSeznama(seznam []float64) (izpis string) {
	izpis = ""

	for _, m := range seznam {
		izpis = izpis + strconv.FormatFloat(m, 'f', -1, 64) + " "
	}
	izpis = strings.Trim(izpis, " ")
	return
}

// pomožna funkcija, za ravnanje besedila v tabeli
func najdaljseImePostaje(postaje map[string]Postaja) (najdaljsa int) {
	najdaljsa = 0

	for _, p := range postaje {
		if len(p.ime) > najdaljsa {
			najdaljsa = len(p.ime)
		}
	}
	return
}

// Za vsako postajo izpiše ime ter najnižjo, najvišjo in povprečno temperaturo.
func IzpisPorocila(postaje map[string]Postaja) {
	maxIme := najdaljseImePostaje(postaje)
	var opis string

	fmt.Printf("%-*s\tmin\tmax\tpovp\topis\n", maxIme, "Postaja")
	println("-------------------------------------------------------------")

	for kratica, p := range postaje {
		min, max, avg, ok := Statistika(postaje, kratica)

		switch {
		case avg >= 25:
			opis = "Vroče"
		case avg >= 15:
			opis = "Toplo"
		case avg >= 5:
			opis = "Hladno"
		default:
			opis = "Mrzlo"
		}

		if ok {
			fmt.Printf("%-*s\t%.1f\t%.1f\t%.1f\t%s\n", maxIme, p.ime, min, max, avg, opis)
		} else {
			// postaje s problemi se izpišejo na konc seznama
			defer fmt.Printf("%-*s\t[!] postaja ima manj kot 3 meritve\n", maxIme, p.ime)
		}
	}
}
