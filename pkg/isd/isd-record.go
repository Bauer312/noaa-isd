package isd

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

/*
Record contains the data in an ISD record
*/
type Record struct {
	Station                          string
	Date                             string
	Time                             string
	DataSourceFlag                   string
	Latitude                         float64
	Longitude                        float64
	SurfaceObservationCode           string
	Elevation                        float64
	CallLetters                      string
	QualityControlProcess            string
	WindDirectionAngle               int
	WindDirectionQualityCode         string
	WindObservationTypeCode          string
	WindSpeed                        float64
	WindSpeedQualityCode             string
	CeilingHeight                    int
	CeilingQualityCode               string
	CeilingDeterminationCode         string
	CavokCode                        string
	Visibility                       int
	VisibilityQualityCode            string
	VisibilityVariability            string
	VisibilityVariabilityQualityCode string
	Temperature                      float64
	TemperatureQualityCode           string
	DewPoint                         float64
	DewPointQualityCode              string
	SeaLevelPressure                 float64
	SeaLevelPressureQualityCode      string
	RelativeHumidity                 float64
	SaturationVaporPressure          float64
	AirDensity                       float64
	Extra                            AdditionalData
}

/*
RecordString returns a string containing delimited data
*/
func (r *Record) RecordString(delim string) string {
	var b strings.Builder

	b.WriteString(r.Station)
	b.WriteString(delim)
	b.WriteString(r.Date)
	b.WriteString(delim)
	b.WriteString(r.Time)
	b.WriteString(delim)
	fmt.Fprint(&b, r.Latitude)
	b.WriteString(delim)
	fmt.Fprint(&b, r.Longitude)
	b.WriteString(delim)
	fmt.Fprint(&b, r.Elevation)
	b.WriteString(delim)
	b.WriteString(r.CallLetters)
	b.WriteString(delim)
	fmt.Fprint(&b, r.WindDirectionAngle)
	b.WriteString(delim)
	b.WriteString(r.WindObservationTypeCode)
	b.WriteString(delim)
	fmt.Fprint(&b, r.WindSpeed)
	b.WriteString(delim)
	fmt.Fprint(&b, r.CeilingHeight)
	b.WriteString(delim)
	b.WriteString(r.CeilingDeterminationCode)
	b.WriteString(delim)
	fmt.Fprint(&b, r.Visibility)
	b.WriteString(delim)
	b.WriteString(r.VisibilityVariability)
	b.WriteString(delim)
	fmt.Fprint(&b, r.Temperature)
	b.WriteString(delim)
	fmt.Fprint(&b, r.DewPoint)
	b.WriteString(delim)
	fmt.Fprint(&b, r.SeaLevelPressure)
	b.WriteString(delim)
	fmt.Fprint(&b, r.RelativeHumidity)
	b.WriteString(delim)
	fmt.Fprint(&b, r.SaturationVaporPressure)
	b.WriteString(delim)
	fmt.Fprint(&b, r.Extra.PressureObservation.StationPressure)
	b.WriteString(delim)
	fmt.Fprint(&b, r.AirDensity)

	return b.String()
}

/*
Parse takes an ISD string and returns a new record
*/
func Parse(line string) Record {
	var rc Record

	rc.Station = fmt.Sprint(line[4:15])
	rc.Date = fmt.Sprint(line[15:23])
	rc.Time = fmt.Sprint(line[23:27])
	rc.DataSourceFlag = fmt.Sprint(line[27:28])

	lat1000, err := strconv.ParseFloat(line[28:34], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.Latitude = lat1000 / 1000.0

	lon1000, err := strconv.ParseFloat(line[34:41], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.Longitude = lon1000 / 1000.0

	rc.SurfaceObservationCode = fmt.Sprint(line[41:46])

	elevation, err := strconv.ParseFloat(line[46:51], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.Elevation = elevation

	rc.CallLetters = fmt.Sprint(line[51:56])

	rc.QualityControlProcess = fmt.Sprint(line[56:60])

	dir, err := strconv.Atoi(line[60:63])
	if err != nil {
		log.Println("Wind Direction Angle")
		log.Fatal(err)
	}
	rc.WindDirectionAngle = dir

	rc.WindDirectionQualityCode = fmt.Sprint(line[63:64])

	rc.WindObservationTypeCode = fmt.Sprint(line[64:65])

	windSpeed10, err := strconv.ParseFloat(line[65:69], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.WindSpeed = windSpeed10 / 10.0

	rc.WindSpeedQualityCode = fmt.Sprint(line[69:70])

	ceil, err := strconv.Atoi(line[70:75])
	if err != nil {
		log.Println("Ceiling Height")
		log.Fatal(err)
	}
	rc.CeilingHeight = ceil

	rc.CeilingQualityCode = fmt.Sprint(line[75:76])

	rc.CeilingDeterminationCode = fmt.Sprint(line[76:77])

	rc.CavokCode = fmt.Sprint(line[77:78])

	vis, err := strconv.Atoi(line[78:84])
	if err != nil {
		log.Println("Visibility")
		log.Fatal(err)
	}
	rc.Visibility = vis

	rc.VisibilityQualityCode = fmt.Sprint(line[84:85])

	rc.VisibilityVariability = fmt.Sprint(line[85:86])

	rc.VisibilityVariabilityQualityCode = fmt.Sprint(line[86:87])

	tempC10, err := strconv.ParseFloat(line[87:92], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.Temperature = tempC10 / 10.0

	rc.TemperatureQualityCode = fmt.Sprint(line[92:93])

	dewPointC10, err := strconv.ParseFloat(line[93:98], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.DewPoint = dewPointC10 / 10.0

	rc.DewPointQualityCode = fmt.Sprint(line[98:99])

	seaLevelPressure10, err := strconv.ParseFloat(line[99:104], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.SeaLevelPressure = convertHectoPascalToMMHG(seaLevelPressure10 / 10.0)

	rc.SeaLevelPressureQualityCode = fmt.Sprint(line[104:105])

	rc.RelativeHumidity = getRelativeHumidity(rc.DewPoint, rc.Temperature)

	rc.SaturationVaporPressure = getSVP(rc.Temperature)

	if line[105:108] == "ADD" {
		rc.Extra = parseAdditionalDataSection(line)
	}

	rc.AirDensity = getAirDensity(rc.Temperature,
		rc.Extra.PressureObservation.StationPressure,
		rc.SaturationVaporPressure,
		rc.RelativeHumidity)

	return rc
}

/*
BasicHeader returns a string containing a header for the basic data output
*/
func BasicHeader(delim string) string {
	var b strings.Builder

	b.WriteString("STATIONID")
	b.WriteString(delim)
	b.WriteString("SV_DATE")
	b.WriteString(delim)
	b.WriteString("SV_TIME")
	b.WriteString(delim)
	b.WriteString("LATITUDE")
	b.WriteString(delim)
	b.WriteString("LONGITUDE")
	b.WriteString(delim)
	b.WriteString("ELEVATION")
	b.WriteString(delim)
	b.WriteString("CALLLETTERS")
	b.WriteString(delim)
	b.WriteString("WINDDIRECTION")
	b.WriteString(delim)
	b.WriteString("WINDCODE")
	b.WriteString(delim)
	b.WriteString("WINDSPEED")
	b.WriteString(delim)
	b.WriteString("CEILING")
	b.WriteString(delim)
	b.WriteString("CEILINGCODE")
	b.WriteString(delim)
	b.WriteString("VISIBILITY")
	b.WriteString(delim)
	b.WriteString("VISIBILITYVARIABILITY")
	b.WriteString(delim)
	b.WriteString("TEMPERATURE")
	b.WriteString(delim)
	b.WriteString("DEWPOINT")
	b.WriteString(delim)
	b.WriteString("SEALEVELPRESSURE")
	b.WriteString(delim)
	b.WriteString("RELATIVEHUMIDITY")
	b.WriteString(delim)
	b.WriteString("SVP")
	b.WriteString(delim)
	b.WriteString("STATIONPRESSURE")
	b.WriteString(delim)
	b.WriteString("AIRDENSITY")

	return b.String()
}

func getSVP(temp float64) float64 {
	/*
		Use Buck's equation from 1996 - the return value is in kilopascals
	*/
	if temp > 0.0 {
		return 0.61121 * math.Exp((18.678-(temp/234.5))*(temp/(257.14+temp)))
	}
	return 0.61115 * math.Exp((23.036-(temp/333.7))*(temp/(279.82+temp)))
}

func getRelativeHumidity(dewPoint, tempC float64) float64 {
	/*
		http://irtfweb.ifa.hawaii.edu/~tcs3/tcs3/Misc/Dewpoint_Calculation_Humidity_Sensor_E.pdf

		H = (log10(RH)-2)/0.4343 + (17.62*T)/(243.12+T);
		17.62Dp = H * (243.12 + Dp);
	*/
	h := (17.62 * dewPoint) / (243.12 + dewPoint)
	return math.Pow(10.0, (0.4343*(h-(17.62*tempC)/(243.12+tempC)) + 2.0))
}

func convertHectoPascalToMMHG(val float64) float64 {
	return val * 0.7500638
}

func getAirDensity(temp, stationPressure, svp, humidity float64) float64 {
	return 1.2929 * (273.0 / (temp + 273.0)) * ((stationPressure - (0.379 * svp * humidity / 100.0)) / 760.0)
}
