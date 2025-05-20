FROM gcc:12.2.0-bullseye

WORKDIR /app

COPY run.sh /run.sh
RUN chmod +x /run.sh

ENTRYPOINT ["/run.sh"]
