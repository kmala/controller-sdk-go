FROM quay.io/deis/go-dev:0.17.0
# This Dockerfile is used to bundle the source and all dependencies into an image for testing.

RUN echo "deb http://packages.cloud.google.com/apt cloud-sdk-jessie main" \
  | tee /etc/apt/sources.list.d/google-cloud-sdk.list \
  && curl https://packages.cloud.google.com/apt/doc/apt-key.gpg \
  | apt-key add - \
  && apt-get update \
  && apt-get install -y google-cloud-sdk \
  --no-install-recommends \
  && rm -rf /var/lib/apt/lists/*

COPY glide.yaml /go/src/github.com/deis/controller-sdk-go/
COPY glide.lock /go/src/github.com/deis/controller-sdk-go/

WORKDIR /go/src/github.com/deis/controller-sdk-go

RUN glide install --strip-vcs --strip-vendor

COPY ./_scripts /usr/local/bin

COPY . /go/src/github.com/deis/controller-sdk-go
