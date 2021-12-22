
build: build-mail build-cron

deploy: build
	scp bin/mail pi@raspy3:/home/pi/sunsetty	
	scp bin/crony pi@raspy3:/home/pi/sunsetty	
	scp run.sh pi@raspy3:/home/pi/sunsetty	
	scp .env pi@raspy3:/home/pi/sunsetty	

build-mail:
	GOOS=linux GOARCH=arm go build -o bin/mail cmd/mail/main.go

build-cron:
	GOOS=linux GOARCH=arm go build -o bin/crony cmd/crony/main.go