FROM golang

RUN mkdir /go/ffmpeg-server
COPY . /go/ffmpeg-server/
WORKDIR /go/bin
RUN go build ffmpeg-server

FROM busybox
RUN adduser -h /home/ffmpeg-server -D ffmpeg-server
WORKDIR /home/ffmpeg-server
COPY --from=0 /go/bin/ffmpeg-server /home/ffmpeg-server/ffmpeg-server
RUN chown ffmpeg-server:ffmpeg-server ffmpeg-server
RUN chmod +x ffmpeg-server

ENTRYPOINT [ "./ffmpeg-server" ]