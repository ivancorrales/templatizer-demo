FROM scratch
COPY {{.RepoName}} /
ENTRYPOINT ["/{{.RepoName}}"]
