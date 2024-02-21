# CasaOS-MessageBus

[![Go Reference](https://pkg.go.dev/badge/github.com/KaySar12/NextZen-MessageBus.svg)](https://pkg.go.dev/github.com/KaySar12/NextZen-MessageBus) [![Go Report Card](https://goreportcard.com/badge/github.com/KaySar12/NextZen-MessageBus)](https://goreportcard.com/report/github.com/KaySar12/NextZen-MessageBus) [![goreleaser](https://github.com/KaySar12/NextZen-MessageBus/actions/workflows/release.yml/badge.svg)](https://github.com/KaySar12/NextZen-MessageBus/actions/workflows/release.yml) [![codecov](https://codecov.io/gh/IceWhaleTech/CasaOS-MessageBus/branch/main/graph/badge.svg?token=U4S4ZSZAL9)](https://codecov.io/gh/IceWhaleTech/CasaOS-MessageBus)

Message bus accepts events and actions from various sources and delivers them to subscribers.

See [openapi.yaml](./api/message_bus/openapi.yaml) for API specification.




## publish api to npm

### edit version in package.json

### run
```bash
yarn

yarn start
```

### publish

Manual publish
```bash
yarn publish
```

Auto publish
```bash 
git push origin dev**
```