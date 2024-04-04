FROM golang:1.22.2 as build

WORKDIR /workspace

# add go modules lockfiles
COPY go.mod go.sum ./
RUN go mod download

# prefetch the binaries, so that they will be cached and not downloaded on each change
RUN go run github.com/steebchen/prisma-client-go prefetch

COPY . ./

# generate the Prisma Client Go client
RUN go run github.com/steebchen/prisma-client-go generate

# build the binary with all dependencies
RUN go build -o /app .

# Expose port 8080 to the outside world
EXPOSE 8080

CMD ["/app"]
 