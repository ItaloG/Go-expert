package main

// func SomaInteiro(m map[string]int) int {
// 	var soma int
// 	for _, v := range m {
// 		soma += v
// 	}
// 	return soma
// }

// func SomaFloat(m map[string]float64) float64 {
// 	var soma float64
// 	for _, v := range m {
// 		soma += v
// 	}
// 	return soma
// }

type MyNumber int

type Number interface {
	~int | ~float64
}

// func Soma[T int | float64](m map[string]T) T {
func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Italo": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Italo": 100.20, "João": 2000.3, "Maria": 30.00}
	m3 := map[string]MyNumber{"Italo": 1000, "João": 2000, "Maria": 3000}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10))
	// println(SomaInteiro(m))
	// println(SomaFloat(m2))
}
