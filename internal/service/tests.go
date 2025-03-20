package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Faturamento struct {
	Dia   int     `json:"dia"`
	Valor float64 `json:"valor"`
}

func N1(log *strings.Builder) {
	fmt.Fprintln(log, "N1")
	INDICE := 13
	SOMA := 0
	K := 0

	for K < INDICE {
		K += 1
		SOMA += K
	}

	fmt.Fprintf(log, "A soma dos %d primeiros números é %d.\n", INDICE, SOMA)
}

func N2(log *strings.Builder) {
	fmt.Fprintln(log, "N2")
	numeros := []int{8, 10, 21}

	for _, num := range numeros {
		if pertenceFibonacci(num) {
			fmt.Fprintf(log, "%d pertence a sequência de Fibonacci.\n", num)
		} else {
			fmt.Fprintf(log, "%d não pertence a sequência de Fibonacci.\n", num)
		}
	}
}

func pertenceFibonacci(n int) bool {
	a, b := 0, 1
	for a <= n {
		if a == n {
			return true
		}
		a, b = b, a+b
	}
	return false
}

func N3(log *strings.Builder) {
	fmt.Fprintln(log, "N3")
	faturamentos := carregarDados("dados.json")

	menor, maior := calcularMenorMaiorFaturamento(faturamentos)
	diasAcimaMedia := calcularDiasAcimaMedia(faturamentos)

	fmt.Fprintf(log, "Menor faturamento diário: %.2f\n", menor)
	fmt.Fprintf(log, "Maior faturamento diário: %.2f\n", maior)
	fmt.Fprintf(log, "Número de dias com faturamento acima da média: %d\n", diasAcimaMedia)
}

func carregarDados(caminho string) []Faturamento {
	arquivo, err := os.Open(caminho)
	if err != nil {
		log.Fatal(err)
	}
	defer arquivo.Close()

	bytes, err := io.ReadAll(arquivo)
	if err != nil {
		log.Fatal(err)
	}

	var faturamentos []Faturamento
	if err := json.Unmarshal(bytes, &faturamentos); err != nil {
		log.Fatal(err)
	}

	return faturamentos
}

func calcularMenorMaiorFaturamento(faturamentos []Faturamento) (float64, float64) {
	var menor, maior float64
	inicializado := false

	for _, fat := range faturamentos {
		if fat.Valor > 0 {
			if !inicializado {
				menor, maior = fat.Valor, fat.Valor
				inicializado = true
			} else {
				if fat.Valor < menor {
					menor = fat.Valor
				}
				if fat.Valor > maior {
					maior = fat.Valor
				}
			}
		}
	}

	return menor, maior
}

func calcularDiasAcimaMedia(faturamentos []Faturamento) int {
	var soma float64
	var diasComFaturamento int

	for _, fat := range faturamentos {
		if fat.Valor > 0 {
			soma += fat.Valor
			diasComFaturamento++
		}
	}

	if diasComFaturamento == 0 {
		return 0
	}

	media := soma / float64(diasComFaturamento)
	diasAcimaMedia := 0

	for _, fat := range faturamentos {
		if fat.Valor > media {
			diasAcimaMedia++
		}
	}

	return diasAcimaMedia
}

func N4(log *strings.Builder) {
	fmt.Fprintln(log, "N4")
	faturamentos := []float64{67836.43, 36678.66, 29229.88, 27165.48, 19849.53}
	estados := []string{"SP", "RJ", "MG", "ES", "Outros"}

	total := calcularTotal(faturamentos)

	fmt.Fprintln(log, "Percentual de representação por estado:")
	for i, valor := range faturamentos {
		percentual := (valor / total) * 100
		fmt.Fprintf(log, "%s: %.2f%%\n", estados[i], percentual)
	}
}

func calcularTotal(valores []float64) float64 {
	total := 0.0
	for _, v := range valores {
		total += v
	}
	return total
}

func N5(log *strings.Builder) {
	fmt.Fprintln(log, "N5")
	str := "Target Sistemas"
	invertida := inverterString(str)

	fmt.Fprintf(log, "Original: %s\nInvertida: %s\n", str, invertida)
}

func inverterString(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}
