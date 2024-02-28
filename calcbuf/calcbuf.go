package main

import "math"

const epleyConst = .0333333
const roundWeightFactor = 2.5 // kgs

func roundWeight(carico float64) float64 {
	return math.Round(carico/roundWeightFactor) * roundWeightFactor
	}

func perc(full, peso float64) float64 {
	return math.Round(peso/full * 100)
	}


type CalcBuf struct {
	OneRepMax	float64
	Buffer		int
	}

func OneRepFurther(carico float64) float64 {
	return carico * epleyConst
	}

func EpleyOneRm(rep int, carico float64) float64 {
	return ((epleyConst * float64(rep)) + 1) * carico
	}

func EpleyReverseOneRm(rep int, carico float64) float64 {
	return carico / ((epleyConst * float64(rep)) + 1)
	}

func (cb *CalcBuf) CalculateBufferSeries(oneRepMax float64, buffer int, repRange []int) map[string]interface{} {
	if len(repRange) == 0 {
	repRange = []int{2,3,4,5,6,7,8,9,10,11,12}
	}

	data := make(map[string]interface{})
	data["buffer"] = buffer
	data["1rm"] = oneRepMax
	data["rep_in_kg"] = roundWeight(OneRepFurther(oneRepMax))
	bufferSeries := make([]map[string]interface{}, 0)

	for _, reps := range repRange {
		bufferSerie := make(map[string]interface{})
		bufferSerie["rep"] = reps
		bufferSerie["peso"] = roundWeight(EpleyReverseOneRm(reps+buffer, oneRepMax))
		bufferSerie["perc"] = perc(oneRepMax, bufferSerie["peso"].(float64))
		bufferSeries = append(bufferSeries, bufferSerie)
		}

	data["buffer_series"] = bufferSeries

	return data
	}


