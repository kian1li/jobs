#!/usr/bin/python

import os
import sys
from os import listdir
from os.path import isfile, join
from os import chdir

import yaml

from pathlib import Path

# modify networkConfig.yaml organizations.identities.certificates.clientPrivateKey
#os.system("docker rmi $(docker images -a -q) -f")
#os.system("docker volume rm $(docker volume ls -q -f 'dangling=true')")
homeDir = str(Path.home())
testNetDir = "/fabric-samples/test-network/"
caliperDir = "/caliper-benchmarks/"
prometheusDir = "/networks/prometheus-grafana/"
peerDir = "/fabric-samples/test-network/organizations/peerOrganizations/"

def reset_docker():
	chdir(homeDir+testNetDir)
	os.system("./network.sh down")
	os.system("docker stop $(docker ps -q -a)")
	os.system("docker rm $(docker ps -a -q)")
	os.system("./network.sh down")
	chdir(homeDir+caliperDir+prometheusDir)
	os.system("docker-compose down --remove-orphans")

def init_docker():
	chdir(homeDir+testNetDir)
	os.system("./network.sh up createChannel -c mychannel -s couchdb -ca")

def init_contract(contract_name):
	os.system("./network.sh deployCC -ccn {} -ccp ~/caliper-benchmarks/src/fabric/api/{}/node/ -ccl javascript -ccep \"OR('Org1MSP.peer','Org2MSP.peer')\"".format(contract_name,contract_name))

def alter_yaml(contract_name):
	orgs = sorted(listdir(homeDir+peerDir))
	orgskifullpaths = []
	with open(homeDir+"/caliper-benchmarks/networks/{}-networkConfig.yaml".format(contract_name)) as f:
		doc = yaml.safe_load(f)
		doc['channels'][0]['contracts'][0]['id'] = contract_name
		print(doc['channels'])
		for i, org in enumerate(orgs):
			keyStore = "/users/User1@"+org+"/msp/keystore/"
			orgskipath = homeDir+peerDir+org+keyStore
			onlyfiles = [f for f in listdir(orgskipath) if isfile(join(orgskipath, f))]
			orgskifullpaths.append(homeDir+peerDir+org+keyStore+onlyfiles[0])
		for i, org in enumerate(orgs):
			doc['organizations'][i]['identities']['certificates'][0]['clientPrivateKey']['path'] = orgskifullpaths[i]
	with open(homeDir+'/caliper-benchmarks/networks/{}-networkConfig.yaml'.format(contract_name), "w") as f:
		yaml.dump(doc, f)
# docker-compose grafana prometheus up network net_test <prometheus_container>
# prometheus dir ~/caliper-benchmarks/networks/prometheus-grafana
def monitor_container_up():
	chdir(homeDir+caliperDir+prometheusDir)
	os.system("docker-compose -f docker-compose-fabric.yaml up")

def main():
	first_contract = sys.argv[1]
	second_contract = sys.argv[2]
	reset_docker()
	init_docker()
	init_contract(first_contract)
	init_contract(second_contract)
	alter_yaml(first_contract)
	alter_yaml(second_contract)
	monitor_container_up()
main()
