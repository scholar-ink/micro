package main

import (
	_ "github.com/micro/go-plugins/broker/rabbitmq"
	_ "github.com/micro/go-plugins/broker/nsq"
	_ "github.com/micro/go-plugins/broker/kafka"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	_ "github.com/micro/go-plugins/registry/etcd"
	_ "github.com/micro/go-plugins/registry/consul"
)
