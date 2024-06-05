FROM golang:1.22

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download
COPY ./ /app/

WORKDIR /app/cmd/api
RUN go build -o /go-docker-backend

# EXPOSE 4000
# Run
CMD ["/go-docker-backend"]