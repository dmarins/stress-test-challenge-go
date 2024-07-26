package entity

type Result struct {
	TotalTime          float64
	TotalRequests      int
	SuccessfulRequests int
	StatusCodeDist     map[int]int
}
