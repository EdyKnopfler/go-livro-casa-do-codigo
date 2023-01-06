package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"
    "encoding/json"
    
    "github.com/edyknopfler/encurtador/database"
    "github.com/edyknopfler/encurtador/url"
)

var (
    porta int
    urlBase string
)

type Headers map[string]string

type Redirecionador struct{
    stats chan string
}

func init() {
    // Inicialização do pacote main
    porta = 8888
    urlBase = fmt.Sprintf("http://localhost:%d", porta)
}

func main() {
    conexao := database.Conectar()
    log.Printf("Conectado à base de dados.")
    defer conexao.Close()

    // TODO criar o novo repositório baseado em SQL
    // url.ConfigurarRepositorio(url.NovoRepositorioSQL(conexao))
    url.ConfigurarRepositorio(url.NovoRepositorioMemoria())

    stats := make(chan string)
    defer close(stats)
    go registrarEstatisticas(stats)
    
    http.HandleFunc("/api/encurtar", Encurtador)
    http.Handle("/r/", &Redirecionador{stats})
    http.HandleFunc("/api/stats/", Visualizador)

    log.Printf("Ouvindo requisições na porta %d.", porta)
    log.Fatal(http.ListenAndServe(
        fmt.Sprintf(":%d", porta), nil))
}

func Encurtador(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        responderCom(w, http.StatusMethodNotAllowed, Headers{
            "Allow": "POST"})
        return
    }
    
    url, nova, err := url.BuscarOuCriarNovaUrl(extrairUrl(r))
    
    if err != nil {
        log.Printf("%s\n", err)
        responderCom(w, http.StatusBadRequest, nil)
        return
    }
    
    var status int
    if nova {
        status = http.StatusCreated
    } else {
        status = http.StatusOK
    }
    
    urlCurta := fmt.Sprintf("%s/r/%s", urlBase, url.Id)
    responderCom(w, status, Headers{
        "Location": urlCurta,
        "Link": fmt.Sprintf("<%s/api/stats/%s>; rel=\"stats\"", urlBase, url.Id),
    })
}

func (red *Redirecionador) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    buscarUrlEExecutar(w, r, func(url *url.Url) {
        red.stats <- url.Id
        http.Redirect(w, r, url.Destino, http.StatusMovedPermanently)
    })
}

func Visualizador(w http.ResponseWriter, r *http.Request) {
    buscarUrlEExecutar(w, r, func(url *url.Url) {
        json, err := json.Marshal(url.Stats())
        
        if err != nil {
            log.Printf("%s\n", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        
        responderComJSON(w, string(json))
    })
}

func responderCom(
    w http.ResponseWriter,
    status int,
    headers Headers,
) {
    for k, v := range headers {
        w.Header().Set(k, v)
    }
    w.WriteHeader(status)
}

func responderComJSON(w http.ResponseWriter, resposta string) {
    responderCom(w, http.StatusOK, Headers{
        "Content-Type": "application/json",
    })
    fmt.Fprintf(w, resposta)  // Até que enfim escrevemos um "corpo" :P
}

func extrairUrl(r *http.Request) string {
    url := make([]byte, r.ContentLength, r.ContentLength)
    r.Body.Read(url)
    return string(url)
}

func buscarUrlEExecutar(
    w http.ResponseWriter,
    r *http.Request,
    executor func(*url.Url),
) {
    caminho := strings.Split(r.URL.Path, "/")
    id := caminho[len(caminho)-1]

    if url := url.Buscar(id); url != nil {
        executor(url)
    } else {
        http.NotFound(w, r)
    }
}

func registrarEstatisticas(ids <-chan string) {
    for id := range ids {
        url.RegistrarClique(id)
        fmt.Printf("Clique registrado para %s.\n", id)
    }
}

