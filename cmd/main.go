package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	//learnArraySliceMapLoop()
	//checkTimePerformance()
	//learnStringRuneByte()
	//learnStruct()
	//learnPointer()
	//learnGoroutines()
	//learnGoroutinesNextLevel()
	learnGenerics()
}

func learnArraySliceMapLoop() {
	//var intArr []int32
	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr)

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("length slice: %v, capacity slice: %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("length slice: %v, capacity slice: %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9, 10}
	intSlice = append(intSlice, intSlice2...)
	fmt.Printf("length slice: %v, capacity slice: %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"vincent": 27, "wuliango": 99}
	fmt.Println(myMap2["vincent"])
	var age, ok = myMap2["vincent"]
	if ok {
		fmt.Println("umur:", age)
	} else {
		fmt.Println("tidak ketemu umurnya")
	}

	for name, age := range myMap2 {
		fmt.Printf("nama: %v, umur: %v\n", name, age)
	}

	for i, v := range intSlice {
		fmt.Printf("key: %v, value: %v\n", i, v)
	}
}

func checkTimePerformance() {
	var n int = 1000_000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("total time without preallocation %v\n", timeLoop(testSlice, n))
	fmt.Printf("total time with preallocation %v\n", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}

	return time.Since(t0)
}

func learnStringRuneByte() {
	var myString = []rune("résumé")
	var indexed = myString[1]
	fmt.Printf("byte: %v, data type: %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("panjang string= %v\n", len(myString))

	var myRune = 'a'
	fmt.Printf("myRune = %v\n", myRune)

	var strSlice = []string{"v", "i", "n", "c", "e", "n", "t"}
	var strBuilder = strings.Builder{}
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var concatenateStr = strBuilder.String()
	fmt.Println(concatenateStr)
}

// Struct

type gasEngine struct {
	mgp     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.mgp * e.gallons
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("aman, pasti sampai ke tujuan")
	} else {
		fmt.Println("perlu isi bensin dulu atuh")
	}
}

func learnStruct() {
	//var myEngine engine = gasEngine{25, 15}
	var myEngine engine = electricEngine{10, 15}
	//fmt.Printf("sisa bensin di tank = %v", myEngine.milesLeft()())
	canMakeIt(myEngine, 200)
}

// Pointers

func learnPointer() {
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("lokasi memori thing1 = %p\n", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("hasil pangkat = %v\n", result)
	fmt.Printf("value dari thing1 = %v\n", thing1)

}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("lokasi memori thing2 = %p\n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}

	return *thing2
}

func learnGoroutines() {
	var wg sync.WaitGroup
	go sayHello(&wg)
	fmt.Println("Hello dari main function!")
	wg.Wait()
}

func sayHello(wg *sync.WaitGroup) {
	wg.Add(1)
	fmt.Println("Hello dari goroutine!")
	wg.Done()
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Pekerja", id, "siap!")
		fmt.Println("Pekerja", id, "mengerjakan job", j)
		time.Sleep(time.Second)
		fmt.Println("Pekerja", id, "sudah menyelesaikan job", j)
		results <- j * 2
	}

}

func learnGoroutinesNextLevel() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)

	fmt.Println("--- Mengambil hasil ---")
	for r := 1; r <= numJobs; r++ {
		result := <-results
		fmt.Println("Hasil diterima:", result)
	}

	fmt.Println("Semua kerjaan sudah selesai")
}

func learnGenerics() {
	//var intSlice = []int{1, 2, 3}
	//fmt.Println(sumSlice(intSlice))
	//
	//var floatSlice = []float32{5.0, 10.0, 15.0}
	//fmt.Println(sumSlice(floatSlice))

	//var contacts []contactInfo = loadJSON[contactInfo]("./contactInfo.json")
	//fmt.Printf("\n%+v", contacts)

	var gasCar = car[gasEngine]{
		carMake:  "Honda",
		carModel: "Civic",
		engine: gasEngine{
			gallons: 12,
			mgp:     40,
		},
	}

	fmt.Println(gasCar)
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T {
	data, _ := os.ReadFile(filePath)

	var loaded = []T{}
	json.Unmarshal(data, &loaded)

	return loaded
}

type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}
