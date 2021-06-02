package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
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
		ileTegoZycia = rand.Int() % ((s.x * s.y) / 3)
	}

	for i := 0; i < ileTegoZycia; i++ {
		x := rand.Intn(s.x)
		y := rand.Intn(s.y)

		s.aktualnaPlansza[y][x] = true
	}
}

func (s *Swiat) NowaGeneracja() {
	for y := 0; y < s.y; y++ {
		for x := 0; x < s.x; x++ {
			s.przyszlaPlansze[y][x] = s.ZycieCzySmierc(x, y)
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
				if s.aktualnaPlansza[sy][sx] {
					ileZywychSasiadow++
				}
			}
		}
	}

	return ileZywychSasiadow == 3 || (ileZywychSasiadow == 2 && s.aktualnaPlansza[y][x])
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

func flagi() (int, int, int64, int, string) {
	// zmienne przy uruchomieniu
	// w - szerokość świata
	// h - wysokość świata
	// s - seed generujący
	// l - ilość startowych komórek żywych
	w := flag.Int("w", 10, "szerokość świata")
	h := flag.Int("h", 10, "wysokość świata")
	s := flag.Int("s", 9, "seed generujący")
	l := flag.Int("l", 0, "ilość startowych komórek żywych")

	f := flag.String("f", "", "plik z mapą startową")

	flag.Parse()

	return *w, *h, int64(*s), *l, *f
}

func WczytajSwiatZPliku(fs string) Swiat {
	f, err := os.Open(fs)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	lx, _, err := reader.ReadLine()
	ly, _, err := reader.ReadLine()

	w, _ := strconv.Atoi(string(lx))
	h, _ := strconv.Atoi(string(ly))
	s := TworzenieSwiata(w, h)

	for {
		lwsp, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		wsp := strings.Split(string(lwsp), ":")
		x, _ := strconv.Atoi(string(wsp[0]))
		y, _ := strconv.Atoi(string(wsp[1]))

		s.aktualnaPlansza[y][x] = true
	}

	return s
}

func main() {
	w, h, s, l, f := flagi()

	var swiat Swiat

	if f != "" {
		swiat = WczytajSwiatZPliku(f)
	} else {
		swiat = TworzenieSwiata(w, h)
		swiat.LosujJakiesDodatkoweZycie(s, l)
	}
	swiat.RysujSwiat()

	// zaczynamy gre
	for {
		czyszczenieKonsoli()
		(&swiat).NowaGeneracja()
		swiat.RysujSwiat()
		time.Sleep(time.Second / 2)
	}
}
