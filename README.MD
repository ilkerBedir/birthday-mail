docker build -t postgres-db .

docker run --name postgres-container -e POSTGRES_PASSWORD=12345 -d -p 5432:5432 postgres-db


//todo: 

make html templates.

create job which triggers at 00.00 clock and start application.

https://support.google.com/accounts/answer/185833