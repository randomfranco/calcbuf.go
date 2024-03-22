package main

import "math"

const epleyConst = .0333333
const roundWeightFactor = 2.5 // kgs

type rep_t	uint8

// converted to this for json
type CalcBuf struct {
	OneRepMax	float64		`json:"onerepmax"`
	BufferRep	rep_t		`json:"bufferrep"`
	RepRange	[]rep_t		`json:"reprange"`
}

func roundWeight(carico float64) float64 {
	return math.Round(carico/roundWeightFactor) * roundWeightFactor
	}

func perc(full, peso float64) float64 {
	return math.Round(peso/full * 100)
	}

func OneRepFurther(carico float64) float64 {
	return carico * epleyConst
	}

func EpleyOneRm(rep rep_t, carico float64) float64 {
	return ((epleyConst * float64(rep)) + 1) * carico
	}

func EpleyReverseOneRm(rep rep_t, carico float64) float64 {
	return carico / ((epleyConst * float64(rep)) + 1)
	}
