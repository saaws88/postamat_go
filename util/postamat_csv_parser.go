package util

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"saaws88/model"
)

func ParseCsv(filepath string) [][]string {

	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Unable to read input file"+filepath, err)
	}
	defer f.Close()

	csvReader := *csv.NewReader(f)
	csvReader.Read()
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to pars file as CSV for "+filepath, err)
	}

	return records

}

func BuildFromLine(line []string) model.Postamat {

	var post model.Postamat = model.Postamat{
		Number:                     int16(ExtractInt(line[0])),
		Owner:                      line[1],
		OfficeCode:                 line[2],
		Region:                     line[3],
		Address:                    line[4],
		FirmwareStatus:             line[5],
		ActivationBlocker:          line[6],
		FirmwareVersion:            line[7],
		TerminalId:                 ExtractInt(line[8]),
		PaymentOn:                  ExtractBool(line[9]),
		OnMaintenance:              ExtractBool(line[10]),
		MaintenanceAgreementNumber: line[11],
		Visible:                    true,
		PostamatCellsNumber:        line[12],
	}

	return post

}

func ExtractNumber(column string) int {

	i, err := strconv.Atoi(column)
	if err != nil {
		fmt.Println("Неверный формат номера ", column)
	}

	return i

}

func ExtractInt(column string) int {
	i, err := strconv.Atoi(column)
	if err != nil {
		return 0
	}

	return i
}

func ExtractBool(column string) bool {

	if column == "TRUE" {
		return true
	} else {
		return false
	}

}
