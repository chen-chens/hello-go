package main

import (
	"fmt"
	// "time"
	// "errors"
	// "sort"
	// "hello/prac"
)

// var sum = 100

// func main() {
// num1 := 100
// num1, num2 := -20, -30
// num2, num3 := 3, 5

// var sum = -60

// sum, num4 := 5050, 2020

// fmt.Println(num1, num2, num3, sum, num4)
// fmt.Printf("num1: %d", num1)
// }

// func main(){
// const (
// 	A = iota           // 0
// 	B                  // 1
// 	_                  // 2 (佔位符也會計算)
// 	C                  // 3
// 	D = iota * 100     // 4 * 100 (往下依序計算)
// 	E                  // 5 * 100
// 	F                  // 6 * 100
// )

// fmt.Println(A, B, C, D, E, F)
// }
// a := 10
// b:=1222
// fmt.Printf("a: %d .....qqq...%s", a,b)

// }

// func main(){
// 	type Info struct{
// 		name string
// 		age int
// 		hobby bool
// 	}
// 	fmt.Printf("Info: %v \n", Info{})
// 	fmt.Printf("Info: %+v \n", Info{})
// 	fmt.Printf("Info: %#v \n", Info{})
// }

// func main(){
// 	a, b := -27, -7
// 	fmt.Println(a%b)
// }
// func main(){
// 	a, b := 13.5, 2
// 	fmt.Println(a + float64(b))
// }

// func main(){
// 	var a string = "隱藏著黑暗力量的鑰匙啊!"
// 	var b string = "在我面前顯示你真正的力量!"
// 	fmt.Print(a, 1234, b)
// }

// func main(){
// 	st1, st2, st3 := "我",123,"Joanna"

// 	var wordsWithSprintln = fmt.Sprintln(st1, st2, st3) // 組合後換行(每個變數有間距)
// 	var wordsWithSprint = fmt.Sprint(st1, st2, st3) // 組合後連再一起
// 	var wordsWithSprintf = fmt.Sprintf("%s%d%s", st1, st2, st3) // 依format格式印出

// 	fmt.Println(wordsWithSprintln, wordsWithSprint, wordsWithSprintf)
// }

// func main(){
// 	num:= 20
// 	if num > 10{
// 		fmt.Println("大")
// 	}else if num >5{
// 		fmt.Println("中")
// 	}else{
// 		fmt.Println("小")
// 	}
// }

// func main(){
// 	num := 13
// 	if num%2 == 0{
// 		fmt.Println("even")
// 	}else{
// 		fmt.Println("odd")
// 	}
// }

// func main(){
// 	num := 0
// 	if num > 0{
// 		fmt.Println("正數")
// 	}else if num < 0{
// 		fmt.Println("負數")
// 	}else{
// 		fmt.Println("0")
// 	}
// }

// func main(){
// 	num := 3
// 	switch num{
// 		case 1, 2:
// 			fmt.Println("1和2 小可")

// 		case 3, 4:
// 			fmt.Println("3和4 小櫻")

// 		default:
// 			fmt.Println("庫洛牌")
// 	}
// }

// func main(){
// 	num := 4
// 	for i :=0; i<4; i++{
// 		if i < num{
// 			fmt.Println("庫洛牌")
// 		}
// 	}
// }

// func main(){
// 	var num int = 6
// 	i :=10
// 	for ; i>0; i--{
// 		if num == i{
// 			continue
// 		}
// 	}
// 	fmt.Println(i)
// }

// func sum(a int, b int) int{
// 	return a+b
// }

// func main(){
// 	fmt.Printf("%d", sum(3,8))
// }

// func mutilple (a int, b int) (int, int){
// 	num1 := a+b
// 	num2 := a-b
// 	return num1, num2
// }
// func main(){
// 	me, _ := mutilple(3,7)
// 	fmt.Println(me)
// }

// func mutilple (a int, b int) (sum int){
// 	sum = a+b
// 	return
// }
// func mutilple2 (a int, b int) (diff int){
// 	diff = a+b
// 	return
// }
// func main(){

// 	fmt.Println(mutilple(4,7))
// 	fmt.Println(mutilple2(19,3))
// }

// func main(){
// 	var arry [6]string
// 	arry[0] = "Joanna"
// 	arry[1] = "David"
// 	arry[2] = "Mia"

// 	fmt.Println(arry[5])
// }

// func arry() [6]string{
// 	var arr = [6]string{
// 		"小櫻",
// 		"小可",
// 		"雪兔哥",
// 	}
// 	return arr
// }
// func main(){
// 	fmt.Println(arry())
// }

// func main(){
// 	nums :=[5]int{1,2,3,4}
// 	result :=fmt.Sprintf("%d-%d-%d-%d", nums[0], nums[1], nums[2], nums[3])

// 	fmt.Println(result)

// }

// func main(){
// 	const INDEX = 5
// 	var nums [INDEX]int
// 	nums[1] = 123
// 	nums[2] = 245

// 	result :=fmt.Sprintf("%d-%d-%d-%d", nums[0], nums[1], nums[2], nums[3])
// 	fmt.Println(result)
// }

// func main(){
// 	var nums = [...]string{
// 		"Iam",
// 		"Joanna",
// 		"HowAreYou",
// 	}
// 	for _, value := range nums{
// 		value = value + value
// 	}
// 	for _, value := range nums{
// 		fmt.Println(value)
// 	}
// }

// func main(){
// 	var nums = make([]string, 3)

// 	fmt.Printf("長度: %d, 容積: %d", len(nums), cap(nums))
// }

// func main(){
// 	info := []string{
// 		"Joanna",
// 		"Mia",
// 		"David",
// 		"Bob",
// 	}
// 	for i:=0; i<len(info); i++{
// 		fmt.Println(info[i])
// 	}
// 	fmt.Printf("長度: %d, 容積: %d\n", len(info), cap(info))
// }

// func main(){
// 	nums := make([]int, 3, 5)
// 	nums[0] = 124
// 	nums[2] = 345
// 	fmt.Println(nums[0], nums[2], len(nums), cap(nums))
// }

// func main(){
// 	sums := []int{1, 2, 3, 4, 5}
// 	fmt.Println(sums[0], sums[2], len(sums), cap(sums))
// }

// func main(){
// 	var sums []string
// 	sums = append(sums, "小可")
// 	sums = append(sums, "小櫻")
// 	sums = append(sums, "桃矢")
// 	fmt.Println(sums[0], sums[2], len(sums), cap(sums))
// }
// func main(){
// 	sums := make(map[string]int)
// 	sums["小可"] = 120
// 	sums["小櫻"] = 156
// 	sums["桃矢"] = 180

// 	fmt.Println(sums["小可"], sums["小櫻"], sums["桃矢"])
// }
// func mutilply(a int, b int) int {
// 	return a*b
// }

// func main(){
// 	for a:=2; a<10; a++{
// 		for b:=2; b<10; b++{
// 			fmt.Printf("%d X %d = %d\n", a, b, mutilply(a, b))
// 		}
// 	}
// }

// func compare(a int, b int) {

// }
// func main(){
// 	arr := [...]int{10, 4, 5, 3, 7, 1, 9, 8}

// 	for i:=len(arr)-1; i>0; i--{
// 		for k:=0; k< len(arr); k++{
// 			// if(arr[k]< arr[k]){
// 			// 	arr[i]=arr[k]
// 			// }
// 		}

// 	}
// 	fmt.Println(arr)

// }

// func swap(arr []int, a int, b int) {
// 	temp := arr[a]
// 	arr[a] = arr[b]
// 	arr[b] = temp
// }

// func main() {
// 	arr := []int{10, 4, 5, 3, 7, 1, 9, 8}

// 	for i := 1; i < len(arr); i++ {
// 		k := i
// 		for k > 0 && arr[k-1] > arr[k] {
// 			swap(arr, k, k-1)
// 			k--
// 		}
// 	}

// 	fmt.Print(arr)
// }

// func main() {
// 	num:= 8
// 	var a *int
// 	a = &num
// 	*a = 100

// 	fmt.Println("pointer: ", a)
// 	fmt.Println("value: ", *a)
// 	fmt.Println("value: ", num)
// }

// func double(num *int) int {
// 	*num = 100
// 	return *num
// }

// func main() {
// 	num:= 8
// 	num = double(&num)

// 	fmt.Println("value: ", num)
// }

// func main(){
// 	numbers :=[]int{
// 		1,
// 		3,
// 		4,
// 	}
// 	fmt.Println(numbers)
// }

// func main(){
// 	var numbers []int
// 	numbers = append(numbers, 100)
// 	numbers = append(numbers, 200)
// 	numbers = append(numbers, 300)

// 	fmt.Println(numbers)
// }

// func main(){
// 	var numbers []int
// 	numbers = append(numbers, 100)
// 	numbers = append(numbers, 200)
// 	numbers = append(numbers, 300)

// 	fmt.Println(numbers)
// }

// func double(nums [4]int) [4]int {
// 	nums[3] = nums[3]*2
// 	return nums
// }

// func main(){
// 	var numbers = [4]int{1, 2, 3, 4}

// 	double(numbers)
// 	fmt.Println(numbers)
// }

// func double(nums []int) []int {
// 	nums[0] = 108
// 	return nums
// }

// func main(){
// 	var numbers = []int{1, 20, 33}

// 	double(numbers)
// 	fmt.Println(numbers)
// }

// func double(nums [5]int) [5]int{
//     for i := 0; i < len(nums); i = i+1{
//         nums[i] = nums[i] * 2
//     }
// 	return nums
// }

// func main(){
//     nums := [5]int{1, 3, 4, 7, 9}
//     double(nums)
//     for _, v := range nums{
//         fmt.Printf("%d ", v)
//     }
// }

// func main(){
// 	type person struct{
// 		name string
// 		age int
// 	}

// 	var info person
// 	info.name = "Joanna"
// 	info.age = 32

// 	fmt.Println(info)
// 	fmt.Printf("%T",info)
// }

// func main(){
// 	info := struct{
// 		name string
// 		age uint
// 	}{
// 		name : "Joanna",
// 		age : 32,
// 	}

// 	fmt.Println(info)
// 	fmt.Printf("%T",info)
// }

// type person struct{
//     name string
//     age uint
// }

// func plusOne(p *person){
//     p.age = p.age + 3
// }

// func main(){
//     sakura := person{
//         name : "小櫻",
//         age  : 10,
//     }
//     fmt.Println(sakura)
//     plusOne(&sakura)
//     fmt.Println(sakura)
// }

// type person struct{
//     name string
//     age uint
// }

// func plusOne(p *person){
//     p.age = p.age + 3
// }

// func main(){
//     sakura := new(person)
// 	sakura.name = "小櫻"
// 	sakura.age = 22

// 	fmt.Println(sakura)
// 	fmt.Printf("%T\n", sakura)

// 	plusOne(sakura)
// 	fmt.Println(sakura)
// }

// type person struct{
// 	name string
// 	age uint
// 	crush *person
// }

// func getCrush(p *person, crush *person){
// 	p.crush = crush
// }
// func getCrushName(p person) string{
// 	return p.crush.name
// }

// func main(){
// 	nameA := person{
// 		name: "Joanna",
// 		age: 32,
// 	}
// 	nameB := person{
// 		name: "David",
// 		age: 37,
// 	}
// 	nameC := person{
// 		name: "Bob",
// 	}
// 	nameD := person{
// 		name: "Amy",
// 	}

// 	getCrush(&nameA, &nameC)
// 	getCrush(&nameB, &nameD)

// 	fmt.Println(getCrushName(nameA))
// }

// type person struct{
// 	name string
// 	age uint
// 	crush *person
// }

// func (p person) getCrushName() string{
// 	if p.crush == nil{
// 		return "沒有"
// 	}
// 	return p.crush.name
// }

// func main(){
// 	nameB := person{
// 		name: "David",
// 		age: 37,
// 	}

// 	fmt.Println(nameB.getCrushName())
// }

// type info struct {
// 	name string
// 	age uint
// }

// func main(){
// 	var myInfo info
// 	myInfo.name = "Joanna"
// 	myInfo.age = 30

// 	fmt.Println("myInfo: ", myInfo)
// }

// func main(){
// 	myInfo := struct{
// 		name string
// 		age uint
// 	}{
// 		name: "Joanna",
// 		age: 30,
// 	}

// 	fmt.Println("myInfo: ", myInfo)
// }

// type rectangle struct{
// 	width float64
// 	height float64
// }

// type circle struct{
// 	radius float64
// }

// func area(r *rectangle) float64{
// 	return r.width * r.height
// }

// func (c *circle) area() float64{
// 	return c.radius * c.radius * 3.1314
// }

// func main(){
// box := &rectangle{width: 8, height: 20}
// cir := &circle{radius: 100}

// fmt.Printf("面積%f: \n", area(box))
// fmt.Printf("面積%f: \n", area(cir))
// }

// type calFunc interface{
// 	area() float64
// }

// type rectangle struct{
// 	width float64
// 	height float64
// }

// type circle struct{
// 	radius float64
// }

// func (r *rectangle) area() float64{
// 	return r.width * r.height
// }

// func (c *circle) area() float64{
// 	return c.radius * c.radius * 3.1314
// }

// func printInfo(result calFunc){
// 	fmt.Println(result.area())
// }

// func main(){
// 	box := &rectangle{width: 8, height: 20}
// 	cir := &circle{radius: 100}

// 	printInfo(box)
// 	printInfo(cir)
// }

/*
	sort 要求的 interface

	type Interface interface {
		Len() int
		Less(i, j int) bool
		Swap(i, j int)
	}
*/

// type mySortValue []int

// func (nums mySortValue) Len() int{
// 	return len(nums)
// }

// func (nums mySortValue) Less (i, j int) bool{
// 	return nums[i] < nums[j]
// }

// func (nums mySortValue) Swap(i, j int){
// 	nums[i], nums[j] = nums[j], nums[i]
// }

// func main(){
// 	nums :=[]int{3,8,2,96,33,42}

// 	// 轉成自定義型別，才能新增方法：
// 	testData := mySortValue(nums)

// 	fmt.Println(testData)
// 	sort.Sort(testData) // 滿足 sort interface 要求，執行 sort.Sort()
// 	fmt.Println(testData)
// }

/*
	sort 要求的 interface

	type Interface interface {
		Len() int
		Less(i, j int) bool
		Swap(i, j int)
	}
*/

// type sortType []int

// func (nums sortType) Len() int{
// 	return len(nums)
// }

// func (nums sortType) Less(i, j int) bool{
// 	return nums[i] < nums[j]
// }

// func (nums sortType) Swap(i, j int){
// 	nums[i], nums[j] = nums[j], nums[i]
// }

// func main(){

// 	numbers := []int{39, 11, 45, 73, 28, -5}

// 	testVal := sortType(numbers)

// 	fmt.Printf("%v\n", testVal)
// 	sort.Sort(testVal)
// 	fmt.Printf("%v\n", testVal)
// 	fmt.Println(prac.Mutil(100, 200))
// }

// func main(){
// 	var num int
// 	fmt.Println("請輸入數字：")

// 	n, err :=fmt.Scanf("%d:", &num)

// 	if err != nil{
// 		fmt.Println("錯誤訊息：", err)
// 		fmt.Println("您輸入錯誤")
// 	}else{
// 		fmt.Println("您輸入的有效參數，數量為：", n)
// 		fmt.Println("您輸入的是：", num)
// 	}
// }

// func main(){
//     var num int
//     fmt.Print("請輸入一個整數：")
//     n, err := fmt.Scanf("%d", &num)
//     if err != nil{    // err != nil 表示有錯誤發生
//         fmt.Println("輸入有效的參數數量為", n)
//         fmt.Println("錯誤訊息", err)
//         fmt.Println("輸入格式有誤")
//     }else{
//         fmt.Println("輸入有效的參數數量為", n)
//         fmt.Println("你輸入的是", num)
//     }
// }

// func handleCal(numA, numB int) (int, error){
// 	if numA == 0 || numB == 0{
// 		return 0, errors.New("無效數字，請確認！")
// 	}else{
// 		return numA/numB, nil
// 	}
// }

// func main(){
//     var numA, numB int

// 	fmt.Println("請輸入兩個整數：")
// 	_, err := fmt.Scanf("%d %d", &numA, &numB)

// 	if err != nil{
// 		fmt.Println("格式錯誤！")
// 		return
// 	}

// 	ans, err := handleCal(numA, numB)

// 	if err!= nil{
// 		fmt.Println("Error: ", err)
// 	}else{
// 		fmt.Printf("%d/%d = %d", numA, numB, ans)
// 	}
// }

// func handleCal(numA, numB int) (int, error){
// 	if numA == 0 || numB == 0{
// 		return 0, fmt.Errorf("無效數字，請確認！")
// 	}else{
// 		return numA/numB, nil
// 	}
// }

// func main(){
//     var numA, numB int

// 	fmt.Println("請輸入兩個整數：")
// 	_, err := fmt.Scanf("%d %d", &numA, &numB)

// 	if err != nil{
// 		fmt.Println("格式錯誤！")
// 		return
// 	}

// 	ans, err := handleCal(numA, numB)

// 	if err!= nil{
// 		fmt.Println("Error: ", err)
// 	}else{
// 		fmt.Printf("%d/%d = %d", numA, numB, ans)
// 	}
// }

// type error interface{
// 	Error() string
// }

// type myError struct{
// 	msg string
// }

// func (my *myError) Error() string{
// 	return my.msg
// }

// func mutilple(a, b uint) (uint, error){
// 	if a > 100 || b > 100{
// 		return 0, &myError{msg:"超過限制範圍！！"}
// 	}else{
// 		return a*b, nil
// 	}
// }

// func main(){
// 	var numA, numB uint
// 	fmt.Println("請輸入兩個數字：")
// 	_, err :=fmt.Scanf("%d %d", &numA, &numB)

// 	if err != nil{
// 		fmt.Printf("%s", err)
// 	}

// 	ans, err := mutilple(numA, numB)
// 	if err != nil{
// 		fmt.Println(err)
// 	}else{
// 		fmt.Printf("%d * %d = %d", numA, numB, ans)
// 	}
// }

// func mutilple(a, b uint) uint{
// 	if a > 100 || b > 100{
// 		panic("強制停止執行！")
// 	}else{
// 		return a*b
// 	}
// }

// func main(){
// 	var numA, numB uint
// 	fmt.Println("請輸入兩個數字：")
// 	_, err :=fmt.Scanf("%d %d", &numA, &numB)

// 	if err != nil{
// 		fmt.Printf("%s", err)
// 	}

// 	fmt.Printf("%d * %d = %d", numA, numB, mutilple(numA, numB))
// }

// func nums(number...uint) uint{
// 	var s uint
// 	for _, value := range number{
// 		s += value
// 	}
// 	return s
// }

// func main(){
// 	fmt.Println(nums(2, 4, 6))
// 	fmt.Println(nums(12, 34, 55, 9, 73, 28))
// }

// func sub(){
// 	for i:=0; i< 10; i++{
// 		fmt.Printf("副執行續 %d\n", i)
// 	}
// }

// func main(){
// 	go sub()
// 	for i:=0; i< 10; i++{
// 		fmt.Printf("主執行續 %d\n", i)
// 	}
// 	time.Sleep(3* time.Second)
// }

func sub(chanValue chan string){
	for i:=0; i< 10; i++{
		fmt.Printf("副執行續 %d\n", i)
	}
	chanValue <- "done"
}

func main(){
	var chanValue = make(chan string)
	go sub(chanValue)
	<- chanValue

	for i:=0; i< 10; i++{
		fmt.Printf("主執行續 %d\n", i)
	}

}