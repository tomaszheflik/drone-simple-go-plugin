FROM alpine
ADD weather-plugin /bin/
RUN apk -Uuv add ca-certificates
ENTRYPOINT /bin/weather-plugin