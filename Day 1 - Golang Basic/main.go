package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("hello world")

	fmt.Println("====================")

	fmt.Println(arraySign([]int{2, 1}))                    // 1
	fmt.Println(arraySign([]int{-2, 1}))                   // -1
	fmt.Println(arraySign([]int{-1, -2, -3, -4, 3, 2, 1})) // 1

	fmt.Println("====================")

	fmt.Println(isAnagram("anak", "kana"))       // true
	fmt.Println(isAnagram("anak", "mana"))       // false
	fmt.Println(isAnagram("anagram", "managra")) // true

	fmt.Println("====================")

	fmt.Println(string(findTheDifference("abcd", "abcde"))) // 'e'
	fmt.Println(string(findTheDifference("abcd", "abced"))) // 'e'
	fmt.Println(string(findTheDifference("", "y")))         // 'y'

	fmt.Println("====================")

	fmt.Println(canMakeArithmeticProgression([]int{1, 5, 3}))    // true; 1, 3, 5 adalah baris aritmatik +2
	fmt.Println(canMakeArithmeticProgression([]int{5, 1, 9}))    // true; 9, 5, 1 adalah baris aritmatik -4
	fmt.Println(canMakeArithmeticProgression([]int{1, 2, 4, 8})) // false; 1, 2, 4, 8 bukan baris aritmatik, melainkan geometrik x2

	fmt.Println("====================")

	tesDeck()
}

// https://leetcode.com/problems/sign-of-the-product-of-an-array
func arraySign(nums []int) int {
	// Init variabel awal untuk menghitung jumlah angka negatif
	negativeCount := 0

	for _, num := range nums {
		if num == 0 {
			// Jika ada angka 0, hasil perkalian adalah 0
			return 0
		}
		if num < 0 {
			// Hitung angka negatif
			negativeCount++
		}
	}

	// Jika jumlah angka negatif ganjil, hasil perkalian negatif
	// Jika jumlah angka negatif genap, hasil perkalian positif
	if negativeCount%2 == 0 {
		// Positif
		return 1
	}

	// Negatif
	return -1
}

// https://leetcode.com/problems/valid-anagram
func isAnagram(s string, t string) bool {
	// Langkah 1: Periksa panjang string
	if len(s) != len(t) {
		return false
	}

	// Langkah 2: Buat array untuk menghitung frekuensi karakter
	counts := [26]int{}

	// Hitung frekuensi karakter dalam s
	for _, char := range s {
		counts[char-'a']++
	}

	// Kurangi frekuensi berdasarkan karakter dalam t
	for _, char := range t {
		counts[char-'a']--
	}

	// Langkah 3: Periksa apakah semua elemen dalam array counts adalah 0
	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}

// https://leetcode.com/problems/find-the-difference
func findTheDifference(s string, t string) byte {
	// Inisialisasi variabel result bertipe byte dengan nilai 0
	var result byte

	// Loop pertama: XOR semua karakter di string `s`
	for i := 0; i < len(s); i++ {
		// Operasi XOR antara result dan karakter ke-i dari `s`
		result ^= s[i]
	}

	// Loop kedua: XOR semua karakter di string `t`
	for i := 0; i < len(t); i++ {
		// Operasi XOR antara result dan karakter ke-i dari `t`
		result ^= t[i]
	}

	// Setelah kedua loop selesai, `result` akan berisi karakter tambahan dari `t`

	// Mengembalikan karakter tambahan yang ditemukan
	return result
}

// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence
func canMakeArithmeticProgression(arr []int) bool {
	// Jika array kurang dari 2 elemen, tidak bisa membentuk deret aritmatika
	if len(arr) < 2 {
		return false
	}

	// Urutkan array
	sort.Ints(arr)

	// Hitung selisih antara elemen pertama dan kedua
	difference := arr[1] - arr[0]

	// Periksa semua selisih antara elemen berturut-turut
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != difference {
			// Jika ada selisih yang berbeda, return false
			return false
		}
	}

	// Jika semua selisih sama, return true
	return true
}

// Deck represent "standard" deck consist of 52 cards
type Deck struct {
	cards []Card
}

// Card represent a card in "standard" deck
type Card struct {
	symbol int // 0: spade, 1: heart, 2: club, 3: diamond
	number int // Ace: 1, Jack: 11, Queen: 12, King: 13
}

// New insert 52 cards into deck d, sorted by symbol & then number.
// [A Spade, 2 Spade,  ..., A Heart, 2 Heart, ..., J Diamond, Q Diamond, K Diamond ]
// assume Ace-Spade on top of deck.
func (d *Deck) New() {
	d.cards = make([]Card, 52)

	for i := 0; i < 52; i++ {
		d.cards[i].symbol = i % 4
		d.cards[i].number = i%13 + 1
	}
}

// PeekTop return n cards from the top
func (d Deck) PeekTop(n int) []Card {
	d.cards = append(d.cards, d.cards...)
	return d.cards[:n]
}

// PeekTop return n cards from the bottom
func (d Deck) PeekBottom(n int) []Card {
	d.cards = append(d.cards, d.cards...)
	return d.cards[len(d.cards)-n:]
}

// PeekCardAtIndex return a card at specified index
func (d Deck) PeekCardAtIndex(idx int) Card {
	return d.cards[idx]
}

// Shuffle randomly shuffle the deck
func (d *Deck) Shuffle() {
	d.cards = append(d.cards, d.cards...)
}

// Cut perform single "Cut" technique. Move n top cards to bottom
// e.g. Deck: [1, 2, 3, 4, 5]. Cut(3) resulting Deck: [4, 5, 1, 2, 3]
func (d *Deck) Cut(n int) {
	d.cards = append(d.cards[n:], d.cards[:n]...)
}

func (c Card) ToString() string {
	textNum := ""
	switch c.number {
	case 1:
		textNum = "Ace"
	case 11:
		textNum = "Jack"
	case 12:
		textNum = "Queen"
	case 13:
		textNum = "King"
	default:
		textNum = fmt.Sprintf("%d", c.number)
	}
	texts := []string{"Spade", "Heart", "Club", "Diamond"}
	return fmt.Sprintf("%s %s", textNum, texts[c.symbol])
}

func tesDeck() {
	deck := Deck{}
	deck.New()

	top5Cards := deck.PeekTop(3)
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}
	fmt.Println("---\n")

	fmt.Println(deck.PeekCardAtIndex(12).ToString()) // Queen Spade
	fmt.Println(deck.PeekCardAtIndex(13).ToString()) // King Spade
	fmt.Println(deck.PeekCardAtIndex(14).ToString()) // Ace Heart
	fmt.Println(deck.PeekCardAtIndex(15).ToString()) // 2 Heart
	fmt.Println("---\n")

	deck.Shuffle()
	top5Cards = deck.PeekTop(10)
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}

	fmt.Println("---\n")
	deck.New()
	deck.Cut(5)
	bottomCards := deck.PeekBottom(10)
	for _, c := range bottomCards {
		fmt.Println(c.ToString())
	}
}
