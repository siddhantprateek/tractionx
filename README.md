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

## Author

This project was authored by Siddhant Prateek Mahanayak. You can find more about the author on their [GitHub profile](github.com/siddhantprateek).

Feel free to explore, contribute, and collaborate on the Tractionx project. If you have any questions or need assistance, don't hesitate to reach out to the author or the project's community.

Happy coding!