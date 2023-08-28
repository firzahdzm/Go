package main

import "fmt"

type kontak struct{
	nama	string
	nomor	int
}

func ListKontak(data []kontak) {
	if len(data) > 0{
		fmt.Println("================")
		fmt.Println("Daftar Kontak")
		fmt.Println("================")
		for _, kontak := range data {
			fmt.Printf("Nama: %s | Nomor: %d\n", kontak.nama, kontak.nomor)
		}
	} else {
		fmt.Println("================")
		fmt.Println("Kontak Kosong")
		fmt.Println("================")
	}
}

func TambahKontak(data *[]kontak) {
	fmt.Println("================")
	fmt.Println("Tambah Kontak")
	fmt.Println("================")
	var nm 	string
	var nmr int
	fmt.Print("Nama :"); fmt.Scanln(&nm)
	fmt.Print("Nomor :"); fmt.Scanln(&nmr)
	*data = append(*data, kontak{nm, nmr})
}

func HapusKontak(data *[]kontak) {
	fmt.Println("================")
	fmt.Println("Hapus Kontak")
	fmt.Println("================")
	var nm string
	fmt.Print("Nama yang akan dihapus: ")
	fmt.Scanln(&nm)

	indexToRemove := -1
	for i, kontak := range *data {
		if kontak.nama == nm {
			indexToRemove = i
			break
		}
	}

	if indexToRemove == -1 {
		fmt.Printf("Kontak dengan nama '%s' tidak ditemukan\n", nm)
	} else {
		*data = append((*data)[:indexToRemove], (*data)[indexToRemove+1:]...)
		fmt.Printf("Kontak dengan nama '%s' telah dihapus\n", nm)
	}
}

func EditKontak(data *[]kontak){
	fmt.Println("================")
	fmt.Println("Edit Kontak")
	fmt.Println("================")
	var nm string
	var nmr, op int
	fmt.Print("Nama yang akan diedit: "); fmt.Scanln(&nm)

	indexToEdit := -1
	for i, kontak := range *data {
		if kontak.nama == nm {
			indexToEdit = i
			break
		}
	}

	if indexToEdit == -1 {
		fmt.Printf("Kontak dengan nama '%s' tidak ditemukan\n", nm)
	} else {
		fmt.Print("Edit Nama(1) atau Nomor(2): ")
		fmt.Scanln(&op)
		if op == 1 {
			fmt.Print("Nama baru: ")
			fmt.Scanln(&nm)
			(*data)[indexToEdit].nama = nm
		} else if op == 2 {
			fmt.Print("Nomor baru: ")
			fmt.Scanln(&nmr)
			(*data)[indexToEdit].nomor = nmr
		}
		fmt.Printf("Kontak dengan nama '%s' telah diubah\n", nm)
	}	
}

func main() {
	data := []kontak{}
	var op int
	for op != 5 {
		fmt.Println("1.Daftar Kontak\n2.Tambah Kontak\n3.Hapus Kontak\n4.Edit Kontak\n5.End Program")
		fmt.Print("Pilih : ")
		fmt.Scanln(&op)
		switch op {
			case 1: ListKontak(data)
			case 2: TambahKontak(&data)
			case 3: HapusKontak(&data)
			case 4: EditKontak(&data)
			case 5: fmt.Println("Program Berakhir...")
			default:fmt.Println("Pilihan tidak valid! masukan pilihan kembali...")
		}
	}		
}