package main

import "fmt"

func main() {
	 //MAIN TYPES
	 //string
	 //bool
	 //int
	 //int int8 int16 int32 int64
	 //uint uint8 uint16 uint32 uint64 uintptr
	 //byte - alias for uint8
	 //rune - alias for int32
	 //float float64
	 //complex64 complex128


	 //using the var keyword
	 var name string = "Daniel"
	 var age int32 = 40
	 const isCool = true 
	 var size float32 = 4.4
      
	 //shorthand 
	 
	 firstName := "daniell"
	 email := "uwezukwechibuzor@gmail.com"

     surName, location := "Chibuzor", "Nigeria"
  
	  fmt.Println(name, age, isCool) 
	  fmt.Println(firstName, email) 
	  fmt.Println(surName, location) 
	  fmt.Printf("%T\n", size) 
 }











