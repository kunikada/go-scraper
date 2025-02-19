FROM golang:1.24-bookworm

# install
# nodejs: for SonarLint
RUN apt-get update && apt-get install -y \
  nodejs \
  && apt-get clean

RUN useradd -m gouser
USER gouser

WORKDIR /workspaces
#RUN go mod tidy

#RUN go install golang.org/x/tools/gopls@latest \
#  && go install honnef.co/go/tools/cmd/staticcheck@latest \

#CMD ["go", "run", "main.go"]
