// ทำงานกับ API
package main

import (
	"errors"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func studentRoute(app *fiber.App, db *gorm.DB) {
	//API Post for Create (Endpoint:/api/students)
	app.Post("/api/students", func(c fiber.Ctx) error {
		//ตัวแปรเก็บข้อมูลที่ส่งมาจาก client-user
		student := new(Student)

		//นำข้อมูลที่ส่งมาเก็บในตัวแปร พร้อมทั้งตรวจสอบ
		if error := c.Bind().Body(student); error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}

		//นำข้อมูลบันทึกลงตารางใน database พร้อมทั้งตรวจสอบ
		if error := createStudent(db, student); error != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}

		//กรณีบันทึกลงเรียบร้อยแล้ว
		return c.Status(201).JSON(fiber.Map{
			"message": "Student created successfully",
			"data":    student,
		})
	})

	//API Get for Read (Endpoint:/api/students)
	app.Get("/api/students", func(c fiber.Ctx) error {
		// ตัวแปรเก็บข้อมูลที่จะส่งกลับไปยัง client-user
		var students []Student

		//-----------------------------------------------------
		// สร้างตัวแปรเก็บข้อมูลที่เป็นเงื่อนไขที่ส่งมาแบบ query string
		StuName := c.Query("StuName")
		StuLevel := c.Query("StuLevel")

		var error error

		if StuName != "" {
			//ดึงข้อมูลเงื่อนไข StuName
			students, error = getStudentByName(db, StuName)
		} else if StuLevel != "" {
			//ดึงข้อมูลเงื่อนไข StuLevel
			StuLevelInt, err := strconv.Atoi(StuLevel)
			if err != nil {
				return c.Status(400).JSON(fiber.Map{
					"message": "Invalid request parameter",
				})
			}
			students, error = getStudentByLevel(db, StuLevelInt)
		} else {
			//ดึงข้อมูลจากตาราง เก็บในตัวแปร
			students, error = getAllStudent(db)
		}
		//-----------------------------------------------------

		// ตรวจสอบว่ามีข้อมูลหรือไม่
		if errors.Is(error, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"message": "Student not found",
			})
		}

		//ตรวจสอบว่ามีการทำงานผิดพลาดหรือไม่
		if error != nil {
			log.Printf("%v", error)
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "Success",
			"data":    students,
		})
	})

	//API Get for Read (Endpoint:/api/students/:id)
	//เงื่อนไขส่งมาแบบ path parameter ส่งต่อท้าย Endpoint
	app.Get("/api/students/:id", func(c fiber.Ctx) error {
		// สร้างตัวเก็บ PK คือ ID ที่ส่งมา
		id := c.Params("id")

		// แปลง id เป็นตัวเลข
		idInt, error := strconv.Atoi(id)
		if error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request parameter",
			})
		}

		// ดึงข้อมูลเก็บในตัวแปร
		student, error := getStudentByID(db, idInt)

		// ตรวจสอบว่ามีข้อมูลหรือไม่
		if errors.Is(error, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"message": "Student not found",
			})
		}

		// มีการทำงานผิดพลาดหรือไม่
		if error != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "Success",
			"data":    student,
		})
	})

	//API Put for Update (Endpoint:/api/students/:id)
	app.Put("/api/students/:id", func(c fiber.Ctx) error {
		// สร้างตัวเก็บ PK คือ ID ที่ส่งมา
		id := c.Params("id")

		// แปลง id เป็นตัวเลข
		idInt, error := strconv.Atoi(id)
		if error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request parameter",
			})
		}

		//ตัวแปรเก็บข้อมูลที่ส่งมาจาก client-user
		student := new(Student)

		//นำข้อมูลที่ส่งมาเก็บในตัวแปร พร้อมทั้งตรวจสอบ
		if error := c.Bind().Body(student); error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}

		//นำข้อมูลบันทึกแก้ไขของตารางใน database พร้อมทั้งตรวจสอบ
		if error := updateStudent(db, idInt, student); error != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}

		//กรณีบันทึกลงเรียบร้อยแล้ว
		return c.Status(200).JSON(fiber.Map{
			"message": "Student updated successfully",
			"data":    student,
		})
	})

	//API Delete for Delete (Endpoint:/api/students/:id)
	app.Delete("/api/students/:id", func(c fiber.Ctx) error {
		// สร้างตัวเก็บ PK คือ ID ที่ส่งมา
		id := c.Params("id")

		// แปลง id เป็นตัวเลข
		idInt, error := strconv.Atoi(id)
		if error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request parameter",
			})
		}

		//ลบข้อมูลออกจากตารางใน database พร้อมทั้งตรวจสอบ
		if error := deleteStudent(db, idInt); error != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}

		//กรณีบันทึกลงเรียบร้อยแล้ว
		return c.Status(200).JSON(fiber.Map{
			"message": "Student delete successfully",
		})
	})
}
