import sys
import os

from pathlib import Path
homeDir = str(Path.home())

def alter_sh(file_path, network_config, benchmark_file, report_file):
    alter_list = ['test-origin.sh', 'network-path','benchmark-path', 'report-path']
	#input file
    #print(os.getcwd())
    fin = open("./test-origin.sh", "rt")
    dir_path = os.path.dirname(file_path)
    shell_file = os.path.basename(file_path)
    #output file to write the result to
    os.makedirs(dir_path, exist_ok=True)
    
    fout = open(file_path, "wt")
    #for each line in the input file
    for line in fin:
	    #read replace the string and write to output file
        fout.write(line.replace('network-path', network_config)
        .replace('benchmark-path', '\"'+benchmark_file+'\"')
        .replace('reportpath', '\"'+report_file+'\"'))
    #close input and output files
    fin.close()
    fout.close()
    os.system("chmod +x {}".format(file_path))
    
# usage
# python3 expr-make-sh.py <network-config-fila> <benchmark-config-file> <report-file-path>
def main():
    file_path = sys.argv[1]
    network_config = sys.argv[2]
    benchmark_file = sys.argv[3]
    report_file = sys.argv[4]
    alter_sh(file_path, network_config, benchmark_file, report_file)
main()