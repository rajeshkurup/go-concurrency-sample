package reports

/**
 * @brief Container to hold aggregated Success Rates
 */
type SuccessRateDo struct {
	SumOfSuccessRates  *float64 `json:"SumOfSuccessRates,omitempty"`
	SuccessRatecount   *float64 `json:"SuccessRatecount,omitempty"`
	AverageSuccessRate *float64 `json:"AverageSuccessRate,omitempty"`
}

/**
 * @brief Constructor for SuccessRateDo
 * @param successRate Success Rate for an application from a host
 * @return An object of SuccessRateDo
 */
func MakeSuccessRateDo(successRate float64, count float64) SuccessRateDo {
	avgSuccessRate := successRate / count
	return SuccessRateDo{&successRate, &count, &avgSuccessRate}
}

/**
 * @brief A public method to removes metadata to generate final report
 * @return Pointer to current SuccessRateDo object
 */
func (successRateDo *SuccessRateDo) Finalize() *SuccessRateDo {
	successRateDo.SuccessRatecount = nil
	successRateDo.SumOfSuccessRates = nil
	return successRateDo
}

/**
 * @brief A public methodd to aggregate Success Rates
 * @param successRate New Success Rate value to be added to previous aggregated Success Rate
 * @return Instance of new SuccessRateDo object
 */
func (successRateDo *SuccessRateDo) AddSuccessRate(successRate float64) SuccessRateDo {
	newSuccessRateSum := *successRateDo.SumOfSuccessRates + successRate
	newSuccessRateCount := *successRateDo.SuccessRatecount + 1
	return MakeSuccessRateDo(newSuccessRateSum, newSuccessRateCount)
}
