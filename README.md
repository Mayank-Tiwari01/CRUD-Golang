# Hospital Management System || A CRUD Application written in Golang

## Overview
This Hospital Management System is built using Golang with the Gin framework for routing and MySQL/Postgres with Gorm as the ORM library. It allows managing doctors and patients, including creating, updating, and retrieving their details. The system also supports searching for doctors and patients by name and retrieving patients by doctor's ID.

## Tech Stack
- **Golang**: Gin framework for routing
- **MySQL/Postgres**: Gorm as ORM library

## Requirements
1. Accept a list of doctors with details: name, address, and phone number.
2. Accept a list of patients with details: name, address, phone number, and assigned doctor’s ID.
3. Allow updating of these details.
4. Allow searching of doctors by name and patients by name.
5. Allow retrieving patients by doctor’s ID/name.
6. Store all records in the database.

## Database Schema

### Doctor
| Column      | Data Type  |
|-------------|------------|
| id          | char(5)    |
| created_at  | timestamp  |
| updated_at  | timestamp  |
| name        | varchar    |
| contact_no  | char(10)   |
| address     | varchar    |

### Patient
| Column      | Data Type  |
|-------------|------------|
| id          | char(5)    |
| created_at  | timestamp  |
| updated_at  | timestamp  |
| name        | varchar    |
| contact_no  | char(10)   |
| address     | varchar(255)|
| doctor_id   | char(5)    |

## API Endpoints

### Doctor Routes

1. **Create Doctor**
   - **POST /doctor/**
   - Request Body:
     ```json
     {
       "name": "Dr. Satoru Gojo",
       "contact_no": "1111111111",
       "address": "Tokyo"
     }
     ```

2. **Get Doctor by ID**
   - **GET /doctor/:id**

3. **Update Doctor**
   - **PATCH /doctor/:id**
   - Request Body:
     ```json
     {
       "name": "Dr. Kento Nanami",
       "contact_no": "2222222222",
       "address": "Shibuya"
     }
     ```

4. **Delete Doctor**
   - **DELETE /doctor/:id**

5. **Search Doctor by Name**
   - **GET /searchDoctorByName**
   - Request:
     ```http
     http://localhost:8080/searchDoctorByName?name=Dr. Satoru Gojo
     ```

### Patient Routes

1. **Create Patient**
   - **POST /patient/**
   - Request Body:
     ```json
     {
       "name": "Yuji Itadori",
       "contact_no": "3333333333",
       "address": "Tokyo Prefecture",
       "doctor_id": "10001"
     }
     ```

2. **Get Patient by ID**
   - **GET /patient/:id**

3. **Update Patient**
   - **PATCH /patient/:id**
   - Request Body:
     ```json
     {
       "name": "Megumi Fushiguro",
       "contact_no": "4444444444",
       "address": "Fukuoka",
       "doctor_id": "10001"
     }
     ```

4. **Delete Patient**
   - **DELETE /patient/:id**

5. **Search Patient by Name**
   - **GET /searchPatientByName**
   - Request:
     ```http
     http://localhost:8080/searchPatientByName?name=Yuji Itadori
     ```

6. **Get Patients by Doctor ID**
   - **GET /fetchPatientByDoctorId/:doctor_id**
   - Request:
     ```http
     http://localhost:8080/fetchPatientByDoctorId/10001
     ```

## Setup and Installation

1. Clone the repository:
   ```bash
   git clone [https://github.com/yourusername/hospital-management-system.git](https://github.com/Mayank-Tiwari01/CRUD-Golang.git)
   cd CRUD-Golang

2. Set up the database:
   - Create a MySQL/Postgres database named `MayankDB`. (or make one of your choice and make the required changes in the code)
     
3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

5. The application will be running on `http://localhost:8080`.

## Testing
Use Postman to test the API endpoints. Examples of requests and responses are provided above.

## Contributing
Feel free to fork the repository and make improvements. Pull requests are welcome.
