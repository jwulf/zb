# Go parameters
	GOCMD=go
	GOBUILD=$(GOCMD) build
	GOCOMPILE=CGO_ENABLED=0 GOOS=linux $(GOBUILD) -a -installsuffix cgo -o main .
	DOCKERCOMPOSE=../../docker-compose.yml
	DOCKERFILE=../../taskworker/Dockerfile
	LOG=   logging:
    driver: "json-file"
	options:
			max-size: "1k"
			max-file: "3"

    all:
		test build
    dockerise:
		echo "version: '2'" > docker-compose.yml && \
		echo "services:" >> docker-compose.yml && \
		cd workers && \
		for d in `ls`; do cd $$d && echo "Building $$d" && $(GOCOMPILE) && echo "Size: `du -h main`" && \
		docker build -t zb-$$d -f $(DOCKERFILE) . && \
		echo "  $$d:" >> $(DOCKERCOMPOSE) && \
		echo "    image: zb-$$d" >> $(DOCKERCOMPOSE) && \
		echo "    container_name: $$d" >> $(DOCKERCOMPOSE) && \
		echo "    network_mode: \"host\"" >> $(DOCKERCOMPOSE) && \
		echo "    environment:" && \
		echo "       - ZEEBE_BROKER_ADDRESS" && \
		echo "    logging:" >> $(DOCKERCOMPOSE) && \
		echo "      driver: \"json-file\"" >> $(DOCKERCOMPOSE) && \
		echo "      options:" >> $(DOCKERCOMPOSE) && \
		echo "        max-size: \"10k\"" >> $(DOCKERCOMPOSE) && \
		echo "        max-file: \"10\"" >> $(DOCKERCOMPOSE) && \
		cd ..; done
