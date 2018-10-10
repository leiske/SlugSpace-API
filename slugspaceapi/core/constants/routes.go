package constants

const (
	LotDataOverTimeFull            = LotDataOverTimeNoID + "/{lotID}"
	LotDataOverTimeNoID            = "/v1/lotdataovertime"
	LotByIDFull                    = LotByIDNoID + "/{lotID}"
	LotByIDNoID                    = "/v1/lot"
	Lots                           = "/v1/lot"
	LotAverageFreespaceByDayNoDate = "/v1/avgfree"
	LotAverageFreespaceByDay       = LotAverageFreespaceByDayNoDate // + "/{day}/{time}"

	RegisterAppInstance = "/v1/register"
)
