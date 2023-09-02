package main

import (
	"assigment1/model"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func studentFinder(){
	//  logic untuk menampilkan/generate data peserta dari cli by name
	if len(os.Args) !=2 {
		fmt.Println("Usage: go run biodata.go <Kode Peserta>")
		return
	}
	codeToFind := os.Args[1]

	// Get the absolute file to JSON file
	jsonPath := filepath.Join("data", "participants.json")

	// Read JSON file content
	data, err := os.ReadFile(jsonPath)
	if err != nil{
		log.Fatal(err)
	}

	// var 	person model.person
	var people model.People

	// Unmarshal JSON data into person struct
	if err := json.Unmarshal(data, &people); err != nil {
		log.Fatal(err)
	}

	// Search Student by Code
	found := false

	for _, student := range people.People{
		if student.Code == codeToFind {
			found = true
			fmt.Printf("ID: %s\n", student.Id)
			fmt.Printf("Code: %s\n", student.Code)
			fmt.Printf("Nama: %s\n", student.Name)
			fmt.Printf("Alamat: %s\n", student.Address)
			fmt.Printf("Pekerjaan: %s\n", student.Occupation)
			fmt.Printf("Alasan: %s\n", student.Reason)		
			break
			
		}
	}
	if !found {
		fmt.Printf("Data Tidak Ditemukan")
}
}

func main(){
	studentFinder()
}