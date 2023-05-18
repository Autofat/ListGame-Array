package main

import "fmt"

const NMAX = 20

type Game struct {
	judul, namaPublisher, platform string
	tahunRilis, jumlahPemain       int
}

type TabGame struct {
	arrGame [NMAX]Game
	n       int
}

func main() {
	var t TabGame
	var tDisplay string
	// var i int
	isiArray(&t)
	// for i := 0; i < t.n; i++ {
	// 	fmt.Println(t.arrGame[i])
	// }
	fmt.Print("Data yang ingin dicari: ")
	fmt.Scan(&tDisplay)
	for i := 0; i < NMAX; i++ {
		displayArray(tDisplay, t, i)
	}
	thnTermuda(t)
	thnTertua(t)

}

func isiArray(t *TabGame) {
	var inJudul, inPlatform, inPublish string
	var i, inTahun, inJumlah int
	fmt.Print("Masukan Judul Game: ")
	fmt.Scan(&inJudul)
	for i < NMAX && inJudul != "XXX" {
		t.arrGame[i].judul = inJudul
		fmt.Print("Tahun perilisan: ")
		fmt.Scan(&inTahun)
		if inTahun > 0 {
			t.arrGame[i].tahunRilis = inTahun
			fmt.Print("Jumlah Pemain: ")
			fmt.Scan(&inJumlah)
		}
		if inJumlah > 0 {
			t.arrGame[i].jumlahPemain = inJumlah
			fmt.Print("Nama Publisher: ")
			fmt.Scan(&inPublish)
		}
		if inPublish != "XXX" {
			t.arrGame[i].namaPublisher = inPublish
			fmt.Print("PC / MOBILE / PS / XBOX: ")
			fmt.Scan(&inPlatform)
		}
		if inPlatform != "XXX" {
			t.arrGame[i].platform = inPlatform
		}
		t.n++
		i++
		fmt.Print("Masukan Judul Game: ")
		fmt.Scan(&inJudul)

	}
}

func displayArray(target string, t TabGame, i int) {
	if target == t.arrGame[i].judul {
		fmt.Println("Judul: ", t.arrGame[i].judul)
		fmt.Println("Tahun Rilis: ", t.arrGame[i].tahunRilis)
		fmt.Println("Jumlah Pemain: ", t.arrGame[i].jumlahPemain)
		fmt.Println("Nama Publisher: ", t.arrGame[i].namaPublisher)
		fmt.Println("Platform: ", t.arrGame[i].platform)
	}
}

func thnTertua(t TabGame) {
	var muda int = 0
	for i := 0; i < t.n; i++ {
		if t.arrGame[muda].tahunRilis > t.arrGame[i].tahunRilis {
			muda = i
		}
	}
	fmt.Println("Game rilisan terlama: ", t.arrGame[muda].judul, "Dirilis pada tahun ", t.arrGame[muda].tahunRilis)
}

func thnTermuda(t TabGame) {
	var tua int = 0
	for i := 0; i < t.n; i++ {
		if t.arrGame[tua].tahunRilis < t.arrGame[i].tahunRilis {
			tua = i
		}
	}
	fmt.Println("Game rilisan terbaru: ", t.arrGame[tua].judul, "Dirilis pada tahun ", t.arrGame[tua].tahunRilis)
}
