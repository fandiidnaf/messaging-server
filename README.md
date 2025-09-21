# Structure Of Project Folder
```doc
myapp/
├── cmd/
│   └── server/              # Entry point (main.go)
│       └── main.go
├── config/                  # Konfigurasi (env, config loader)
│   └── config.go
├── internal/                # Source code utama (biasa pakai internal biar tidak di-import keluar)
│   ├── app/
│   │   ├── handler/         # Handler untuk Gin (controller)
│   │   │   └── user_handler.go
│   │   ├── service/         # Business logic (use case)
│   │   │   └── user_service.go
│   │   ├── repository/      # Akses DB (data access layer)
│   │   │   └── user_repository.go
│   │   ├── model/           # Struct model/DTO
│   │   │   └── user.go
│   │   └── router/          # Routing setup untuk Gin
│   │       └── router.go
│   └── pkg/                 # Utilitas umum (helper, middleware, dll)
│       ├── middleware/
│       │   └── auth.go
│       └── response/
│           └── response.go
├── migrations/              # (opsional) file migrasi database
├── go.mod
└── go.sum

🛠️ Penjelasan Tiap Bagian

cmd/server/main.go
Entry point aplikasi, isinya setup config, koneksi DB, inisialisasi Gin, dan start server.

config/
Tempat untuk load konfigurasi (misal pakai .env, atau Viper). Contoh: config.go baca DB_URL, JWT_SECRET, dll.

internal/app/handler/
Controller / handler yang berhubungan langsung dengan HTTP request.
Contoh: user_handler.go berisi fungsi GetUsers, CreateUser.

internal/app/service/
Berisi business logic. Handler akan memanggil service, bukan langsung query DB.

internal/app/repository/
Abstraksi untuk database (misal pakai GORM/pgx/sqlx). Service memanggil repository.

internal/app/model/
Struct untuk representasi data (User, Product, DTO request/response).

internal/app/router/
Definisi semua route Gin (/users, /auth, dsb).

internal/pkg/
Utility: middleware, helper, response formatter, error wrapper, dll.

migrations/
Kalau kamu pakai DB (Postgres/MySQL), taruh file migrasi SQL di sini.
```

# GENERATE SWAGGER
```bin
swag init -g cmd/server/main.go
```
