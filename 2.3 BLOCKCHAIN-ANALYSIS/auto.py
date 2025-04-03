#!/usr/bin/python

from os import listdir
from os.path import isfile, join

import docker
import yaml

from pathlib import Path

# modify networkConfig.yaml organizations.identities.certificates.clientPrivateKey
homeDir = str(Path.home())
peerDir = "/fabric-samples/test-network/organizations/peerOrganizations/"
orgs = sorted(listdir(homeDir+peerDir))
orgskifullpaths = []
with open(homeDir+"/caliper-benchmarks/networks/networkConfig.yaml") as f:
	doc = yaml.safe_load(f)
	for i, org in enumerate(orgs):
		keyStore = "/users/User1@"+org+"/msp/keystore/"
		orgskipath = homeDir+peerDir+org+keyStore
		onlyfiles = [f for f in listdir(orgskipath) if isfile(join(orgskipath, f))]
		orgskifullpaths.append(homeDir+peerDir+org+keyStore+onlyfiles[0])
	for i, org in enumerate(orgs):
		doc['organizations'][i]['identities']['certificates'][0]['clientPrivateKey']['path'] = orgskifullpaths[i]
with open(homeDir+'/caliper-benchmarks/networks/networkConfig.yaml', "w") as f:
	yaml.dump(doc, f)
# docker network net_test <prometheus_container>
client = docker.from_env()
print(client.containers.list())
