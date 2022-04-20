package service

import (
	"ec2/model"
)

type PersonService interface {
	Add(person model.Person) (model.Person, error)
	Get() ([]model.Person, error)
}
