package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// センチネルエラー（事前に定義されたエラーで先頭はErr~で始めるのが慣習的）
var ErrInvalidId = errors.New("invalid ID")

type ErrEmptyField struct {
	EmptyField string
}

func (e ErrEmptyField) Error() string {
	return fmt.Sprintf("%v is empty", e.EmptyField)
}

func main() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	for d.More() {
		count++
		var emp Employee
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %v\n", count, err)
			continue
		}
		err = ValidateEmployee(emp)
		if err != nil {
			switch err := err.(type) { // 型switch
			case interface{ Unwrap() []error }: // /opt/homebrew/Cellar/go/1.26.3/libexec/src/errors/join.go : Unwrap() []error
				// ValidateEmployee()でerrors.Join()で返却されている場合
				allErrs := err.Unwrap()
				var messages []string
				for _, e := range allErrs {
					messages = append(messages, e.Error())
				}
				fmt.Printf("record %d: %+v errors: %s\n", count, emp, strings.Join(messages, ", "))
				continue
			default:
				// ValidateEmployee()で単一のエラーが返却されている場合
				fmt.Printf("record %d: %+v error: %v\n", count, emp, err)
				continue
			}
		}

		fmt.Printf("record %d: %+v\n", count, emp)
	}
}

const data = `
{
	"id": "ABCD-123",
	"first_name": "Bob",
	"last_name": "Bobson",
	"title": "Senior Manager"
}
{
	"id": "XYZ-123",
	"first_name": "Mary",
	"last_name": "Maryson",
	"title": "Vice President"
}
{
	"id": "BOTX-263",
	"first_name": "",
	"last_name": "Garciason",
	"title": "Manager"
}
{
	"id": "HLXO-829",
	"first_name": "Pierre",
	"last_name": "",
	"title": "Intern"
}
{
	"id": "MOXW-821",
	"first_name": "Franklin",
	"last_name": "Watanabe",
	"title": ""
}
{
	"id": "",
	"first_name": "Shelly",
	"last_name": "Shellson",
	"title": "CEO"
}
{
	"id": "YDOD-324",
	"first_name": "",
	"last_name": "",
	"title": ""
}
`

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

var (
	validID = regexp.MustCompile(`\w{4}-\d{3}`)
)

func ValidateEmployee(e Employee) error {
	var errs []error
	if len(e.ID) == 0 {
		errs = append(errs, ErrEmptyField{EmptyField: "id"})
	}
	if !validID.MatchString(e.ID) {
		errs = append(errs, ErrInvalidId)
	}
	if len(e.FirstName) == 0 {
		errs = append(errs, ErrEmptyField{EmptyField: "first_name"})
	}
	if len(e.LastName) == 0 {
		errs = append(errs, ErrEmptyField{EmptyField: "last_name"})
	}
	if len(e.Title) == 0 {
		errs = append(errs, ErrEmptyField{EmptyField: "title"})
	}

	switch len(errs) {
	case 0:
		return nil
	case 1:
		return errs[0]
	default:
		return errors.Join(errs...)
	}
}
