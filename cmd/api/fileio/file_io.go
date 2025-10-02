package fileio

import (
	"bufio"
	"fmt"
	"os"
)

func WriteFile(path string, content []byte) error {
	// Write data to file (creates if not exists, overwrites if exists)
	err := os.WriteFile(path, content, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}

	fmt.Println("File written successfully!")
	return nil
}

func CreateWriteFile(path string, content []byte) error {
	file, err := os.Create(path) // creates or truncates file
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(content))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	fmt.Println("File written successfully!")
	return nil
}

func ReadFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	fmt.Println("\nFile content:")
	fmt.Println(string(data))
	return data, nil
}

func ReadFileLineByLine(path string) error {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	fmt.Println("\nFile content line by line:")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return nil
}

func Run() {
	// -- Write data to file
	WriteFile("./fileio/example.txt", []byte("Hello, Go!\nThis is a file write example."))

	// -- Create and write data to file
	CreateWriteFile("./fileio/example2.txt", []byte("Hello, Go! Using file.Write"))

	// -- Read data from file
	ReadFile("./fileio/example.txt")

	// -- Read data from file line by line
	ReadFileLineByLine("./fileio/example.txt")
}
