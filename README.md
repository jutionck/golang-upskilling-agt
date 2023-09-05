# Golang Upskilling

## Config

`Note`:

- Pada folder `environment` silahkan buat file baru dengan nama `dev.env` yang di duplikat dari `.env.example` kemudian isi pada bagian berikut:

```env
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_DRIVER=
API_HOST=
API_PORT=
MIGRATION=
FILE_PATH=
TOKEN_NAME=
TOKEN_KEY=
TOKEN_EXPIRES=
```

## Run Apps

Jika sudah jalankan program dengan perintah berikut:

```bash
make run
```

Jika sudah berjalan silahkan akses swagger nya di alamat berikut:

```
http://localhost:8888/swagger/index.html
```

## Run Testing

Untuk menjalankan unit testing:

```bash
make coverage
```
