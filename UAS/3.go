package main

import "fmt"

const NMAX = 1024

type RekTiket struct {
	ID        int
	RootCause string
	Jam       int
	Menit     int
}

type ListTiket struct {
	info   [NMAX]RekTiket
	Jumlah int
}

func Durasi(J, M int) int {
	return (J*60 + M) * 60
}

func Sort(T *ListTiket) {
	var i, j, d int

	var temp RekTiket
	j = 1
	for j < T.Jumlah {
		i = j
		temp = T.info[j]
		d = Durasi(temp.Jam, temp.Menit)
		for i > 0 && d > Durasi(T.info[i-1].Jam, T.info[i-1].Menit) {
			T.info[i] = T.info[i-1]
			i = i - 1
		}
		T.info[i] = temp
		j = j + 1
	}
}

func main() {
	a := ListTiket{
		Jumlah: 3,
		info: [NMAX]RekTiket{
			{ID: 0, RootCause: "apaan", Jam: 2, Menit: 3},
			{ID: 0, RootCause: "apaan", Jam: 3, Menit: 4},
			{ID: 0, RootCause: "apaan", Jam: 1, Menit: 6},
		},
	}

	Sort(&a)
	for x := 0; x < a.Jumlah; x++ {
		fmt.Print(a.info[x])
	}

}
