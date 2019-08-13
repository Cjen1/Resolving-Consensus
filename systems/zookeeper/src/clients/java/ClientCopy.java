package clients.java;
import org.zeromq.ZMQ;
import org.zeromq.SocketType;
import org.apache.zookeeper.*;
import org.apache.zookeeper.data.*;
import java.util.*;
import OpWire.*;

public class Client implements org.apache.zookeeper.Watcher {
	public static final int CLIENT_PORT = 2181;
	public static final int QUORUM_PORT = 2888;
	public static final int ELECTION_PORT = 3888;
	public static final int SESSION_TIMEOUT = 10000; //Time to session timeout in ms
	public static final Client CLIENT = new Client();

	final java.util.concurrent.CountDownLatch connectedSignal = new java.util.concurrent.CountDownLatch(1);

	public void process(org.apache.zookeeper.WatchedEvent we){ 
		try{
			if (we.getState() == Watcher.Event.KeeperState.SyncConnected) {
				connectedSignal.countDown();
				java.util.concurrent.TimeUnit.SECONDS.sleep(2);
			}
		} catch(Exception e) {
			if (we.getState() == Watcher.Event.KeeperState.SyncConnected) {
				connectedSignal.countDown();
			}
		}
	}
	
	public Client(){ ; }

	public OpWire.Message.Operation receiveOp(ZMQ.Socket socket) throws Exception{
		byte[] payload = socket.recv(0);
		OpWire.Message.Operation op = OpWire.Message.Operation.parseFrom(payload);
		return op;
	}


	public OpWire.Message.Response put(ZooKeeper client, OpWire.Message.Operation opr, int clientId){
		String err = "None";
		String path = "/" + opr.getPut().getKey();
		String data = opr.getPut().getValue().toString();
		long start = 0L, start2 = 0L;
		long stop = 0L, stop2 = 0L;
		try{
			start = System.currentTimeMillis();
			start2 = System.nanoTime();
			client.setData(path, data.getBytes(), -1);
			stop2 = System.nanoTime();
			stop = System.currentTimeMillis();
		} catch(KeeperException.NoNodeException n) {
			start = System.currentTimeMillis();
			start2 = System.nanoTime();
			try{
				client.create(path, data.getBytes(), ZooDefs.Ids.OPEN_ACL_UNSAFE, CreateMode.PERSISTENT);
			} catch(Exception e) {
				err = "Caught Exception: " + e.getMessage();
			}
			stop2 = System.nanoTime();
			stop = System.currentTimeMillis();
		} catch(Exception k) {
			err = "Caught Exception: " + k.getMessage();
		} 
		double duration = (stop2 - start2 + 0.0) / 1E9; // Duration in seconds.
		OpWire.Message.Response resp = OpWire.Message.Response.newBuilder()
							     .setResponseTime(duration)
							     .setErr(err)
							     .setClientStart(start / 1E3)
							     .setQueueStart(opr.getPut().getStart())
							     .setEnd(stop / 1E3)
							     .setClientid(clientId)
							     .setOptype("Write")
							     .setTarget("N/A")
							     .build();
		return resp;
	}


	public OpWire.Message.Response get(ZooKeeper client, OpWire.Message.Operation opr, int clientId){
		String err = "None";
		String path = "/" + opr.getGet().getKey();
		long start = System.currentTimeMillis();
		long start2 = System.nanoTime();
		try{
			client.getData(path, false, client.exists(path, false));
		} catch(Exception e) {
			err = "Caught Exception: " + e.getMessage();
		}
		long stop2 = System.nanoTime();
		long stop = System.currentTimeMillis();
		double duration = (stop2 - start2 + 0.0) / 1E9; // Duration in seconds.
		OpWire.Message.Response resp = OpWire.Message.Response.newBuilder()
							     .setResponseTime(duration)
							     .setErr(err)
							     .setClientStart(start / 1E3)
							     .setQueueStart(opr.getGet().getStart())
							     .setEnd(stop / 1E3)
							     .setClientid(clientId)
							     .setOptype("Read")
							     .setTarget("N/A")
							     .build();
		return resp;
	}


	public static void main(String[] args) throws Exception {
		System.out.println("CLIENT: STARTING");
		Client mainClient = new Client();
		String[] endpoints = args[0].split(",");
		int clientId = Integer.parseInt(args[1]);
		String address = args[2];

		ZMQ.Context context = ZMQ.context(1);			// Creates a context with 1 IOThread.
		ZMQ.Socket socket = context.socket(SocketType.REQ);
		
		String quorum = String.join(":" + CLIENT_PORT + ",", endpoints) + ":" + CLIENT_PORT;

		ZooKeeper cli = new ZooKeeper(quorum, SESSION_TIMEOUT, new Client());
		

		String binding;
		boolean quits = false;
		OpWire.Message.Response resp = null;
		byte[] payload;
		System.out.println("CLIENT: Connecting to socket.");
		socket.connect("ipc://" + address);
		System.out.println("CLIENT: COnnected to socket.");
		socket.send("", 0);
		while(!quits){
			OpWire.Message.Operation opr = mainClient.receiveOp(socket);
			System.out.println("CLIENT: Received Operation.");
			switch(opr.getOpTypeCase().getNumber()){
				case 1: 	//1 = Put
					resp = mainClient.put(cli, opr, clientId);
					break;			
				case 2:		//2 = Get
					resp = mainClient.get(cli, opr, clientId);
					break;				
				case 3:		//3 = Quit
					socket.close();
					context.term();
					return;
				default:
					resp = OpWire.Message.Response.newBuilder()
					      .setResponseTime(-100000.0)
					      .setErr("Operation was not found / supported")
					      .setClientStart(0.0)
					      .setQueueStart(0.0)
					      .setEnd(0.0)
					      .setClientid(clientId)
					      .setOptype("N/A")
					      .setTarget("N/A")
					      .build();
					quits = true;
					break;
			}
			payload = resp.toByteArray();
			socket.send(payload, 0);
		}
	
		socket.close();
		context.term();
	}
}