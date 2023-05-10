package gocdds

/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -lddsc
#include "ddsc/dds.h"
*/
import "C"
import (
	"time"
	"unsafe"
)

type QoS C.dds_qos_t

func CreateQoS() *QoS {
	return (*QoS)(C.dds_create_qos())
}

func (qos *QoS) SetDurability(dur Durability) {
	C.dds_qset_durability((*C.dds_qos_t)(qos), C.dds_durability_kind_t(dur))
}

func (qos *QoS) SetHistory(his History, depth int32) {
	C.dds_qset_history((*C.dds_qos_t)(qos), C.dds_history_kind_t(his), C.int32_t(int32(depth)))
}

func (qos *QoS) SetOwnership(own Ownership) {
	C.dds_qset_ownership((*C.dds_qos_t)(qos), C.dds_ownership_kind_t(own))
}

func (qos *QoS) SetLiveliness(liv Liveliness, n time.Duration) {
	C.dds_qset_liveliness((*C.dds_qos_t)(qos), C.dds_liveliness_kind_t(liv), C.int64_t(int64(n)))
}

func (qos *QoS) SetReliability(rel Reliability, n time.Duration) {
	C.dds_qset_reliability((*C.dds_qos_t)(qos), C.dds_reliability_kind_t(rel), C.int64_t(int64(n)))
}

func (qos *QoS) SetDestinationOrder(desord DestinationOrder) {
	C.dds_qset_destination_order((*C.dds_qos_t)(qos), C.dds_destination_order_kind_t(desord))
}

func (qos *QoS) SetPresentationAccess(preacc PresentationAccess, coherent_access, ordered_access bool) {
	C.dds_qset_presentation((*C.dds_qos_t)(qos), C.dds_presentation_access_scope_kind_t(preacc), C.bool(coherent_access), C.bool(ordered_access))
}

func (qos *QoS) SetIgnoreLocal(ignloc IgnoreLocal) {
	C.dds_qset_ignorelocal((*C.dds_qos_t)(qos), C.dds_ignorelocal_kind_t(ignloc))
}

func (qos *QoS) SetConsistency(con Consistency, ignore_sequence_bounds, ignore_string_bounds, ignore_member_names, prevent_type_widening, force_type_validation bool) {
	C.dds_qset_type_consistency((*C.dds_qos_t)(qos), C.dds_type_consistency_kind_t(con), C.bool(ignore_sequence_bounds), C.bool(ignore_string_bounds), C.bool(ignore_member_names), C.bool(prevent_type_widening), C.bool(force_type_validation))
}

func (qos *QoS) SetWriterDataLifecycle(autoDispose bool) {
	C.dds_qset_writer_data_lifecycle((*C.dds_qos_t)(qos), C.bool(autoDispose))
}

func (qos *QoS) SetPartition(num int, partitions *string) {
	C.dds_qset_partition((*C.dds_qos_t)(qos), C.uint32_t(num), (**C.char)(unsafe.Pointer(partitions)))

}

func (qos *QoS) delete() {
	C.dds_delete_qos((*C.dds_qos_t)(qos))
}
