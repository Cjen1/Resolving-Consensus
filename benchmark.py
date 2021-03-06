import argparse

from src.client_runner import run_test
from src.distributions import register_ops_args
from src.distributions import get_ops_provider
from src.failures import register_failure_args, get_failure_provider
from src.topologies import register_topo_args, get_topology_provider
from systems import register_system_args, get_system

import logging
logging.basicConfig(
    format="%(asctime)s %(message)s", datefmt="%I:%M:%S %p", level=logging.DEBUG
)


# ------- Parse arguments --------------------------
parser = argparse.ArgumentParser(
    description="Runs a benchmark of a local fault tolerant datastore"
)

register_system_args(parser)
register_topo_args(parser)
register_ops_args(parser)
register_failure_args(parser)

arg_group = parser.add_argument_group("benchmark")
arg_group.add_argument("-d",action="store_true",help="Debug mode")
arg_group.add_argument("--rate", type=float, default=10)
arg_group.add_argument("--duration", type=float, default=120)
arg_group.add_argument("--result-location", default="test.res")

args = parser.parse_args()

system = get_system(args)

net, cluster, clients = get_topology_provider(args).setup()

failure_provider = get_failure_provider(args)

restarters, stoppers = system.start_nodes(cluster)
if args.d:
    from mininet.cli import CLI

    CLI(net)
else:
    ops_provider = get_ops_provider(args)
    failures = failure_provider.get_failures(cluster, system, restarters, stoppers)

    logging.info("BENCHMARK: Starting Test")
    run_test(
        args.result_location,
        clients,
        ops_provider,
        args.rate,
        args.duration,
        system,
        cluster,
        failures,
    )

    for stopper in stoppers.values():
        stopper()

    logging.info("Finished Test")
