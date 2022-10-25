# CSV ファイル便利ツール

## commands

### csv2json

csvファイルをjsonに変換する。
--row/--colオプションでフィールドをキーにした連想配列にすることができる。
オプション指定がない場合は２次元配列で出力される。

### transcsv

csvファイルを読み込んで列と行を入れ替えて出力する。

### trimcsv

csvファイルの一部を切り出す。

### dumpcsv

csvファイルを読み込んで出力する。特に変換はしない。

## build

```bash
$ make
```

## functions

### Reaad
csvtools.Read(reader io.Reader) ([][]string, error)

CSVファイルの読み込み
### Dump
csvtools.Dump(records [][]string)

CSVを出力
### MaxCol
csvtools.MaxCol(records [][]string)

列数の最大を返す
### MaxRow
csvtools.MaxRow(records [][]string)

行数の最大を返す
### Transpose
csvtools.Transpose(records [][]string) [][]string

行と列を入れ替える
### ColHash
csvtools.ColHash(records [][]string, key int, value int) map[string]string

列毎に連想配列化する

### RowHash
csvtools.RowHash(records [][]string, key int, value int) map[string]string

行毎に連想配列化する
### Json
csvtools.Json(data interface{}) string

JSON文字列に変換する

### Trim
Trim(records [][]string, col int, row int, colsize int, rowsize int) [][]string

CSVの一部を切り出す