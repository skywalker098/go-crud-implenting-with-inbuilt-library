package controllers

import (
	"encoding/json"
	"fmt"
	"os"
)

// -------------Save to JSON --------------------------------
func (t Blogstore) SavetoJson() {
	// create files
	data, err := json.Marshal(t.Blogs)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile("db.json", data, 0o644); err != nil {
		panic(err)
	}
}

// ---------------Checck file --------------------------------
func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

// ---------------Add ID to todo --------------------------------
func (t Blogstore) AddnewId() int {
	if len(t.Blogs) == 0 {
		return 1
	}
	// return t.Blogs[len(t.Blogs)-1].Id + 1
	return t.Blogs[len(t.Blogs)-1].Id + 1

}

// --------------Load from JSON --------------------------------
func (t *Blogstore) LoadFromJson() {
	// Check if the file exists
	if _, err := os.Stat("db.json"); os.IsNotExist(err) {
		os.Create("db.json")
	}

	data, err := os.ReadFile("db.json")
	if err != nil {
		panic(err)
	}

	if len(data) > 0 {
		// Unmarshal the data from the file to t.todoStore
		err = json.Unmarshal(data, &t.Blogs)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("No data found")
	}
}
