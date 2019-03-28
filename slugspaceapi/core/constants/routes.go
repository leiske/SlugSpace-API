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
	LotAvailabilities                        = "/v1/lotavailability"
	LotAvailabilityByID                     = LotAvailabilities + "/{id}"
	LotPredictedFreespace       = "/v1/predictfreespace" //needs ?id=1&datetime="XXXX-XX-XX X:X:X" -> nasty that I use multiple styles of parameters here. TODO make parameters consistently ?paramterName=paramterVal
	TrackedLotFullInfo = "/v1/trackedlotfullinfo/{id}"


	RegisterAppInstance = "/v1/register"
)
