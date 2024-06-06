package main

import (
	"fmt"
	"reflect"
)

/*
Package Reflect
Dalam bahasa pemograman, biasanya ada fitu reflection, dimana kita bisa melihat struktur kode kita
pada saat aplikasi sedang berjalan
Hal ini bisa dilakukan di Go-Lang dengan menggunakan package reflect
Relection sangat berguna ketika kita ingin membuat library yang general sehinggi mudah digunakan
*/

type Sample struct {
	Name string `required:"true" max:"10"` // StructTag
}

// fitur struct tag bisa digunakan sebagai validation
type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println(structField.Name, "With Type", structField.Type)
		fmt.Println(structField.Tag.Get("required")) // Mengambil struct tag required
		fmt.Println(structField.Tag.Get("max"))      // Mengambil struct tag max
	}
}

// Memanfaatkan Struct Tag sebagai validasi
func isValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface() // interface() disini adalah datanya
			fmt.Println(data)
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

func main() {
	readField(Sample{"Varrel"})
	readField(Person{"Varrel", "123", "123"})

	person := Person{
		Name:    "Varrel",
		Address: "Bandung",
		Email:   "test",
	}
	fmt.Println(isValid(person))

}
