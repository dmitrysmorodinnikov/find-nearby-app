.PHONY: all
all: docker.start

docker.start:
	docker build -t find-nearby-frontend . && docker run -it --name find-nearby-react --rm -v ${PWD}:/app -v /app/node_modules -p 3003:3000 find-nearby-frontend