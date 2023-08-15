# Tractionx

TractionX is a sample real estate transaction platform that uses the Hyperledger Fabric Raft protocol for secure and efficient property transactions.


## Tech Stack

- `Hyperledger Fabric`: A blockchain framework for developing enterprise-grade decentralized applications.
- `Go`: The programming language used for developing the application logic.
- `Kubernetes`: An open-source container orchestration platform for automating the deployment, scaling, and management of containerized applications.
- `Docker`: A platform that automates the deployment of applications inside lightweight, portable containers.

## Setting up the Development Environment

Before you can start working on the Tractionx project, ensure that you have the following tools installed in your development environment:

- `kubectl`: Kubernetes command-line tool.
- `jq`: A lightweight and flexible command-line JSON processor.
- [envsubst](https://www.gnu.org/software/gettext/manual/html_node/envsubst-Invocation.html): A command-line tool for substituting environment variables in files.
- `KIND`: Kubernetes in Docker - used for creating local Kubernetes clusters.
- `Docker`: Containerization platform.

Follow these steps to set up your development environment:

### Install Required Tools

```bash
brew install kubectl jq
brew install gettext  # This includes envsubst
brew install kind
```

### Create a KIND Cluster

Navigate to the project's `test-network-k8s` directory and run the following commands to set up a KIND cluster:
- To start test network using k8s

```bash
cd test-network-k8s
./network kind
./network cluster init
```

### Launch the Network and Create a Channel

Launch the network and create a channel using the provided scripts:

```bash
./network up
./network channel create
```

```bash
peer lifecycle chaincode package <chaincode-package-name>.tgz \
--path ./path/to/chaincode --lang golang --label <chaincode-package-label>

peer lifecycle chaincode install <chaincode-package-name>.tgz

# or
peer lifecycle chaincode package property.tgz \ 
--path ./contracts/ --lang golang --label property_1

peer lifecycle chaincode install property.tgz
```

## Setting Up a Hyperledger Fabric Network with Microfab

### Configure Microfab

Create a file named `microfab-config.json` and define your network configuration:

```json
{
  "port": 8080,
  "endorsing_organizations": [
    {
      "name": "ProducersOrg"
    },
    {
      "name": "SellersOrg"
    }
  ],
  "channels": [
    {
      "name": "property-channel",
      "endorsing_organizations": [
        "ProducersOrg",
        "SellersOrg"
      ]
    }
  ]
}
```

### Start Microfab

Run the following commands to start Microfab using Docker Compose:

```bash
docker-compose up -d
docker run -e MICROFAB_CONFIG -p 8080:8080 -d ibmcom/ibp-microfab
```

### Set Up Weft

Run the following command to retrieve component information from Microfab and configure Weft:

```bash
curl -s http://console.127-0-0-1.nip.io:8080/ak/api/v1/components | weft microfab -w ./_wallets -p ./_gateways -m ./_msp -f
```

### Download Hyperledger Fabric

Run the following command to download the Hyperledger Fabric binaries:

```bash
curl -sSL https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh | bash -s -- binary
```

Add Fabric binaries to your PATH and set the configuration path:

```bash
export PATH=$PATH:${PWD}/bin
export FABRIC_CFG_PATH=${PWD}/config
```

### Set Environment Variables

Set environment variables for your peer organization:

```bash
export CORE_PEER_LOCALMSPID=ProducersOrgMSP
export CORE_PEER_MSPCONFIGPATH=${PWD}/_msp/ProducersOrg/producersorgadmin/msp
export CORE_PEER_ADDRESS=producersorgpeer-api.127-0-0-1.nip.io:8080
```

```bash
./test-network-k8s/network up createChannel -c propertyChannel -ca

# or

./test-network/network.sh up createChannel -c propertych -ca
```

## Deploying Chaincode 

```bash
 ./test-network-k8s/network chaincode deploy property-transfer ./contracts

#  or
./test-network/network.sh deployCC -ccn propertych -c propertych  
-ccp $PWD/contracts -ccl go

-ccn: chaincode name
-c: channel name
-ccl: chaincode language 
-ccp: chaincode package
```

```bash
./test-network-k8s/network chaincode invoke property-transfer '{"function":"CreateProperty", "Args": ["1", "125", "Martin", "Jeff", "50000", "55000"]}'

export PAYLOAD='{
  "function": "CreateProperty",
  "Args": [
    "1",
    "125",
    "Martin",
    "Jeff",
    "50000",
    "55000"
  ]
}'
```

- Retrieve current chaincode version in Hyperledger Fabric
```bash
./test-network-k8s/network peer chaincode list --instantiated -C test-network 
```

## Author

This project was authored by Siddhant Prateek Mahanayak. You can find more about the author on their [GitHub profile](github.com/siddhantprateek).

Feel free to explore, contribute, and collaborate on the Tractionx project. If you have any questions or need assistance, don't hesitate to reach out to the author.

Happy coding!