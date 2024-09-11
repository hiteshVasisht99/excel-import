package employeeDao

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hiteshVasisht99/excel-import/models"
	"github.com/redis/go-redis/v9"
)

type EmployeeDao struct {
	Db  *sql.DB
	Rdb *redis.Client
}

func (e EmployeeDao) InsertRecords(employeeData *models.Employee) error {
	query := `INSERT INTO EMPLOYEE (ID,first_name, last_name,company_name, address, city, country, postal, phone, email, web) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?,?,?,?)`
	_, err := e.Db.Exec(query,
		employeeData.ID,
		employeeData.FirstName,
		employeeData.LastName,
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

func (e EmployeeDao) FindEmployee(employeeID string) (*models.Employee, error) {

	query := `SELECT * FROM EMPLOYEE WHERE ID = ?`
	employee := &models.Employee{}
	err := e.Db.QueryRow(query, employeeID).Scan(
		&employee.ID,
		&employee.FirstName,
		&employee.LastName,
		&employee.CompanyName,
		&employee.Address,
		&employee.City,
		&employee.Country,
		&employee.Postal,
		&employee.Phone,
		&employee.Email,
		&employee.Web)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return employee, nil
}

func (e EmployeeDao) FindDataFromRedis(ctx context.Context, employeeID string) (*models.Employee, error) {
	result, err := e.Rdb.Get(context.Background(), employeeID).Result()
	if err != nil || result == "" {
		log.Println("Failed to get data from redis")
		return nil, err
	} else {
		var employee models.Employee
		err = json.Unmarshal([]byte(result), &employee)
		if err != nil {
			log.Println("Failed to unmarshall data received from redis")
			return nil, err
		}
		return &employee, nil

	}
}

func (e EmployeeDao) UpdateEmployeeByID(employee *models.Employee) error {
	query := `
        UPDATE EMPLOYEE
        SET first_name = ?, 
            last_name = ?, 
            company_name = ?, 
            address = ?, 
            city = ?, 
            country = ?, 
            postal = ?, 
            phone = ?, 
            email = ?, 
            web = ?
        WHERE ID = ?`

	result, err := e.Db.Exec(query,
		employee.FirstName,
		employee.LastName,
		employee.CompanyName,
		employee.Address,
		employee.City,
		employee.Country,
		employee.Postal,
		employee.Phone,
		employee.Email,
		employee.Web,
		employee.ID,
	)
	if err != nil {
		return err
	}

	// Check if any rows were affected (i.e., if the employee exists)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no employee found with ID %s", employee.ID)
	}

	//Update the cache with the new data
	employeeJSON, err := json.Marshal(employee)
	if err != nil {
		return fmt.Errorf("error marshalling employee data for cache: %w", err)
	}
	err = e.Rdb.Set(context.Background(), employee.ID, employeeJSON, 5*time.Minute).Err()
	if err != nil {
		return fmt.Errorf("error setting employee in cache: %w", err)
	}

	return nil
}
