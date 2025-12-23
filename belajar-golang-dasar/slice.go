package main

import "fmt"

func main() {
	// konsep slice[low:high]
	// jika merubah referensi arraynya dari slice akan merubah arayy utamanya juga

	names := [...]string{"ilham", "king", "uji", "lana", "anjay", "audah"}

	slice1 := names[4:6]
	fmt.Println(slice1) // [anjay audah]

	slice2 := names[:3]
	fmt.Println(slice2) // [ilham king uji]

	slice3 := names[3:]
	fmt.Println(slice3) // [lana anjay audah]

	var slice4 []string = names[:]
	fmt.Println(slice4) // [ilham king uji lana anjay audah]

	// method slice
	fmt.Println(len(slice4)) // panjang slice
	fmt.Println(cap(slice4)) // mendapatkan capasitas slice

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	fmt.Println(days)

	daysSlice1 := days[5:] // Sabtu, Minggu
	daysSlice1[0] = "Sabtu baru"
	daysSlice1[1] = "Minggu baru"
	fmt.Println(daysSlice1)
	fmt.Println(days)

	// method append slice
	// append akan membuat array baru jika sudah melebihi kapasitasnya
	// jika belum penuh kapasitasnya tidak akan membuat array baru

	daysSlice2 := append(daysSlice1, "Libur baru")
	daysSlice2[0] = "Sabtu lama"
	fmt.Println(daysSlice1)
	fmt.Println(daysSlice2)
	fmt.Println(days)

	// method make slice

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "king"
	newSlice[1] = "king"
	// newSlice[2] = "king" harusnya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "king")

	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "lana"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// method copy slice

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	fmt.Println(toSlice)

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// perbedaan array & slice
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
