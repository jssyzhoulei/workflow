GOOS 		= linux
CGO_ENABLED = 0
ROOT_DIR    = $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))/
DIST_DIR 	= $(ROOT_DIR)dist/
POD_NAME	= org

IMAGE_NAME_API := 192.168.2.190:5000/library/org-api:1.0.0
IMAGE_NAME_SVC := 192.168.2.190:5000/library/org-svc:1.0.0
#IMAGE_NAME := 10.182.240.76:30180/draco/draco:1.0.0

.PHONY: release
release: dist_dir build docker push-image  del clean

.PHONY: docker
docker:
	@echo ========== current docker build image is: $(IMAGE_NAME) ==========
	cp Dockerfile-api $(DIST_DIR)
	cp Dockerfile-svc $(DIST_DIR)
	docker build -f $(DIST_DIR)Dockerfile-api  -t $(IMAGE_NAME_API) $(DIST_DIR)
	docker build -f $(DIST_DIR)Dockerfile-svc  -t $(IMAGE_NAME_SVC) $(DIST_DIR)

.PHONY: push-image
push-image:
	@echo ========== current docker push image is: $(IMAGE_NAME) ==========
	docker push $(IMAGE_NAME_SVC)
	docker push $(IMAGE_NAME_API)

.PHONY: build
build:
	@echo ========== go build ==========
#	go mod vendor
	protoc -I $(ROOT_DIR)src/proto/user/v1/ --go_out=plugins=grpc:./src/proto/user/v1/ $(ROOT_DIR)src/proto/user/v1/*.proto
	env CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) go build -installsuffix cgo -o $(DIST_DIR)org-api  $(ROOT_DIR)cmd/org-api/org-api.go
	env CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) go build -installsuffix cgo -o $(DIST_DIR)org-svc  $(ROOT_DIR)cmd/org-svc/org-svc.go

.PHONY: dist_dir
dist_dir: ; $(info ======== prepare distribute dir:)
	mkdir -p $(DIST_DIR)

.PHONY: del
del:
	ssh root@192.168.2.190 "/root/org/delete-pod.sh $(POD_NAME)"
.PHONY: clean
clean: ; $(info ======== clean all:)
	rm -rf $(DIST_DIR)*
	docker rmi $(IMAGE_NAME_API)
	docker rmi $(IMAGE_NAME_SVC)