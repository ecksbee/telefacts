package renderables

import (
	"time"

	"github.com/klauspost/lctime"
)

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

const iso8601 = "2006-01-02"
const terseUS = "Jan 2, 2006"
const terseStrf = "%b %d, %Y"

func formatDate(lang Lang, date string) string {
	//todo duration context
	switch lang {
	case PureLabel:
		return date
	case English:
		t, err := time.Parse(iso8601, date)
		if err != nil {
			return date
		}
		return t.Format(terseUS)
	case Español:
		t, err := time.Parse(iso8601, date)
		if err != nil {
			return date
		}
		txt, err := lctime.StrftimeLoc("es_ES", terseStrf, t)
		if err != nil {
			return date
		}
		return txt
	default:
		return date
	}
}
