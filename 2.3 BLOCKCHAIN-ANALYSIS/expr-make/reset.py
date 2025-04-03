import os
expr_chaincode_list = ["fixed-asset", "index-asset"]

expr_function_list = ["create", "get", "update", "delete"]

expr_variable_list = ["txnum", "tps", "bytesize"]

for i, chaincode in enumerate(expr_chaincode_list):
    for j, function in enumerate(expr_function_list):
        os.system("rm ../expr-result/{}/{}/*".format(chaincode, function))
        for k, variable in enumerate(expr_variable_list):
            os.system("rm -rf ../{}/{}/{}-{}".format(chaincode, function, function, variable))