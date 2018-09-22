package main

import (
	"fmt"
	"sort"
)

type people struct {
	Name string
	Age  int
}

type students []string

func (p people) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

//ByAge uses Interface interface to sort a struct of people by age
type ByAge []people

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

//ByName uses Interface interface to sort a struct of people by name
type ByName []people

func (b ByName) Len() int           { return len(b) }
func (b ByName) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByName) Less(i, j int) bool { return b[i].Name < b[j].Name }

func main() {
	studyGroup := []people{
		{"Isaac Muthui", 30},
		{"John Nderitu", 32},
		{"Daniel Mahinda", 28},
		{"Caroline Atieno", 35},
		{"Azezue Zunma", 24},
	}

	class := students{"Zezu", "Michael", "Kenn", "Chomba", "Kiki"}
	topStudents := []string{"Ojode", "Luseno", "Karinga", "Byrone"}
	fmt.Println("Unsorted strings class")
	fmt.Println(class)

	sort.Strings(class)
	fmt.Println("Sorted strings class")
	fmt.Println(class)

	fmt.Println("Unsorted strings top students")
	fmt.Println(topStudents)
	sort.Strings(topStudents)
	fmt.Println("Sorted strings top students")
	sort.Strings(topStudents)
	fmt.Println(topStudents)
	//in reverse

	fmt.Println("In reverse")
	fmt.Println(topStudents)
	sort.Sort(sort.Reverse(sort.StringSlice(topStudents)))
	fmt.Println(topStudents)
	//Sort by Age
	fmt.Println(studyGroup)
	sort.Sort(ByAge(studyGroup))
	fmt.Println("Sort by Age")
	fmt.Println(studyGroup)
	fmt.Println("***********************************************")
	//Sort by Name
	fmt.Println(studyGroup)
	sort.Sort(ByName(studyGroup))
	fmt.Println("Sort")
	fmt.Println(studyGroup)
	fmt.Println(studyGroup)
	fmt.Println("Sort in reverse")
	sort.Slice(studyGroup, func(i, j int) bool {
		return studyGroup[i].Name > studyGroup[j].Name
	})
	fmt.Println(studyGroup)
	n := []int{34, 34, 56341, 4234, 123, 41, 4, 14}
	fmt.Println(n)
	sort.Ints(n)
	fmt.Println(n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)
}
