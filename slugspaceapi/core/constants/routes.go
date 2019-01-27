package constants

const (
	LotDataOverTimeFull            = LotDataOverTimeNoID + "/{id}"
	LotDataOverTimeNoID            = "/v1/lotdataovertime"
	LotByID                        = Lots + "/{id}"
	Lots                           = "/v1/lot"
	UntrackedLots                  = "/v1/untrackedlot"
	UntrackedLotsByID              = UntrackedLots + "/{id}"
	Permits                        = "/v1/permit"
	PermitByID                     = Permits + "/{id}"
	PayStations                        = "/v1/paystation"
	PayStationByID                     = PayStations + "/{id}"
	LotAverageFreespaceByDayNoDate = "/v1/avgfree"
	LotAverageFreespaceByDay       = LotAverageFreespaceByDayNoDate // + "/{day}/{time}"

	RegisterAppInstance = "/v1/register"
)
