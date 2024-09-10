package exercises

import "fmt"

//struct er et objekt.
type Student struct {
	Name  string
	Age   int
	House string
	Color string
}

// * er en go poinyer, der finder pladsen i hukkomelsen.

func (student *Student) SetAge(age int) {
	student.Age = age

}

func TestStudents() {
	// harry := Student{
	// 	Name:  "harry",
	// 	Age:   10,
	// 	House: "SLG",
	// 	Color: "Black",
	// }

	ron := Student{
		Name:  "ron",
		Age:   12,
		House: "NV",
	}
	ron.SetAge(19)
	fmt.Println(ron)
}
