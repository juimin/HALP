# add the necessary instructions
# to create a Docker container image
# for your Go API gateway server
# Put your docker build instructions here!
FROM alpine
COPY gateway gateway

RUN apk update && apk add ca-certificates
EXPOSE 443
ENTRYPOINT [ "gateway"]