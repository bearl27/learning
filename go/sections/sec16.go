package sections

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

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

func f_time() {
	t := time.Now()
	fmt.Println(t)

	t2 := time.Date(2020, 6, 10, 10, 10, 10, 0, time.Local)
	fmt.Println(t2)
	fmt.Println(t2.Year())
	fmt.Println(t2.YearDay())
	fmt.Println(t2.Month())
	fmt.Println(t2.Weekday())
	fmt.Println(t2.Day())
	fmt.Println(t2.Hour())
	fmt.Println(t2.Minute())
	fmt.Println(t2.Second())
	fmt.Println(t2.Nanosecond())
	fmt.Println(t2.Zone())
}
func f_duration() {
	fmt.Println(time.Hour)
	fmt.Printf("%T\n", time.Hour)
	fmt.Println(time.Minute)
	fmt.Println(time.Second)
	fmt.Println(time.Millisecond)
	fmt.Println(time.Microsecond)
	fmt.Println(time.Nanosecond)

	//string
	d, _ := time.ParseDuration("2h30m")
	fmt.Println(d) // 2h30m0s

	//time
	t := time.Now()
	t2 := t.Add(2*time.Hour - 15*time.Minute)
	fmt.Println(t2) // 2025-03-11 12:16:00.407823 +0900 JST m=+8100.000468251
}
func f_time_comp() {
	t1 := time.Date(2020, 7, 17, 0, 0, 0, 0, time.Local)
	t2 := time.Date(2020, 7, 18, 0, 0, 0, 0, time.Local)
	fmt.Println(t1.Sub(t2))    // -24h0m0s
	fmt.Println(t1.Before(t2)) // true
	fmt.Println(t1.Equal(t2))  // false
	fmt.Println(t1.After(t2))  // false
}
func f_sleep() {
	time.Sleep(3 * time.Second)
	fmt.Println("3秒間sleep")
}
func f_fprint() {
	f, _ := os.Create("fprint.txt")
	defer f.Close()
	fmt.Fprint(f, "Hello")
	fmt.Fprintf(f, "Hello")
	fmt.Fprintln(f, "Hello")
}
func f_logs() {
	//ログの出力先
	log.SetOutput(os.Stdout)
	log.Print("Log")
	// log.Fatal("log")

	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Print("Log")
	log.SetFlags(log.LstdFlags)
	log.SetPrefix("[LOG]")
	log.Print("Log")

	logger := log.New(os.Stdout, "[LOG]", log.Ldate|log.Ltime|log.Llongfile)
	logger.Print("Log")
	log.Print("log")
}

func f_flag() {
	var (
		max int
		msg string
		x   bool
	)
	flag.IntVar(&max, "n", 32, "処理数の最大値")
	flag.StringVar(&msg, "m", "", "処理メッセージ")
	flag.BoolVar(&x, "x", false, "拡張オプション")
	flag.Parse()

	fmt.Println("処理数の最大値 = ", max)
	fmt.Println("処理メッセージ = ", msg)
	fmt.Println("拡張オプション = ", x)
}
func f_strconv() {
	bt := true
	fmt.Printf("%T\n", strconv.FormatBool(bt))

	fmt.Println(strconv.FormatFloat(123.456, 'E', -1, 64))
	fmt.Println(strconv.FormatFloat(123.456, 'e', -1, 64))
	fmt.Println(strconv.FormatFloat(123.456, 'f', -1, 64))
	fmt.Println(strconv.FormatFloat(123.456, 'f', 2, 64))
	fmt.Println(strconv.FormatFloat(123.456, 'g', -1, 64))
	fmt.Println(strconv.FormatFloat(123.456789123, 'f', 4, 64))
	fmt.Println(strconv.FormatFloat(123.456789123, 'G', 4, 64))
}
func f_strings() {
	// Join
	s := []string{"A", "B", "C"}
	fmt.Println(strings.Join(s, ","))                    // A,B,C
	fmt.Println(strings.Index("ABCDE", "A"))             // 0
	fmt.Println(strings.Index("ABCDE", "Z"))             // -1 ない
	fmt.Println(strings.LastIndex("ABCDABCD", "BC"))     // 6 最後のBC
	fmt.Println(strings.IndexAny("ABCDABCD", "AD"))      // 0 どれか探す
	fmt.Println(strings.LastIndexAny("ABCDABCD", "AD"))  // 7
	fmt.Println(strings.HasPrefix("ABCDE", "AB"))        // true  それで始まるか
	fmt.Println(strings.HasSuffix("ABCDE", "DE"))        // true  それで終わるか
	fmt.Println(strings.Contains("ABCDE", "AD"))         // false その形で入っているから
	fmt.Println(strings.ContainsAny("ABCDE", "AD"))      // true  形は違えど入っているか
	fmt.Println(strings.Count("ABCABC", "B"))            // 2　数
	fmt.Println(strings.Repeat("ABC", 3))                // ABCABCABC　繰り返す
	fmt.Println(strings.Repeat("ABC", 0))                // 空　マイナスはエラー
	fmt.Println(strings.Replace("AAAA", "A", "B", 1))    // BAAA 1回
	fmt.Println(strings.Replace("AAAA", "A", "B", -1))   // BBBB 全部
	fmt.Println(strings.Split("A,B,C,D,E", ","))         // [A B C D E]　	,なし
	fmt.Println(strings.SplitAfter("A,B,C,D,E", ","))    // [A, B, C, D, E]	,あり
	fmt.Println(strings.ToLower("ABC"))                  // abc
	fmt.Println(strings.ToUpper("abc"))                  // ABC
	fmt.Println(strings.TrimSpace("       aaa      "))   // aaa
	fmt.Println(strings.TrimSpace("    ~   aaa   ~   ")) // ~  aaa  ~ 前後の空白のみ
	strs := strings.Fields("    ~   aaa   ~   ")
	fmt.Println(strs)             // [~ aaa ~] 空白で区切られたを取り出す
	fmt.Println(strs[0], strs[1]) // ~ aaa

}
func f_bufio() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if scanner.Text() == "exit" {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
func f_ioutil() {
	// ioutil
	f, _ := os.Open("test1.txt")
	// bs, _ := ioutil.ReadAll(f)
	// fmt.Println(string(bs))

	// if err := ioutil.WriteFile("test_write.txt", bs, 0666); err != nil {
	// 	log.Fatalln(err)
	// }
	bs, _ := io.ReadAll(f)
	fmt.Println(string(bs))

	if err := os.WriteFile("test_write.txt", bs, 0666); err != nil {
		log.Fatalln(err)
	}
}
func f_regexp() {
	//regexp
	match, _ := regexp.MatchString("A", "ABC")
	fmt.Println(match)
	//compile
	re1, _ := regexp.Compile("A")
	match = re1.MatchString("ABC")
	fmt.Println(match)
	//mustcompile
	re2 := regexp.MustCompile("A")
	match = re2.MatchString("ABC")
	fmt.Println(match)
}

func isValidPassword(password string) bool {
	// 正規表現パターン（先読みを使用せずに条件を直接指定）
	re := regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()_+{}\[\]:;<>,.?/~\\|` + "`" + `=+-]{8,16}$`)

	// パスワードが指定した文字セットと長さ条件を満たすかチェック
	if !re.MatchString(password) {
		return false
	}

	// 小文字、大文字、数字、特殊文字がそれぞれ含まれているかを確認
	hasLower := false
	hasUpper := false
	hasDigit := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= '0' && char <= '9':
			hasDigit = true
		case char == '!' || char == '@' || char == '#' || char == '$' || char == '%' || char == '^' ||
			char == '&' || char == '*' || char == '(' || char == ')' || char == '_' || char == '+' ||
			char == '{' || char == '}' || char == '[' || char == ']' || char == ':' || char == ';' ||
			char == '<' || char == '>' || char == ',' || char == '.' || char == '?' || char == '/' ||
			char == '~' || char == '\\' || char == '|' || char == '`' || char == '=' || char == '-':
			hasSpecial = true
		}
	}

	// すべての条件を満たす場合にのみ有効
	return hasLower && hasUpper && hasDigit && hasSpecial
}

func pass_regexp() {
	// テストするパスワード
	passwords := []string{
		"Password123!",    // 有効
		"password123",     // 無効（大文字、特殊文字がない）
		"PASSWORD123!",    // 無効（小文字がない）
		"Pass123",         // 無効（長さが足りない）
		"ValidPassword1!", // 有効
		"12345678",        // 無効（特殊文字がない）
		"Abc12345!",       // 有効
	}

	// 各パスワードを検証
	for _, password := range passwords {
		if isValidPassword(password) {
			fmt.Printf("パスワード '%s' は有効です\n", password)
		} else {
			fmt.Printf("パスワード '%s' は無効です\n", password)
		}
	}
}

var st struct{ A, B, C int }
var mutex *sync.Mutex

func UpdateAndPrint(n int) {
	// Lock
	mutex.Lock()

	st.A = n
	time.Sleep(time.Microsecond)
	st.B = n
	time.Sleep(time.Microsecond)
	st.C = n
	time.Sleep(time.Microsecond)
	fmt.Println(st)

	// unLock
	mutex.Unlock()
}

func f_sync() {
	mutex = &sync.Mutex{}
	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				UpdateAndPrint(i)
			}
		}()
	}

	for {

	}
}

func f_sync2() {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st")
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2nd")
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3rd")
		}
		wg.Done()
	}()
	wg.Wait()
}

type A struct {
}

type Student struct {
	ID      int       `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Age     int       `json:"age"`
	Created time.Time `json:"-"`
	A       A
}

func f_json() {

	s := new(Student)
	s.ID = 1
	s.Name = "test"
	s.Age = 2
	s.Created = time.Now()

	bs, err := json.Marshal(s)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}

type Entry struct {
	Key   string
	Value string
}

func f_sort() {
	i := []int{5, 3, 2, 8, 1}
	sort.Ints(i)

	el := []Entry{
		{"a", "A"},
		{"c", "C"},
		{"b", "B"},
	}
	sort.Slice(el, func(i, j int) bool { return el[i].Key < el[j].Key })
	fmt.Println(el)
}

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker: Canceled")
			return
		default:
			fmt.Println("Worker: Processing...")
			time.Sleep(1 * time.Second)
		}
	}
}

func f_context() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(3 * time.Second) // 3秒待ってからキャンセル
	fmt.Println("Main: Canceling worker...")
	cancel() // ゴルーチンの処理をキャンセル

	time.Sleep(1 * time.Second) // 終了確認のため少し待つ
	fmt.Println("Main: Done")
}

// type MyHandler struct{}

// func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello, World!")
// }

func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("top.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, "This is Top")
}

func f_server() {
	http.HandleFunc("/tip", top)
	http.ListenAndServe(":8080", nil)
}

func Sec16() {
	// f_os()
	//f_defer()
	//f_log()
	//f_args()
	//f_open()å
	//f_create()
	//f_read()
	//f_openfile()
	//f_time()
	//f_duration()
	//f_time_comp()
	//f_sleep()
	//f_fprint()
	// f_logs()
	// f_flag()
	// f_strconv()
	// f_strings()
	// f_bufio()
	// f_ioutil()
	// f_regexp()
	// pass_regexp()
	// f_sync()
	// f_sync2()
	// f_json()
	// f_sort()
	// f_context()
	f_server()
}
