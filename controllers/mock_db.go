package controllers

import (
	"CRUD-Golang/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

// Mock database functions
var mockDoctors = make(map[string]models.Doctor)
var mockPatients = make(map[string]models.Patient)

func MockDBCreateDoctor(doctor *models.Doctor) *gorm.DB {
	doctor.ID = "mock_id_" + doctor.ID
	doctor.CreatedAt = time.Now()
	doctor.UpdatedAt = time.Now()
	mockDoctors[doctor.ID] = *doctor
	return &gorm.DB{}
}

func MockDBFindDoctorByID(id string) (models.Doctor, error) {
	if doctor, ok := mockDoctors[id]; ok {
		return doctor, nil
	}
	return models.Doctor{}, errors.New("doctor not found")
}

func MockDBSaveDoctor(doctor *models.Doctor) *gorm.DB {
	mockDoctors[doctor.ID] = *doctor
	return &gorm.DB{}
}

func MockDBDeleteDoctor(id string) error {
	if _, ok := mockDoctors[id]; ok {
		delete(mockDoctors, id)
		return nil
	}
	return errors.New("doctor not found")
}

func MockDBSearchDoctorsByName(name string) []models.Doctor {
	var doctors []models.Doctor
	for _, doctor := range mockDoctors {
		if doctor.Name == name {
			doctors = append(doctors, doctor)
		}
	}
	return doctors
}

// Similarly, create mock functions for patients
func MockDBCreatePatient(patient *models.Patient) *gorm.DB {
	patient.ID = "mock_id_" + patient.ID
	patient.CreatedAt = time.Now()
	patient.UpdatedAt = time.Now()
	mockPatients[patient.ID] = *patient
	return &gorm.DB{}
}

func MockDBFindPatientByID(id string) (models.Patient, error) {
	if patient, ok := mockPatients[id]; ok {
		return patient, nil
	}
	return models.Patient{}, errors.New("patient not found")
}

func MockDBSavePatient(patient *models.Patient) *gorm.DB {
	mockPatients[patient.ID] = *patient
	return &gorm.DB{}
}

func MockDBDeletePatient(id string) error {
	if _, ok := mockPatients[id]; ok {
		delete(mockPatients, id)
		return nil
	}
	return errors.New("patient not found")
}

func MockDBSearchPatientsByName(name string) []models.Patient {
	var patients []models.Patient
	for _, patient := range mockPatients {
		if patient.Name == name {
			patients = append(patients, patient)
		}
	}
	return patients
}
