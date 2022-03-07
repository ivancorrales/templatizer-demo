build:
	docker build -it {{.Application}} .

run: 
	docker run {{.Application}}