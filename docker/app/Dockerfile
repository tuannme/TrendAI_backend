# Using golang image for build only
FROM golang:1.11.5 AS builder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/trend-ai/TrendAI_mobile_backend
RUN pwd
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./

# Install gin for live-reloading Go app
RUN go get github.com/githubnemo/CompileDaemon
ENV CGO_ENABLED=0
ENTRYPOINT CompileDaemon -exclude-dir=.git -exclude-dir=.idea -build="go build -a -installsuffix nocgo -o /app ." -command="/app"

# Use for build production
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .
# Using an empty image (scratch) for run app app
# FROM scratch
# COPY --from=builder /app ./
# ENTRYPOINT ["./app"]