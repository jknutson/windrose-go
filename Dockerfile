FROM golang:1.24-bookworm

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV APP_USER=app
ENV APP_HOME=/go/src/windrose-go
ENV GROUP_ID=1001
ENV USER_ID=1001

RUN groupadd --gid $GROUP_ID app && \
  useradd -m -l --uid $USER_ID --gid $GROUP_ID $APP_USER && \
  mkdir -p $APP_HOME && \
  chown -R $APP_USER:$APP_USER $APP_HOME

USER $APP_USER
WORKDIR $APP_HOME
COPY . .

ENV GOFLAGS='-buildvcs=false -trimpath'
RUN go build -o ./windrose-go .
EXPOSE 8080
CMD ["./windrose-go"]
