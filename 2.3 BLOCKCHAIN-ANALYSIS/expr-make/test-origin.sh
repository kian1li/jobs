#!/bin/bash
# Set workspace as caliper-benchmarks root
WORKSPACE=$(cd "$(dirname '$1')" &>/dev/null && printf "%s/%s" "$PWD" "${1##*/}")
cd ${WORKSPACE}

npx caliper bind --caliper-bind-sut fabric:2.2
# Nominate a target network
NETWORK=network-path

# Enable tests to use existing caliper-core package
#export NODE_PATH=$(which node)

# Build config for target network

#./generate.sh
#cd ${WORKSPACE}

# Available benchmarks
BENCHMARK=benchmark-path
# Available phases
PHASES=("caliper-flow-only-start" "caliper-flow-only-install" "caliper-flow-only-test" "caliper-flow-only-end")

# Execute Phases
function runBenchmark () {
    PHASE=$1
    npx caliper launch manager \
    --caliper-workspace ${WORKSPACE} \
    --caliper-benchconfig ${BENCHMARK} \
    --caliper-networkconfig ${NETWORK} \
    --caliper-fabric-gateway-enabled \
    --caliper-report-path reportpath \
    --${PHASE}

    sleep 5s
} 

# Repeat for PHASES
for PHASE in ${PHASES[@]}; do	
    runBenchmark ${PHASE}
done