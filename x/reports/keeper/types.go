package keeper

import "github.com/desmos-labs/desmos/x/reports/types"

func NewWrappedReports(reports []types.Report) WrappedReports {
	return WrappedReports{Reports: reports}
}

func (wrapped WrappedReports) AppendIfMissing(report types.Report) (reports WrappedReports, appended bool) {
	for _, existing := range wrapped.Reports {
		if existing.Equal(report) {
			return wrapped, false
		}
	}
	return NewWrappedReports(append(wrapped.Reports, report)), true
}
