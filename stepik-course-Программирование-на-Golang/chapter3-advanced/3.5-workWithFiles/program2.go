package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

type notCsvError string

func (e notCsvError) Error() string {
	return fmt.Sprintf("parse %v: internal error", string(e))
}

const dir = "./task"

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data, err := csv.NewReader(file).ReadAll()
	if err != nil {
		// 1) это создание пустой структуры (ошибка это структура)
		// 2) нужно создать именно ссылку на структуру, так как интерфейс Error работает именно с ней.
		return notCsvError(path)
	}

	if len(data) == 1 && len(data[0]) == 1 {
		return nil
	}

	fmt.Println(file.Name(), data[4][2])

	return nil
}

func Program2() {
	if err := filepath.Walk(dir, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}
}
