import time
import docker
import os

client = docker.from_env()
prometheus_container = client.containers.list(filters={'name': "prometheus"})
print(prometheus_container[0].id)
test_network = client.networks.get(client.networks.list(filters={'name': "net_test"})[0].id)
#print(net_containers)
#print item['Name'][:26] for item in net_containers
#print(test_network.attrs[u'Containers'].values())
print("prometheus connect net_test")
test_network.connect(prometheus_container[0].id)
time.sleep(5)
os.system("docker network inspect net_test")
containers_list = client.containers.list()
print()
#net_containers = test_network.attrs[u'Containers'].values()
#for item in net_containers:
#    print(item[u'Name'][:26])