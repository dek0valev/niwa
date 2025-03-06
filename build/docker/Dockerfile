# Build the application from source
FROM golang:1.24.0 AS build

RUN apt update && apt install -y unzip && rm -rf /var/lib/apt/lists/*

RUN curl -fsSL https://bun.sh/install | bash -s "bun-v1.2.4"

ENV PATH="/root/.bun/bin:$PATH"

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY package.json bun.lock ./
RUN bun install --frozen-lockfile

COPY . ./

RUN bun run tailwind

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app ./cmd/niwa

# Deploy the application binary into a lean image
FROM gcr.io/distroless/static-debian12

WORKDIR /

COPY --from=build /go/bin/app /
COPY --from=build /usr/src/app/content /content
COPY --from=build /usr/src/app/web /web
COPY --from=build /usr/src/app/configs/app.yaml /configs/app.yaml

ENV CONFIG_PATH='./configs/app.yaml'

ENTRYPOINT ["/app"]