# Excel-import api

## Overview

This project provides  RESTful APIs for managing employee data. It allows users to upload employee data from an Excel file, retrieve employee data by ID, and update employee records.
The application uses Go for the backend, with integration for a MySQL database and Redis cache.

## Features

- **Upload Employee Data**: Upload employee data from an Excel file and process it asynchronously.
- **Get Employee Data by ID**: Retrieve employee details using the employee ID.
- **Update Employee Data by ID**: Update existing employee records using the employee ID.

## Technologies Used

- Go
- Gorilla Mux (for routing)
- MySQL 
- Redis
- UUID (for generating unique IDs)
- Excel parsing library (`common.ParseExcelFile` function)

## Setup and Installation

### Prerequisites

- Go 1.18 or higher
- MySQL (or any compatible SQL database)
- Redis

### After DB setup , please run this query to create a table:
CREATE TABLE EMPLOYEE (
    ID VARCHAR(255) NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    company_name VARCHAR(255),
    address TEXT,
    city VARCHAR(100),
    country VARCHAR(100),
    postal VARCHAR(20),
    phone VARCHAR(20),
    email VARCHAR(255),
    web VARCHAR(255),
    PRIMARY KEY (ID)
);

### You are ready to Go now....
