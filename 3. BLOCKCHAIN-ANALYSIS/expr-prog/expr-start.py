import subprocess
import os
from pathlib import Path

home = str(Path.home())
benchmarkpath = home + "/caliper-benchmarks"
exprpath = "/benchmarks/api/fabric"

def checkDF():
    pipe = subprocess.Popen("df -h | awk 'NR == 4 {print $5}'", shell=True, stdout=subprocess.PIPE)
    pipe.wait()
    result = pipe.stdout.read()
    result = int(result[0 : len(result) - 2])
    #print(os.getcwd())
    if (result > 75):
        #fabric restart
        #os.system("rm {}/caliper.log".format(benchmarkpath))
        os.system("rm {}/{}/expr-prog/caliper.log".format(benchmarkpath, exprpath))
        os.system("python3 {}/networkdownup.py fixed-asset index-asset".format(exprpath))

def runExpr(function, variable, number):
    home = str(Path.home())
    #print(home)
    os.chdir(benchmarkpath)
    #print(caliperexpr)
    os.system(".{}/{}/{}/{}-{}/{}-{}-{}.sh > ./{}/expr-result/{}/{}/{}-{}-{}.txt & .{}/{}/{}/{}-{}/{}-{}-{}.sh > ./{}/expr-result/{}/{}/{}-{}-{}.txt".\
    format(exprpath, \
    "fixed-asset", \
    function, \
    function, \
    variable, \
    function, \
    variable, \
    number, \
    exprpath, \
    "fixed-asset", \
    function, \
    function, \
    variable, \
    number, \
    exprpath, \
    "index-asset", \
    function, \
    function, \
    variable, \
    function, \
    variable, \
    number, \
    exprpath, \
    "index-asset", \
    function, \
    function, \
    variable, \
    number, \
    ))

expr_number_list = [1, 2, 3, 4, 5, 6, 7, 8, 9,
                    10, 20, 30 , 40, 50, 60 ,70, 80, 90,
                    100, 200, 300, 400, 500, 600, 700, 800, 900,
                    1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000,
                    10000, 20000, 30000, 40000, 50000, 60000, 70000, 80000, 90000,
                    100000]

expr_chaincode_list = ["fixed-asset", "index-asset"]

expr_function_list = ["create", "get", "update", "delete"]

expr_variable_list = ["txnum", "tps", "bytesize"]

for i, number in enumerate(expr_number_list):
    for j, function in enumerate(expr_function_list):
        for k, variable in enumerate(expr_variable_list):
            checkDF()
            if variable == "tps":
                if (number >= 1 and number <= 200):
                    runExpr(function, variable, number)
            elif variable == "bytesize": 
                if (number >= 80 and number <= 30000):
                    runExpr(function, variable, number)
            elif variable == "txnum":
                runExpr(function, variable, number)