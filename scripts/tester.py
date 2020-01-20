from subprocess import call
import numpy as np
from itertools import product

abs_path = '/root/mounted/Resolving-Consensus'

for rate in [1,4,16,32,64,128,256]:
    for n in [3,7,11,15, 19, 23, 25]:
        for system in ['etcd_go']:
            call(
                    [
                        'python',
                        'benchmark.py',
                        system,
                        'tree',
                        '--topo_args', 'n={0},nc={1}'.format(n, 30),
                        'uniform',
                        'none',
                        '--benchmark_config', 'rate={0},'.format(rate) + 
                        'duration=480,'+
                            'dest=../results/{0}_{1}_{2}.res,'.format(n, rate, system)+
                            'logs={0}_{1}_{2}_'.format(n,rate,system),
                        abs_path,
                        #"-d"
                    ]
                )
            call(['bash', 'clean.sh'])

'''
for rate in [4**exp for exp in [0,1,2,3,4,5,6]]:
    for n in [2**exp + 1 for exp in [1,2,3,4,5,6,7]]:
        for system in ['ocaml-paxos_ocaml','etcd_go','zookeeper_java']:
            call(
                    [
                        'python',
                        'benchmark.py',
                        system,
                        'simple',
                        '--topo_args', 'n={0},nc={1}'.format(n, 30),
                        'uniform',
                        'leader',
                        '--benchmark_config', 'rate={0},'.format(rate) + 
                        'duration=480,'+
                            'dest=../results/lf_{0}_{1}_{2}.res'.format(n, rate, system),
                        abs_path,
                        '-d'
                    ]
                )
            call(['bash', 'clean.sh'])
            '''
