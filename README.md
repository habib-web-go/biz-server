# biz-server

## Install proto-compiler

### mac

```sh
brew install protobuf
```

### linux

```sh
apt install -y protobuf-compiler
```

## Compile Proto

```
protoc --go_out=. --go-grpc_out=.  grpc/sqlpb.proto
```

## Create Random User And Insert To DB

### Regenerate Sample User

```sh
python -m venv venv
source venv/bin/activate
pip install -r requirements.txt
cd postges
python random_users.py
```

### Insert Users to DB

```sh
docker compose exec -it postgres psql -U ${user} -d ${database} -f /scripts/insert_users.sql
```
### Delete All Rows
```sh
docker compose exec -it postgres psql -U ${user} -d ${database} -f /scripts/delete_all.sql
```
