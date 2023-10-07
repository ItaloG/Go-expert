apt-get install -y sqlite3

# TABELAS SQLITE
# create table categories (id string, name string, description string)
# create table courses (id string, name string, description string, category_id string)

go run github.com/99designs/gqlgen generate

go mod tidy