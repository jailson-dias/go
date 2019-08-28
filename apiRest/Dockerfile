from golang:alpine3.10

ENV GOPATH /proj
ENV PATH="/proj/bin:${PATH}"

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

RUN go get -u github.com/Masterminds/glide/...
RUN go get github.com/githubnemo/CompileDaemon

WORKDIR /proj/src/go/apiRest

COPY glide.yaml .

RUN glide install

COPY main.go .


CMD CompileDaemon -log-prefix=false -build="go build ./main.go" -command="./main"