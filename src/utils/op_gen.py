import link
import message_pb2 as msg_pb
from numpy import random as rand
import math


class Operation:
    def __init__(self, opType):
        self.type = opType

    @staticmethod
    def system_failure(failure_fn):
        res = Operation(Operation.SYSTEMFAILURE)
        res.fn = failure_fn
        return res

    @staticmethod
    def system_recovery(recovery_fn):
        res = Operation(Operation.SYSTEMRECOVERY)
        res.fn = recovery_fn
        return res

    @staticmethod
    def standard(op):
        res = Operation(Operation.STANDARD)
        res.op = op
        return res

    STANDARD = "standard"
    SYSTEMFAILURE = "systemfailure"
    SYSTEMRECOVERY = "systemrecovery"


# Number of keys to write, size of data in bytes
# Return (<operations>, [])
def sequential_keys(num_ops, data_size, limit=256 ** 4):
    num_keys = num_ops
    if num_keys > limit:
        num_keys = limit - 1

    ops = [Op_put(key, gen_payload(data_size)) for key in range(num_keys)]

    return (ops, [])


## Number of keys to write, size of data in bytes, proportion of the operations to be reads
## Return (<operations>, <prerequisite operations>)
# def mixed_ops(num_ops, num_keys, data_size, split, limit = 256**4):
#    if num_keys > limit:
#        num_keys = limit - 1
#
#    num_reads = int(num_ops * split)
#    num_writes =num_ops - num_reads
#
#    writes = [Op_put(key, gen_payload(data_size)) for key in rand.randint(0, num_keys, num_writes)]
#    reads = [Op_get(key) for key in rand.randint(0, num_keys, num_reads)]
#
#    ops = []
#    ops.extend(writes)
#    ops.extend(reads)
#    rand.shuffle(ops)
#
#    return (ops, [Op_put(key, "") for key in range(num_keys)])

# Dynamically generate a mixed operation
# ratio: The reads/total
def mixed_ops(num_keys=100, data_size=1024, ratio=0.5):
    return (
        lambda: Op_get(rand.randint(0, num_keys))
        if rand.rand() < ratio
        else Op_put(rand.randint(0, num_keys), rand.bytes(data_size)),
        [Op_put(key, rand.bytes(data_size)) for key in range(num_keys)],
    )


# Dynamically generate a write operation
def write_ops(num_keys=100, data_size=1024):
    return (lambda: Op_put(rand.randint(0, num_keys), rand.bytes(data_size)), [])


# Dynamically generate a read operation
def read_ops(num_keys=100, data_size=1024):
    return (
        lambda: Op_get(rand.randint(0, num_keys)),
        [Op_put(key, rand.bytes(data_size)) for key in range(num_keys)],
    )


def gen_payload(num_bytes):
    return rand.bytes(num_bytes)


def Op_put(key, value):
    op = msg_pb.Operation()
    op.put.key = key
    op.put.value = value

    # payload = op.SerializeToString()
    payload = op

    return payload


def Op_get(key):
    op = msg_pb.Operation()
    op.get.key = key

    # payload = op.SerializeToString()
    payload = op

    return payload


def Op_quit():
    op = msg_pb.Operation()
    op.quit.msg = "Quit"

    # payload = op.SerializeToString()
    payload = op

    return payload
