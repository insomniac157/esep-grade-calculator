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

	gradeCalculator.AddGrade("open source assignment", 40, Assignment)
	gradeCalculator.AddGrade("exam 1", 30, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 20, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF2(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB_Incomplete(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected := "C"

	gc := NewGradeCalculator()

	gc.AddGrade("assignment", 75, Assignment)
	gc.AddGrade("exam", 75, Exam)
	gc.AddGrade("essay", 75, Essay)

	if got := gc.GetFinalGrade(); got != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}

func TestGetGradeD(t *testing.T) {
	expected := "D"

	gc := NewGradeCalculator()

	gc.AddGrade("assignment", 65, Assignment)
	gc.AddGrade("exam", 65, Exam)
	gc.AddGrade("essay", 65, Essay)

	if got := gc.GetFinalGrade(); got != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}

func TestGradeTypeString(t *testing.T) {

	if Assignment.String() != "assignment" {
		t.Errorf("Assignment.String() incorrect: %s", Assignment.String())
	}
	if Exam.String() != "exam" {
		t.Errorf("Exam.String() incorrect: %s", Exam.String())
	}
	if Essay.String() != "essay" {
		t.Errorf("Essay.String() incorrect: %s", Essay.String())
	}
}

func TestComputeAverageSingle(t *testing.T) {
	gc := NewGradeCalculator()

	gc.AddGrade("assignment", 90, Assignment)
	
	if got := computeAverage(gc.assignments); got != 90 {
		t.Errorf("Expected 90, got %d", got)
	}
}