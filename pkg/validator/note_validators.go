package validator

import (
	"fmt"
	"notes/pkg/request"
	"notes/pkg/xerror"
)

const (
	titleMin int = 1
	titleMax int = 50
	textMin  int = 1
	textMax  int = 500
)

func ValidateCreateNoteRequest(input request.CreateNoteRequest) error {
	e := Title(input.Title)
	if e != nil {
		return e
	}

	e = Text(input.Text)
	if e != nil {
		return e
	}

	return nil
}

func ValidateUpdateNoteRequest(input request.UpdateNoteRequest) error {
	e := ValidateId(input.Id)
	if e != nil {
		return e
	}

	e = Title(input.Title)
	if e != nil {
		return e
	}

	e = Text(input.Text)
	if e != nil {
		return e
	}

	return nil
}

func ValidateId(id int) error {
	if id <= 0 {
		return &xerror.ValidationError{Message: fmt.Sprintf("invalid id")}
	}

	return nil
}

func Title(text string) error {
	if len(text) < titleMin {
		return &xerror.ValidationError{Message: fmt.Sprintf("length of the title should be no less than %d and no more than %d characters", titleMin, titleMax)}
	}
	return nil
}

func Text(text string) error {
	if len(text) < titleMin {
		return &xerror.ValidationError{Message: fmt.Sprintf("length of the text should be no less than %d and no more than %d characters", textMin, textMax)}
	}
	return nil
}
