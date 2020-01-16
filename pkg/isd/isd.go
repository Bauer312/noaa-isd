package isd

import (
	"fmt"
	"log"
	"strconv"
)

func parseAdditionalDataSection(line string) AdditionalData {
	var extra AdditionalData
	var rec string
	var err error
	var qty int
	//The additional data section starts at 108
	cur := 108
	//fmt.Println(line[108:])
	for cur < len(line) {

		if line[cur:(cur+3)] == "REM" {
			rec = "REM"
			qty = 1
		} else {
			rec = fmt.Sprint(line[cur:(cur + 2)])
			qty, err = strconv.Atoi(line[(cur + 2):(cur + 3)])
			if err != nil {
				log.Println("ADDITIONAL DATA SECTION", cur, line)
				log.Fatal(err)
			}
		}
		cur += 3
		switch rec {
		case "AA":
			extra.LiquidPrecipitation[qty-1] = parseAA(cur, line)
			cur += 8
		case "AB":
			extra.MonthlyPrecipitation = parseAB(cur, line)
			cur += 7
		case "AD":
			extra.MonthlyPrecipitationRecord = parseAD(cur, line)
			cur += 19
		case "AE":
			extra.PrecipitationDays = parseAE(cur, line)
			cur += 12
		case "AH":
			extra.MonthlyPrecipitationShortDuration[qty-1] = parseAH(cur, line)
			cur += 15
		case "AI":
			extra.MonthlyPrecipitationShortDurationMax[qty-1] = parseAI(cur, line)
			cur += 15
		case "AU":
			extra.PresentWeather[qty-1] = parseAU(cur, line)
			cur += 8
		case "AK":
			cur += 12
		case "AJ":
			extra.SnowDepth = parseAJ(cur, line)
			cur += 14
		case "AL":
			extra.SnowAccumulation[qty-1] = parseAL(cur, line)
			cur += 7
		case "AM":
			cur += 18
		case "AN":
			extra.PeriodSnowAccumulation = parseAN(cur, line)
			cur += 9
		case "AW":
			cur += 3
		case "AT":
			cur += 9
		case "AX":
			cur += 6
		case "ED":
			cur += 8
		case "GA":
			extra.SkyCover[qty-1] = parseGA(cur, line)
			cur += 13
		case "GD":
			extra.SkyCoverSummation[qty-1] = parseGD(cur, line)
			cur += 12
		case "GE":
			extra.SkyCondition = parseGE(cur, line)
			cur += 19
		case "GF":
			extra.SkyConditionObservation = parseGF(cur, line)
			cur += 23
		case "GJ":
			cur += 5
		case "GK":
			cur += 4
		case "MA":
			extra.PressureObservation = parseMA(cur, line)
			cur += 12
		case "MD":
			extra.PressureChange = parseMD(cur, line)
			cur += 11
		case "OC":
			extra.WindGust = parseOC(cur, line)
			cur += 5
		case "OD":
			extra.WindObservation[qty-1] = parseOD(cur, line)
			cur += 11
		case "KA":
			extra.ExtremeTemp[qty-1] = parseKA(cur, line)
			cur += 10
		case "KB":
			extra.AverageAirTemperature[qty-1] = parseKB(cur, line)
			cur += 10
		case "KC":
			extra.MonthlyExtremeTemp[qty-1] = parseKC(cur, line)
			cur += 14
		case "KD":
			extra.HeatingCoolingDays[qty-1] = parseKD(cur, line)
			cur += 9
		case "KE":
			extra.CriteriaDaysTemp = parseKE(cur, line)
			cur += 12
		case "KG":
			cur += 11
		case "MF":
			cur += 12
		case "MG":
			extra.StationPressureDay = parseMG(cur, line)
			cur += 12
		case "MH":
			cur += 12
		case "MK":
			cur += 24
		case "MW":
			cur += 3
		case "MV":
			cur += 3
		case "OE":
			extra.WindSummaryDay[qty-1] = parseOE(cur, line)
			cur += 16
		case "RH":
			cur += 9
		case "WA":
			cur += 6
		case "REM":
			extra.Remark = parseREM(cur, line)
			cur = len(line)
		default:
			fmt.Printf("[%d]%s\n", cur-3, line[(cur-3):cur])
		}
	}
	return extra
}

func parseAA(cur int, line string) LiquidPrecipitation {
	var rc LiquidPrecipitation

	period, err := strconv.Atoi(line[cur:(cur + 2)])
	if err != nil {
		log.Println("AA")
		log.Fatal(err)
	}
	rc.Period = period

	depth10, err := strconv.ParseFloat(line[(cur+2):(cur+6)], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.Depth = depth10 / 10.0

	rc.Condition = fmt.Sprint(line[(cur + 6):(cur + 7)])
	rc.Quality = fmt.Sprint(line[(cur + 7):(cur + 8)])

	return rc
}

func parseAB(cur int, line string) MonthlyPrecipitation {
	var rc MonthlyPrecipitation

	depth10, err := strconv.ParseFloat(line[cur:(cur+5)], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.Depth = depth10 / 10.0

	rc.Condition = fmt.Sprint(line[(cur + 5):(cur + 6)])
	rc.Quality = fmt.Sprint(line[(cur + 6):(cur + 7)])

	return rc
}

func parseAD(cur int, line string) MonthlyPrecipitationRecord {
	var rc MonthlyPrecipitationRecord

	depth10, err := strconv.ParseFloat(line[cur:(cur+5)], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.Depth = depth10 / 10.0

	rc.Date1 = fmt.Sprint(line[(cur + 5):(cur + 9)])
	rc.Date2 = fmt.Sprint(line[(cur + 9):(cur + 13)])
	rc.Date3 = fmt.Sprint(line[(cur + 13):(cur + 17)])
	rc.Condition = fmt.Sprint(line[(cur + 17):(cur + 18)])
	rc.Quality = fmt.Sprint(line[(cur + 18):(cur + 19)])

	return rc
}

func parseAE(cur int, line string) PrecipitationDays {
	var rc PrecipitationDays

	days, err := strconv.Atoi(line[cur:(cur + 2)])
	if err != nil {
		log.Println("AE")
		log.Fatal(err)
	}
	rc.HundredthDays = days
	rc.HundredthDaysQC = fmt.Sprint(line[(cur + 2):(cur + 3)])

	days, err = strconv.Atoi(line[(cur + 3):(cur + 5)])
	if err != nil {
		log.Println("AE")
		log.Fatal(err)
	}
	rc.TenthDays = days
	rc.TenthDaysQC = fmt.Sprint(line[(cur + 5):(cur + 6)])

	days, err = strconv.Atoi(line[(cur + 6):(cur + 8)])
	if err != nil {
		log.Println("AE")
		log.Fatal(err)
	}
	rc.HalfDays = days
	rc.HalfDaysQC = fmt.Sprint(line[(cur + 8):(cur + 9)])

	days, err = strconv.Atoi(line[(cur + 9):(cur + 11)])
	if err != nil {
		log.Println("AE")
		log.Fatal(err)
	}
	rc.OneDays = days
	rc.OneDaysQC = fmt.Sprint(line[(cur + 11):(cur + 12)])

	return rc
}

func parseAH(cur int, line string) MonthlyPrecipitationShortDuration {
	var rc MonthlyPrecipitationShortDuration

	minutes, err := strconv.Atoi(line[cur:(cur + 3)])
	if err != nil {
		log.Println("AH")
		log.Fatal(err)
	}
	rc.Minutes = minutes

	depth10, err := strconv.ParseFloat(line[(cur+3):(cur+7)], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.Depth = depth10 / 10.0

	rc.Condition = fmt.Sprint(line[(cur + 7):(cur + 8)])
	rc.EndingDate = fmt.Sprint(line[(cur + 8):(cur + 14)])
	rc.Quality = fmt.Sprint(line[(cur + 14):(cur + 15)])

	return rc
}

func parseAI(cur int, line string) MonthlyPrecipitationShortDurationMax {
	var rc MonthlyPrecipitationShortDurationMax

	minutes, err := strconv.Atoi(line[cur:(cur + 3)])
	if err != nil {
		log.Println("AI")
		log.Fatal(err)
	}
	rc.Minutes = minutes

	depth10, err := strconv.ParseFloat(line[(cur+3):(cur+7)], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.Depth = depth10 / 10.0

	rc.Condition = fmt.Sprint(line[(cur + 7):(cur + 8)])
	rc.EndingDate = fmt.Sprint(line[(cur + 8):(cur + 14)])
	rc.Quality = fmt.Sprint(line[(cur + 14):(cur + 15)])

	return rc
}

func parseKB(cur int, line string) AverageAirTemperature {
	var rc AverageAirTemperature

	hours, err := strconv.Atoi(line[cur:(cur + 3)])
	if err != nil {
		log.Println("KB")
		log.Fatal(err)
	}
	rc.Hours = hours

	rc.Code = fmt.Sprint(line[(cur + 3):(cur + 4)])

	temp100, err := strconv.ParseFloat(line[(cur+4):(cur+9)], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.Temperature = temp100 / 100.0

	rc.Quality = fmt.Sprint(line[(cur + 9):(cur + 10)])

	return rc
}

func parseKC(cur int, line string) MonthlyExtremeTemp {
	var rc MonthlyExtremeTemp

	rc.Code = fmt.Sprint(line[cur:(cur + 1)])
	rc.Condition = fmt.Sprint(line[(cur + 1):(cur + 2)])

	temp100, err := strconv.ParseFloat(line[(cur+2):(cur+7)], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.Temperature = temp100 / 100.0

	rc.Date = fmt.Sprint(line[(cur + 7):(cur + 13)])
	rc.Quality = fmt.Sprint(line[(cur + 13):(cur + 14)])

	return rc
}

func parseKD(cur int, line string) HeatingCoolingDays {
	var rc HeatingCoolingDays

	hours, err := strconv.Atoi(line[cur:(cur + 3)])
	if err != nil {
		log.Println("KD")
		log.Fatal(err)
	}
	rc.Hours = hours

	rc.Code = fmt.Sprint(line[(cur + 3):(cur + 4)])

	qty, err := strconv.Atoi(line[(cur + 4):(cur + 8)])
	if err != nil {
		log.Println("KD")
		log.Fatal(err)
	}
	rc.Quantity = qty
	rc.Quality = fmt.Sprint(line[(cur + 13):(cur + 14)])

	return rc
}

func parseKE(cur int, line string) CriteriaDaysTemp {
	var rc CriteriaDaysTemp

	qty, err := strconv.Atoi(line[cur:(cur + 2)])
	if err != nil {
		log.Println("KE")
		log.Fatal(err)
	}
	rc.MaxLow = qty
	rc.MaxLowQC = fmt.Sprint(line[(cur + 2):(cur + 3)])

	qty, err = strconv.Atoi(line[(cur + 3):(cur + 5)])
	if err != nil {
		log.Println("KE")
		log.Fatal(err)
	}
	rc.MaxHigh = qty
	rc.MaxHighQC = fmt.Sprint(line[(cur + 5):(cur + 6)])

	qty, err = strconv.Atoi(line[(cur + 6):(cur + 8)])
	if err != nil {
		log.Println("KE")
		log.Fatal(err)
	}
	rc.MinLow = qty
	rc.MinLowQC = fmt.Sprint(line[(cur + 8):(cur + 9)])

	qty, err = strconv.Atoi(line[(cur + 9):(cur + 11)])
	if err != nil {
		log.Println("KE")
		log.Fatal(err)
	}
	rc.MinReallyLow = qty
	rc.MinReallyLowQC = fmt.Sprint(line[(cur + 11):(cur + 12)])

	return rc
}

func parseGA(cur int, line string) SkyCover {
	var rc SkyCover

	rc.Code = fmt.Sprint(line[cur:(cur + 2)])
	rc.CodeQC = fmt.Sprint(line[(cur + 2):(cur + 3)])

	height, err := strconv.Atoi(line[(cur + 3):(cur + 9)])
	if err != nil {
		log.Println("GA")
		log.Fatal(err)
	}
	rc.BaseHeight = height

	rc.BaseHeightQC = fmt.Sprint(line[(cur + 9):(cur + 10)])
	rc.CloudType = fmt.Sprint(line[(cur + 11):(cur + 12)])
	rc.CloudTypeQC = fmt.Sprint(line[(cur + 12):(cur + 13)])

	return rc
}

func parseGD(cur int, line string) SkyCoverSummation {
	var rc SkyCoverSummation

	rc.Code = fmt.Sprint(line[cur:(cur + 1)])
	rc.Code2 = fmt.Sprint(line[(cur + 1):(cur + 3)])
	rc.CodeQC = fmt.Sprint(line[(cur + 3):(cur + 4)])

	height, err := strconv.Atoi(line[(cur + 4):(cur + 10)])
	if err != nil {
		log.Println("GD")
		log.Fatal(err)
	}
	rc.Height = height

	rc.HeightQC = fmt.Sprint(line[(cur + 10):(cur + 11)])
	rc.Characteristic = fmt.Sprint(line[(cur + 11):(cur + 12)])

	return rc
}

func parseGE(cur int, line string) SkyCondition {
	var rc SkyCondition

	rc.ConvectiveCloud = fmt.Sprint(line[cur:(cur + 1)])
	rc.VerticalDatum = fmt.Sprint(line[(cur + 1):(cur + 7)])

	height, err := strconv.Atoi(line[(cur + 7):(cur + 13)])
	if err != nil {
		log.Println("GE")
		log.Fatal(err)
	}
	rc.BaseHeightUpperRange = height

	height, err = strconv.Atoi(line[(cur + 13):(cur + 19)])
	if err != nil {
		log.Println("GE")
		log.Fatal(err)
	}
	rc.BaseHeightLowerRange = height

	return rc
}

func parseAU(cur int, line string) PresentWeather {
	var rc PresentWeather

	rc.Intensity = fmt.Sprint(line[cur:(cur + 1)])
	rc.Descriptor = fmt.Sprint(line[(cur + 1):(cur + 2)])
	rc.Precipitation = fmt.Sprint(line[(cur + 2):(cur + 4)])
	rc.Obscuration = fmt.Sprint(line[(cur + 4):(cur + 5)])
	rc.OtherPhenomena = fmt.Sprint(line[(cur + 5):(cur + 6)])
	rc.Combination = fmt.Sprint(line[(cur + 6):(cur + 7)])
	rc.Quality = fmt.Sprint(line[(cur + 7):(cur + 8)])

	return rc
}

func parseGF(cur int, line string) SkyConditionObservation {
	var rc SkyConditionObservation

	rc.Coverage = fmt.Sprint(line[cur:(cur + 2)])
	rc.OpaqueCoverage = fmt.Sprint(line[(cur + 2):(cur + 4)])
	rc.CoverageQC = fmt.Sprint(line[(cur + 4):(cur + 5)])
	rc.LowestCloud = fmt.Sprint(line[(cur + 5):(cur + 7)])
	rc.LowestCloudQC = fmt.Sprint(line[(cur + 7):(cur + 8)])
	rc.LowCloudGenus = fmt.Sprint(line[(cur + 8):(cur + 10)])
	rc.LowCloudGenusQC = fmt.Sprint(line[(cur + 10):(cur + 11)])

	height, err := strconv.Atoi(line[(cur + 11):(cur + 16)])
	if err != nil {
		log.Println("GF")
		log.Fatal(err)
	}
	rc.LowCloudBaseHeight = height

	rc.LowCloudBaseHeightQC = fmt.Sprint(line[(cur + 16):(cur + 17)])
	rc.MidCloudGenus = fmt.Sprint(line[(cur + 17):(cur + 19)])
	rc.MidCloudGenusQC = fmt.Sprint(line[(cur + 19):(cur + 20)])
	rc.HighCloudGenus = fmt.Sprint(line[(cur + 20):(cur + 22)])
	rc.HighCloudGenusQC = fmt.Sprint(line[(cur + 22):(cur + 23)])

	return rc
}

func parseMA(cur int, line string) PressureObservation {
	var rc PressureObservation

	alt10, err := strconv.ParseFloat(line[cur:(cur+5)], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.AltimiterSetting = convertHectoPascalToMMHG(alt10 / 10.0)

	rc.AltimiterQC = fmt.Sprint(line[(cur + 5):(cur + 6)])

	stat10, err := strconv.ParseFloat(line[(cur+6):(cur+11)], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.StationPressure = convertHectoPascalToMMHG(stat10 / 10.0)

	rc.StationPressureQC = fmt.Sprint(line[(cur + 11):(cur + 12)])

	return rc
}

func parseKA(cur int, line string) ExtremeTemp {
	var rc ExtremeTemp

	hrs10, err := strconv.ParseFloat(line[cur:(cur+3)], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.Hours = hrs10 / 10.0

	rc.Code = fmt.Sprint(line[(cur + 3):(cur + 4)])

	temp10, err := strconv.ParseFloat(line[(cur+4):(cur+9)], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.Temperature = temp10 / 10.0

	rc.TemperatureQC = fmt.Sprint(line[(cur + 9):(cur + 10)])

	return rc
}

func parseMD(cur int, line string) PressureChange {
	var rc PressureChange

	rc.Tendency = fmt.Sprint(line[cur:(cur + 1)])
	rc.TendencyQC = fmt.Sprint(line[(cur + 1):(cur + 2)])

	three10, err := strconv.ParseFloat(line[(cur+2):(cur+5)], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.ThreeHour = three10 / 10.0

	rc.ThreeHourQC = fmt.Sprint(line[(cur + 5):(cur + 6)])

	twoFour10, err := strconv.ParseFloat(line[(cur+6):(cur+10)], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.TwentyFourHour = twoFour10 / 10.0

	rc.TwentyFourHourQC = fmt.Sprint(line[(cur + 10):(cur + 11)])

	return rc
}

func parseOC(cur int, line string) WindGust {
	var rc WindGust

	gust10, err := strconv.ParseFloat(line[cur:(cur+4)], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.Speed = gust10 / 10.0

	rc.SpeedQC = fmt.Sprint(line[(cur + 4):(cur + 5)])

	return rc
}

func parseOD(cur int, line string) WindObservation {
	var rc WindObservation

	rc.ObservationType = fmt.Sprint(line[cur:(cur + 1)])
	hrs, err := strconv.Atoi(line[(cur + 1):(cur + 3)])
	if err != nil {
		log.Println("OD")
		log.Fatal(err)
	}
	rc.Hours = hrs

	speed10, err := strconv.ParseFloat(line[(cur+3):(cur+7)], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.Speed = speed10 / 10.0

	rc.SpeedQC = fmt.Sprint(line[(cur + 7):(cur + 8)])

	dir, err := strconv.Atoi(line[(cur + 8):(cur + 11)])
	if err != nil {
		log.Println("OD")
		log.Fatal(err)
	}
	rc.Direction = dir

	return rc
}

func parseAJ(cur int, line string) SnowDepth {
	var rc SnowDepth

	depth, err := strconv.Atoi(line[cur:(cur + 4)])
	if err != nil {
		log.Println("AJ")
		log.Fatal(err)
	}
	rc.Depth = depth

	rc.DepthCode = fmt.Sprint(line[(cur + 4):(cur + 5)])
	rc.DepthQC = fmt.Sprint(line[(cur + 5):(cur + 6)])

	equiv10, err := strconv.ParseFloat(line[(cur+6):(cur+12)], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.EquivalentWaterDepth = equiv10 / 10.0

	rc.EquivalentWaterDepthCode = fmt.Sprint(line[(cur + 12):(cur + 13)])
	rc.EquivalentWaterDepthQC = fmt.Sprint(line[(cur + 13):(cur + 14)])

	return rc
}

func parseAL(cur int, line string) SnowAccumulation {
	var rc SnowAccumulation

	hrs, err := strconv.Atoi(line[cur:(cur + 2)])
	if err != nil {
		log.Println("AL")
		log.Fatal(err)
	}
	rc.Hours = hrs

	depth, err := strconv.Atoi(line[(cur + 2):(cur + 5)])
	if err != nil {
		log.Println("AL")
		log.Fatal(err)
	}
	rc.Depth = depth

	rc.Condition = fmt.Sprint(line[(cur + 5):(cur + 6)])
	rc.Quality = fmt.Sprint(line[(cur + 6):(cur + 7)])

	return rc
}

func parseAN(cur int, line string) PeriodSnowAccumulation {
	var rc PeriodSnowAccumulation

	hrs, err := strconv.Atoi(line[cur:(cur + 3)])
	if err != nil {
		log.Println("AN")
		log.Fatal(err)
	}
	rc.Hours = hrs

	depth10, err := strconv.ParseFloat(line[(cur+3):(cur+7)], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.Depth = depth10 / 10.0

	rc.Condition = fmt.Sprint(line[(cur + 7):(cur + 8)])
	rc.Quality = fmt.Sprint(line[(cur + 8):(cur + 9)])

	return rc
}

func parseMG(cur int, line string) StationPressureDay {
	var rc StationPressureDay

	sp10, err := strconv.ParseFloat(line[cur:(cur+5)], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.AvgStationPressure = sp10 / 10.0

	rc.AvgStationPressureQC = fmt.Sprint(line[(cur + 5):(cur + 6)])

	min10, err := strconv.ParseFloat(line[(cur+6):(cur+11)], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.MinSeaLevelPressure = min10 / 10.0

	rc.MinSeaLevelPressureQC = fmt.Sprint(line[(cur + 11):(cur + 12)])

	return rc
}

func parseOE(cur int, line string) WindSummaryDay {
	var rc WindSummaryDay

	rc.Code = fmt.Sprint(line[cur:(cur + 1)])

	hrs, err := strconv.Atoi(line[(cur + 1):(cur + 3)])
	if err != nil {
		log.Println("OE")
		log.Fatal(err)
	}
	rc.Hours = hrs

	speed, err := strconv.ParseFloat(line[(cur+3):(cur+8)], 64)
	if err != nil {
		log.Fatal(err)
	}
	rc.Speed = speed / 100.0

	dir, err := strconv.Atoi(line[(cur + 8):(cur + 11)])
	if err != nil {
		log.Println("OE")
		log.Fatal(err)
	}
	rc.Direction = dir

	rc.UTC = fmt.Sprint(line[(cur + 11):(cur + 15)])
	rc.Quality = fmt.Sprint(line[(cur + 15):(cur + 16)])

	return rc
}

func parseREM(cur int, line string) Remark {
	var rc Remark

	rc.Type = fmt.Sprint(line[cur:(cur + 3)])

	length, err := strconv.Atoi(line[(cur + 3):(cur + 6)])
	if err != nil {
		log.Println("REM")
		log.Fatal(err)
	}
	rc.Length = length

	rc.Remark = fmt.Sprint(line[(cur + 6):])

	return rc
}
