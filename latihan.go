
package main
import "fmt"
const NMax int = 10
type pemain struct{
	nama string
	gol, assist int
}
type tabBola [NMax] pemain
func main() {
	var A, temp tabBola
	var i, gol, assist, j, idx int
	var namaPemain string
	for i = 0; i < 3; i++{
		fmt.Scan(&namaPemain, &gol, &assist)
		A[i].nama = namaPemain
		A[i].gol = gol
		A[i].assist = assist
	}
	for i = 0; i < 3; i++{
		gol = A[i].gol
		for j = i+1; j < 3; j++{
			if gol > A[i].gol{
				idx = j
			}
		}
		temp[0] = A[i]
		A[i] = A[idx]
		A[idx] = temp[0]
	}
	fmt.Println(A[idx].nama)
	for i = 0; i < 3; i++{
		assist = A[0].assist
		for j = i+1; j < 3; j++{
			if assist > A[i].assist{
				idx = j
			}
		}
		temp[0] = A[i]
		A[i] = A[idx]
		A[idx] = temp[0]
	}
	fmt.Println(A[idx].nama)
}