package main

import (
	"fmt"
	"strconv"
)

func conversion()  {
	fmt.Println("=============INT TO FLOAT=============")
	// utk ubah tipe data ke tipe lain, strig => int
	// harus eksplisit utk menentukan tipe data tujuan
	// data itu fleksibel, konversi buat jdi sesuai kebutuhan
	var a int = 10
	// a = 10.5 //error
	var b float64 = float64(a) //int to float
	b = 10.5
	
	fmt.Println("Nilai a:", a)
	fmt.Println("Nilai b:", b)


	fmt.Println("=============INT TO STRING=============")
	var score int = 95
	var text string = strconv.Itoa(score) //int to string
	
	fmt.Println("Nilai score:", score)
	fmt.Println("Nilai text:", text + "text")


	fmt.Println("=============STRING TO INT=============")
	// parsing
	var text2 string = "100"
	// var text2 string = "100b" //error
	number, err := strconv.Atoi(text2) //string to int
	
	if err != nil {
	fmt.Println("Pesan Error: ", err.Error())	
	} else {
		fmt.Println("Angka:", number)
	}


	fmt.Println("=============BOOL TO STRING=============")
	// bool to string
	truth := true
	str := strconv.FormatBool(truth)
	fmt.Println("Boolean ke String:", str + "fix")


	fmt.Println("=============STRING TO BOOL=============")
	// string to bool
	val, _ := strconv.ParseBool(str)
	fmt.Println("String ke Boolean:",  val)	
}