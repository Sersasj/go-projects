package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

// Aluno struct
type Aluno struct {
	gorm.Model
	Nome  string `json:"nome" validate:"nonzero"`
	Idade int    `json:"idade" validate:"min=0, max=100"`
	Email string `json:"email" validate:"regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
}

var Alunos []Aluno

func ValidateAluno(aluno *Aluno) error {
	if err := validator.Validate(aluno); err != nil {
		return err
	}
	return nil
}
