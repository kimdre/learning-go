package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
} 

// type displayer interface {
// 	Display()
// }

type outputable interface {
	saver
	// displayer
	Display()
}

// type outputable interface {
// 	Save() error
// 	Display()
// }

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("hello")

	title, content := getNoteData()
	todoText := getUserInput("Todo Text:")

	todo, err := todo.New(todoText)

	printSomething(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}
	

	err = outputData(userNote)

	if err != nil {
		return
	}
}

func printSomething(value any) {
	intVal, isInt := value.(int)

	if isInt {
		fmt.Println("Interger:", intVal)
		return
	}

	floatVal, isType := value.(float64)

	if isType {
		fmt.Println("Float:", floatVal)
		return
	}

	stringVal, isType := value.(string)

	if isType {
		fmt.Println("String:", stringVal)
		return
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Interger:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println("String:", value)
	// }
}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ",prompt)
	
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}