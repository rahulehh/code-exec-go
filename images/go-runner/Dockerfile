FROM golang:1.21-alpine
WORKDIR /app
COPY run.sh /run.sh
RUN chmod +x /run.sh
ENTRYPOINT ["/run.sh"]
