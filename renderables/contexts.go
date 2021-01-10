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
			ret = append(ret, explicitMember.Text+"<"+explicitMember.Dimension+"<segment")
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

func getRelevantContexts(schemedEntity string, instance *xbrl.Instance,
	schema *xbrl.Schema, hrefs []string) ([]RelevantContext, int) {
	factuaHrefs := make([]string, 0, len(hrefs))
	for _, href := range hrefs {
		var c *xbrl.Concept
		_, c, _ = xbrl.HashQuery(schema, href) //todo catch errors
		if !c.Abstract {
			factuaHrefs = append(factuaHrefs, href)
		}
	}
	if len(factuaHrefs) <= 0 {
		return []RelevantContext{}, 0
	}

	maxDepth := 0
	contextRefTaken := make(map[string]bool)
	ret := make([]RelevantContext, 0, len(instance.Context)>>1)
	for _, factualEdgeHref := range factuaHrefs {
		for _, fact := range instance.Facts { //todo parallelize nlogn
			if _, taken := contextRefTaken[fact.ContextRef]; taken {
				continue
			}
			factualHref, _, err := xbrl.NameQuery(schema, fact.XMLName.Space, fact.XMLName.Local)
			if err != nil {
				continue
			}
			if factualEdgeHref == factualHref {
				var context *xbrl.Context
				context = getContext(instance, fact.ContextRef)
				if len(context.Entity.Identitifier) <= 0 {
					continue
				}
				contextualSchemedEntity := context.Entity.Identitifier[0].Scheme + "/" + context.Entity.Identitifier[0].Text
				if contextualSchemedEntity != schemedEntity {
					continue
				}
				dommems := domainMembersString(context)
				if len(dommems) > maxDepth {
					maxDepth = len(dommems)
				}
				ret = append(ret, RelevantContext{
					ContextRef:          context.ID,
					PeriodHeader:        periodString(context),
					DomainMemberHeaders: dommems,
				})
				contextRefTaken[fact.ContextRef] = true
			}
		}
	}
	sortContexts(ret)
	return ret, maxDepth
}
