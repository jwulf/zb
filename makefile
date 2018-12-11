# Go parameters
	GOCMD=go
	GOBUILD=$(GOCMD) build
	GOCOMPILE=CGO_ENABLED=0 GOOS=linux $(GOBUILD) -a -installsuffix cgo -o main .
	DOCKERCOMPOSE=../../docker-compose.yml
	DOCKERFILE=../../taskworker/Dockerfile

    all:
	    make build
	    make dockerise
    build:
		cd workers && \
		for d in `ls`; do cd $$d && echo "Building $$d" && $(GOCOMPILE) && echo "Size: `du -h main`" && cd ..; done && \
		cd ..
    dockerise:
		echo "version: '2'" > docker-compose.yml && \
		echo "services:" >> docker-compose.yml && \
		cd workers && \
		for d in `ls`; do cd $$d && \
		docker build -t sitapati/zb-$$d -f $(DOCKERFILE) . && \
		echo "  $$d:" >> $(DOCKERCOMPOSE) && \
		echo "    image: sitapati/zb-$$d" >> $(DOCKERCOMPOSE) && \
		echo "    container_name: $$d" >> $(DOCKERCOMPOSE) && \
		echo "    network_mode: \"host\"" >> $(DOCKERCOMPOSE) && \
		echo "    environment:" >> $(DOCKERCOMPOSE) && \
		echo "       - ZEEBE_BROKER_ADDRESS" >> $(DOCKERCOMPOSE) && \
		echo "    logging:" >> $(DOCKERCOMPOSE) && \
		echo "      driver: \"json-file\"" >> $(DOCKERCOMPOSE) && \
		echo "      options:" >> $(DOCKERCOMPOSE) && \
		echo "        max-size: \"10k\"" >> $(DOCKERCOMPOSE) && \
		echo "        max-file: \"10\"" >> $(DOCKERCOMPOSE) && \
		cd ..; done
    publish:
		cd workers && \
		for d in `ls`; do docker push sitapati/zb-$$d; done