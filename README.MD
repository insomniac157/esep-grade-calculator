# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

Runtime requirements:
Golang 1.24+

Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```

If not assigned any values for an entry, the default is 0. So, students with no assignments, exams, and/or essays will get a 0 for the respective sections by default.

Grading Scale:
A >= 90
B >= 80
C >= 70
D >= 60
F < 60

Default weightage is 50% for assignments, 35% for exams, and 15% for essays.