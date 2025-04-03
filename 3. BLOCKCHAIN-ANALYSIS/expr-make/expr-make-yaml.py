import sys
import os
import yaml

from pathlib import Path
homeDir = str(Path.home())

def alter_yaml(file_path, test_name, label_name, chaincode_name, tx_number, tps_number, module_file, create_size):
    with open('./test-origin.yaml') as f:
        doc = yaml.safe_load(f)
        #print(doc)
        doc['test']['name'] = test_name
        doc['test']['rounds'][0]['chaincodeID'] = chaincode_name
        doc['test']['rounds'][0]['label'] = label_name
        doc['test']['rounds'][0]['txNumber'] = int(tx_number)
        doc['test']['rounds'][0]['rateControl']['opts']['tps'] = int(tps_number)
        doc['test']['rounds'][0]['workload']['module'] = module_file
        doc['test']['rounds'][0]['workload']['arguments']['assets'] = int(tx_number)
        doc['test']['rounds'][0]['workload']['arguments']['byteSize'] = int(create_size)
        doc['test']['rounds'][0]['workload']['arguments']['chaincodeID'] = chaincode_name
        create_sizes = [int(create_size)]
        doc['test']['rounds'][0]['workload']['arguments']['create_sizes'] = create_sizes
    with open(file_path, "w+") as f:
        yaml.dump(doc, f, default_flow_style=False,sort_keys=False, explicit_start=True)
    
# usage
# python3 expr-make-sh.py <network-config-fila> <benchmark-config-file> <report-file-path>
def main():
    file_path = sys.argv[1]
    test_name = sys.argv[2]
    label_name = sys.argv[3]
    chaincode_name = sys.argv[4]
    tx_number = sys.argv[5]
    tps_number = sys.argv[6]
    module_file = sys.argv[7]
    create_size = sys.argv[8]
    alter_yaml(file_path, test_name, label_name, chaincode_name, tx_number,
    tps_number, module_file, create_size)
main()