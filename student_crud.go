// ทำงานกับ DB-CRUD
package main

import "gorm.io/gorm"

// ฟังก์ชันเพิ่มข้อมูลลงตาราง
func createStudent(db *gorm.DB, student *Student) error {
	return db.Create(student).Error
}

// ฟังก์ชันดึงข้อมูลทั้งหมดจากตาราง
func getAllStudent(db *gorm.DB) ([]Student, error) {
	//สร้างตัวแปรเก็บข้อมูลที่ได้การดึง
	var students []Student

	result := db.Find(&students)

	return students, result.Error
}

// ฟังก์ชันดึงข้อมูลจากตาราง เงื่อนไข PK คือ ID
func getStudentByID(db *gorm.DB, id int) (*Student, error) {
	//สร้างตัวแปรเก็บข้อมูลที่ได้การดึง
	var student Student

	result := db.First(&student, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &student, nil
}

// ฟังก์ชันดึงข้อมูลจากตาราง เงื่อนไขไม่ใช่ PK เช่น ชื่อ ตรงกับที่ส่งมาบางส่วน
func getStudentByName(db *gorm.DB, StuName string) ([]Student, error) {
	//สร้างตัวแปรเก็บข้อมูลที่ได้การดึง
	var students []Student

	result := db.Where("stu_name LIKE ?", "%"+StuName+"%").Find(&students)

	return students, result.Error
}

// ฟังก์ชันดึงข้อมูลจากตาราง เงื่อนไขไม่ใช่ PK เช่น ชั้นปี ตรงกับที่ส่งมา
func getStudentByLevel(db *gorm.DB, StuLevel int) ([]Student, error) {
	//สร้างตัวแปรเก็บข้อมูลที่ได้การดึง
	var students []Student

	result := db.Where("stu_level = ?", StuLevel).Find(&students)

	return students, result.Error
}

// ฟังก์ชัน แก้ไขข้อมูลในตาราง เงื่อนไข PK คือ ID
func updateStudent(db *gorm.DB, id int, student *Student) error {
	return db.Model(&Student{}).Where("id = ?", id).Updates(student).Error
}

// ฟังก์ชัน ลบข้อมูลในตาราง เงื่อนไข PK คือ ID
func deleteStudent(db *gorm.DB, id int) error {
	return db.Delete(&Student{}, id).Error
}
