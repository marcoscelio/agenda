# Agenda App

Start Mongo DB

sudo ~/mongodb-macos-x86_64-4.2.7/bin/mongod --dbpath ~/mongodb-macos-x86_64-4.2.7/data/contacts

# Set environment variables

1 READ_TIMEOUT -> value in seconds
Ex: 
export READ_TIMEOUT=3

2 WRITE_TIMEOUT -> value in seconds
Ex: 
export WRITE_TIMEOUT=3

3 PORT -> integer value
Ex: 
export PORT=3000

# Start App

cd <app base dir>
go run server.go

# Build Docker image

docker build --build-arg PORT=${PORT} --build-arg READ_TIMEOUT=${READ_TIMEOUT} --build-arg WRITE_TIMEOUT=${WRITE_TIMEOUT} -t agenda-golang .

# Run docker image

docker run -p $PORT:$PORT agenda-golang  

# Run Docker compose

docker-compose up

# Testing
Unit tests:

1) Enter folder to test
2) Run command: go test
3) For coverage report: go test -cover


