FROM golang:alpine
RUN apk add git
COPY . /go/src/ffmpeg-server/
WORKDIR /go/src/ffmpeg-server
RUN go get -u github.com/gobuffalo/packr/v2/packr2
RUN go get -u github.com/gobuffalo/packr/v2/...
RUN go get -v -t -d ./...

WORKDIR /go/bin
RUN packr2 build ffmpeg-server
RUN adduser -h /home/ffmpeg-server -D ffmpeg-server
WORKDIR /home/ffmpeg-server
COPY --from=0 /go/bin/ffmpeg-server /home/ffmpeg-server/ffmpeg-server
RUN chown ffmpeg-server:ffmpeg-server ffmpeg-server
RUN chmod +x ffmpeg-server

ENTRYPOINT [ "./ffmpeg-server" ]