package services

import (
	"context"
	"fmt"
	"log"

	"github.com/hiteshVasisht99/excel-import/employeeDao"
	"github.com/hiteshVasisht99/excel-import/models"
)

type EmployeeService struct {
	Dao employeeDao.EmployeeDao
}

func (s EmployeeService) InsertRecord(employeeData *models.Employee) error {
	//excluding the transactions functionality for now, we can add transactions here to make sure that data remains consistent on mysql and redis
	err := s.Dao.InsertRecords(employeeData)
	if err != nil {
		log.Println("Error while inserting Employee data", employeeData)
		return err
	}
	log.Println("Inserted Employee Record in Employee table", employeeData)

	err = s.Dao.CacheData(employeeData)
	if err != nil {
		log.Println("Error while caching Employee data", employeeData)
		return err
	}
	return nil
}

func (s EmployeeService) FindEmployee(empID string) (*models.Employee, error) {
	log.Println("Getting data from redis", empID)

	employee, err := s.Dao.FindDataFromRedis(context.Background(), empID)
	if err != nil {
		log.Println("Failed to get data from Redis, forwarding to mysql")

		employee, err = s.Dao.FindEmployee(empID)
		if err != nil {
			return nil, fmt.Errorf("error while finding employee with ID : %s", empID)
		}
		return employee, err
	}
	return employee, nil
}

func (s EmployeeService) UpdateEmployeeByID(employee *models.Employee) error {
	err := s.Dao.UpdateEmployeeByID(employee)
	if err != nil {
		return fmt.Errorf("error updating employee with ID %s: %w", employee.ID, err)
	}
	return nil
}
