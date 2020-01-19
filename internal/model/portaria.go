package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Pessoa  model
type Pessoa struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Nome       string             `json:"nome" bson:"nome"`
	Contatos   []Contato          `json:"contatos" bson:"contatos"`
	Documentos []Documento        `json:"documentos" bson:"documentos"`
	Enderecos  []Endereco         `json:"enderecos" bson:"enderecos"`
}

// Visitante model
type Visitante Pessoa

// Documento model
type Documento struct {
	Tipo   string `json:"tipo" bson:"tipo"`
	Numero string `json:"numero" bson:"numero"`
}

// Endereco model
type Endereco struct {
	CEP         string `json:"cep" bson:"cep"`
	UF          string `json:"uf" bson:"uf"`
	Localidade  string `json:"localidade" bson:"localidade"`
	Bairro      string `json:"bairro" bson:"bairro"`
	Logradouro  string `json:"logradouro" bson:"logradouro"`
	Complemento string `json:"complemento" bson:"complemento"`
	Numero      string `json:"numero" bson:"numero"`
}

// Contato  model
type Contato struct {
	Tipo      string `json:"tipo" bson:"tipo"`
	Descricao string `json:"descricao" bson:"descricao"`
}

// Veiculo model
type Veiculo struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Marca  string             `json:"value" bson:"value"`
	Modelo string             `json:"dataType" bson:"dataType"`
	Placa  string             `json:"dataType" bson:"dataType"`
	Cor    string             `json:"dataType" bson:"dataType"`
}

// EventoPortaria model
type EventoPortaria struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,om∏itempty"`
	Usuario    string             `json:"usuario" bson:"usuario"`
	Tipo       string             `json:"tipo" bson:"tipo"`
	DataHora   time.Time          `json:"dataType" bson:"dataType"`
	Visitantes []Visitante        `json:"visitantes" bson:"visitantes"`
	Veiculos   []Veiculo          `json:"veiculos" bson:"veiculos"`
}
