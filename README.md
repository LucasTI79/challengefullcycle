### Instruções para executar


Baixe os submódulos do projeto
```sh
git submodule update --init --recursive
```

Suba o ambiente
```sh
docker compose up -d
```

Abra o ambiente do servidor grpc
```sh
docker exec -it server bash
```

Execute o ambiente do servidor grpc
```sh
go run main.go grpc
```

Abra o ambiente do cliente grpc
```sh
docker exec -it client bash
```

Instale as dependências do cliente dentro do ambiente

```npm i``` ou ```yarn``` ou ```pnpm install```

Execute o ambiente do cliente grpc
```sh
npm start:dev
```


