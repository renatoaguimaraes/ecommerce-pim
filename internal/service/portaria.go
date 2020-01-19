package service

import (
	"github.com/renatoaguimaraes/ecm-pim/internal/model"
	"github.com/renatoaguimaraes/ecm-pim/internal/repo"
)

// PortariaService service
type PortariaService struct {
	VisitanteRepo      repo.VisitanteRepo
	VeiculoRepo        repo.VeiculoRepo
	EventoPortariaRepo repo.EventoPortariaRepo
}

func (s PortariaService) InsertVisitnte(p *model.Visitante) error {
	return s.VisitanteRepo.Insert(p)
}

func (s PortariaService) UpdateVisitnte(p *model.Visitante) error {
	return s.VisitanteRepo.Update(p)
}

func (s PortariaService) FindVisitnteByID(ID string) *model.Visitante {
	return s.VisitanteRepo.FindByID(ID)
}

func (s PortariaService) ListVisitntes(filter string) []model.Visitante {
	return s.VisitanteRepo.List(filter)
}

func (s PortariaService) InsertVeiculo(p *model.Veiculo) error {
	return s.VeiculoRepo.Insert(p)
}

func (s PortariaService) UpdateVeiculo(p *model.Veiculo) error {
	return s.VeiculoRepo.Update(p)
}

func (s PortariaService) FindVeiculoByID(ID string) *model.Veiculo {
	return s.VeiculoRepo.FindByID(ID)
}

func (s PortariaService) ListVeiculos(filter string) []model.Veiculo {
	return s.VeiculoRepo.List(filter)
}

func (s PortariaService) InsertEventoPortaria(p *model.EventoPortaria) error {
	return s.EventoPortariaRepo.Insert(p)
}

func (s PortariaService) UpdateEventoPortaria(p *model.EventoPortaria) error {
	return s.EventoPortariaRepo.Update(p)
}

func (s PortariaService) FindEventoPortariaByID(ID string) *model.EventoPortaria {
	return s.EventoPortariaRepo.FindByID(ID)
}

func (s PortariaService) ListEventoPortaria(filter string) []model.EventoPortaria {
	return s.EventoPortariaRepo.List(filter)
}