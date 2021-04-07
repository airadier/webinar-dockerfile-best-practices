all: step1 step2 step3 step4

step1: src/example_app
	docker build -f Dockerfile1 -t test1_1 .

step2: src/example_app
	docker build -f Dockerfile1 -t test1_2 .

step3: src/example_app
	docker build -f Dockerfile1 -t test1_3 .

step4: src/example_app
	docker build -f Dockerfile1 -t test1_3 .

src/example_app:
	make -C src build-linux