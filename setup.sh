
# setting up network config
export MICROFAB_CONFIG='{
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
}'

# starting docker-compose
docker-compose up -d 
docker run -e MICROFAB_CONFIG -p 8080:8080 -d ibmcom/ibp-microfab

# Setting up weft

curl -s http://console.127-0-0-1.nip.io:8080/ak/api/v1/components | weft microfab -w ./_wallets -p ./_gateways -m ./_msp -f


# Download fabric 
curl -sSL https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh | bash -s -- binary

export PATH=$PATH:${PWD}/bin
export FABRIC_CFG_PATH=${PWD}/config

export CORE_PEER_LOCALMSPID=ProducersOrgMSP

export CORE_PEER_MSPCONFIGPATH=${PWD}/_msp/ProducersOrg/producersorgadmin/msp

export CORE_PEER_ADDRESS=producersorgpeer-api.127-0-0-1.nip.io:8080