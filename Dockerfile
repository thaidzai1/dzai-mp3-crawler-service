FROM golang:1.15

RUN apt-get update
RUN apt-get install -y dejavu-* pulseaudio
RUN wget https://dl.google.com/linux/direct/google-chrome-beta_current_amd64.deb
RUN dpkg -i google-chrome*.deb || apt-get install -f -y

WORKDIR /go/src/app
COPY . .

RUN go install ./...

CMD ["/bin/bash"]

EXPOSE 4001