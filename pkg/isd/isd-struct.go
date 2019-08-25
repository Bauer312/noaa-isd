package isd

/*
LiquidPrecipitation contains information about liquid precipitation
*/
type LiquidPrecipitation struct {
	Period    int
	Depth     float64
	Condition string
	Quality   string
}

/*
SkyCover contains information about sky cover
*/
type SkyCover struct {
	Code         string
	CodeQC       string
	BaseHeight   int
	BaseHeightQC string
	CloudType    string
	CloudTypeQC  string
}

/*
SkyCoverSummation contains information about sky cover summations
*/
type SkyCoverSummation struct {
	Code           string
	Code2          string
	CodeQC         string
	Height         int
	HeightQC       string
	Characteristic string
}

/*
SkyCondition contains information about sky conditions
*/
type SkyCondition struct {
	ConvectiveCloud      string
	VerticalDatum        string
	BaseHeightUpperRange int
	BaseHeightLowerRange int
}

/*
SkyConditionObservation contains information about sky conditions
*/
type SkyConditionObservation struct {
	Coverage             string
	OpaqueCoverage       string
	CoverageQC           string
	LowestCloud          string
	LowestCloudQC        string
	LowCloudGenus        string
	LowCloudGenusQC      string
	LowCloudBaseHeight   int
	LowCloudBaseHeightQC string
	MidCloudGenus        string
	MidCloudGenusQC      string
	HighCloudGenus       string
	HighCloudGenusQC     string
}

/*
PressureObservation contains information about atmospheric pressure observations
*/
type PressureObservation struct {
	AltimiterSetting  float64
	AltimiterQC       string
	StationPressure   float64
	StationPressureQC string
}

/*
StationPressureDay contains station pressure information for the day
*/
type StationPressureDay struct {
	AvgStationPressure    float64
	AvgStationPressureQC  string
	MinSeaLevelPressure   float64
	MinSeaLevelPressureQC string
}

/*
PressureChange contains information about atmospheric pressure changes
*/
type PressureChange struct {
	Tendency         string
	TendencyQC       string
	ThreeHour        float64
	ThreeHourQC      string
	TwentyFourHour   float64
	TwentyFourHourQC string
}

/*
WindGust contains information about wind gusts
*/
type WindGust struct {
	Speed   float64
	SpeedQC string
}

/*
WindObservation contains information about wind observations
*/
type WindObservation struct {
	ObservationType string
	Hours           int
	Speed           float64
	SpeedQC         string
	Direction       int
}

/*
WindSummaryDay contains wind summary information for a day
*/
type WindSummaryDay struct {
	Code      string
	Hours     int
	Speed     float64
	Direction int
	UTC       string
	Quality   string
}

/*
ExtremeTemp contains information about extreme temperature observations
*/
type ExtremeTemp struct {
	Hours         float64
	Code          string
	Temperature   float64
	TemperatureQC string
}

/*
PresentWeather contains information about present weather conditions
*/
type PresentWeather struct {
	Intensity      string
	Descriptor     string
	Precipitation  string
	Obscuration    string
	OtherPhenomena string
	Combination    string
	Quality        string
}

/*
SnowDepth contains information about snow
*/
type SnowDepth struct {
	Depth                    int
	DepthCode                string
	DepthQC                  string
	EquivalentWaterDepth     float64
	EquivalentWaterDepthCode string
	EquivalentWaterDepthQC   string
}

/*
SnowAccumulation contains information about snow accumulation
*/
type SnowAccumulation struct {
	Hours     int
	Depth     int
	Condition string
	Quality   string
}

/*
PeriodSnowAccumulation contains information about snow accumulation for the day or month
*/
type PeriodSnowAccumulation struct {
	Hours     int
	Depth     float64
	Condition string
	Quality   string
}

/*
Remark contains remarks section in ISD record
*/
type Remark struct {
	Type   string
	Length int
	Remark string
}

/*
MonthlyPrecipitation contains monthly precipitation data
*/
type MonthlyPrecipitation struct {
	Depth     float64
	Condition string
	Quality   string
}

/*
MonthlyPrecipitationRecord contains monthly precipitation 24-hour recorddata
*/
type MonthlyPrecipitationRecord struct {
	Depth     float64
	Condition string
	Date1     string
	Date2     string
	Date3     string
	Quality   string
}

/*
MonthlyPrecipitationShortDuration contains information about monthly short duration precipitation
*/
type MonthlyPrecipitationShortDuration struct {
	Minutes    int
	Depth      float64
	Condition  string
	EndingDate string
	Quality    string
}

/*
MonthlyPrecipitationShortDurationMax contains information about monthly short duration maximum precipitation
*/
type MonthlyPrecipitationShortDurationMax struct {
	Minutes    int
	Depth      float64
	Condition  string
	EndingDate string
	Quality    string
}

/*
PrecipitationDays contains the number of days that certain precipitation amounts occurred
*/
type PrecipitationDays struct {
	HundredthDays   int
	HundredthDaysQC string
	TenthDays       int
	TenthDaysQC     string
	HalfDays        int
	HalfDaysQC      string
	OneDays         int
	OneDaysQC       string
}

/*
AverageAirTemperature contains information about average air temperature
*/
type AverageAirTemperature struct {
	Hours       int
	Code        string
	Temperature float64
	Quality     string
}

/*
MonthlyExtremeTemp contains information about extreme temperatures for a month
*/
type MonthlyExtremeTemp struct {
	Code        string
	Condition   string
	Temperature float64
	Date        string
	Quality     string
}

/*
HeatingCoolingDays contains information about the number of heating and cooling days in a month
*/
type HeatingCoolingDays struct {
	Hours    int
	Code     string
	Quantity int
	Quality  string
}

/*
CriteriaDaysTemp contains information about the number of days in which the high or low is higher or lower than a max or min value
*/
type CriteriaDaysTemp struct {
	MaxLow         int
	MaxLowQC       string
	MaxHigh        int
	MaxHighQC      string
	MinLow         int
	MinLowQC       string
	MinReallyLow   int
	MinReallyLowQC string
}

/*
AdditionalData contains all of the optional additional data in an ISD record
*/
type AdditionalData struct {
	LiquidPrecipitation                  [4]LiquidPrecipitation
	SkyCover                             [6]SkyCover
	SkyCoverSummation                    [6]SkyCoverSummation
	SkyCondition                         SkyCondition
	SkyConditionObservation              SkyConditionObservation
	PressureObservation                  PressureObservation
	PressureChange                       PressureChange
	WindGust                             WindGust
	Remark                               Remark
	WindObservation                      [3]WindObservation
	ExtremeTemp                          [4]ExtremeTemp
	PresentWeather                       [9]PresentWeather
	SnowDepth                            SnowDepth
	SnowAccumulation                     [4]SnowAccumulation
	PeriodSnowAccumulation               PeriodSnowAccumulation
	StationPressureDay                   StationPressureDay
	WindSummaryDay                       [3]WindSummaryDay
	MonthlyPrecipitation                 MonthlyPrecipitation
	MonthlyPrecipitationRecord           MonthlyPrecipitationRecord
	PrecipitationDays                    PrecipitationDays
	MonthlyPrecipitationShortDuration    [6]MonthlyPrecipitationShortDuration
	MonthlyPrecipitationShortDurationMax [6]MonthlyPrecipitationShortDurationMax
	AverageAirTemperature                [3]AverageAirTemperature
	MonthlyExtremeTemp                   [2]MonthlyExtremeTemp
	HeatingCoolingDays                   [2]HeatingCoolingDays
	CriteriaDaysTemp                     CriteriaDaysTemp
}
