package main

import (
	"encoding/json"
	"fmt"

	r "gopkg.in/gorethink/gorethink.v4"
)

var session *r.Session

// Reptile estructura para mappear del go a rethink
type Reptile struct {
	ID     string  `gorethink:"id,omitempty"`
	Name   string  `gorethink:"name"`
	Lenght float32 `gorethink:"length"`
	Weight float32 `gorethink:"weight"`
}

func init() {
	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "test",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {

	createTable()

	rep := Reptile{
		Name:   "Komodo",
		Lenght: 1.5,
		Weight: 45.6,
	}

	id := insertReptile(rep)

	updateReptil(id)

	rep = Reptile{
		Name:   "Python",
		Lenght: 17.0,
		Weight: 70.7,
	}
	insertReptile(rep)

	getAllReptiles()

}

// Crea la tabla persona
func createTable() {
	result, err := r.DB("test").TableCreate("reptiles").RunWrite(session)
	if err != nil {
		fmt.Println(err)
	}

	printResult("Create Table", result)
}

// Inserta un reptil
func insertReptile(reptile Reptile) string {

	result, err := r.Table("reptiles").Insert(reptile).RunWrite(session)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	printResult("Insert", result)

	return result.GeneratedKeys[0]
}

// Actualiza un reptil
func updateReptil(id string) {
	var data = map[string]interface{}{
		"Name": "Komodo dragon",
	}

	result, err := r.Table("reptiles").Get(id).Update(data).RunWrite(session)
	if err != nil {
		fmt.Println(err)
		return
	}

	printResult("Update", result)
}

// Obtiene todos los reptiles
func getAllReptiles() {
	cursor, err := r.Table("reptiles").Run(session)
	if err != nil {
		fmt.Println(err)
		return
	}

	var reptiles []Reptile
	cursor.All(&reptiles)
	cursor.Close()

	printResult("Get All", nil)
	for _, r := range reptiles {
		printObject(r)
	}
}

// Borra un reptil
func deleteReptile(id string) {
	result, err := r.Table("reptiles").Get(id).Delete().Run(session)
	if err != nil {
		fmt.Println(err)
		return
	}

	printResult("Delete", result)
}

//
func printResult(method string, v interface{}) {
	fmt.Println(fmt.Sprintf("%s successful...", method))
	if v != nil {
		printObject(v)
	}
}

//
func printObject(v interface{}) {
	vBytes, _ := json.Marshal(v)
	fmt.Println(string(vBytes))
}
