package main

import "fmt"

type Barang struct {
	nama       string
	stok       int
	hargaModal int
}

func tampilkanBarang(data []Barang) {
	fmt.Println("\nDaftar Barang:")
	for i := 0; i < len(data); i++ {
		fmt.Println(
			i, "-", data[i].nama,
			"| Stok:", data[i].stok,
			"| Harga Modal:", data[i].hargaModal,
		)
	}
}

func main() {
	var inventory []Barang
	var menu int

	for {
		fmt.Println("=== APLIKASI TRACK INVENTORY ===")
		fmt.Println("1. Input Barang Baru")
		fmt.Println("2. Tambah Stok Barang")
		fmt.Println("3. Kurangi Stok Barang")
		fmt.Println("4. Lihat Daftar Barang")
		fmt.Println("5. Hitung Total Nilai Inventaris")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&menu)

		if menu == 0 {
			fmt.Println("Program selesai.")
			break
		}

		switch menu {

		case 1:
			for {
				var b Barang

				fmt.Print("Nama barang: ")
				fmt.Scan(&b.nama)

				fmt.Print("Stok awal: ")
				fmt.Scan(&b.stok)

				fmt.Print("Harga modal: ")
				fmt.Scan(&b.hargaModal)

				inventory = append(inventory, b)

				var lanjut string
				fmt.Print("Tambah barang lagi? (y/n): ")
				fmt.Scan(&lanjut)

				if lanjut != "y" {
					break
				}
			}

		case 2:
			if len(inventory) == 0 {
				fmt.Println("Belum ada barang.")
				break
			}

			tampilkanBarang(inventory)

			var index, jumlah int
			fmt.Print("Pilih index barang: ")
			fmt.Scan(&index)

			if index < 0 || index >= len(inventory) {
				fmt.Println("Index tidak valid!")
				break
			}

			fmt.Print("Jumlah stok masuk: ")
			fmt.Scan(&jumlah)

			if jumlah > 0 {
				inventory[index].stok += jumlah
			} else {
				fmt.Println("Jumlah harus lebih dari 0.")
			}

		case 3:
			if len(inventory) == 0 {
				fmt.Println("Belum ada barang.")
				break
			}

			tampilkanBarang(inventory)

			var index, keluar int
			fmt.Print("Pilih index barang: ")
			fmt.Scan(&index)

			if index < 0 || index >= len(inventory) {
				fmt.Println("Index tidak valid!")
				break
			}

			fmt.Print("Jumlah barang keluar: ")
			fmt.Scan(&keluar)

			if keluar <= 0 {
				fmt.Println("Jumlah harus lebih dari 0.")
			} else if inventory[index].stok >= keluar {
				inventory[index].stok -= keluar
			} else {
				fmt.Println("Stok tidak mencukupi!")
			}

		case 4:
			if len(inventory) == 0 {
				fmt.Println("Inventory kosong.")
			} else {
				tampilkanBarang(inventory)
			}

		case 5:
			total := 0
			for i := 0; i < len(inventory); i++ {
				total += inventory[i].stok * inventory[i].hargaModal
			}
			fmt.Println("Total nilai inventaris:", total)

		default:
			fmt.Println("Menu tidak tersedia.")
		}
	}
}
