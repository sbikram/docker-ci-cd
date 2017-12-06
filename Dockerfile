FROM golang


ADD . /Users/bikramsingh/docker-ci-cd
RUN go install docker-ci-cd
CMD /go/bin/docker-ci-cd

EXPOSE 8080
