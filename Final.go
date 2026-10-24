package main
import "fmt"
const NMAX_AKUN int = 100
const NMAX_PERANGKAT int = 50

type Perangkat struct {
	NamaPerangkat string
	DayaWatt      int
	DurasiJam     float64
	TotalKWh      float64
}
type TabPerangkat [NMAX_PERANGKAT]Perangkat
type Akun struct {
	Username        string
	Password        string
	DaftarPerangkat TabPerangkat
	JumlahPerangkat int
}
type TabAkun [NMAX_AKUN]Akun

// I.S. nAkun terdefinisi dan jumlah akun belum melebihi kapasitas maksimum
// F.S. Data akun baru tersimpan pada array TA dan nAkun bertambah 1
func Register(TA *TabAkun, nAkun *int) {
	var u, p string
	var i int
	i = *nAkun
	if *nAkun >= NMAX_AKUN {
		fmt.Println("\n[ERROR] Pendaftaran gagal! Kapasitas memori akun sudah penuh.")
	} else {
		fmt.Print("Masukkan Username Baru (Tanpa Spasi): ")
		fmt.Scan(&u)
		fmt.Print("Masukkan Password Baru (Tanpa Spasi): ")
		fmt.Scan(&p)
		TA[i].Username = u
		TA[i].Password = p
		TA[i].JumlahPerangkat = 0
		*nAkun = *nAkun + 1
		fmt.Println("\nAkun berhasil terdaftar! Silakan login.")
	}
}

// I.S. TA terdefinisi, nAkun berisi jumlah akun yang tersimpan
// F.S. Jika username dan password ditemukan maka user berhasil login,
//      idxAkun berisi indeks akun yang login dan isLoggedIn bernilai true
func Login(TA TabAkun, nAkun int, idxAkun *int, isLoggedIn *bool) {
	var u, p string
	var i int
	var ketemu bool
	fmt.Print("Masukkan Username: ")
	fmt.Scan(&u)
	fmt.Print("Masukkan Password: ")
	fmt.Scan(&p)
	i = 0
	ketemu = false
	for i < nAkun && !ketemu {
		if TA[i].Username == u && TA[i].Password == p {
			ketemu = true
			*idxAkun = i
			*isLoggedIn = true
		} else {
			i = i + 1
		}
	}
	if ketemu {
		fmt.Printf("\n[SUKSES] Selamat datang, %s!\n", u)
	} else {
		fmt.Println("\n[ERROR] Username atau Password salah.")
	}
}
// =========================================================================
// SUBPROGRAM: PENGELOLAAN DATA PERANGKAT (SETELAH LOGIN)
// =========================================================================

// I.S. User sudah login dan data akun terdefinisi
// F.S. Data perangkat baru tersimpan pada akun yang aktif atau proses dihentikan
func TambahPerangkat(TA *TabAkun, idxAkun int) {
	var lanjut bool
	var nama string
	var slotBaru, watt int
	var jam float64
	lanjut = true
	for lanjut {
		slotBaru = TA[idxAkun].JumlahPerangkat
		if slotBaru >= NMAX_PERANGKAT {
			fmt.Println("\n[PERINGATAN] Memori penuh! Tidak bisa menambah perangkat lagi.")
			lanjut = false
		} else {
			fmt.Print("\nMasukkan Nama Perangkat (Ketik 'stop' untuk kembali): ")
			fmt.Scan(&nama)
			if nama == "stop" {
				lanjut = false
			} else {
				fmt.Print("Masukkan Daya Perangkat (Watt): ")
				fmt.Scan(&watt)
				fmt.Print("Masukkan Durasi Penggunaan (Jam/Hari): ")
				fmt.Scan(&jam)
				TA[idxAkun].DaftarPerangkat[slotBaru].NamaPerangkat = nama
				TA[idxAkun].DaftarPerangkat[slotBaru].DayaWatt = watt
				TA[idxAkun].DaftarPerangkat[slotBaru].DurasiJam = jam
				TA[idxAkun].DaftarPerangkat[slotBaru].TotalKWh = (float64(watt) * jam) / 1000.0
				TA[idxAkun].JumlahPerangkat = TA[idxAkun].JumlahPerangkat + 1
				fmt.Println("[OK] Data perangkat berhasil dicatat.")
			}
		}
	}
}
func TampilPerangkat(TA TabAkun, idxAkun int) {
	var n, i int
	var p Perangkat
	n = TA[idxAkun].JumlahPerangkat
	if n == 0 {
		fmt.Println("\n[INFO] Belum ada data perangkat.")
	} else {
		fmt.Printf("\n=== Daftar Perangkat Listrik Akun: %s ===\n", TA[idxAkun].Username)
		for i = 0; i < n; i++ {
			p = TA[idxAkun].DaftarPerangkat[i]
			fmt.Printf("%d. Nama Perangkat: %s\n", i+1, p.NamaPerangkat)
			fmt.Printf("   Daya Listrik  : %d Watt\n", p.DayaWatt)
			fmt.Printf("   Durasi Pakai  : %.2f Jam/Hari\n", p.DurasiJam)
			fmt.Printf("   Total Konsumsi: %.4f kWh\n", p.TotalKWh)
			fmt.Println("   -----------------------------------")
		}
	}
}

// =========================================================================
// SUBPROGRAM: ALGORITMA PENGURUTAN (SORTING)
// =========================================================================

// I.S. Data perangkat pada akun aktif terdefinisi
// F.S. Data perangkat terurut menaik berdasarkan NamaPerangkat
func InsertionSortNama(TA *TabAkun, idxAkun int) {
	var n, i, j int
	var temp Perangkat
	n = TA[idxAkun].JumlahPerangkat
	for i = 1; i < n; i = i + 1 {
		temp = TA[idxAkun].DaftarPerangkat[i]
		j = i - 1
		for j >= 0 &&
			TA[idxAkun].DaftarPerangkat[j].NamaPerangkat > temp.NamaPerangkat {
			TA[idxAkun].DaftarPerangkat[j+1] = TA[idxAkun].DaftarPerangkat[j]
			j = j - 1
		}
		TA[idxAkun].DaftarPerangkat[j+1] = temp
	}
}

// I.S. Data perangkat pada akun aktif terdefinisi
// F.S. Data perangkat terurut menaik berdasarkan DayaWatt
func SelectionSortDaya(TA *TabAkun, idxAkun int) {
	var n, i, j, idxMin int
	var temp Perangkat
	n = TA[idxAkun].JumlahPerangkat
	for i = 0; i < n-1; i = i + 1 {
		idxMin = i
		for j = i + 1; j < n; j = j + 1 {
			if TA[idxAkun].DaftarPerangkat[j].DayaWatt <
				TA[idxAkun].DaftarPerangkat[idxMin].DayaWatt {
				idxMin = j
			}
		}
		temp = TA[idxAkun].DaftarPerangkat[i]
		TA[idxAkun].DaftarPerangkat[i] = TA[idxAkun].DaftarPerangkat[idxMin]
		TA[idxAkun].DaftarPerangkat[idxMin] = temp
	}
}

// I.S. Data perangkat pada akun aktif terdefinisi
// F.S. Data perangkat terurut menurun (descending)
//      berdasarkan TotalKWh menggunakan Selection Sort
func SelectionSortKWhMax(TA *TabAkun, idxAkun int) {
	var n, i, j, idxMax int
	var temp Perangkat
	n = TA[idxAkun].JumlahPerangkat
	for i = 0; i < n-1; i = i + 1 {
		idxMax = i
		for j = i + 1; j < n; j = j + 1 {
			if TA[idxAkun].DaftarPerangkat[j].TotalKWh > TA[idxAkun].DaftarPerangkat[idxMax].TotalKWh {
				idxMax = j
			}
		}
		temp = TA[idxAkun].DaftarPerangkat[i]
		TA[idxAkun].DaftarPerangkat[i] = TA[idxAkun].DaftarPerangkat[idxMax]
		TA[idxAkun].DaftarPerangkat[idxMax] = temp
	}
}

// =========================================================================
// SUBPROGRAM: PENCARIAN & ANALISIS (SEQUENTIAL SEARCH)
// =========================================================================

// I.S. Data perangkat pada akun aktif terdefinisi
// F.S. Menampilkan data perangkat yang sesuai dengan keyword pencarian
//      menggunakan Sequential Search
func CariPerangkat(TA TabAkun, idxAkun int) {
	var n, i int
	var keyword string
	var ketemu bool
	var p Perangkat
	n = TA[idxAkun].JumlahPerangkat
	if n == 0 {
		fmt.Println("\n[INFO] Data kosong, tidak ada yang bisa dicari.")
	} else {
		fmt.Print("Masukkan Nama Perangkat yang Dicari: ")
		fmt.Scan(&keyword)
		i = 0
		ketemu = false
		fmt.Println("\n--- Hasil Pencarian (Sequential Search) ---")
		for i < n {
			if TA[idxAkun].DaftarPerangkat[i].NamaPerangkat == keyword {
				p = TA[idxAkun].DaftarPerangkat[i]
				fmt.Printf("Ditemukan pada Indeks ke-%d -> Watt: %d, Durasi: %.2f Jam, Total: %.4f kWh\n", i+1, p.DayaWatt, p.DurasiJam, p.TotalKWh)
				ketemu = true
			}
			i = i + 1
		}
		if !ketemu {
			fmt.Println("Maaf, perangkat dengan nama tersebut tidak ditemukan.")
		}
	}
}

// I.S. Data perangkat pada akun aktif terdefinisi
// F.S. Menampilkan estimasi biaya listrik harian dan bulanan
//      berdasarkan total penggunaan energi seluruh perangkat
func HitungEstimasiBiaya(TA TabAkun, idxAkun int) {
	var n, i int
	var totalKWh, tarifPerKWh, biayaBulanan, biayaHarian float64
	tarifPerKWh = 1500.0
	n = TA[idxAkun].JumlahPerangkat
	totalKWh = 0
	if n == 0 {
		fmt.Println("\n[INFO] Data kosong, estimasi biaya adalah Rp 0.")
	} else {
		for i = 0; i < n; i = i + 1 {
			totalKWh = totalKWh + TA[idxAkun].DaftarPerangkat[i].TotalKWh
		}
		biayaHarian = totalKWh * tarifPerKWh
		biayaBulanan = biayaHarian * 30.0
		fmt.Println("\n====================================================")
		fmt.Printf("        LAPORAN ESTIMASI BIAYA LISTRIK: %s\n", TA[idxAkun].Username)
		fmt.Println("====================================================")
		fmt.Printf("Total Perangkat Aktif : %d Perangkat\n", n)
		fmt.Printf("Total Akumulasi Daya  : %.4f kWh / Hari\n", totalKWh)
		fmt.Printf("Estimasi Biaya Harian : Rp %.2f\n", biayaHarian)
		fmt.Printf("Estimasi Biaya Bulanan: Rp %.2f\n", biayaBulanan)
		fmt.Println("====================================================")
	}
}

// =========================================================================
// MAIN PROGRAM
// =========================================================================
func main() {
	var DataAkun TabAkun
	var TotalAkun, IndexAktif int
	var StatusLogin bool
	var MenuUtama, MenuApp string
	TotalAkun = 0
	IndexAktif = -1
	StatusLogin = false
	MenuUtama = ""

	for MenuUtama != "3" {
		if !StatusLogin {
			fmt.Println("\n=====================================")
			fmt.Println("      APLIKASI POWERLOG (MP-14)      ")
			fmt.Println("=====================================")
			fmt.Println("1. Register Akun Baru")
			fmt.Println("2. Login Pengguna")
			fmt.Println("3. Keluar Aplikasi")
			fmt.Println("=====================================")
			fmt.Print("Pilih tindakan (1-3): ")
			fmt.Scan(&MenuUtama)
			if MenuUtama == "1" {
				Register(&DataAkun, &TotalAkun)
			} else if MenuUtama == "2" {
				Login(DataAkun, TotalAkun, &IndexAktif, &StatusLogin)
			}
		} else {
			MenuApp = ""
			for StatusLogin && MenuApp != "7" {
				fmt.Println("\n====================================================")
				fmt.Printf("   DASHBOARD UTAMA POWERLOG | USER: %s\n", DataAkun[IndexAktif].Username)
				fmt.Println("====================================================")
				fmt.Println("1. Tambah Data Perangkat Listrik Baru")
				fmt.Println("2. Tampilkan Perangkat (Urutan Abjad A-Z)")
				fmt.Println("3. Tampilkan Perangkat (Daya Listrik Terendah)")
				fmt.Println("4. Tampilkan Perangkat Paling Boros (kWh Tertinggi)")
				fmt.Println("5. Cari Data Perangkat Berdasarkan Nama")
				fmt.Println("6. Analisis & Estimasi Total Biaya Bulanan")
				fmt.Println("7. Logout Aplikasi")
				fmt.Println("====================================================")
				fmt.Print("Pilih menu (1-7): ")
				fmt.Scan(&MenuApp)
				if MenuApp == "1" {
					TambahPerangkat(&DataAkun, IndexAktif)
				} else if MenuApp == "2" {
					InsertionSortNama(&DataAkun, IndexAktif)
					TampilPerangkat(DataAkun, IndexAktif)
				} else if MenuApp == "3" {
					SelectionSortDaya(&DataAkun, IndexAktif)
					TampilPerangkat(DataAkun, IndexAktif)
				} else if MenuApp == "4" {
					SelectionSortKWhMax(&DataAkun, IndexAktif)
					TampilPerangkat(DataAkun, IndexAktif)
				} else if MenuApp == "5" {
					CariPerangkat(DataAkun, IndexAktif)
				} else if MenuApp == "6" {
					HitungEstimasiBiaya(DataAkun, IndexAktif)
				} else if MenuApp == "7" {
					StatusLogin = false
					IndexAktif = -1
					fmt.Println("\n[INFO] Anda telah logout. Kembali ke Menu Utama.")
				}
			}
		}
	}

	fmt.Println("\nAplikasi Ditutup. Terima Kasih telah menggunakan PowerLog!")
}