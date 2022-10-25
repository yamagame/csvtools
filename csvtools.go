package csvtools

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
	"sort"
)

func Read(reader io.Reader) ([][]string, error) {
	r := csv.NewReader(reader)
	r.FieldsPerRecord = -1
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func Dump(records [][]string) {
	writer := csv.NewWriter(os.Stdout)
	writer.WriteAll(records)
}

func MaxCol(records [][]string) int {
	max := 0
	for _, col := range records {
		if max < len(col) {
			max = len(col)
		}
	}
	return max
}

func MaxRow(records [][]string) int {
	return len(records)
}

func Transpose(records [][]string) [][]string {
	maxRow := MaxRow(records)
	maxCol := MaxCol(records)
	retval := make([][]string, maxCol)
	for y := 0; y < maxCol; y++ {
		r := make([]string, maxRow)
		for x := 0; x < maxRow; x++ {
			if len(records) > x && len(records[x]) > y {
				r[x] = records[x][y]
			} else {
				r[x] = ""
			}
		}
		retval[y] = r
	}
	return retval
}

func RowHash(records [][]string, key int, value int) map[string]string {
	retval := make(map[string]string, MaxRow(records))
	for _, record := range records {
		if len(record) > value {
			retval[record[key]] = record[value]
		} else {
			retval[record[key]] = ""
		}
	}
	return retval
}

func ColHash(records [][]string, key int, value int) map[string]string {
	retval := make(map[string]string, MaxCol(records))
	if len(records) > key {
		for x, field := range records[key] {
			if len(records[value]) > x {
				retval[field] = records[value][x]
			} else {
				retval[field] = ""
			}
		}
	}
	return retval
}

func Json(data interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return string(jsonData)
}

func Trim(records [][]string, col int, row int, colsize int, rowsize int) [][]string {
	maxrow := MaxRow(records)
	maxcol := MaxCol(records)
	if rowsize <= 0 {
		rowsize = maxrow - row
	}
	if colsize <= 0 {
		colsize = maxcol - col
	}
	retval := make([][]string, rowsize)
	for r := 0; r < rowsize; r++ {
		record := make([]string, colsize)
		retval[r] = record
		if row+r < maxrow {
			for c := 0; c < colsize; c++ {
				if col+c < maxcol {
					retval[r][c] = records[row+r][col+c]
				} else {
					retval[r][c] = ""
				}
			}
		} else {
			for c := 0; c < colsize; c++ {
				retval[r][c] = ""
			}
		}
	}
	return retval
}

func Sort(records [][]string, key int) [][]string {
	maxrow := MaxRow(records)
	type Record struct {
		key    string
		record []string
	}
	// keyで並べ替え
	result := make([]Record, maxrow)
	for i, record := range records {
		result[i] = Record{key: record[key], record: record}
	}
	sort.SliceStable(result, func(i, j int) bool { return result[i].key < result[j].key })
	// 並べ替えたレコードを２次元配列に戻す
	retval := make([][]string, maxrow)
	for i, record := range result {
		retval[i] = record.record
	}
	return retval
}
