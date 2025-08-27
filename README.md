# variables env
cp .env.example .env
# editar .env con tus datos DB

# build local
go build -o robotapp main.go

# ejecutar local (requiere env vars)
export $(cat .env | xargs)
./robotapp

# o usando docker-compose (solo app; DB en la nube)
docker compose up --build

# crear tablas y seed (usa las mismas env vars)
go run ./cmd/migrate - # para que use env y haga AutoMigrate
# con seed: SEED=1 go run ./cmd/migrate

# si usas el binary dentro del container:
docker compose run --rm app /bin/robotapp   # o configurar ENTRYPOINT para migrar al arrancar
