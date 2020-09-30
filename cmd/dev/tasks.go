package main

import (
	"strconv"

	"github.com/akhripko/gremlin-ent/src/model"
)

type Service interface {
	AddPersona(p *model.Persona) error
	AddGender(p model.Gender) error
}

func fillDB(s Service) error {
	var err error

	err = s.AddGender(model.Male)
	if err != nil {
		return err
	}
	err = s.AddGender(model.Female)
	if err != nil {
		return err
	}

	// add users
	for i := 0; i < 10; i++ {
		err = s.AddPersona(&model.Persona{
			ID:     i,
			Name:   "user" + strconv.Itoa(i),
			Age:    10 + i,
			Gender: getGender(i),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func getGender(i int) model.Gender {
	if i%2 == 0 {
		return model.Male
	}
	return model.Female
}
