package models

// Student is a model for a student in the database.
// a student has many courses, and a course has many students.
// a student has many scores, and a score belongs to a student on a task.
type Student struct {
	User
	//student's student no
	StudentNo string  `gorm:"uniqueIndex" json:"studentNo,omitempty"`
	Wechat    *string `json:"wechat,omitempty"`
	//student's password in hash
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// This is a "has many" relationship. A student has many courses.
	// This is the "one" side of the relationship.
	Courses []Course `gorm:"many2many:student_courses;" json:"courses,omitempty"`
}
