package main

import (
	"fmt"
	"math"
)

// "reflect"

// #1
// func main() {
// 	const message string = "text"

// Uncomment this line to define and use the variable message1
// var message1 = 17

// fmt.Println(reflect.TypeOf(message))
// Uncomment this line to print the value of message1
// fmt.Println(message1)
// }

// #2
// func main() {
// 	// var message string
// 	// var number float32
// 	// var b bool

// 	// b = true
// 	// // // number = 12.2
// 	// var a byte = 97
// 	var a rune = 'a'

// 	// // message := []byte("asd")
// 	// Printing default zero values for string and int
// 	// fmt.Println(message) // Output will be an empty string
// 	// fmt.Println(number)  // Output will be 0
// 	// fmt.Println(b)
// 	fmt.Println(a)
// }

// #3
// var a, b, c int

// func main() {
// 	a, b, c = 1, 2, 3

// 	// a, b = b, a

// 	// d, g, e := 2, 6, 8
// 	// d, _, e := 2, 5, 6

// 	println(a, b, c)
// 	// println(d, e)
// }

// func print() {
// 	// a, b, c := 4, 5, 6
// 	println(a, b, c)
// }

// func main() {
// 	// var message string
// 	// message = sayHello("Абдр", 21)
// 	// printMessage(message)
// 	// printMessage("Вызов 1")
// 	// printMessage("Вызов 2")
// 	// printMessage("Вызов 3")
// 	// message, entered := enterTheClub(18)
// 	// fmt.Println(message)

// 	// message, err := enterTheClub(12)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	// fmt.Println(preduction("пн"))
// }

// func printMessage(message string) {
// 	fmt.Println(message)
// }

// func sayHello(name string, age int) string {
// 	result := fmt.Sprintf("Привет, %s! Тебе %d лет", name, age)
// 	return result
// }

// #4

// func enterTheClub(age int) (string, error) {
// 	if age >= 18 && age <= 45 {
// 		response := "Входи"
// 		return response, nil
// 	} else if age >= 45 && age <= 65 {
// 		return "Вам точно понравиться музыка", nil
// 	} else if age >= 65 {
// 		return "Это уже слишком", errors.New("You are too old")
// 	}

// 	return "Уходи", errors.New("You are too young or too old")
// }

// # 5

// func preduction(dayOfWeek string) (string, error) {
// 	switch dayOfWeek {
// 	case "пн":
// 		return "Хорошего пн", nil
// 	case "вт":
// 		return "Хорошего вт", nil
// 	case "ср":
// 		return "Хорошего ср", nil
// 	default:
// 		return "Невалидный день недели", errors.New("indalid day of week")
// 	}
// }

// #6

// func main() {
// 	fmt.Println(findMin(1, 435, 345345, 34535, 12234, 234524, 121))
// }

// func findMin(numbers ...int) int {
// 	if len(numbers) == 0 {
// 		return 0
// 	}
// 	min := numbers[0]

// 	for _, i := range numbers {
// 		if i > min {
// 			min = i
// 		}
// 	}

// 	return min
// }

// #7. Анонимные функции && init

// var msg string

// func init() {
// 	msg = "from init()"
// }

// func main() {
// 	inc := increment() // создаём новое замыкание
// 	fmt.Println(msg)
// 	fmt.Println(inc()) // выводит 1
// 	fmt.Println(inc()) // выводит 2
// 	fmt.Println(inc()) // выводит 3
// 	fmt.Println(inc()) // выводит 4

// 	// вызовы increment2
// 	fmt.Println(increment2()) // вызывает бесконечную рекурсию и вызывает паник
// 	fmt.Println(increment2())
// 	fmt.Println(increment2())
// 	fmt.Println(increment2())
// }

// func increment() func() int {
// 	count := 0 // переменная count определена в этой функции
// 	return func() int {
// 		count++      // увеличиваем count
// 		return count // возвращаем текущее значение count
// 	}
// }

// func increment2() func() int {
// 	count := 0
// 	count++
// 	return increment2() // вызывает бесконечную рекурсию и вызывает паник
// }

// #6Указатели и ссылки

// func main() {
// 	number := 5
// 	var p *int
// 	p = &number

// 	fmt.Println(p)
// 	fmt.Println(&number)

// 	*p = 10

// 	fmt.Println(number)
// 	// nil - пустой указатель

// 	message := "Скоро я стану хокаге"

// 	fmt.Println(&message)

// 	changeMessage(&message) // referencing
// 	// & принимает область памяти
// 	// fmt.Println(&message)
// }

// func changeMessage(message *string) {
// 	//*message += " (из функции printMessage()) " // dereference - обращаемся по
// 	fmt.Println(&message)
// }

//  #7 massive && slice

// func main() {
// 	message := make([]string, 100)
// 	fmt.Println((message))    // выведет слайс с 100 пустыми строками
// 	fmt.Println(len(message)) // длина слайса - 100
// 	fmt.Println(cap(message)) // ёмкость слайса - 100
// 	message = append(message, "101")
// 	fmt.Println(len(message)) // длина слайса - 101
// 	fmt.Println(cap(message)) // ёмкость слайса - увеличилась до 215
// }

// func printMessage(message []string) error {
// 	if len(message) == 0 {
// 		return errors.New("empty array")
// 	}

// 	message[1] = "5"

// 	fmt.Println(message)

// 	return nil
// }

// #7 matrics && for cycle

// func main() {

// x := 0
// for true {
// 	x++
// 	fmt.Println(x)
// }

// matrix := make([][]int, 10)

// counter := 0
// for x := 0; x < 10; x++ {
// 	matrix[x] = make([]int, 10)
// 	for y := 0; y < 10; y++ {
// 		counter++
// 		matrix[x][y] = counter
// 	}
// 	fmt.Println(matrix[x])

// message := []string{
// 	"message 1",
// 	"message 2",
// 	"message 3",
// 	"message 4",
// }

// for i := 0; i < len(message); i++ {
// 	fmt.Println(message[i])
// }

// for _, message := range message {
// 	fmt.Println(message)
// }

// 	counter := 0
// 	for {
// 		if counter == 10 {
// 			break
// 		}
// 		counter++
// 		fmt.Println(counter)
// 	}
// }

// #8 panic, defer, recover

// func main() {
// 	defer handlePanic()

// 	fmt.Println("main()")
// 	message := []string{
// 		"message 1",
// 		"message 2",
// 		"message 3",
// 		"message 4",
// 	}

// 	message[4] = "message 5"

// 	fmt.Println(message)

// 	// panic("AAAAAAAAAAAA")
// }

// func handlePanic() {
// 	if r := recover(); r != nil {
// 		fmt.Println(r)
// 	}

// 	fmt.Println("handlePanic() worked succesfully")
// }

// #8 maps

// func main() {
// 	users := map[string]int{
// 		"A": 35,
// 		"B": 36,
// 		"C": 34,
// 	}

// 	// age, exists := users["Kostya"]

// 	// if exists {
// 	// 	fmt.Println("Kostya", age)
// 	// } else {
// 	// 	fmt.Println("Person doesn't exist in list")
// 	// }

// 	// var users map[string]int
// 	// users = make(map[string]int)

// 	users["AB"] = 15
// 	// delete(users, "B")

// 	for key, value := range users {
// 		fmt.Println(key, value)
// 	}

// 	defer func() {
// 		users["D"] = 25
// 	}()

// 	fmt.Println(len(users))
// }

// #10 structure

// type Age int

// type User struct {
// 	name   string
// 	age    Age
// 	sex    string
// 	weight int
// 	height int
// }

// type Databases struct {
// 	m map[string]string
// }

// func NewDatabase() *Databases {
// 	return &Databases{
// 		m: make(map[string]string),
// 	}
// }

// func (a Age) isAdult() bool {
// 	return a >= 18
// }

// func (u *User) setName(name string) {
// 	u.name = name
// }

// func (u *User) getName() string {
// 	return u.name
// }

// func NewUser(name, sex string, age, weight, height int) User {
// 	return User{
// 		name:   name,
// 		age:    Age(age),
// 		sex:    sex,
// 		weight: weight,
// 		height: height,
// 	}
// }

// func main() {
// 	// Create users using the NewUser function and the User struct directly
// 	user := NewUser("A", "Male", 23, 75, 189)
// 	user1 := User{"B", 25, "Male", 87, 234}

// 	// Print user information
// 	fmt.Println(user.age.isAdult())

// 	fmt.Println(user.getName())
// 	fmt.Println(user1.getName())
// }

// # 11 interface

type Shape interface {
	ShapeWithArea
	ShapeWithPerimetr
}

type Square struct {
	sideLenght float32
}

type ShapeWithArea interface {
	Area() float32
}

type ShapeWithPerimetr interface {
	Perimetr() float32
}

func (s Square) Area() float32 {
	return s.sideLenght * s.sideLenght
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func main() {
	// square := Square{5}
	// circle := Circle{8}

	// fmt.Println(square)
	// fmt.Println(circle)

	// printInterface(square)
	// printInterface(circle)
	// printInterface("qwe")
	// printInterface(22)
	// printInterface(true)

	printInterface(("hofweoihhwegohwguwe"))

}

func printShapeArea(shape Shape) {
	fmt.Println(shape.Area())
}

func printInterface(i interface{}) {
	// switch value := i.(type) {
	// case int:
	// 	fmt.Println("int", value)
	// case bool:
	// 	fmt.Println("bool", value)
	// default:
	// 	fmt.Println("unknown type", value)
	// }
	// // fmt.Printf("%+v\n", i)

	str, ok := i.(string)
	if !ok {
		fmt.Println("Interface is not string")
	}
	fmt.Println(len(str))
}
