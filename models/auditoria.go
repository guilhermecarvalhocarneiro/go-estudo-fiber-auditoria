package models

import (
	"github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

type Auditoria struct {
	gorm.Model
	IpAddress    string         `json:"ip_address" gorm:"index"`
	NomeSistema  string         `json:"nome_sistema" gorm:"index"`
	AppName      string         `json:"app_name" gorm:"index"`
	Modulo       string         `json:"modulo" gorm:"index"`
	Operacao     string         `json:"operacao" gorm:"index"`
	Usuario      postgres.Jsonb `json:"usuario" gorm:"index"`
	Grupo        postgres.Jsonb `json:"grupo"`
	Permissao    postgres.Jsonb `json:"permissao"`
	DadoAnterior postgres.Jsonb `json:"dado_anterior"`
	DadoAtual    postgres.Jsonb `json:"dado_atual"`
}
