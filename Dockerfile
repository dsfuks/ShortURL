FROM golang:latest
WORKDIR src/ShortURL
COPY ./ ./
RUN go build -o main .
ENTRYPOINT ["powershell.exe"]
CMD ["go run . memory"]