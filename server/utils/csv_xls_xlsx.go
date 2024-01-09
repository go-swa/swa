package utils

import (
	"encoding/csv"
	"github.com/extrame/xls"
	"github.com/pkg/errors"
	"github.com/tealeg/xlsx"
	"os"
)

func ReadCsv(filePath string) (res [][]string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return res, errors.Errorf("打开csv文件(%s)错误,err:%v", filePath, err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true
	record, read_err := reader.ReadAll()
	if read_err != nil {
		return res, errors.Errorf("读csv(%s)错误,err:%v", filePath, err)
	}
	return record, err
}

func ReadXls(filePath string) (res [][]string, err error) {
	if xlFile, err := xls.Open(filePath, "utf-8"); err == nil {

		sheet := xlFile.GetSheet(0)
		if sheet.MaxRow != 0 {
			temp := make([][]string, sheet.MaxRow+1)
			for i := 0; i <= int(sheet.MaxRow); i++ {
				row := sheet.Row(i)
				data := make([]string, 0)
				if row.LastCol() > 0 {
					for j := 0; j < row.LastCol(); j++ {
						col := row.Col(j)
						data = append(data, col)
					}
					temp[i] = data
				}
			}
			res = append(res, temp...)
		}
	} else {
		return res, errors.Errorf("打开xls文件(%s)错误,err:%v", filePath, err)
	}
	return res, err
}

func ReadXlsx(filePath string) (res [][]string, err error) {
	if xlFile, err := xlsx.OpenFile(filePath); err == nil {
		for index, sheet := range xlFile.Sheets {
			if index == 0 {
				temp := make([][]string, len(sheet.Rows))
				for k, row := range sheet.Rows {
					var data []string
					for _, cell := range row.Cells {
						data = append(data, cell.Value)
					}
					temp[k] = data
				}
				res = append(res, temp...)
			}
		}
	} else {
		return res, errors.Errorf("打开xlsx文件(%s)错误,err:%v", filePath, err)
	}
	return res, err
}

func ValidUTF8(buf []byte) bool {
	nBytes := 0
	for i := 0; i < len(buf); i++ {
		if nBytes == 0 {
			if (buf[i] & 0x80) != 0 {
				for (buf[i] & 0x80) != 0 {
					buf[i] <<= 1
					nBytes++
				}

				if nBytes < 2 || nBytes > 6 {
					return false
				}

				nBytes--
			}
		} else { // 处理多字节字符
			if buf[i]&0xc0 != 0x80 {
				return false
			}
			nBytes--
		}
	}
	return nBytes == 0
}
