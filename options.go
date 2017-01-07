package main

import "fmt"

type BenchOptions struct {
	PrePrimeCount int // number of records in the datastore before the test
	RecordSize    int // size of one record
}

func (opt BenchOptions) TestDesc() string {
	return fmt.Sprintf("preprime=%d-size=%d", opt.PrePrimeCount, opt.RecordSize)
}

var DefaultBenchOpts = OptionsRange2pow(BenchOptions{1, 10 << 10}, BenchOptions{1 << 20, 10 << 10}, 11)

func OptionsRange2pow(start, end BenchOptions, countPerAxis int) []BenchOptions {
	res := []BenchOptions{start}

	if start == end {
		return res
	}

	axis := make([]float64, countPerAxis)
	maxN := 1 << uint(countPerAxis)
	for i := 0; i < countPerAxis-1; i++ {
		axis[i] = float64(uint(1)<<uint(i+1)) / float64(maxN)
	}
	axis[countPerAxis-1] = 1

	if start.PrePrimeCount != end.PrePrimeCount {
		bRes := res[:]
		res = make([]BenchOptions, 0, countPerAxis*len(bRes))
		for _, opt := range bRes {
			for _, scale := range axis {
				opt.PrePrimeCount = int(float64(end.PrePrimeCount-start.PrePrimeCount)*scale) + start.PrePrimeCount
				res = append(res, opt)
			}
		}

	}
	if start.RecordSize != end.RecordSize {
		bRes := res[:]
		res = make([]BenchOptions, 0, countPerAxis*len(bRes))
		for _, opt := range bRes {
			for _, scale := range axis {
				opt.RecordSize = int(float64(end.RecordSize-start.RecordSize)*scale) + start.RecordSize
				res = append(res, opt)
			}
		}

	}

	return res
}
