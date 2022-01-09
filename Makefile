generate: generate/server

generate/server:
	 docker run --rm -it -v /$(shell pwd)://project -w //project quay.io/goswagger/swagger generate server \
	 --spec //project/swagger.json \
	 --server-package src/server \
	 --model-package src/server/models \
	 --exclude-main \
	 --main-package server