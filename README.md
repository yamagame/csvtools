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

### csvtools.Read(reader io.Reader) ([][]string, error)
### csvtools.Dump(records [][]string)
### csvtools.MaxCol(records [][]string)
### csvtools.MaxRow(records [][]string)
### csvtools.Transpose(records [][]string) [][]string
### csvtools.RowHash(records [][]string, key int, value int) map[string]string
### csvtools.ColHash(records [][]string, key int, value int) map[string]string
### csvtools.Json(data interface{}) string
### Trim(records [][]string, col int, row int, colsize int, rowsize int) [][]string

