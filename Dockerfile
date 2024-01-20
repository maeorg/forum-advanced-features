FROM golang:latest

# Set destination for COPY
WORKDIR /app

# Copy the source code
COPY . .

# Build
RUN go build -v -o /docker-forum

# Run
CMD ["/docker-forum"]