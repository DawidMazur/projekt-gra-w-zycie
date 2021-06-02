package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

//czyszczenie ekranu
func czyszczenieKonsoli() {
	cmd := exec.Command("cmd", "/c", "cls", "clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

type Swiat struct {
	x, y                             int      // wielkości planszy/swiata
	aktualnaPlansza, przyszlaPlansze [][]bool // plansza komórek żywych lub martwych i dodatkowa plansza do wyliczania przyszłego pokolenia
}

func TworzenieSwiata(x, y int) Swiat {
	plansza := make([][]bool, y)
	for i := range plansza {
		plansza[i] = make([]bool, x)
	}
	plansza2 := make([][]bool, y)
	for i := range plansza2 {
		plansza2[i] = make([]bool, x)
	}
	return Swiat{
		x:               x,
		y:               y,
		aktualnaPlansza: plansza,
		przyszlaPlansze: plansza2,
	}
}

func (s *Swiat) LosujJakiesDodatkoweZycie(seed int64, ileTegoZycia int) {
	rand.Seed(seed)
	if ileTegoZycia == 0 {
		ileTegoZycia = rand.Int() % ((s.x * s.y) / 4)
	}

	for i := 0; i < ileTegoZycia; i++ {
		s.aktualnaPlansza[rand.Intn(s.x)][rand.Intn(s.y)] = true
	}
}

func (s *Swiat) NowaGeneracja() {
	for y := 0; y < s.y; y++ {
		for x := 0; x < s.x; x++ {
			s.przyszlaPlansze[x][y] = s.ZycieCzySmierc(x, y)
		}
	}
	s.aktualnaPlansza, s.przyszlaPlansze = s.przyszlaPlansze, s.aktualnaPlansza
}

func (s Swiat) ZycieCzySmierc(x, y int) bool {
	ileZywychSasiadow := 0
	for wiersz := 0; wiersz < 3; wiersz++ {
		for kolumna := 0; kolumna < 3; kolumna++ {
			if wiersz == 1 && kolumna == 1 {
				continue
			}

			sx := x + wiersz - 1
			sy := y + kolumna - 1

			if sx < s.x && sy < s.y && sx >= 0 && sy >= 0 {
				if s.aktualnaPlansza[sx][sy] {
					ileZywychSasiadow++
				}
			}
		}
	}

	return ileZywychSasiadow == 3 || (ileZywychSasiadow == 2 && s.aktualnaPlansza[x][y])
}

func (s Swiat) RysujSwiat() {
	for i := 0; i < s.x+2; i++ {
		fmt.Print("*")
	}
	fmt.Println()

	for ix := 0; ix < s.y; ix++ {
		fmt.Print("*")
		for iy := 0; iy < s.x; iy++ {
			if s.aktualnaPlansza[ix][iy] {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("*")
		fmt.Println()
	}

	for i := 0; i < s.x+2; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	// zmienne przy uruchomieniu

	swiat := TworzenieSwiata(10, 10)
	swiat.LosujJakiesDodatkoweZycie(9, 17)
	swiat.RysujSwiat()

	// zaczynamy gre
	for {
		czyszczenieKonsoli()
		(&swiat).NowaGeneracja()
		swiat.RysujSwiat()
		time.Sleep(time.Second / 2)
	}
}
