FROM golang:1.21

# Set destination for COPY
WORKDIR /app

# Generate self-signed cert
RUN openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes -subj "/C=US/ST=Ottawa/L=Ottawa/O=Wedding backend/CN=*.brent.click"


# Download Go modules
COPY ./ ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /wedding-backend

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

# Run
CMD ["/wedding-backend"]