#DEV

build-dev:
	docker build -t meet_harbor -f containes/images/Dockerflie . && docker build -t turn -f containers/images/Dockerflie.turn .

clean-dev:
	docker-compose -f containers/composes/dc.dev.yml down

run-dev:
	docker-compose -f containers/composes/dc.dev.yml up

