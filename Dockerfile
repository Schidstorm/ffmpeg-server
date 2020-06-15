FROM golang

COPY . /go/src/ffmpeg-server/
WORKDIR /go/bin
RUN go get -v -t -d /go/src/ffmpeg-server/...
RUN go build ffmpeg-server

FROM alpine
RUN apk add alpine
RUN adduser -h /home/ffmpeg-server -D ffmpeg-server
WORKDIR /home/ffmpeg-server
COPY --from=0 /go/bin/ffmpeg-server /home/ffmpeg-server/ffmpeg-server
RUN chown ffmpeg-server:ffmpeg-server ffmpeg-server
RUN chmod +x ffmpeg-server

ENTRYPOINT [ "./ffmpeg-server" ]