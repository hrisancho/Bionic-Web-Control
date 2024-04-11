# Для начало работы требуется скачать protobuf-compiler
# Дальше требуется установить плагин для работы с golang
# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# Не стоит забывать что плагин компилятора должен находиться в $PATH(переменную среды)
#	1) Перенести скаченный бинарный файл в переменную среды
#		$ sudo cp $GOPATH/bin/protoc-gen-go /usr/bin
#	2) Добавить файл $GOPATH в переменную среды
#		$ export PATH=$PATH:$GOPATH/bin
protobuf:
	protoc --proto_path=./proto/ --go_out=./proto/ --go_opt=paths=source_relative ./proto/*/*.proto