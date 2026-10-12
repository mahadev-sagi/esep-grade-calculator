package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 50, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 50, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"
	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("assignment", 75, Assignment)
	gradeCalculator.AddGrade("exam", 75, Exam)
	gradeCalculator.AddGrade("essay", 75, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"
	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("assignment", 65, Assignment)
	gradeCalculator.AddGrade("exam", 65, Exam)
	gradeCalculator.AddGrade("essay", 65, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}


func TestAssignmentString(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Errorf("Expected Assignment to return 'assignment'; got '%s'", Assignment.String())
	}
	if Exam.String() != "exam" {
		t.Errorf("Expected Exam to return 'exam'; got '%s'", Exam.String())
	}
	if Essay.String() != "essay" {
		t.Errorf("Expected Essay to return 'essay'; got '%s'", Essay.String())
	}
}

