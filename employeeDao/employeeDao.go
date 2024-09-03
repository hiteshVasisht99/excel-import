package employeeDao

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/hiteshVasisht99/excel-import/models"
	"github.com/redis/go-redis/v9"
)

type EmployeeDao struct {
	Db  *sql.DB
	Rdb *redis.Client
}

func (e EmployeeDao) InsertRecords(employeeData *models.Employee) error {
	query := `INSERT INTO Employee (ID,first_name, company_name, address, city, country, postal, phone, email, web) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?,?,?)`
	_, err := e.Db.Exec(query,
		employeeData.ID,
		employeeData.FirstName,
		employeeData.CompanyName,
		employeeData.Address,
		employeeData.City,
		employeeData.Country,
		employeeData.Postal,
		employeeData.Phone,
		employeeData.Email,
		employeeData.Web,
	)
	return err
}

func (e EmployeeDao) CacheData(employeeData *models.Employee) error {
	bytesData, err := json.Marshal(employeeData)
	if err != nil {
		return err
	}
	err = e.Rdb.Set(context.Background(), employeeData.ID, bytesData, 5*time.Minute).Err()
	if err != nil {
		return err
	}
	return err
}
