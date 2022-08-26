# pack it

This is a project to test the [goreleaser](https://goreleaser.com) project. 

## How to install goreleaser

I installed the go tool

```sh
go1.18.3 install github.com/goreleaser/goreleaser@latest
```

## Steps

1. run goreleaser init to create a `.goreleaser.yaml` file.

```sh
goreleaser init
```

2. check if the configuration is valid.

```sh
goreleaser check
```

3. generate binaries just to test the build operation.

```sh
goreleaser release --snapshot --rm-dist
```

4. we can build it only for one target

```sh
goreleaser build --single-target --snapshot --rm-dist
```

5. In case you want to release it to github, you must provide a GITHUB token.

```sh
export GITHUB_TOKEN="YOUR_GH_TOKEN"
```

please notice that goreleaser will use the lastest git tag of our repository.