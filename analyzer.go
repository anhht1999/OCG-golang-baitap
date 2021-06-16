package main

import (
    // "fmt"
	"sort"
    "time"
	"strconv"
	"strings"
)

type element struct {
    Key   string
    Value int
}
//nhóm theo thành phó
func GroupPeopleByCity(p []Person) (result map[string][]Person) {
	result = make(map[string][]Person)
	for _, person := range p {
		result[person.City] = append(result[person.City], person)
	}
	return result
}
//Đếm số người theo Job 
func GroupPeopleByJob(p []Person) (result map[string]int) {

	result = make(map[string]int)
	for _, person := range p {
		result[person.Job]++
	}
	return result
}

//Tìm 5 nghề có nhiều người làm nhất, đếm từ cao xuống thấp
func Top5JobsByNumer(values map[string]int) []string {
    var arr []element
    for k, v := range values {
        arr = append(arr, element{k, v})
    }
    sort.Slice(arr, func(i, j int) bool {
        return arr[i].Value > arr[j].Value
    })
    ranked := make([]string, len(values))
    for i, element := range arr {
        ranked[i] = element.Key
    }
    return ranked[:5]
}
//Tìm 5 thành phố có nhiều người trong danh sách ở nhất, đếm từ cao xuống thấp
func Top5CitiesByNumber(p map[string][]Person) []string {

    result := make(map[string]int)
    for k,v:= range p {
        result[k]= len(v)
    }
    var arr []element
    for k, v := range result {
        arr = append(arr, element{k, v})
    }
    sort.Slice(arr, func(i, j int) bool {
        return arr[i].Value > arr[j].Value
    })
    ranked := make([]string, len(result))
    for i, element := range arr {
        ranked[i] = element.Key
    }
    return ranked[:5]
}

//Trong mỗi thành phố, hãy tìm ra nghề nào được làm nhiều nhất
func TopJobByNumerInEachCity(p []Person) (result map[string]string){
	result = make(map[string]string)
	peopleByCity := GroupPeopleByCity(p)
	for city, people := range peopleByCity{
		peopleByJob := GroupPeopleByJob(people)
		max := 0
		var jobMax string
		for job, count := range peopleByJob{
			if count > max{
				jobMax = job
				max = count
			}
		}
		result[city] = jobMax
	}
	return result
}

//Ứng với một nghề, hãy tính mức lương trung bình
func AverageSalaryByJob(p []Person,t map[string]int) (result map[string]int) {

    result = make(map[string]int)
	for _, values := range p {
        result[values.Job] += values.Salary
	}
	for key := range t {
		result[key] = result[key]/t[key]
	}
    return result
}

//Năm thành phố có mức lương trung bình cao nhất
func CountPersonByCity(p []Person) (result map[string]int) {
	result = make(map[string]int)
	for _, person := range p {
		result[person.City]++
	}
	return result
}

func SalaryEachCity(p []Person) (result map[string]int) {
	result = make(map[string]int)
	for _, person := range p {
		result[person.City] += person.Salary
	}
	return result
}


func FiveCitiesHasTopAverageSalary(number map[string]int, salary map[string]int) (result []element)  {
    tmpMap := make(map[string]int)
	for key := range number {
		tmpMap[key] = salary[key] / number[key]
	}
	var arr []element
	for key, value := range tmpMap {
		arr = append(arr, element{key, value})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Value > arr[j].Value
	})
	result = arr[0:5]
	return result
}

//Năm thành phố có mức lương trung bình của developer cao nhất
func CountDeveloperByCity(p []Person) (result map[string]int) {
	result = make(map[string]int)
	job := "developer"
	for _, person := range p {
		if person.Job == job {
			result[person.City]++
		}
	}
	return result
}

func SalaryDeveloperByCity(p []Person) (result map[string]int) {
	result = make(map[string]int)
	job := "developer"
	for _, person := range p {
		if person.Job == job {
			result[person.City] += person.Salary
		}
	}
	return result
}

func FiveCitiesHasTopSalaryForDeveloper(number map[string]int, salary map[string]int) (result []element) {
	tmpMap := make(map[string]int)
	for key := range number {
		tmpMap[key] = salary[key] / number[key]
	}
	var arr []element
	for key, value := range tmpMap {
		arr = append(arr, element{key, value})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Value > arr[j].Value
	})
	result = arr
	return result
}

//Tuổi trung bình từng nghề nghiệp
func CountAge(birthdayString string) int {
	birthdaySlice := strings.Split(birthdayString, "-")
	by, _ :=strconv.Atoi(birthdaySlice[0])
	bm, _ :=strconv.Atoi(birthdaySlice[1])
	bd, _ :=strconv.Atoi(birthdaySlice[2])
	birthday := time.Date(by, time.Month(bm), bd, 0, 0, 0, 0, time.UTC)
	now := time.Now()	
	years := now.Year() - birthday.Year()
	if now.YearDay() < birthday.YearDay(){
		years--
	}
	return years
}

func AverageAgePerJob(p []Person) (result map[string]int){
	result = make(map[string]int)
	peopleNumberByJob := GroupPeopleByJob(p)
	sumOfPeopleAgeByJob := make(map[string]int)
	for _, person := range p{
		sumOfPeopleAgeByJob[person.Job]+=CountAge(person.Birthday)
	}
	for job, sumAge := range sumOfPeopleAgeByJob{
		result[job] = (sumAge) / (peopleNumberByJob[job]) 
	}
	return result
}

//Tuổi trung bình ở từng thành phố
func AverageAgePerCity(p []Person) (result map[string]int){
	result = make(map[string]int)
	peopleNumberByCity := CountPersonByCity(p)
	sumOfPeopleAgeByCity := make(map[string]int)
	for _, person := range p{
		sumOfPeopleAgeByCity[person.City]+=CountAge(person.Birthday)
	}
	for city, sumAge := range sumOfPeopleAgeByCity{
		result[city] = (sumAge) / (peopleNumberByCity[city]) 
	}
	return
}