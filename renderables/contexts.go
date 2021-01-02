package renderables

import (
	"sort"

	"ecks-bee.com/telefacts/xbrl"
)

func getContext(instance *xbrl.Instance, contextRef string) *xbrl.Context {
	for _, context := range instance.Context {
		if context.ID == contextRef {
			return &context
		}
	}
	return nil
}

type RelevantContext struct {
	ContextRef          string
	PeriodHeader        string
	DomainMemberHeaders []string
}

func sortContexts(relevantContexts []RelevantContext) {
	sort.SliceStable(relevantContexts, func(i int, j int) bool {
		if relevantContexts[i].PeriodHeader == relevantContexts[j].PeriodHeader {
			if len(relevantContexts[i].DomainMemberHeaders) == len(relevantContexts[j].DomainMemberHeaders) {
				for c := 0; c < len(relevantContexts[i].DomainMemberHeaders); c++ {
					if relevantContexts[i].DomainMemberHeaders[c] == relevantContexts[j].DomainMemberHeaders[c] {
						continue
					}
					return relevantContexts[i].DomainMemberHeaders[c] < relevantContexts[j].DomainMemberHeaders[c]
				}
			} else {
				return len(relevantContexts[i].DomainMemberHeaders) < len(relevantContexts[j].DomainMemberHeaders)
			}
		}
		return relevantContexts[i].PeriodHeader < relevantContexts[j].PeriodHeader
	})
}

func periodString(context *xbrl.Context) string {
	if len(context.Period.EndDate) > 0 {
		return context.Period.StartDate + "/" + context.Period.EndDate
	}
	return context.Period.Instant
}

func domainMembersString(context *xbrl.Context) []string {
	if len(context.Entity.Segment.ExplicitMember) > 0 {
		ret := make([]string, 0, len(context.Entity.Segment.ExplicitMember))
		for _, explicitMember := range context.Entity.Segment.ExplicitMember {
			ret = append(ret, explicitMember.Text+"<"+explicitMember.Dimension+"<segment") //use href through namequery
		}
		sort.SliceStable(ret, func(i int, j int) bool {
			return ret[i] < ret[j]
		})
		return ret
	}
	return []string{}
}

func dedupEntities(instance *xbrl.Instance) []string {
	if len(instance.Context) <= 0 {
		return []string{}
	}
	entities := func(i *xbrl.Instance) []string {
		ret := []string{}
		for _, e := range i.Context {
			if len(e.Entity.Identitifier) <= 0 {
				continue
			}
			ret = append(ret, e.Entity.Identitifier[0].Scheme+"/"+e.Entity.Identitifier[0].Text)
		}
		return ret
	}(instance)
	uniques := dedup(entities)
	return uniques
}

func sortedEntities(instance *xbrl.Instance) []string {
	schemedEntities := dedupEntities(instance)
	sort.Strings(schemedEntities)
	return schemedEntities
}
