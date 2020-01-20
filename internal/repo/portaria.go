package repo

import (
	"github.com/renatoaguimaraes/ecommerce-pim/internal/model"
)

// VisitanteRepo interface
type VisitanteRepo interface {
	Insert(p *model.Visitante) error
	Update(p *model.Visitante) error
	FindByID(ID string) *model.Visitante
	List(filter string) []model.Visitante
}

// VisitanteRepoMongo repo mongo
type VisitanteRepoMongo struct{}

func (r VisitanteRepoMongo) Insert(p *model.Visitante) error {
	return nil
}

func (r VisitanteRepoMongo) Update(p *model.Visitante) error {
	return nil
}

func (r VisitanteRepoMongo) FindByID(ID string) *model.Visitante {
	return nil
}

func (r VisitanteRepoMongo) List(filter string) []model.Visitante {
	return []model.Visitante{}
}

// VeiculoRepo interface
type VeiculoRepo interface {
	Insert(p *model.Veiculo) error
	Update(p *model.Veiculo) error
	FindByID(ID string) *model.Veiculo
	List(filter string) []model.Veiculo
}

// VeiculoRepoMongo repo mongo
type VeiculoRepoMongo struct{}

func (r VeiculoRepoMongo) Insert(m *model.Veiculo) error {
	return nil
}

func (r VeiculoRepoMongo) Update(m *model.Veiculo) error {
	return nil
}

func (r VeiculoRepoMongo) FindByID(ID string) *model.Veiculo {
	return nil
}

func (r VeiculoRepoMongo) List(filter string) []model.Veiculo {
	return nil
}

// EventoPortariaRepo interface
type EventoPortariaRepo interface {
	Insert(p *model.EventoPortaria) error
	Update(p *model.EventoPortaria) error
	FindByID(ID string) *model.EventoPortaria
	List(filter string) []model.EventoPortaria
}

// EventoPortariaRepoMongo repo mongo
type EventoPortariaRepoMongo struct{}

func (r EventoPortariaRepoMongo) Insert(m *model.EventoPortaria) error {
	return nil
}

func (r EventoPortariaRepoMongo) Update(m *model.EventoPortaria) error {
	return nil
}

func (r EventoPortariaRepoMongo) FindByID(ID string) *model.EventoPortaria {
	return nil
}

func (r EventoPortariaRepoMongo) List(filter string) []model.EventoPortaria {
	return nil
}
