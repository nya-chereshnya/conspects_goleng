## Start postgres
sudo systemctl start postgresql

## Restart postgres
sudo systemctl restart postgresql

## User creation
sudo -u postgres createuser --interactive

## Add password 
sudo -u postgres psql
> password <имя_пользователя>

## DB creation
sudo -u postgres createdb <имя_пользователя>

## Connect to DB
psql -U <имя_пользователя> -d <имя_базы_данных>

## Изменение настроек аутентификации
cd /etc/postgresql/<версия>/main
pg_hba.conf -> METHOD md5

## Restore from dump
psql -U <имя_пользователя> -d <имя_базы_данных> -f dump.sql

