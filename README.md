# go-livro-casa-do-codigo

* Estudando a linguagem Go pelo [livro](https://www.casadocodigo.com.br/products/livro-google-go) da editora
* Acrescentando melhorias como uso de banco de dados
  * https://aprendagolang.com.br/2022/06/02/como-conectar-e-fazer-crud-em-um-banco-postgresql/
  * https://koho.dev/understanding-go-and-databases-at-scale-connection-pooling-f301e56fa73

Para facilitar o banco PostgreSQL está configurado com Docker e docker-compose, bastando executar:

```bash
docker-compose up  # -d se desejar em background
```

É provido um script para subir um servidor conectado ao database:

```bash
./run.sh
```
