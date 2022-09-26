package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Gerar strings aleatórias de acordo com o tamanho e amostra passados como parâmetros
func Test_Random(t *testing.T) {
	var stringTeste string
	for j := 0; j < 5; j++ {
		for _, v := range m {
			stringTeste = Random(j, v)
		}
		require.Equal(t, len(stringTeste), j)
		if j == 0 {
			require.Empty(t, stringTeste)
		} else {
			require.NotEmpty(t, stringTeste)
		}
	}

}

// Gerar string aleatória de acordo com o tamanho da string passado como parâmetro
func Test_RandomString(t *testing.T) {
	var stringTeste string
	for j := 0; j < 5; j++ {
		stringTeste = RandomString(j)
		require.Equal(t, len(stringTeste), j)
		if j == 0 {
			require.Empty(t, stringTeste)
		} else {
			require.NotEmpty(t, stringTeste)
		}
	}
}

// // gerar email aleatório
// func RandomEmail() string {
// 	email := Random(10, alphabet) + "@" + Random(5, alphabet) + ".com"
// 	return email
// }

// // gerar senha aleatória com n dígitos
// func RandomSenha(n int) string {
// 	senha := Random(n, lettersWithSimbols)
// 	return senha
// }

// // gerar nome aleatório com n dígitos
// func RandomName(n int) string {
// 	name := Random(n, alphabetWithCapital)
// 	return name
// }

// // gerar número aleatório com n dígitos
// func RandomNumber(n int) int {
// 	number, err := strconv.Atoi(Random(n, numbers))
// 	if err != nil {
// 		panic(err)
// 	}
// 	return number
// }

// // gerar número aletório com n digitos e d casas decimais
// func RandomFloatString(n, d int) string {
// 	var stringOfNumber string
// 	if n > 1 {
// 		primeiroDigito := Random(1, numbers)
// 		if primeiroDigito != "0" {
// 			stringOfNumber = primeiroDigito + Random(n-1, numbers) + "." + Random(d, numbers)
// 		} else {
// 			stringOfNumber = Random(n-1, numbers) + "." + Random(d, numbers)
// 		}
// 	} else {
// 		stringOfNumber = Random(n, numbers) + "." + Random(d, numbers)
// 	}
// 	return stringOfNumber
// }

// // gerar número aletório com n digitos e d casas decimais
// func RandomFloat(n, d int) float64 {
// 	stringOfNumber := Random(n, numbers) + "." + Random(d, numbers)
// 	number, err := strconv.ParseFloat(stringOfNumber, 64)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return number
// }

// // gerar ano aleatório com 4 dígitos até 2029
// func RandomYear() int {
// 	number, err := strconv.Atoi("202" + Random(1, numbers))
// 	if err != nil {
// 		panic(err)
// 	}
// 	return number
// }

// // gerar mês aleatório
// func RandomMonth() int {
// 	primeiroDigito := Random(1, binaries)
// 	var segundoDigito string
// 	if primeiroDigito == "0" {
// 		segundoDigito = Random(1, numbersLessZero)
// 	} else {
// 		segundoDigito = Random(1, zeroTwo)
// 	}
// 	number, err := strconv.Atoi(primeiroDigito + segundoDigito)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return number
// }

// // gerar dia aleatório
// func RandomDay() int {
// 	primeiroDigito := Random(1, zeroThree)
// 	var segundoDigito string
// 	switch {
// 	case primeiroDigito == "0":
// 		segundoDigito = Random(1, numbersLessZero)
// 	case primeiroDigito == "3":
// 		segundoDigito = Random(1, binaries)
// 	default:
// 		segundoDigito = Random(1, numbers)
// 	}
// 	number, err := strconv.Atoi(primeiroDigito + segundoDigito)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return number
// }

// // gerar hora aleatória
// func RandomHour() int {
// 	primeiroDigito := Random(1, zeroTwo)
// 	var segundoDigito string
// 	if primeiroDigito == "0" {
// 		segundoDigito = Random(1, numbersLessZero)
// 	} else {
// 		segundoDigito = Random(1, zeroTwo)
// 	}
// 	number, err := strconv.Atoi(primeiroDigito + segundoDigito)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return number
// }

// // gerar hora aleatória
// func RandomMinutes() int {
// 	primeiroDigito := Random(1, zeroSix)
// 	var segundoDigito string
// 	switch {
// 	case primeiroDigito == "0":
// 		segundoDigito = Random(1, numbersLessZero)
// 	case primeiroDigito == "6":
// 		segundoDigito = "0"
// 	default:
// 		segundoDigito = Random(1, numbers)
// 	}
// 	number, err := strconv.Atoi(primeiroDigito + segundoDigito)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return number
// }

// // Gerar data aleatória no formato time.Time
// func RandomDate() time.Time {
// 	timeZone, _ := time.LoadLocation("America/Sao_Paulo")
// 	return time.Date(RandomYear(), time.Month(RandomMonth()), RandomDay(), 0, 0, 0, 0, timeZone)

// }

// // Gerar data aleatória no formato time.Time completa
// func RandomFullDate() time.Time {
// 	timeZone, _ := time.LoadLocation("America/Sao_Paulo")
// 	return time.Date(RandomYear(), time.Month(RandomMonth()), RandomDay(), RandomHour(), RandomMinutes(), RandomMinutes(), 0, timeZone)
// }

// // func main() {
// // 	fmt.Println(RandomDate())
// // 	fmt.Println(RandomFullDate())
// // 	fmt.Println(RandomEmail())
// // 	fmt.Println(RandomName(15))
// // 	fmt.Println(RandomSenha(15))
// // 	fmt.Println(RandomFloat(2, 2))
// // 	fmt.Println(RandomFloatString(1, 2))
// // 	fmt.Println(RandomFloatString(2, 2))
// // 	fmt.Println(RandomFloatString(3, 2))
// // 	fmt.Println(RandomFloatString(4, 2))
// // }
