package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

type Book struct {
	Title string `xml:"title"`
	Year  int    `xml:"year"`
}
type Library struct {
	Books []Book `xml:"book"`
}

func main() {
	// task1()
	task2()
}
func task2() {
	var lib Library
	data, err := os.ReadFile("books.xml")
	if err != nil {
		log.Fatal(err)
	} else {
		err := xml.Unmarshal(data, &lib)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(lib)
	for i := 0; i < len(lib.Books); i++ {
		lib.Books[i].incrementYear()
	}
	fmt.Println(lib)
	//writeXml("books.xml",lib)
}
func (b *Book) incrementYear() {
	b.Year++
}

func task1() {
	people := make([]Person, 0)
	data, err := os.ReadFile("users.json")
	if err != nil {
		log.Fatal(err)
	} else {
		err := json.Unmarshal(data, &people)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(people)
	for i := 0; i < len(people); i++ {
		people[i].incrementAge()
	}
	fmt.Println(people)
	//writeJson("users.json",people)

}
func (p *Person) incrementAge() {
	p.Age++
}
func writeXml(path string, l Library) {
	data, err := xml.Marshal(l)
	if err != nil {
		log.Fatal(err)
	} else {
		err = os.WriteFile(path, data, 0644)
	}
}

func writeJson(path string, people []Person) {
	data, err := json.Marshal(people)
	if err != nil {
		log.Fatal(err)
	} else {
		err = os.WriteFile(path, data, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}
