package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang-clinic-app/config"
	"golang-clinic-app/models"
)

// Create patient (Receptionist only)
func CreatePatient(c *gin.Context) {
	role := c.GetString("role")
	if role != "receptionist" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only receptionists can create patients"})
		return
	}

	var p models.Patient
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	p.CreatedBy = c.GetInt("userID")

	_, err := config.DB.Exec(`
		INSERT INTO patients (name, age, gender, diagnosis, created_by)
		VALUES ($1, $2, $3, $4, $5)`,
		p.Name, p.Age, p.Gender, p.Diagnosis, p.CreatedBy)

	if err != nil {
		fmt.Println("DB ERROR:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "DB error",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient created"})
}

// Get all patients (Both roles)
func GetPatients(c *gin.Context) {
	var patients []models.Patient
	err := config.DB.Select(&patients, "SELECT * FROM patients")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch patients"})
		return
	}
	c.JSON(http.StatusOK, patients)
}

// Update patient (Both roles)
func UpdatePatient(c *gin.Context) {
	id := c.Param("id")
	var p models.Patient
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := config.DB.Exec(`
		UPDATE patients SET name=$1, age=$2, gender=$3, diagnosis=$4 WHERE id=$5`,
		p.Name, p.Age, p.Gender, p.Diagnosis, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient updated"})
}

// Delete patient (Receptionist only)
func DeletePatient(c *gin.Context) {
	role := c.GetString("role")
	if role != "receptionist" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only receptionists can delete patients"})
		return
	}

	id := c.Param("id")
	_, err := config.DB.Exec("DELETE FROM patients WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted"})
}
