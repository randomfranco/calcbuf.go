package main

// create a map of possible buffer series on RM
func (cb CalcBuf) RMBufferSeries(oneRepMax float64, buffer int, repRange []int) map[string]interface{} {
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


