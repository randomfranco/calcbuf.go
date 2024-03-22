package main

import "math"

const epleyConst = .0333333
const roundWeightFactor = 2.5 // kgs

type CalcBuf struct {
	OneRepMax	float64
	Buffer		int
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

func EpleyOneRm(rep int, carico float64) float64 {
	return ((epleyConst * float64(rep)) + 1) * carico
	}

func EpleyReverseOneRm(rep int, carico float64) float64 {
	return carico / ((epleyConst * float64(rep)) + 1)
	}
