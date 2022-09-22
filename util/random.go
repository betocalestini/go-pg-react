package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const lettersWithSimbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ/;.,~][!@#$%*()_+}{:?><|1234567890"
const alphabet = "abcdefghijklmnopqrstuvwxyz"
const alphabetWithCapital = "abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"

// Gerar strings aleatórias de acordo com o tamanho e amostra passados como parâmetros
func Random(n int, amostra string) string {
	tempoAleatorio := rand.NewSource(time.Now().UnixNano())
	numeroAleatorio := rand.New(tempoAleatorio)
	tamanhoDaAmostra := len(amostra)
	textoFinal := make([]byte, n)
	for i := range textoFinal {
		textoFinal[i] = amostra[numeroAleatorio.Intn(tamanhoDaAmostra)]
	}
	return string(textoFinal)
}

// Gerar string aleatória de acordo com o tamanho da string passado como parâmetro
func RandomString(number int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < number; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// gerar email aleatório
func RandomEmail() string {
	email := Random(10, alphabet) + "@" + Random(5, alphabet) + ".com"
	return email
}

// gerar senha aleatória com n dígitos
func RandomSenha(n int) string {
	senha := Random(n, lettersWithSimbols)
	return senha
}

// gerar nome aleatório com n dígitos
func RandomName(n int) string {
	name := Random(n, alphabetWithCapital)
	return name
}

// gerar número aleatório com n dígitos
func RandomNumber(n int) int {
	number, err := strconv.Atoi(Random(n, numbers))
	if err != nil {
		panic(err)
	}
	return number
}
