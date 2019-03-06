package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"struct/defer/handleFile"
)

func main () {
	//tryDefer()
	//writeFile("a.txt")

	http.HandleFunc("/", errWrapper(handleFile.HandleFile))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}


type appHandler func (writer http.ResponseWriter, request *http.Request) error

func errWrapper( handler appHandler) func( http.ResponseWriter,  *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic: %v", r)
				http.Error(writer,http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)	//意料之中用error 意料之外使用panic
			}
		}()
		err := handler(writer, request)
		if err != nil {
			log.Printf("Error occurred handling request: %s", err.Error())
			if userError, ok := err.(userError); ok {		//用来判断是否为userError
				http.Error(writer, userError.Message(), http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)

		}
	}
}

type userError interface {
	error
	Message() string
}


func tryDefer() {
	for i := 0; i < 100;i++  {
		defer fmt.Println(i)
		if i == 30 {
			panic("panic too many")
		}
	}

	return
}


func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	defer file.Close()
	//err = errors.New("hehhehe")
	if err != nil {
		//panic(err)
		if pathError11, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError11.Op, pathError11.Path, pathError11.Err)
		}
		return
	}
	writer := bufio.NewWriter(file)
	writer.WriteString("ertyuiowewewew")
	defer writer.Flush()





}