PREFIX = laher
ARTIFACT = kdisco

build:
		CGO_ENABLED=0 go build -o kdisco

image: TAG ?= latest
image: build 
		docker build -t $(PREFIX)/${ARTIFACT}:$(TAG) .

image-and-push: TAG ?= latest
image-and-push: build 
		sudo docker build -t $(PREFIX)/${ARTIFACT}:$(TAG) .
		sudo docker push $(PREFIX)/${ARTIFACT}:$(TAG)


