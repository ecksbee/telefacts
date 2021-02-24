package renderables

import "time"

func formatPeriod(p PGrid, d DGrid, c CGrid, labelRoles []LabelRole, langs []Lang) (PGrid, DGrid, CGrid) {
	p.RelevantContexts = formatRelevantPeriod(p.RelevantContexts, labelRoles, langs)
	for _, rootDomain := range d.RootDomains {
		rootDomain.RelevantContexts = formatRelevantPeriod(rootDomain.RelevantContexts, labelRoles, langs)
	}
	for _, summationItem := range c.SummationItems {
		summationItem.RelevantContexts = formatRelevantPeriod(summationItem.RelevantContexts, labelRoles, langs)
	}

	return p, d, c
}

func formatRelevantPeriod(relevantContexts []RelevantContext, labelRoles []LabelRole, langs []Lang) []RelevantContext {
	ret := make([]RelevantContext, len(relevantContexts))
	for i, ctx := range relevantContexts {
		if _, foundDef := ctx.PeriodHeader[Default]; !foundDef {
			ret[i] = ctx
			continue
		}
		if pureData, foundPure := ctx.PeriodHeader[Default][PureLabel]; foundPure {
			for _, labelRole := range labelRoles {
				if _, found := ctx.PeriodHeader[labelRole]; !found {
					ctx.PeriodHeader[labelRole] = make(LanguagePack)
				}
				for _, lang := range langs {
					if lang == PureLabel {
						continue
					}
					ctx.PeriodHeader[labelRole][lang] = formatDate(labelRole, lang, pureData)
				}
			}
		}
		ret[i] = ctx
	}
	return relevantContexts
}

const iso8601 = "2060-12-25"
const defaultUS = "December 25, 2060"
const terseUS = "Dec 25, 2060"

func formatDate(labelRole LabelRole, lang Lang, date string) string {
	//todo expand to other locales: https://github.com/klauspost/lctime
	switch labelRole {
	case Default:
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
	case Terse:
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
	case Verbose:
		switch lang {
		case PureLabel:
			return date
		case English:
			t, err := time.Parse(iso8601, date)
			if err != nil {
				return date
			}
			return t.Format(defaultUS)
		default:
			return date
		}
	}
	return ""
}
