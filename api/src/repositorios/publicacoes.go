package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(Publicacao modelos.Publicacao) (uint64, error) {
	statement, err := repositorio.db.Prepare(
		"insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	resultado, err := statement.Exec(Publicacao.Titulo, Publicacao.Conteudo, Publicacao.AutorID)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil
}
