package main 
import "fmt"
func main(){
	var jenis,nominal,transaksi,tagihan,admin,total,pakai,pajak float64
	var keterangan,status string
	var tanggal bool
	fmt.Println("=====kalkulator listrik=====\n")
	
	status = "y"
	
	for status == "y"{
		fmt.Printf("Transaksi ke-%v\n", transaksi+1)
		fmt.Print("pilih jenis transaksi (1=token, 2=pascabayar):")
		fmt.Scan(&jenis)
		switch jenis{
			case 1:
			fmt.Print("masukkan nominal:")
			fmt.Scan(&nominal)
			
			if nominal >= 100000{
				keterangan = "nominal sangat besar"
				admin = 5000
				pajak = 0.07
			}else if nominal >= 50000{
				keterangan = "nominal besar"
				admin = 3000
				pajak = 0.05
			}else if nominal >= 20000{
				keterangan = "nominal sedang"
				admin = 2000
				pajak = 0.03
			}else{
				keterangan = "nominal  kecil"
				admin = 1000
				pajak = 0.02
			}
			total = nominal + admin + (nominal * pajak)
			fmt.Printf("keterangan: %v\nadmin: Rp %v \npajak: %.0f%% \ntotal: Rp %v\n\n", keterangan, admin, pajak * 100, total)
			
			
			
			case 2:
			fmt.Print("masukkan pemakaian (kwh):")
			fmt.Scan(&pakai)
			fmt.Print("bayar lewat tanggal 20? (true/false) :")
			fmt.Scan(&tanggal)
			
			if pakai <= 50{
				pakai = pakai * 1500
			}else if pakai <= 100{
				pakai = ((pakai - 50)*2000) + 75000
			}else if pakai <= 150{
				pakai = ((pakai - 100)*2500) + 175000
			}else{
				pakai = ((pakai - 150)*3000) + 300000
			}
			
			if tanggal == true{
				pakai = pakai + 5000
			}
			
			total = pakai
			fmt.Println("total : Rp", total, "\n")
			
			
		}
		
		transaksi++
		tagihan = tagihan + total
		fmt.Print("hitung lagi? (y/n):")
		fmt.Scan(&status)
		fmt.Print("\n")
	}
	
	fmt.Printf("=====rekap====\ntotal transaksi:%v\ntotal tagihan : Rp %.0f", transaksi, tagihan)
}