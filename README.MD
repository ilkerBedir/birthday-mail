docker build -t postgres-db .

docker run --name postgres-container -e POSTGRES_PASSWORD=12345 -d -p 5432:5432 postgres-db

https://support.google.com/accounts/answer/185833