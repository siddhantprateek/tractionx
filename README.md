# Tractionx

## Tech Stack

- `Hyperledger Fabric`
- `Go`
- `Kubernetes`
- `Docker`


## Setting up Development environment

- `kubectl`
- jq `brew install jq`
- [envsubst](https://www.gnu.org/software/gettext/manual/html_node/envsubst-Invocation.html) `brew install gettext`
- `KIND` `brew install kind`
- `Docker`

### Create a KIND cluster
```bash
cd test-network-k8s
./network kind
./network cluster init
```

- Launch the network and create a channel
```bash
./network up

./network channel create
```

## Author

- [Siddhant Prateek Mahanayak](github.com/siddhantprateek)