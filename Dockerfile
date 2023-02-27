FROM golang:1.19.1
WORKDIR /
COPY . .
RUN go install .
ENTRYPOINT ["gen-tools"]