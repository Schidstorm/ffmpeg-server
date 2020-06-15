FROM golang:alpine
RUN apk add git
COPY . /go/src/ffmpeg-server/
WORKDIR /go/src/ffmpeg-server
RUN go get -u github.com/gobuffalo/packr/v2/packr2
RUN go get -u github.com/gobuffalo/packr/v2/...
RUN go get -v -t -d ./...
RUN packr2 clean
RUN packr2
WORKDIR /go/bin
RUN go build ffmpeg-server

FROM alpine
RUN apk add alpine
RUN adduser -h /home/ffmpeg-server -D ffmpeg-server
WORKDIR /home/ffmpeg-server
COPY --from=0 /go/bin/ffmpeg-server /home/ffmpeg-server/ffmpeg-server
RUN chown ffmpeg-server:ffmpeg-server ffmpeg-server
RUN chmod +x ffmpeg-server

ENTRYPOINT [ "./ffmpeg-server" ]