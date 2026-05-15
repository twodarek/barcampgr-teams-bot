FROM golang:1.25 AS base

ENV GO111MODULE=on

# test and build our app
WORKDIR /go/src/twodarek/barcampgr-teams-bot
COPY . .
RUN go build


FROM docker pull git.hedgewizards.org/twodarek/ubuntu-base-image:latest

RUN mkdir -p /public/front-end
COPY --from=base /go/src/twodarek/barcampgr-teams-bot/front-end /public/front-end
COPY --from=base /go/src/twodarek/barcampgr-teams-bot/barcampgr-teams-bot /usr/local/bin

ENTRYPOINT ["/usr/local/bin/barcampgr-teams-bot"]
