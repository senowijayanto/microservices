FROM golang:1.16

RUN mkdir -p /app

COPY ./ /app

RUN apt-get update && apt-get install -y --no-install-recommends \
	git \
	&& rm -rf /var/lib/apt/lists/* \
	&& cd /app \
	&& rm -rf ./docker \
	&& rm -rf ./.git \
	&& go mod download \
  && go build main.go

EXPOSE 3000

CMD ["/app/main"]