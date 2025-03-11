package sections

import (
	"fmt"
	"log"
	"os"
)

func Sec16() {
	// f_os()
	//f_defer()
	//f_log()
	//f_args()
	//f_open()
	//f_create()
	//f_read()
	f_openfile()
}

func f_os() {
	os.Exit(1)
	fmt.Println("Start")
}

func f_defer() {
	defer func() {
		fmt.Println("defer")
	}()
	os.Exit(0)
}
func f_log() {
	_, err := os.Open("A.txt")
	if err != nil {
		log.Fatalln(err)
	}
}
func f_args() {
	fmt.Printf("length=%d\n", len(os.Args))
	for i, v := range os.Args {
		fmt.Println(i, v)
	}
}

func f_open() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
}

func f_create() {
	f, _ := os.Create("test1.txt")
	f.Write([]byte("Hello\n"))
	f.WriteAt([]byte("Golang"), 6)
	f.Seek(0, os.SEEK_END)
	f.WriteString("Yaah")
}

func f_read() {
	f, err := os.Open("test1.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	bs := make([]byte, 128)
	n, err := f.Read(bs)    // nはバイト数
	fmt.Println(n)          // 16
	fmt.Println(string(bs)) //ないよう
	bs2 := make([]byte, 128)
	n2, err := f.ReadAt(bs2, 5) //オフセット 10バイト目から読み込む
	fmt.Println(n2)
	fmt.Println(string(bs2))
}

func f_openfile() {
	// O_RDONLY　読み込み専用
	// O_WRONLY　書き込み専用
	// O_RDWR　読み書き両方
	// O_APPEND ファイル末尾に追記
	// O_CREATE ファイルがなければ作成
	// O_TRUNC　ファイルの内容をオープン時に空にする
	f, err := os.OpenFile("foo.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	bs := make([]byte, 128)
	n, err := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))
}
