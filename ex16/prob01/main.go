package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type Person struct {
	Title      string `minStrlen:"1"`
	FirstName  string `minStrlen:"5"`
	MiddleName string
	LastName   string `minStrlen:"6"`
	Age        int
}

func main() {
	p := Person{
		Title:      "Mr.",
		FirstName:  "Jake",
		MiddleName: "Bobberick",
		LastName:   "Doe",
		Age:        25,
	}

	err := validateStringLength(p)
	if err != nil {
		fmt.Println("Validation error:", err)
	} else {
		fmt.Println("Validation passed")
	}
}

func validateStringLength(target any) error {
	// 構造体が渡されたことを検証する
	t := reflect.TypeOf(target)
	if t == nil {
		return errors.New("target is nil")
	}
	if t.Kind() != reflect.Struct {
		return errors.New("target is not a struct")
	}

	v := reflect.ValueOf(target)
	var foundErrors []error
	for i := 0; i < t.NumField(); i++ {
		currentField := t.Field(i)
		tag, ok := currentField.Tag.Lookup("minStrlen")
		// 構造体タグが付いていない、もしくは構造体タグの値が空文字であればスキップ
		if !ok || tag == "" {
			continue
		}

		// 構造体タグで指定された最低文字数を満たしていない場合、エラー扱いとする
		minLength, err := strconv.Atoi(tag)
		if err != nil {
			return errors.New("invalid minStrlen tag value")
		}
		if len(v.Field(i).String()) < minLength {
			foundErrors = append(foundErrors, errors.New(
				fmt.Errorf("min length of %s is %d", currentField.Name, minLength).Error(),
			))
		}
	}

	if len(foundErrors) > 0 {
		return errors.Join(foundErrors...)
	}

	// 問題なければnilを返す
	return nil
}
