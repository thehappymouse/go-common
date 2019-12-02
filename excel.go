package utils


import (
	"encoding/csv"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"os"
	"strings"
)

// 读取批定的csv文件
func LoadCsvData(source string) ([][]string, error) {
	cntb,err := ioutil.ReadFile(source)
	if err != nil {
		return nil, err
	}
	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	return r2.ReadAll()
}
// 写入数据到指定的csv文件
func Write2CsvFile(records [][]string, filename string) error  {
	nf, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer nf.Close()

	nw := csv.NewWriter(nf)
	err = nw.WriteAll(records)
	if err != nil {
		return err
	}
	nw.Flush()
	return nil
}

// 读取Excel指定sheet的数据
func LoadExcel(xlsxFile string, sheet string) [][]string {
	xlsx, err := excelize.OpenFile(xlsxFile)
	CheckError(err)
	rows := xlsx.GetRows(sheet)
	rows = rows[1:]
	return rows
}

func LoadSheet1(xlsxFile string) [][]string {

	xlsx, err := excelize.OpenFile(xlsxFile)
	CheckError(err)
	sheet := xlsx.GetSheetName(1)
	rows := xlsx.GetRows(sheet)
	return rows
}