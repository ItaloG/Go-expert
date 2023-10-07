# REF https://github.com/golang-migrate/migrate
# REF https://github.com/jmoiron/sqlx
# REF https://sqlc.dev/

curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -

echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ (lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list

apt-get update

apt-get install -y migrate
