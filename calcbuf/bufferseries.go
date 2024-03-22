package main

// create a map of possible buffer series on RM
func (cb CalcBuf) RMBufferSeries() map[string]interface{} {
	if len(cb.RepRange) == 0 {
		cb.RepRange = []rep_t{2,3,4,5,6,7,8,9,10,11,12}
	}

	data := make(map[string]interface{})

	data["buffer"] = cb.BufferRep
	data["1rm"] = cb.OneRepMax
	data["rep_in_kg"] = roundWeight(OneRepFurther(cb.OneRepMax))
	bufferSeries := make([]map[string]interface{}, 0)

	for _, reps := range cb.RepRange {
		bufferSerie := make(map[string]interface{})
		bufferSerie["rep"] = reps
		bufferSerie["peso"] = roundWeight(EpleyReverseOneRm(reps + cb.BufferRep, cb.OneRepMax))
		bufferSerie["perc"] = perc(cb.OneRepMax, bufferSerie["peso"].(float64))
		bufferSeries = append(bufferSeries, bufferSerie)
		}

	data["buffer_series"] = bufferSeries

	return data
	}


