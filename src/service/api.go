package service

import (
	"github.com/akhripko/gremlin-ent/src/model"
)

func (s *Service) AddPersona(p *model.Persona) error {
	s.client.Persona.Create().
		SetID(p.ID).
		SetName(p.Name).
		SetAge(p.Age).
		SetGender(string(p.Gender))
	return nil
}

func (s *Service) AddGender(p model.Gender) error {
	s.client.Gender.Create().
		SetValue(string(p))
	return nil
}
