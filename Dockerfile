FROM golang:1.14 as build
ENV PORT 8080 
ENV HOST 0.0.0.0
WORKDIR /go/src/app
COPY . .
RUN go build -v -o /app .
# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=build /app /app
CMD ["/app"]