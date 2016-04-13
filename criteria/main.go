/*
See: https://en.wikipedia.org/wiki/Criteria_Pattern
See: http://www.tutorialspoint.com/design_pattern/filter_pattern.htm
*/

package main

import (
	"fmt"
	"strings"
)

type person struct {
	name          string
	gender        string
	maritalStatus string
}

func (p person) String() string {
	return fmt.Sprintf("Person : [Name : %s, Gender : %s, Marital Status : %s]\n", p.name, p.gender, p.maritalStatus)
}

type criteria interface {
	MeetCriteria([]person) []person
}

type maleCriteria struct{}

func (m maleCriteria) MeetCriteria(people []person) []person {
	males := []person{}

	for _, person := range people {
		if strings.ToUpper(person.gender) == "MALE" {
			males = append(males, person)
		}
	}

	return males
}

type femaleCriteria struct{}

func (f femaleCriteria) MeetCriteria(people []person) []person {
	females := []person{}

	for _, person := range people {
		if strings.ToUpper(person.gender) == "FEMALE" {
			females = append(females, person)
		}
	}

	return females
}

type singleCriteria struct{}

func (s singleCriteria) MeetCriteria(people []person) []person {
	singles := []person{}

	for _, person := range people {
		if strings.ToUpper(person.maritalStatus) == "SINGLE" {
			singles = append(singles, person)
		}
	}

	return singles
}

type andCriteria struct {
	criteria criteria
	other    criteria
}

func (a andCriteria) MeetCriteria(people []person) []person {
	firstCriteria := a.criteria.MeetCriteria(people)
	return a.other.MeetCriteria(firstCriteria)
}

type orCriteria struct {
	criteria criteria
	other    criteria
}

func (o orCriteria) MeetCriteria(people []person) []person {
	firstCriteria := o.criteria.MeetCriteria(people)
	otherCriteria := o.other.MeetCriteria(people)

	for _, p := range otherCriteria {
		if !contains(firstCriteria, p) {
			firstCriteria = append(firstCriteria, p)
		}
	}

	return firstCriteria
}

func contains(people []person, aPerson person) bool {
	for _, p := range people {
		if p.name == aPerson.name {
			return true
		}
	}

	return false
}

func main() {
	people := []person{
		{"Robert", "Male", "Single"},
		{"John", "Male", "Married"},
		{"Laura", "Female", "Married"},
		{"Diana", "Female", "Single"},
		{"Mike", "Male", "Single"},
		{"Bobby", "Male", "Single"},
	}

	male := maleCriteria{}
	female := femaleCriteria{}
	single := singleCriteria{}
	singleMale := andCriteria{single, male}
	singleOrFemale := orCriteria{single, female}

	fmt.Println("Males: ")
	fmt.Println(male.MeetCriteria(people))
	fmt.Println()

	fmt.Println("Females: ")
	fmt.Println(female.MeetCriteria(people))
	fmt.Println()

	fmt.Println("Single Males: ")
	fmt.Println(singleMale.MeetCriteria(people))
	fmt.Println()

	fmt.Println("Single Or Females: ")
	fmt.Println(singleOrFemale.MeetCriteria(people))
	fmt.Println()
}
