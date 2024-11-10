# keeper
Менеджер секретностей

# Как запустить

1. Создать сертификаты x509: `make certs`
    Они будут созданы в папке `./data`
2. Запустить сервер:
```bash
TLS_PUBLIC_KEY_PATH=./data/server_cert.pem TLS_PRIVATE_KEY_PATH=./data/server_key.pem go run ./cmd/server/
```
3. Запустить клиента:
```bash
TLS_PRIVATE_KEY_PATH=./data/server_key.pem TLS_PUBLIC_KEY_PATH=./data/server_cert.pem go run ./cmd/client/ ui
```