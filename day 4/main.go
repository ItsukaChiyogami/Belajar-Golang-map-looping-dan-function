package main

import (
	"fmt"
)

func main() {

	// tipe data map
	// pada array or slice, untuk mengakses data, perlu menggunakan index dimulai dari 0
	// map adalah tipe data lain yang berisikan sekumpulan data yang sama,namun kita bisa menentukan jenis tipe data index yang akan kita gunakan.
	// sederhananya, map adalah tipe data kumpulan key-value (kata kunci - nilai), dimana kata kuncinya bersifak unique dan tidak sama
	// berbeda dengan array dan slice, jumlah datat yang kita masukkan ke map boleh sebanyak banyaknya,asalkan kata kuncinya berbeda
	// jika menggunakan kata kunci yang sama data yang lama akan terbaharui yang artinya yang lama ilang yang baru datang kacaw

	wife := map[string]string{ //map [tipe key] dan string
		"name":    "Itsuka Chiyogami",
		"age":     "8th",
		"country": "japan",
		"city":    "hokkaido",
		"birth":   "2015/11/11",
		"gender":  "WOMAN",
	}

	// function map
	// len, map[key], map[key] = value, make[map(typekey)value], delete

	fmt.Println()
	fmt.Println(wife)
	fmt.Println(len(wife))
	fmt.Println(wife["name"])
	delete(wife, "gender")
	fmt.Println(wife)

	manga := make(map[string]string)
	manga["title"] = "date a live"

	fmt.Println(manga)

	// if else if dan else expression
	// digunakan jika kondisi benar maka lanjutkan

	var weaboo = 1

	if weaboo == 1 {

		fmt.Println("aku wibu aku bangga")
	} else if weaboo == 2 {
		fmt.Println("aku bukan wibu")
	} else {
		fmt.Println("apa itu wibu?")
	}

	// if with short statement
	// if mendukung short statement sebelum kondisi benar atau salah
	// hal ini sangat cocok ntuk membuat statement yang sederhana sebelum melakukan pengecekan
	// terhadap kondisi

	if length := len(wife["name"]); length >= 5 {
		fmt.Println("waifumu kawaii mas")
	} else {
		fmt.Println("nama waifumu tidak panjang")
	}

	// switch expression
	// selain if untuk melakukan percabangan, kita juga bisa menggunakan Switch Expression
	// Switch Expression sangat sederhana dibandingkan if
	// biasanya Switch Expression digunakan untuk melakukan pengecekan ke kondisi dalam satu variabel

	var input string
	fmt.Println("masukkan nama")
	fmt.Scanln(&input)

	switch {
	case input != "":
		fmt.Println("selamat datang", input)
	default:
		fmt.Println("halo")
	}

	ayank := "origami"

	switch ayank {
	case "origami":
		fmt.Println("selamat datang Sayank")
	default:
		fmt.Println("kamu siapa??")
	}

	// Switch Short Statement

	switch lenght := len(wife["name"]); lenght > 5 {
	case true:
		fmt.Println("panjang")
		fmt.Println(wife["name"])
	case false:
		fmt.Println("pendek")

	}

	// switch tanpa kondisi
	// kondisi switch expression tidak wajib
	// jika kita tidak menggunakan kondisi switch expression, kita bisa menambahkan kondisi tersebut di setiap casenya

	var umur uint

	fmt.Scan(&umur)
	switch {
	case umur <= 17:

		fmt.Println("Umur kamu =", umur, "gak boleh nonton Anime")

	case umur >= 18:

		fmt.Println("Umur kamu = ", umur, "kamu boleh nonton Anime")

	default:

		fmt.Println("masukkan umur")

	}

	// for loop
	// perulangan menggunakan for
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke ", counter)
		counter++
	}

	fmt.Printf("selesai")
	fmt.Printf("")
	fmt.Printf("")

	// for dengan statement
	// dalam for, kita bisa menambahkan statement, dimana terdapat 2 statement yang bisa tambahkan di for
	// init statement, yaitu statement sebelum for di eksekusi
	// Post statement, yaitu statement yang akan selalu di eksekusi di akhir perulangan

	for HealtPoint := 1; HealtPoint < 10; HealtPoint++ {
		fmt.Println("Perulangan Ke = ", HealtPoint)
	} // gak rapi

	HealtPoint := 1
	for HealtPoint <= 10 {
		fmt.Println("perulangan ke ", HealtPoint)
		HealtPoint++
	} // rapi

	fmt.Println()

	// for range
	// for bisa digunakan untuk melakukan iterasi terhadap semua data collection
	// data collection Array, slice dan map

	Idol := []string{"Hatsune Miku", "JKT48", "AKB48"}
	for i := 0; i < len(Idol); i++ {
		fmt.Println(Idol[i])
	}

	fmt.Println()

	for index, Idols := range Idol {
		fmt.Println("index", index, "=", Idols)
	}

	fmt.Println()

	for _, Idols := range Idol {
		fmt.Println(Idols)
	}

	// break dan continue
	// break khusus berhenti dan continue melanjutkan
	// sering dipakai di looping

	for i := 0; i < 10; i++ {

		if i == 5 {
			break

		}

		fmt.Println("kata ini ke == ", i)
	}

	for i := 0; i < 10; i++ {

		if i%2 == 0 {
			continue //jika genap atau jelasnya i = 2,4,6,... dia akan di skip,continue sering dibuat skip
		}
		fmt.Println("ini ke = ", i)

	}

	// Function
	// function adalah sebuah blok kode yang sengaja dibuat agar bisa digunakan berulang ulang
	// cara agar membuat func sangat sederhana, hanya menggunakan kata func lalu diikuti nama function dan kode blok isi functionnya
	// setelah membuat function, kita bisa mengeksekusi function tersebut dengan memanggilnya menggunakan kata kunci nama function beserta tanda kutip
	sayhello()

	// function parameter
	// function parameter ialah data dari luar yang mana disebut parameter,kadang kadang kita membutuhkannya
	// kita bisa menambahkan parameter di func lebih dari satu
	// parameter tidak wajib jdi bisa membuat func tanpa parameter
	// namun jika kita menambahkan parameter di func, maka ketika dipanggil funcnya kita wajib masukkaan data ke parameternya
	Greets("Origami")
	Greets("Chiyogami")

	spasi()

	// function return value
	// func bisa mengembalikan data
	// untuk memberitahu bahwa func mengembalikan data, kita harus menulis tipe data kembalian dari func tersebut
	// jika func tersebut kita deklarasikan dengan tipe data pengembalian, maka wajib di dalam func kita harus mengembalikan data
	// untuk mengembalikan data dari func, kita bisa menggunakan kata kunci return, diikuti dengan datanya

	result := gethello("ayank Chiyogami")
	fmt.Println(result)

	fmt.Println(gethello("sayank Origami"))

	spasi()

	// returning multiple values
	// func tidak hanya dapat mengembalikan satu value, tapi juga bisa multiple value
	// untuk memberitahu jika func mengembalikan multiple value, kita harus menulis semua tipe data return valuenya di func
	firstname, lastname := fullname()
	fmt.Println(firstname, lastname)
	fmt.Println(lastname, firstname)

	spasi()

	// menghiraukan return value
	// multple return wajib ditangkap semua valuenya
	// jika kita ingin menghiraukan return value tersebut, kita bisa menggunakan tanda _
	__, lastname := fullname()
	fmt.Println(lastname)
	fmt.Println(lastname, __)

	// named return values
	// biasanya saat kita memberitahu bahwa sebuah func mengembalikan value, maka kita hanya mendeklarasikan tipe data return value di dalam func
	// namun kita juga bisa membuat variable secara lansung di tipe data return func
	name, age := getwaifu()
	fmt.Println(name, age)

	// variadic func
	// parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs
	// varargs artinya datanya bisa menerima lebih dari 1 input, atau semacam array.
	// apa bedanya parameter biasa dan tipe data array?
	// jika parameter tipe array, kita wajib membuat array terlebih dahulu sebelum mengirimkan function
	// jika parameter menggunakan varargs, kita bisa lansung mengirim datanya, jika lebih dari 1, cukup gunakan tanda koma

	total := sumALL(10, 10, 10, 10, 10, 10)
	fmt.Println(total)

	spasi()
}

func sumALL(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func getwaifu() (name, age string) {
	name = "Tobiichi Chiyogami"
	age = "8th old"

	return name, age
}

func fullname() (string, string) {
	return "Tobiichi", "Origami"
}

func gethello(name string) string {
	return "hello " + name
}

func sayhello() {
	fmt.Println("Say Hello")
}

func Greets(ayank string) {
	fmt.Println("Selamat Pagi ayank", ayank)
}

func spasi() {
	fmt.Println()
}

/* Day 3

// Type Data Array

// Mendeklarasikan Anime menjadi 4 elemen {array 4 elemen}
var Anime [4]string

Anime[0] = "Naruto"
Anime[1] = "Date A Live"
Anime[2] = "To Love Ru"
Anime[3] = "Shelter"

fmt.Println(Anime)

// array deklarasi lansung
var waifu = [...]string{
	"Tobiichi Origami",
	"Itsuka Chiyogami",
	"Menhera",
}

fmt.Println(waifu[2])

fmt.Println()

// for loop
for i := 0; i < 3; i++ {
	fmt.Printf("%d.%s\n", i+1, waifu[i])
}

// function ke array
// len (array)
// array[index]
// array[index] = value

fmt.Println(len(waifu))
fmt.Println()

// tipe data slice
// potongan dari data array
// mirip seperti array tapi ukuran slice dapat berubah
// slice dan array selalu terkoneksi, dimana slice adalah data yang bisa mengakses sebagian atau seluruh data di array

// type data slice
// pointer, lenght, capacity

// cara membuat slice dari Array

slices := waifu[:]
fmt.Println("mengambil semua dalam array/slices = ", slices)

slices2 := waifu[:2]
fmt.Println("mengambil dari 0 sampai 1", slices2)

slices3 := waifu[1:]
fmt.Println("mengambil dari 1 sampai terakhir", slices3)

var slices4 []string = waifu[0:2]
fmt.Println("mengambil dari 0 sampai sampai 1", slices4)

// function slice
// len, cap, append, make, copy

day := [...]string{
	"senin",
	"selasa",
	"rabu",
	"kamis",
	"jumat",
	"sabtu",
	"minggu",
}

fmt.Println()
fmt.Println(day)

dayslice1 := day[5:]

dayslice1[0] = "minggu"
dayslice1[1] = "sabtu"
fmt.Println(day)

dayslice2 := append(dayslice1, "libur")

fmt.Println(dayslice2)

dayslice3 := append(day[:], "liburan")
fmt.Println(dayslice3)
fmt.Println(cap(dayslice3))
fmt.Println(cap(day))
fmt.Println(day)

Character := make([]string, 2, 5)
Character[0] = "Naruto"
Character[1] = "Tobiichi Origami"
// Character[2] = "Tobiichi Origami" error akibat gak bisa di tambahkan lagi max 2 dari 5 cap dan wajib pakai append

fmt.Println()
fmt.Println(Character)
fmt.Println(len(Character))
fmt.Println(cap(Character))

Character2 := append(Character, "Itsuka Chiyogami")
fmt.Println(Character2)
fmt.Println(len(Character2))
fmt.Println(cap(Character2))

// mengubah slice 1
Character2[0] = "Menhera"
fmt.Println(Character2)
fmt.Println(Character)
fmt.Println()

// membuat copy slice
toCharacter := Character2[:]
CopyCharacter := make([]string, len(toCharacter), cap(Character2))

copy(CopyCharacter, toCharacter)
fmt.Println(CopyCharacter)

iniarray := [...]int{1, 2, 3, 4, 5}
inislice := []int{1, 2, 3, 4, 5}
fmt.Println(iniarray)
fmt.Println(inislice)

*/

/* Day 2

// Constant Variabel Tidak dapat diubah setelah diberi nilai pertama kali
// Cara Pembuatan Constant mirip dengan Variabel dan wajib lansung menginisialisasikan datanya

const (
	nama = "Tobiichi Origami"
	umur = ""
)

fmt.Println(nama, umur)

// konversi tipe data numeric (1)
// jika tipe data overflow = kembali ke belakang
var (
	nilai1 int64 = 98500
	nilai2 int32 = int32(nilai1)
	nilai3 int16 = int16(nilai2)
)

fmt.Println(nilai1)
fmt.Println(nilai2)
fmt.Println(nilai3)
fmt.Println()

// konversi tipe data string (2)
var (
	name          = "Origami"
	e       uint8 = name[6]   // e adalah byte
	estring       = string(e) // konversi byte ke string
)

fmt.Println(e)       // Akan menampilkan nilai byte dari karakter 'O'
fmt.Println(estring) // Akan menampilkan karakter 'O'
fmt.Println()

// type Declarations
// Type Declarations membuat ulang data baru dari tipe yang sudah ada
// Type Declarations biasanya digunakan untuk membuat alias terhadap data yang sudah ada,dengan tujuan agar lebih mudah dimengerti.

type Umur int
var UmurOrigami Umur = 18
var tahunlalu int = 17
var dulu Umur = Umur(tahunlalu)
fmt.Println(UmurOrigami)
fmt.Println(Umur(8))
fmt.Println(Umur(dulu))
fmt.Println()

// Operasi Matematika

var (
	a = 10
	b = 20
	c = 30
	d = 40
	f = a + b - c/d
)
fmt.Println(f)
fmt.Println()

// Augmented Assigment
var (
	i = 10
)
i += 10
fmt.Println(i)

// Unary Operator
var (
	G = 1
)
G++
fmt.Println(G)
G--
fmt.Println(G)
fmt.Println()

// operasi perbandingan
// operasi perbandingan adalah operasi yang membandikan dua buah data
// biasa menghasilkan boolean (True OR False)
var (
	number1 int = 1
	number2 int = 2

	anime1 = "DateALive"
	anime2 = "DateALive"

	waifu1 = "Tobiichi Origami"
	waifu2 = "Itsuka Chiyogami"
)

var result2 bool = anime1 == anime2
var result bool = number1 >= number2
var result3 bool = waifu1 != waifu2

fmt.Println("1 lebih besar dari 2 = ", result)
fmt.Println("Judul Animenya Apakah sama = ", result2)
fmt.Println("Kedua Waifuku Apakah beda = ", result3)
fmt.Println()

// Operasi Boolean
// && = dan
// || = OR
// ! = kebalikan

var (
	grade_A = 90
	exam    = 100

	nilaisiswa1 bool = 100 >= grade_A
	examsiswa1  bool = 100 >= exam

	nilaisiswa2 bool = 100 > grade_A
	absensiswa2 bool = 80 > exam

	// nilaisiswa3 bool = 100 > grade_A
	// absensiswa3 bool = 1 > absen

	pass1 = nilaisiswa1 && examsiswa1
	pass2 = nilaisiswa2 || absensiswa2
)

fmt.Println("Apakah Siswa 1 Lulus? = ", pass1)
fmt.Println("Apakah Siswa 2 Lulus? = ", pass2)

*/

/* DAY 1

// about integer and data floating
// complex64 float32 int8 unint8

fmt.Println("satu = ", 1)
fmt.Println(`dua = `, 2)
fmt.Println(`Tiga koma lima = `, 3.5)

fmt.Println()
// boolean
fmt.Println("true = ", true)
fmt.Println("false = ", false)

// string
fmt.Println("")
fmt.Println("String")

fmt.Println("")

// function for String
fmt.Println(len("wkwkwkwk")) //menghitung jumlah string
fmt.Println("nani"[1])       // menghitung ASCII string

// Variable
var name string
name = "Chiyogami"
fmt.Println(name)

name = "Tobiichi Origami"
fmt.Println(name)

Country := "Indonesia +62" //menambahkan : di := akan membuat variable gak bisa diulang contoh nama := dan nama := akan error
fmt.Println(Country)

Country = "japan +81"
fmt.Println(Country)

// Deklarasi Multiple Variable
var (
	Nama = "Tobiichi Origammi"
	Umur = 18
	// wajib deklarasikan dan gunakan variablenya atau akan error
)

fmt.Println(Nama)
fmt.Println(Umur)

*/
