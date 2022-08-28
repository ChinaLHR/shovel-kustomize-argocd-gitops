FROM docker:20-dind
RUN apk update && apk add make && apk add git && apk add curl && apk add openssh-client

# Install kustomize
RUN  curl -L --output /tmp/kustomize_v4.5.5_linux_amd64.tar.gz https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize/v4.5.5/kustomize_v4.5.5_linux_amd64.tar.gz \
  && tar -xvzf /tmp/kustomize_v4.5.5_linux_amd64.tar.gz -C /usr/local/bin \
  && chmod +x /usr/local/bin/kustomize
