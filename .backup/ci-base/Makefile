## build release docker image
.PHONY: build-release-image
build-release-image:
	docker build . \
		--no-cache \
		--force-rm \
		-t $(release-image) \
		-f Dockerfile

## login docker registry
.PHONY: login-docker-registry
login-docker-registry:
	docker login -u $(registry-user) -p $(registry-password) $(registry-host)

## push release docker image
.PHONY: push-release-image
push-release-image:
	docker push $(release-image)