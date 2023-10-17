package main

import (
	"encoding/csv"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

type Card struct {
	Number string
	Cpf    string
	Holder string
	Unit   string
}

func parseRows(rows [][]string) []Card {
	var cards []Card

	var firstLine = true

	for _, row := range rows {
		if firstLine {
			firstLine = false
			continue
		}

		var holder string

		if len(row[3]) < 35 {
			holder = row[3]
		} else {
			holder = row[3][:35]
		}

		cards = append(cards, Card{
			Number: row[1],
			Cpf:    strings.Replace(strings.Replace(row[2], ".", "", -1), "-", "", -1),
			Holder: strings.ToUpper(holder),
			Unit:   row[4],
		})
	}

	return cards
}

func main() {
	file, err := excelize.OpenFile("cartoes.xlsx")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	rows, err := file.GetRows("Results")
	if err != nil {
		panic(err)
	}

	cards := parseRows(rows)

	f, err := os.Create("_cartoes_.csv")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	writer := csv.NewWriter(f)
	writer.Comma = ';'
	defer writer.Flush()

	header := []string{"Numero de Serie", "CPF", "Valor da Carga", "Observacao"}
	writer.Write(header)

	for _, card := range cards {
		row := []string{card.Number, card.Cpf, " ", card.Holder}
		writer.Write(row)
	}
}
