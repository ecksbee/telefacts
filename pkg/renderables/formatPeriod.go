package renderables

import (
	"strings"
	"time"
)

func formatPeriod(p PGrid, d DGrid, c CGrid, langs []Lang) (PGrid, DGrid, CGrid) {
	p.PeriodHeaders = formatRelevantPeriod(p.PeriodHeaders, langs)
	for _, rootDomain := range d.RootDomains {
		rootDomain.PeriodHeaders = formatRelevantPeriod(rootDomain.PeriodHeaders, langs)
	}
	for _, summationItem := range c.SummationItems {
		summationItem.PeriodHeaders = formatRelevantPeriod(summationItem.PeriodHeaders, langs)
	}

	return p, d, c
}

func formatRelevantPeriod(periodHeaders PeriodHeaders, langs []Lang) PeriodHeaders {
	ret := make(PeriodHeaders, len(periodHeaders))
	for i, ctxPtr := range periodHeaders {
		ctx := *ctxPtr
		if _, foundPure := ctx[PureLabel]; !foundPure {
			ret[i] = &ctx
			continue
		}
		if pureData, foundPure := ctx[PureLabel]; foundPure {
			for _, lang := range langs {
				if lang == PureLabel {
					continue
				}
				ctx[lang] = formatDate(lang, pureData)
			}
		}
		ret[i] = &ctx
	}
	return periodHeaders
}

const iso8601 = "2006-01-02"

func defaultDate(start string, end string) string {
	isDuration := start != ""
	if isDuration {
		return start + "/" + end
	}
	return end
}

func formatDate(lang Lang, date string) string {
	i := strings.IndexRune(date, '/')
	var start, end string
	if i < 0 {
		end = date
	} else {
		start = date[:i]
		end = date[i+1:]
	}
	switch lang {
	case PureLabel:
		return date
	case English:
		return formatEnglishDates(start, end)
	case Español:
		return formatSpanishDates(start, end)
	default:
		return date
	}
}

func monthCount(start time.Time, end time.Time) int {
	//https://yourbasic.org/golang/days-between-dates/
	days := end.Sub(start).Hours() / 24
	//per accounting practice, there are 30 days in a month
	if days < 30 {
		return 0
	}
	//https://flaviocopes.com/golang-count-months-since-date/
	months := 1
	month := start.Month()
	for start.Before(end) {
		start = start.Add(time.Hour * 24)
		nextMonth := start.Month()
		if nextMonth != month {
			months++
		}
		month = nextMonth
	}
	return months
}
