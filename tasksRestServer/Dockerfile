# This will most probably not work because of the invalid paths
FROM goland
ADD . /Users/jrumjantsev/.go/src/awesomeProject/tasksRestServer
WORKDIR .go/src/awesomeProject/tasksRestServer
RUN go get github.com/tools/godep
RUN godep restore
RUN go install github.com/jackinf/awesomeProject/tasksRestServer
ENTRYPOINT /go/bin/tasksRestServer
EXPOSE 8080

# docker build -t taskRestServer
# docker run --publish 80:8080 --name taskmanager_api --rm taskRestServer