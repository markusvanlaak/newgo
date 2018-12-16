FROM debian:7-slim
COPY test .
RUN test
