FROM golang AS build-stage
COPY . src/
RUN cd src && go build .

FROM scratch AS export-stage
COPY --from=build-stage /go/src/1brc /