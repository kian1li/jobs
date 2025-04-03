import os
"""
expr-make-sh.py
argv[0]: python3 
argv[1]: expr-make-sh.py 
argv[2]: <network-config-fila> 
argv[3]: <benchmark-config-file> 
argv[4]: <report-file-path>
test-case: create/read/update/delete
chaincode: fixed-asset/index-asset

expr-make-yaml.py
argv[0]: python3 
argv[1]: expr-make-yaml.py 
argv[2]: yaml_file_path: ../test-asset/create/create-txnum/create-txnum-1000.yaml 
argv[3]: test_name: test-asset-test 
argv[4]: label_name: test-asset-create-txnum-1000 
argv[5]: chaincode_name: test-asset 
argv[6]: tx_number: 1000 
argv[7]: tps_number: 20 
argv[8]: moduel_file_path: test-asset-lib/create.js 
argv[9]: create_size: 100

chaincode: fixed-asset/index-asset
function: create/get/update/delete
number-case: 
1 - 9
10 - 90
100 - 900
1000 - 9000
10000 - 90000
100000
variable-value: 
txNumber: minimum(1) ~ maximum(100000)
bytesize: minimum(80) ~ maximum(100000)
send_rate: minimum(1) ~ maximum(100000)
"""
expr_number_list = [1, 2, 3, 4, 5, 6, 7, 8, 9,
                    10, 20, 30 , 40, 50, 60 ,70, 80, 90,
                    100, 200, 300, 400, 500, 600, 700, 800, 900,
                    1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000,
                    10000, 20000, 30000, 40000, 50000, 60000, 70000, 80000, 90000,
                    100000]

expr_chaincode_list = ["fixed-asset", "index-asset"]

expr_function_list = ["create", "get", "update", "delete"]

expr_variable_list = ["txnum", "tps", "bytesize"]

def alter_sh(chaincode, function, variable, number):
    os.system("python3 expr-make-sh.py \
                ../{}/{}/{}-{}/{}-{}-{}.sh \
                networks/{}-networkConfig.yaml \
                benchmarks/api/fabric/{}/{}/{}-{}/{}-{}-{}.yaml \
                result/{}/{}/{}-{}-{}.html".\
                format(chaincode, \
                function, \
                function, \
                variable, \
                function, \
                variable, \
                number, \
                chaincode, \
                chaincode, \
                function, \
                function, \
                variable, \
                function, \
                variable, \
                number, \
                chaincode, \
                function, \
                function, \
                variable, \
                number))

def alter_yaml(chaincode, function, variable, tx_number, tps_number, create_size):
    os.system("python3 expr-make-yaml.py \
                ../{}/{}/{}-{}/{}-{}-{}.yaml  \
                {}-test \
                {}-{}-{}-{} \
                {} \
                {} \
                {} \
                benchmarks/api/fabric/{}-lib/{}-assets.js \
                {}".\
                format(chaincode, \
                function, \
                function, \
                variable, \
                function, \
                variable, \
                number, \
                chaincode, \
                chaincode, \
                function, \
                variable, \
                number, \
                chaincode, \
                tx_number, \
                tps_number, \
                chaincode, \
                function, \
                create_size
                ))

def alter_txnum(chaincode, function, variable, number, tps_number = 20, create_size = 100):
    alter_yaml(chaincode, function, variable, number, tps_number, create_size)
def alter_tps_number(chaincode, function, variable, number, tx_number = 100, create_size = 100):
    alter_yaml(chaincode, function, variable, tx_number, number, create_size)
def alter_create_size(chaincode, function, variable, number, tx_number = 100, tps_number = 100):
    alter_yaml(chaincode, function, variable, tx_number, tps_number, number)
"""
python3 expr-make-sh.py 
../test-asset/create/create-txnum/create-txnum-1000.sh 
networks/test-asset-networkConfig.yaml 
benchmarks/api/fabric/test-asset/create/create-txnum/create-txnum-1000.yaml 
result/test-asset/create/create-txnum-1000.html

python3 expr-make-yaml.py 
../test-asset/create/create-txnum/create-txnum-1000.yaml 
test-asset-test 
test-asset-create-txnum-1000 
test-asset 
txNumber 1000 
send_rate 20 
module test-asset-lib/create.js 
create_size 100
"""
for i, chaincode in enumerate(expr_chaincode_list):
    for j, function in enumerate(expr_function_list):
        for k, variable in enumerate(expr_variable_list):
            for l, number in enumerate(expr_number_list):
                if variable == "txnum":
                    alter_sh(chaincode, function, variable, number)
                    alter_txnum(chaincode, function, variable, number)
                elif variable == "tps":
                    if (number >= 1 and number <= 200):
                        alter_sh(chaincode, function, variable, number)
                        alter_tps_number(chaincode, function, variable, number)
                elif variable == "bytesize": 
                    if (number >= 80 and number <= 30000):
                        alter_sh(chaincode, function, variable, number)
                        alter_create_size(chaincode, function, variable, number)