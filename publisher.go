package gocdds

/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -lddsc
#include "ddsc/dds.h"
*/
import "C"

type Publisher Participant

func (p *Publisher) CreateWriter(topic interface{}, qos *QoS, listener *Listener) (*Writer, error) {
	return (*Participant)(p).CreateWriter(topic, qos, listener)
}

func (p *Publisher) Delete() error {
	return (*Participant)(p).Delete()
}
