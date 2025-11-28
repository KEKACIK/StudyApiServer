package validation

import (
	"StudyApiServer/config"
	"errors"
)

var (
	ErrValidationNameEmpty = errors.New("invalid name: empty")
)

func NameValidation(name string) error {
	if name == "" {
		return ErrValidationNameEmpty
	}

	return nil
}

var (
	ErrValidationAgeTooSmall = errors.New("invalid age: too small")
	ErrValidationAgeTooBig   = errors.New("invalid age: too big")
)

func AgeValidation(age int) error {
	if age < config.StudyStudentMinAge {
		return ErrValidationAgeTooSmall
	}
	if age > config.StudyStudentMaxAge {
		return ErrValidationAgeTooBig
	}

	return nil
}

var (
	ErrValidationSexInvalid = errors.New("invalid sex")
)

func SexValidation(sex string) error {
	sexList := [2]string{
		config.StudyStudentSexMan,
		config.StudyStudentSexWoman,
	}
	sexInList := false
	for _, value := range sexList {
		if sex == value {
			sexInList = true
			break
		}
	}

	if !sexInList {
		return ErrValidationSexInvalid
	}

	return nil
}

var (
	ErrValidationCourseTooSmall = errors.New("invalid course: too small")
	ErrValidationCourseTooBig   = errors.New("invalid course: too big")
)

func CourseValidation(course int) error {
	if course < config.StudyStudentMinCourse {
		return ErrValidationCourseTooSmall
	}
	if course > config.StudyStudentMaxCourse {
		return ErrValidationCourseTooBig
	}

	return nil
}
