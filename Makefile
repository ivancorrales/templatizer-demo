build:
	docker build -t {{.Application}} .

run: 
	docker run {{.Application}}