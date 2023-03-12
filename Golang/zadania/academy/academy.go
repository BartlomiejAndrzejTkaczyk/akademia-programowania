package academy

import (
	"math"
)

type Student struct {
	Name      string
	Grades    []int
	Project   int
	Attendace []bool
}

// AverageGrade returns an average grade given a
// slice containing all grades received during a
// semester, rounded to the nearest integer.
func AverageGrade(grades []int) int {
	length := len(grades)
	if length == 0 {
		return 0
	}

	sum := 0

	for _, grade := range grades {
		sum += grade
	}

	return int(
		math.Round(float64(sum) / float64(length)),
	)
}

// AttendancePercentage returns a percentage of class
// attendance, given a slice containing information
// whether a student was present (true) of absent (false).
//
// The percentage of attendance is represented as a
// floating-point number ranging from  0 to 1,
// with 2 digits of precision.
func AttendancePercentage(attendance []bool) float64 {
	length := len(attendance)

	if length == 0 {
		return 0
	}

	present := 0

	for _, value := range attendance {
		if value {
			present++
		}
	}

	presentPercent := float64(present) / float64(length)

	return math.Round(presentPercent*1000) / 1000
}

// FinalGrade returns a final grade achieved by a student,
// ranging from 1 to 5.
//
// The final grade is calculated as the average of a project grade
// and an average grade from the semester, with adjustments based
// on the student's attendance. The final grade is rounded
// to the nearest integer.
//
// If the student's attendance is below 80%, the final grade is
// decreased by 1. If the student's attendance is below 60%, average
// grade is 1 or project grade is 1, the final grade is 1.
func FinalGrade(s Student) int {
	averageGrade := AverageGrade(s.Grades)
	attendance := AttendancePercentage(s.Attendace)

	if attendance < 0.6 || averageGrade == 1 || s.Project == 1 {
		return 1
	}

	finalGrade := AverageGrade([]int{averageGrade, s.Project})

	if attendance < 0.8 {
		finalGrade--
	}

	return finalGrade
}

// GradeStudents returns a map of final grades for a given slice of
// Student structs. The key is a student's name and the value is a
// final grade.
func GradeStudents(students []Student) map[string]uint8 {
	studentToGarde := make(map[string]uint8)

	for _, s := range students {
		studentToGarde[s.Name] = uint8(FinalGrade(s))
	}

	return studentToGarde
}
