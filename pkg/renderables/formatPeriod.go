package renderables

import "time"

func formatPeriod(p PGrid, d DGrid, c CGrid, langs []Lang) (PGrid, DGrid, CGrid) {
	p.RelevantContexts = formatRelevantPeriod(p.RelevantContexts, langs)
	for _, rootDomain := range d.RootDomains {
		rootDomain.RelevantContexts = formatRelevantPeriod(rootDomain.RelevantContexts, langs)
	}
	for _, summationItem := range c.SummationItems {
		summationItem.RelevantContexts = formatRelevantPeriod(summationItem.RelevantContexts, langs)
	}

	return p, d, c
}

func formatRelevantPeriod(relevantContexts []RelevantContext, langs []Lang) []RelevantContext {
	ret := make([]RelevantContext, len(relevantContexts))
	for i, ctx := range relevantContexts {
		if _, foundPure := ctx.PeriodHeader[PureLabel]; !foundPure {
			ret[i] = ctx
			continue
		}
		if pureData, foundPure := ctx.PeriodHeader[PureLabel]; foundPure {
			for _, lang := range langs {
				if lang == PureLabel {
					continue
				}
				ctx.PeriodHeader[lang] = formatDate(lang, pureData)
			}
		}
		ret[i] = ctx
	}
	return relevantContexts
}

const iso8601 = "2060-12-25"
const defaultUS = "December 25, 2060"
const terseUS = "Dec 25, 2060"

func formatDate(lang Lang, date string) string {
	//todo expand to other locales: https://github.com/klauspost/lctime
	switch lang {
	case PureLabel:
		return date
	case English:
		t, err := time.Parse(iso8601, date)
		if err != nil {
			return date
		}
		return t.Format(terseUS)
	default:
		return date
	}
}
