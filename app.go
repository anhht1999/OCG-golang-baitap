package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/* https://stackoverflow.com/questions/45303326/how-to-parse-non-standard-time-format-from-json
"name":"Dee Leng",
"email":"dleng0@cocolog-nifty.com",
"job":"developer",
"gender":"Female",
"city":"London",
"salary":9662,
"birthdate":"2007-09-30" */
type Person struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Job      string `json:"job"`
	City     string `json:"city"`
	Salary   int    `json:"salary"`
	Birthday string `json:"birthdate"`
}

func (p *Person) String() string {
	return fmt.Sprintf("name: %s, email: %s, job: %s, city: %s, salary: %d, birthday: %s",
		p.Name, p.Email, p.Job, p.City, p.Salary, p.Birthday)
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("personsmall.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened person.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var people []Person

	json.Unmarshal(byteValue, &people)

	/*
		for i := 0; i < 10; i++ {
			fmt.Println(&people[i])
		}
	*/

	fmt.Println("______Nhóm theo thành phố______")
	peopleByCity := GroupPeopleByCity(people)
	for key, value := range peopleByCity {
		fmt.Println(key)
		for _, person := range value {
			fmt.Println("  ", (&person).Name)
		}
	}

	fmt.Println("______2.2 Đếm số CV theo Job ______")
	peopleByJob := GroupPeopleByJob(people)
	for key, value := range peopleByJob {
		fmt.Println( key ,"-",value)
	}

	fmt.Println("______2.3 Tìm 5 nghề có nhiều người làm nhất, đếm từ cao xuống thấp________")
	Top5JobsByNumer := Top5JobsByNumer(peopleByJob)[:5]
	for _, index := range Top5JobsByNumer {
		fmt.Printf("%s -> %d \n",index,peopleByJob[index])
	}

	fmt.Println("______2.4 Tìm 5 thành phố có nhiều người trong danh sách ở nhất, đếm từ cao xuống thấp________")	
	Top5CitiesByNumber := Top5CitiesByNumber(peopleByCity)[:5]
	for _, index := range Top5CitiesByNumber {
		fmt.Printf("%s - %d \n",index,len(index))
	}

	fmt.Println("______2.5 Trong mỗi thành phố, hãy tìm ra nghề nào được làm nhiều nhất________")
	TopJobByNumerInEachCity :=TopJobByNumerInEachCity(people)
	for key, city := range TopJobByNumerInEachCity{
		fmt.Println(key,"-",city)
	}

	fmt.Println("______2.6 Ứng với một nghề, hãy tính mức lương trung bình________")
	AverageSalaryByJob := AverageSalaryByJob(people, peopleByJob)
	for key := range AverageSalaryByJob{
		fmt.Println(key,"-",AverageSalaryByJob[key])
	}

	fmt.Println("______2.7 Năm thành phố có mức lương trung bình cao nhất________")
	numberEachCity := CountPersonByCity(people)
	salaryEachCity := SalaryEachCity(people)
	averageSalaryByCity := FiveCitiesHasTopAverageSalary(numberEachCity, salaryEachCity)
	for _, a := range averageSalaryByCity {
		fmt.Printf("%s : %d \n", a.Key, a.Value)
		fmt.Println("")
	}

	fmt.Println("______2.8 Năm thành phố có mức lương trung bình của developer cao nhất________")
	numberDeveloperEachCity := CountDeveloperByCity(people)
	salaryDeveloperEachCity := SalaryDeveloperByCity(people)
	averageSalaryDeveloperByCity := FiveCitiesHasTopSalaryForDeveloper(numberDeveloperEachCity, salaryDeveloperEachCity)
	for _, a := range averageSalaryDeveloperByCity {
		fmt.Printf("%s : %d \n", a.Key, a.Value)
	}

	fmt.Println("______2.9 Tuổi trung bình từng nghề nghiệp________")
	AverageAgePerJob := AverageAgePerJob(people)
	for key,age := range AverageAgePerJob{
		fmt.Print(key,"-",age ,"\n")
	}

	fmt.Println("______2.10 Tuổi trung bình ở từng thành phố________")
	AverageAgePerCity := AverageAgePerCity(people)
	for key,age := range AverageAgePerCity{
		fmt.Print(key,"-",age ,"\n")
	}

}
