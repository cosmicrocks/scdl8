package constants

// PostgreSQL specific constants
const (
	EventStreamCRDApiVersion     = "cosmic.rocks/v1"
	EventStreamCRDKind           = "FabricEventStream"
	EventStreamCRDName           = "fabriceventstreams.cosmic.rocks"
	EventStreamSourcePGType      = "PostgresLogicalReplication"
	EventStreamSourceSlotPrefix  = "fes"
	EventStreamSourcePluginType  = "pgoutput"
	EventStreamSourceAuthType    = "DatabaseAuthenticationSecret"
	EventStreamFlowPgGenericType = "PostgresWalToGenericNakadiEvent"
	EventStreamSinkNakadiType    = "Nakadi"
)
