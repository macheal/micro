# nohup ./micro --registry=etcd --registry_address=127.0.0.1:2379 api --namespace=" "&
# nohup ./micro --registry=etcd --registry_address=127.0.0.1:2379 web &

nohup ./micro api --namespace=" "&

nohup ./micro web &


