package services

import (
	"log"

	"github.com/hiteshVasisht99/excel-import/employeeDao"
	"github.com/hiteshVasisht99/excel-import/models"
)

type EmployeeService struct {
	Dao employeeDao.EmployeeDao
}

func (s EmployeeService) InsertRecord(employeeData *models.Employee) error {
	err := s.Dao.InsertRecords(employeeData)
	if err != nil {
		log.Println("Error while inserting Employee data", employeeData)
		return err
	}
	log.Println("Inserted Employee Record in Employee table", employeeData)
	return nil
}

func (s EmployeeService) CacheData(employeeData *models.Employee) error {
	err := s.Dao.CacheData(employeeData)
	if err != nil {
		return err
	}
	log.Println("successfully cached data for key", employeeData.ID)
	return nil
}
