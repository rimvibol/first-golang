package Models

import (
	"first-api/Config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//GetEmployees ... GetEmployees
func GetEmployees(employee *[]Employee) (err error) {
	if err = Config.DB.Find(employee).Error; err != nil {
		return err
	}
	return nil
}

// AddEmployee ... add new employee
func AddEmployee(employee *Employee) (erro error) {
	if erro = Config.DB.Create(employee).Error; erro != nil {
		return erro
	}
	return nil
}

//FindEmployeeByID ... find Employee by Id
func FindEmployeeByID(employee *Employee, id string) (erro error) {
	if erro = Config.DB.Where("id = ?", id).First(employee).Error; erro != nil {
		return erro
	}
	return nil
}

//UpdateEmployee ... Udpate Employee
func UpdateEmployee(employee *Employee, id string) (err error) {
	fmt.Println(employee)
	Config.DB.Save(employee)
	return nil
}
