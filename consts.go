package gocdds

/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -lddsc
#include "ddsc/dds.h"
*/
import "C"

const DomainDefault = C.DDS_DOMAIN_DEFAULT

type Durability C.dds_durability_kind_t
type History C.dds_history_kind_t
type Ownership C.dds_ownership_kind_t
type Liveliness C.dds_liveliness_kind_t
type Reliability C.dds_reliability_kind_t
type DestinationOrder C.dds_destination_order_kind_t
type PresentationAccess C.dds_presentation_access_scope_kind_t
type IgnoreLocal C.dds_ignorelocal_kind_t
type Consistency C.dds_type_consistency_kind_t

// Durability QoS: Applies to Topic, DataReader, DataWriter
const (
	DurabilityVolatile       Durability = C.DDS_DURABILITY_VOLATILE
	DurabilityTransientLocal Durability = C.DDS_DURABILITY_TRANSIENT_LOCAL
	DurabilityTransient      Durability = C.DDS_DURABILITY_TRANSIENT
	DurabilityPersistent     Durability = C.DDS_DURABILITY_PERSISTENT
)

// History QoS: Applies to Topic, DataReader, DataWriter
const (
	HistoryKeepLast History = C.DDS_HISTORY_KEEP_LAST
	HistoryKeepAll  History = C.DDS_HISTORY_KEEP_ALL
)

// Ownership QoS: Applies to Topic, DataReader, DataWriter
const (
	OwnershipShared    Ownership = C.DDS_OWNERSHIP_SHARED
	OwnershipExclusive Ownership = C.DDS_OWNERSHIP_EXCLUSIVE
)

// Liveliness QoS: Applies to Topic, DataReader, DataWriter
const (
	LivelinessAutomatic           Liveliness = C.DDS_LIVELINESS_AUTOMATIC
	LivelinessManualByParticipant Liveliness = C.DDS_LIVELINESS_MANUAL_BY_PARTICIPANT
	LivelinessManualByTopic       Liveliness = C.DDS_LIVELINESS_MANUAL_BY_TOPIC
)

// Reliability QoS: Applies to Topic, DataReader, DataWriter
const (
	ReliabilityBestEffort Reliability = C.DDS_RELIABILITY_BEST_EFFORT
	ReliabilityReliable   Reliability = C.DDS_RELIABILITY_RELIABLE
)

// DestinationOrder QoS: Applies to Topic, DataReader, DataWriter
const (
	DestOrderReception DestinationOrder = C.DDS_DESTINATIONORDER_BY_RECEPTION_TIMESTAMP
	DestOrderSource    DestinationOrder = C.DDS_DESTINATIONORDER_BY_SOURCE_TIMESTAMP
)

// Presentation QoS: Applies to Publisher, Subscriber
const (
	PresentationInstance PresentationAccess = C.DDS_PRESENTATION_INSTANCE
	PresentationTopic    PresentationAccess = C.DDS_PRESENTATION_TOPIC
	PresentationGroup    PresentationAccess = C.DDS_PRESENTATION_GROUP
)

// Ignore-local QoS: Applies to DataReader, DataWriter
const (
	IgnoreLocalNone        IgnoreLocal = C.DDS_IGNORELOCAL_NONE
	IgnoreLocalParticipant IgnoreLocal = C.DDS_IGNORELOCAL_PARTICIPANT
	IgnoreLocalProcess     IgnoreLocal = C.DDS_IGNORELOCAL_PROCESS
)

const (
	ConsistencyDisallow Consistency = C.DDS_TYPE_CONSISTENCY_DISALLOW_TYPE_COERCION
	ConsistencyAllow    Consistency = C.DDS_TYPE_CONSISTENCY_ALLOW_TYPE_COERCION
)

type CommunicationStatus C.uint32_t

const (
	CommunicationNil   CommunicationStatus = 0
	PublicationMatched CommunicationStatus = C.DDS_PUBLICATION_MATCHED_STATUS
)

type ReadConditionState uint32

const (
	ReadSampleState    ReadConditionState = C.DDS_READ_SAMPLE_STATE
	NotReadSampleState ReadConditionState = C.DDS_NOT_READ_SAMPLE_STATE
	AnySampleState     ReadConditionState = C.DDS_ANY_SAMPLE_STATE

	NewViewState    ReadConditionState = C.DDS_NEW_VIEW_STATE
	NotNewViewState ReadConditionState = C.DDS_NOT_NEW_VIEW_STATE
	AnyViewState    ReadConditionState = C.DDS_ANY_VIEW_STATE

	AliveInstanceState             ReadConditionState = C.DDS_ALIVE_INSTANCE_STATE
	NotAliveDisposedInstanceState  ReadConditionState = C.DDS_NOT_ALIVE_DISPOSED_INSTANCE_STATE
	NotAliveNoWritersInstanceState ReadConditionState = C.DDS_NOT_ALIVE_NO_WRITERS_INSTANCE_STATE
	AnyInstanceState               ReadConditionState = C.DDS_ANY_INSTANCE_STATE

	AnyState ReadConditionState = C.DDS_ANY_STATE
)

type InstanceState uint32

const (
	IstAlive             InstanceState = C.DDS_ALIVE_INSTANCE_STATE
	IstNotAliveDisposed  InstanceState = C.DDS_NOT_ALIVE_DISPOSED_INSTANCE_STATE
	IstNotAliceNoWriters InstanceState = C.DDS_NOT_ALIVE_NO_WRITERS_INSTANCE_STATE
)
