package url

import (
	"database/sql"
    "log"
)

type repositorioDatabase struct {
	conexao *sql.DB
}

func logarErro(acao string, err error) {
    if err != nil {
        log.Printf("Erro ao %s: %s\n", acao, err)
    }
}

func NovoRepositorioDatabase(conexao *sql.DB) *repositorioDatabase {
    return &repositorioDatabase{conexao}
}

func (r repositorioDatabase) IdExiste(id string) (existe bool) {
    row := r.conexao.QueryRow("SELECT id FROM url WHERE id = $1", id)
    err := row.Scan()  // TODO posso passar sem argumentos só para ver se a linha existe?
    existe = err == nil || err != sql.ErrNoRows

    if err == sql.ErrNoRows {
        err = nil
    }

    logarErro("verificar id", err)
    return existe
}

func (r repositorioDatabase) BuscarPorId(id string) (*Url) {
    row := r.conexao.QueryRow(
        "SELECT id, criacao, destino FROM url WHERE id = $1", id)
    url := Url{}
    err := row.Scan(&url.Id, &url.Criacao, &url.Destino)

    if err == sql.ErrNoRows {
        return nil
    }

    logarErro("buscar por id", err)
    return &url
}

func (r repositorioDatabase) BuscarPorUrl(urlDestino string) (*Url) {
    row := r.conexao.QueryRow(
        "SELECT id, criacao, destino FROM url WHERE destino = $1", urlDestino)
    url := Url{}
    err := row.Scan(&url.Id, &url.Criacao, &url.Destino)
    
    if err == sql.ErrNoRows {
        return nil
    }

    logarErro("buscar por URL", err)
    return &url
}

func (r *repositorioDatabase) Salvar(url Url) error {
    sql := "INSERT INTO url (id, criacao, destino) VALUES ($1, $2, $3)"
    _, err := r.conexao.Exec(sql, url.Id, url.Criacao, url.Destino)
    logarErro("inserir URL", err)
    return err
}

func (r *repositorioDatabase) RegistrarClique(id string) {
    sql := `
        INSERT INTO clique (url_id, contagem) VALUES ($1, 1)
        ON CONFLICT (url_id) DO
            UPDATE SET contagem = clique.contagem + 1
            WHERE clique.url_id = $1
    `
    _, err := r.conexao.Exec(sql, id)
    logarErro("registrar clique", err)
}

func (r *repositorioDatabase) BuscarCliques(id string) (contagem int) {
    row := r.conexao.QueryRow(
        "SELECT contagem FROM clique WHERE url_id = $1", id)
    err := row.Scan(&contagem)

    if err == sql.ErrNoRows {
        return 0
    }

    logarErro("buscar cliques", err)
    return  // contagem! ;)
}
