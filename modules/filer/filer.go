package filer

import (
	"errors"
	"fmt"
	"os"
	"time"

	local_types "github.com/heikkilavaste/go_subnet/modules/types"
	"github.com/xuri/excelize/v2"
	"gopkg.in/yaml.v3"
)

func WriteToConsole(Sets []local_types.AddressSet) {
	for _, s := range Sets {
		fmt.Println("Subnet : ", s.Subnet)
		fmt.Println("Gateway IP : ", s.GW)
		fmt.Println("Broadcast address is : ", s.BC)
		fmt.Println("First usable address is : ", s.First)
		fmt.Println("Last usable address is : ", s.Last)
	}
}

func WriteToYaml(Sets []local_types.AddressSet, to string) bool {
	output, _ := os.OpenFile(to+".yaml", os.O_RDWR|os.O_CREATE, os.ModePerm)
	enc := yaml.NewEncoder(output)
	_ = enc.Encode(Sets)
	return true

}

func WriteToCSV(Sets []local_types.AddressSet, file string, sheetname string) bool {
	filename := file + ".xlsx"
	RowIndex := 2
	f := excelize.NewFile()
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		f = excelize.NewFile()
	} else {
		f, _ = excelize.OpenFile(filename)
	}

	sheetName := sheetname
	index, err := f.NewSheet(sheetName)
	if err != nil {
		fmt.Println(err)
	}
	f.SetActiveSheet(index)
	headers := []string{"subnet", "gateway", "Broadcast ip", "first ip", "last ip"}
	for i, header := range headers {
		cell := fmt.Sprintf("%s1", string('A'+i))
		f.SetCellValue(sheetName, cell, header)
	}
	for rowIndex, d := range Sets {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", rowIndex+RowIndex), d.Subnet)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", rowIndex+RowIndex), d.GW)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", rowIndex+RowIndex), d.BC)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", rowIndex+RowIndex), d.First)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", rowIndex+RowIndex), d.Last)
	}

	f.SaveAs(filename, excelize.Options{})

	return true
}

func GetSheetName() string {
	out := time.Now()
	h, m, _ := out.Clock()
	date := fmt.Sprintf("%d-%d", h, m)
	return "subnets-" + date
}

func GetFileName() string {
	out := time.Now()
	Y, M, D := out.Date()
	date := fmt.Sprintf("%d_%d_%d", D, M, Y)
	return "reports/legacy_voice_report_" + date + ".xlsx"
}
