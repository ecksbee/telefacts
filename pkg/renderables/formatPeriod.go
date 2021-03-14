package renderables

import (
	"strconv"
	"strings"
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

func formatEnglishDates(start string, end string) string {
	isDuration := start != ""
	endTime, err := time.Parse(iso8601, end)
	if err != nil {
		return defaultDate(start, end)
	}
	const strfUS = "January 2, 2006"
	date := endTime.Format(strfUS)
	if isDuration {
		startTime, err := time.Parse(iso8601, start)
		if err != nil {
			return defaultDate(start, end)
		}
		months := monthCount(startTime, endTime)
		if months < 1 {
			return defaultDate(start, end)
		} else if months == 1 {
			return "1 month ended " + date
		} else {
			m := strconv.Itoa(months)
			return m + " months ended " + date
		}
	}
	return "as of " + date
}

func formatSpanishDates(start string, end string) string {
	isDuration := start != ""
	endTime, err := time.Parse(iso8601, end)
	if err != nil {
		return defaultDate(start, end)
	}
	date, err := lctime.StrftimeLoc("es_ES", "%d %B %Y", endTime)
	if err != nil {
		return defaultDate(start, end)
	}
	if isDuration {
		startTime, err := time.Parse(iso8601, start)
		if err != nil {
			return defaultDate(start, end)
		}
		months := monthCount(startTime, endTime)
		if months < 1 {
			return defaultDate(start, end)
		} else if months == 1 {
			return "un mes terminado al " + date
		} else {
			m := strconv.Itoa(months)
			return m + " meses terminados al " + date
		}
	}
	return "al " + date
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
