package graph

import "fmt"

func main() {
	g := NewGraph(5)

	// Tambah edge arah: from -> to
	g.AddEdge(0, 1, 10)
	g.AddEdge(0, 2, 5)
	g.AddEdge(1, 2, 3)
	g.AddEdge(2, 3, 2)
	g.AddEdge(3, 4, 7)

	// Cek degree
	fmt.Println("OutDegree dari 0:", g.OutDegree(0))     // 2
	fmt.Println("InDegree dari 2:", g.InDegree(2))       // 2 (dari 0 dan 1)
	fmt.Println("TotalDegree dari 2:", g.TotalDegree(2)) // 3

	// Cek edge
	fmt.Println("Apakah ada edge dari 0 ke 2?", g.HasEdge(0, 2)) // true
	fmt.Println("Apakah ada edge dari 2 ke 0?", g.HasEdge(2, 0)) // false

	// Tetangga dari node 0
	fmt.Println("Tetangga dari 0:", g.GetNeighboor(0)) // [2 1] (karena AddFirst, urutan terbalik)

	// Hapus edge 0 -> 2
	g.RemoveEdge(0, 2)

	// Cek lagi
	fmt.Println("Apakah ada edge dari 0 ke 2 setelah dihapus?", g.HasEdge(0, 2)) // false
	fmt.Println("OutDegree dari 0 setelah dihapus:", g.OutDegree(0))             // 1
}
