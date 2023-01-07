# go-livro-casa-do-codigo

* Estudando a linguagem Go pelo [livro](https://www.casadocodigo.com.br/products/livro-google-go) da editora
* Acrescentando melhorias como uso de banco de dados

Para facilitar o banco PostgreSQL está configurado com Docker e docker-compose, bastando executar:

```bash
docker-compose up  # -d se desejar em background
```

É provido um script para subir um servidor conectado ao database:

```bash
./run.sh
```
