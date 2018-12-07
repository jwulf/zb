# Go parameters
	GOCMD=go
	GOBUILD=$(GOCMD) build
	GOCOMPILE=CGO_ENABLED=0 GOOS=linux $(GOBUILD) -a -installsuffix cgo -o main .
	DOCKERCOMPOSE=../../docker-compose.yml
	DOCKERFILE=../../taskworker/Dockerfile

    all:
		test build
    dockerise:
		echo "version: '3'" > docker-compose.yml && \
		echo "services:" >> docker-compose.yml && \
		cd workers && \
		for d in `ls`; do cd $$d && echo "Building $$d" && $(GOCOMPILE) && echo "Size: `du -h main`" && \
		docker build -t zb-$$d -f $(DOCKERFILE) . && \
		echo "  $$d:" >> $(DOCKERCOMPOSE) && \
		echo "    image: zb-$$d" >> $(DOCKERCOMPOSE) && \
		echo "    container_name: $$d" >> $(DOCKERCOMPOSE) && \
		cd ..; done
