module "{{.ModuleName}}"

go 1.24.4

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/joho/godotenv v1.5.1
	github.com/rafa-mori/logz v1.3.0
	{{- if .DatabaseType}}
	{{- if eq .DatabaseType "postgres"}}
	github.com/lib/pq v1.10.9
	gorm.io/driver/postgres v1.5.4
	{{- else if eq .DatabaseType "mysql"}}
	github.com/go-sql-driver/mysql v1.7.1
	gorm.io/driver/mysql v1.5.2
	{{- else if eq .DatabaseType "mongodb"}}
	go.mongodb.org/mongo-driver v1.13.1
	{{- else if eq .DatabaseType "sqlite"}}
	gorm.io/driver/sqlite v1.5.4
	{{- end}}
	{{- if ne .DatabaseType "mongodb"}}
	gorm.io/gorm v1.25.5
	{{- end}}
	{{- end}}
	{{- if .CacheType}}
	{{- if eq .CacheType "redis"}}
	github.com/redis/go-redis/v9 v9.3.0
	{{- end}}
	{{- end}}
	{{- if hasFeature "Authentication"}}
	github.com/golang-jwt/jwt/v5 v5.2.0
	{{- end}}
	{{- if hasFeature "API Documentation"}}
	github.com/swaggo/swag v1.16.2
	github.com/swaggo/gin-swagger v1.6.0
	github.com/swaggo/files v1.0.1
	{{- end}}
	{{- if hasFeature "Rate Limiting"}}
	github.com/ulule/limiter/v3 v3.11.2
	{{- end}}
)
