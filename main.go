package main

import (
	"aa/school"
	"aa/utill"
	"bytes"
	// "encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http" //import package net-http

	// "strconv"
	"strings"
	"time"

	"aa/book"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func main() {
	var a bool
	var b int
	var c, d string                         //decalr both in the same line
	var e, f bool = false, true             //declar type with value
	var h, i, j = false, 10, "เรียนภาษา Go" //autu declar TYPE [bool, int, string]
	k, l, m := false, 10, "เรียนภาษา Go"    //shorthand

	c, d = "c", "d"
	b = 2  //assign new value
	g := 3 //declar with value

	fmt.Printf("from fmt.Printf %v %v %s\n", a, b, c)                      //allow to format
	fmt.Println(`from fmt.Println`, a, b, c, d, e, f, g, h, i, j, k, l, m) // function from fmt    , same as fmt.Print but append newline \n
	fmt.Print("from fmt.Print ", "aaa ", "avx", a, b)                      //no format //no space between value
	println("Hellow")
	utill.TestFunction()
	//utill.testFunction()  this can't access cause private func

	if a == true && c == "c" {
		fmt.Println("a = true")
	} else {
		fmt.Println("a = false")
	}

	switch a {
	case true:
		fmt.Println("a = true")
	case false:
		fmt.Println("a = false")
	default:
		fmt.Println("a = nu")
	}

	sum := 0
	vA := 1
	vB := 4
	if sum = vA + vB; sum == 5 {
		println("k มีค่าเท่ากับ 5")
	}

	//array
	var array [5]int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	fmt.Println("array 5 arg: ", array)

	//array
	var name = [3]string{"Chaiyarin", "Atikom", "Kritsana"}
	fmt.Println("array with init value 'name':", name, "len(name): ", len(name))

	//slice
	nameArray := []string{}
	nameArray = append(nameArray, "Chaiyarin")
	nameArray = append(nameArray, "Atikom")
	nameArray = append(nameArray, "Kritsana")

	fmt.Println(nameArray)

	mapM := make(map[string]int)
	mapM["chaiyarin"] = 1
	mapM["atikom"] = 2
	mapM["kritsana"] = 3
	fmt.Println(mapM)
	delete(mapM, "kritsana")
	fmt.Println("after delate map: ", mapM)

	schoolAddress := school.GetSchoolAddress()
	fmt.Println(schoolAddress)
	fmt.Println(school.SchooleName)
	resultCode, resultAddress := getSchoolAddressWith2()
	fmt.Println(resultCode, resultAddress)

	for i := 1; i < 9; i++ {
		fmt.Println(i)
	}

	//do while loop
	round := 1
	for {
		if round == 2 {
			println(round)
			break
		}
		round++
	}

	//while loop
	iround := 1
	for iround <= 5 {
		fmt.Println(iround)
		iround = iround + 1
	}

	for i := range [5]int{} {
		fmt.Println(i)
	}

	//defer
	printFirst()
	defer printFinish() // มีการเพิ่ม Defer มาที่ Function นี้ will do lasttime after main()
	printSecond()
	//expect first second close

	printFirst()
	defer printThird()  // -> ลำดับที่ 3
	defer printFourth() // -> ลำดับที่ 2
	defer printFinish() // -> ลำดับที่ 1
	printSecond()
	//expect first second close fourth third

	cus := customer{
		firstname: "Chaiyarin",
		lastname:  "Niamsuwan",
		code:      111990,
		phone:     "085661234",
	} // การกำหนดค่าเริ่มต้น ให้ customer struct
	cus.firstname = "Atikom"
	fmt.Println(cus)
	//[{Chaiyarin Niamsuwan 111990 085661234} {Atikom Sombutjalearn 111991 085664321} {Kritsana Punyaphon 111992 085662344}]


	//go routine     purpose to do thread  help do work 
	//แบ่งไปให้ CPU อีก core ช่วยทำงาน
	fmt.Println("------------DO GO ROUTINE")
	go fmt.Println("ซื้อแว่น ที่ เซเว่น")  // ใส่ go ลงไป
	go fmt.Println("ซื้อนาฬิกา ที่ เซ็นทรัล")  // ใส่ go ลงไป
	fmt.Println("ซื้อผลไม้ ที่ สยามพารากอน")
	fmt.Println("ซื้อรถ ที่ ศูนย์ Toyota")
	time.Sleep(1 * time.Second)


	//Channel    collect the result from routine to main core
	fmt.Println("--------------DO CHANNEL")
	ch := make(chan string); //for sent string type to main core
	go fmt.Println("ซื้อแว่น ที่ เซเว่น")
	go fmt.Println("ซื้อนาฬิกา ที่ เซ็นทรัล")
	go sendToMisterA(ch); //sent to ch
	fmt.Println("ซื้อผลไม้ ที่ สยามพารากอน")
	fmt.Println("ซื้อรถ ที่ ศูนย์ Toyota")
	messageFromMisterB := <-ch
	fmt.Println(messageFromMisterB)
	
	// ch <- "get item from other CORE"
	// misterA := <- ch


	//split
	substrings := strings.Split("a,b,c", ",")
	fmt.Println("do substring: ",substrings)

	//do slice
	slice := []int{1, 2, 3, 4, 5}
    sliced := slice[1:4] // Extract elements at index 1, 2, and 3
    fmt.Println("do slice: ",sliced)  // Output: [2 3 4]
	handleRequest()
}

func sendToMisterA(message chan<- string) {
	time.Sleep(1 * time.Second)
	message <- "กำลังส่งของให้ นาย A"
 }

func printFirst() {
	fmt.Println("First")
}
func printSecond() {
	fmt.Println("Second")
}
func printFinish() {
	fmt.Println("Close")
}
func printThird() {
	fmt.Println("Third")
}
func printFourth() {
	fmt.Println("Fourth")
}

type customer struct { // การประกาศโครงสร้าง struct
	firstname string
	lastname  string
	code      int
	phone     string
}

func homePage(w http.ResponseWriter, r *http.Request) { // (1)
	fmt.Printf("request: %v", r)
	fmt.Fprint(w, "Welcome to the HomePage!") // (2)
}

func getAddressBookAll(c *fiber.Ctx) error {

	addBook := addressBook{
		Firstname: "Chaiyarin",
		Lastname:  "Niamsuwan",
		Code:      1993,
		Phone:     "0870940955",
	}
	return c.JSON(addBook)
}

func getSchoolAddressWith2() (int, string) {
	code := 1993
	address := "กรุงเทพ"
	//can return more than one
	return code, address
}

type Sess struct {
	sessi string `json:"session"`
}

func postUrl(c *fiber.Ctx) error {
	validator := validator.New()
	payload := new(Post)
	// contentType := c.Get("Content-Type")
	// Or get all headers
	// headers := c.Request().Header
	fmt.Println(c.GetReqHeaders())
	
	//session = A
	

	// Do something with the headers
	// For example, print them out
	// for headerName, headerValues := range headers {
	// 	for _, headerValue := range headerValues {
	// 		println(headerName, ":", headerValue)
	// 	}
	// }

	if err := c.BodyParser(&payload); err != nil {
		log.Fatal(err)
        return err
    }
	err := validator.Struct(payload)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("fine")
		errors.New("a")
    }
	fmt.Println(payload)
	posturl := "https://jsonplaceholder.typicode.com/posts"

	body := []byte(`{
		"title": "Post title",
		"body": "Post description",
		"userId": 3
	}`)
	
	r, err := http.NewRequest(http.MethodPost, posturl, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal("error: ",err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Fatal("error2: ", err)
	}

	defer res.Body.Close()

	post := &Post{}
	// derr := json.NewDecoder(res.Body).Decode(post)
	// if derr != nil {
	// 	log.Fatal("error3: ", derr)
	// }

	// if res.StatusCode != http.StatusCreated {
	// 	log.Fatal(res.Status)
	// }

	return c.JSON(post)
}

type Post struct {
	Id     int    `json:"id"`
	Title  string `json:"title" validate:"required,oneof=A B"`
	Body   string `json:"body"`
	UserId int    `json:"userId" validate:"required,number,max=6"`
}

type addressBook struct {
	Firstname string
	Lastname  string
	Code      int
	Phone     string
}

type A interface {
	hello()
}

type AAA struct {
	name string `json:"name" validate:"required,oneof=JOHN SMITE"`
}

type BBB struct {
	name string
}

func (p AAA) hello() {
	fmt.Println(p.name)
}

func (p BBB) hello() {
	fmt.Println(p.name)
}

func (p *AAA) changeName(a string) {
	p.name = a
}

func changeName(p *AAA, a string) {
	p.name = a
}

func handleRequest() { // (3)
	// var a = 1
	// var b int
	// var c []int
	// var d [3]int
	// var dd bool
	// e := make(map[string]string)
	// var f map[string]string

	aaa := AAA{name: "abc"}
	// aaa.name = "ccc"
	aaa.changeName("ccc")
	changeName(&aaa, "ddd")
	bbb := BBB{name: "def"}

	cc := []A {aaa, bbb}

	for i:=0;i<len(cc);i++ {
		cc[i].hello()
	}

	for a,b  := range cc{
		fmt.Println(a, b)
	}
	// myRouter := mux.NewRouter().StrictSlash(true)
	// myRouter.HandleFunc("/", homePage).Methods(http.MethodPost)
	// myRouter.HandleFunc("/getAddress", getAddressBookAll).Methods(http.MethodGet)
	// myRouter.HandleFunc("/postUrl", postUrl).Methods(http.MethodPost)
	// http.ListenAndServe(":8080", myRouter)

	app := fiber.New()

    app.Get("/getAddress", getAddressBookAll)
	app.Post("/postUrl", postUrl)

	// CRUD routes
	app.Get("/book", book.GetBooks)
	app.Get("/book/:id", book.GetBook)
	app.Post("/book", book.CreateBook)
	app.Put("/book/:id", book.UpdateBook)
	app.Delete("/book/:id", book.DeleteBook)

	app.Post("/test", book.Test)
	log.Fatal(app.Listen(":8080"))
}