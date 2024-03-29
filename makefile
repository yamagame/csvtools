all: \
	csv2json \
	dumpcsv \
	transcsv \
	trimcsv \
	sortcsv \
	joincsv \
	gencode

csv2json: cmd/csv2json/main.go
	go build -o bin/csv2json cmd/csv2json/main.go

dumpcsv: cmd/dumpcsv/main.go
	go build -o bin/dumpcsv cmd/dumpcsv/main.go

transcsv: cmd/transcsv/main.go
	go build -o bin/transcsv cmd/transcsv/main.go

trimcsv: cmd/trimcsv/main.go
	go build -o bin/trimcsv cmd/trimcsv/main.go

sortcsv: cmd/sortcsv/main.go
	go build -o bin/sortcsv cmd/sortcsv/main.go

joincsv: cmd/sortcsv/main.go
	go build -o bin/joincsv cmd/joincsv/main.go

gencode: cmd/gencode/main.go
	go build -o bin/gencode cmd/gencode/main.go
