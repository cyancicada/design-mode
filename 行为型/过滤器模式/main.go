package main

import (
	"fmt"

	"design-mode/行为型/过滤器模式/filter"
)

func main() {
	var persons []filter.Person
	persons = append(persons, filter.GetPerson("Robert", "Male", "Single"))
	persons = append(persons, filter.GetPerson("John", "Male", "Married"))
	persons = append(persons, filter.GetPerson("Laura", "Female", "Married"))
	persons = append(persons, filter.GetPerson("Diana", "Female", "Single"))
	persons = append(persons, filter.GetPerson("Mike", "Male", "Single"))
	persons = append(persons, filter.GetPerson("Bobby", "Male", "Single"))
	male := new(filter.CriteriaMale)
	fmt.Println(male.MeetCriteria(persons))
	female := new(filter.CriteriaFemale)
	fmt.Println(female.MeetCriteria(persons))
	single := new(filter.CriteriaSingle)
	fmt.Println(single.MeetCriteria(persons))
	singleMale := new(filter.AndCriteria)
	singleMale.AndCriteria(single, male)
	fmt.Println(singleMale.MeetCriteria(persons))
	singleFemale := new(filter.AndCriteria)
	singleFemale.AndCriteria(single, female)
	fmt.Println(singleFemale.MeetCriteria(persons))
}
