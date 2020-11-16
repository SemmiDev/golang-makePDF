package main
import ("github.com/jung-kurt/gofpdf";"log")

/**
* @param1 : Orientasi potrait(P) or landscape (L)
* @param2 : Satuan doc milimeter (mm)
* @param3 : Ukuran doc, ex A4 A5 etc
* @param4 : Folder font
* addPage : Membuat halaman baru (default tidak mempunyai hal)
* setFont : fontFamily, fontStyle, fontSize
* text    :
* 		@param1 : x mm dari kiri
* 		@param1 : y mm dari atas
* image   :
* 		@param1 ke-1 : path image.
* 		@param2 ke-2 : x offset. Nilai 56 artinya 56mm dari kiri.
* 		@param3 ke-3 : y offset. Nilai 40 artinya 40mm dari atas.
* 		@param4 ke-4 : width gambar. Jika diisi dengan nilai lebih dari 0 maka
* 		@param4 ke-4 : width gambar. Jika diisi dengan nilai lebih dari 0 maka
* 		 	gambar akan di-resize secara proporsional sesuai angka. Jika di-isi 0 ,
* 		 	maka gambar akan muncul sesuai ukuran aslinya.
*		@param5      : height gambar
*/

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("arial", "B", 14)
	pdf.Text(40,10,"Hello Sam")
	pdf.Image("./sam.jpg", 56,40,100,0,false,"",0,"")
	err := pdf.OutputFileAndClose("./sam.pdf")
	if err != nil {
		log.Println("ERROR", err.Error())
	}
}