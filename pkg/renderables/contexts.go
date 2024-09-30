package renderables

import (
	"sort"

	"ecksbee.com/telefacts/internal/graph"
	"ecksbee.com/telefacts/pkg/hydratables"
	myarcs "github.com/joshuanario/arcs"
)

func getContext(instance *hydratables.Instance, contextRef string) *hydratables.Context {
	for _, context := range instance.Contexts {
		if context.ID == contextRef {
			return &context
		}
	}
	return nil
}

type relevantContext struct {
	ContextRef   string
	PeriodHeader LanguagePack
	Members      []RelevantMember
}

type RelevantMember struct {
	Dimension
	ExplicitMember *ExplicitMember
	TypedMember    string
	TypedDomain    *TypedDomain
	IsSegment      bool
}

type ExplicitMember struct {
	Href  string
	Label LabelPack
}

type Dimension struct {
	Href  string
	Label LabelPack
}

type TypedDomain struct {
	Href  string
	Label LabelPack
}

func sortContexts(relevantContexts []relevantContext) {
	sort.SliceStable(relevantContexts, func(i int, j int) bool {
		if relevantContexts[i].PeriodHeader[PureLabel] == relevantContexts[j].PeriodHeader[PureLabel] {
			if len(relevantContexts[i].Members) == len(relevantContexts[j].Members) {
				for c := 0; c < len(relevantContexts[i].Members); c++ {
					a := relevantContexts[i].Members[c]
					b := relevantContexts[j].Members[c]
					if a.IsSegment && !b.IsSegment {
						return true
					}
					if !a.IsSegment && b.IsSegment {
						return false
					}
					if a.TypedDomain != nil && b.TypedDomain == nil {
						return true
					}
					if a.TypedDomain == nil && b.TypedDomain != nil {
						return false
					}
					if a.TypedDomain != nil && b.TypedDomain != nil {
						if a.TypedDomain.Href == b.TypedDomain.Href {
							return a.TypedMember < b.TypedMember
						}
						return a.TypedDomain.Href < b.TypedDomain.Href
					}
					if a.ExplicitMember != nil && b.ExplicitMember != nil {
						return a.ExplicitMember.Href < b.ExplicitMember.Href
					}
					if a.ExplicitMember == nil && b.ExplicitMember != nil {
						return true
					}
					if a.ExplicitMember != nil && b.ExplicitMember == nil {
						return false
					}
					return i < j
				}
			} else {
				return len(relevantContexts[i].Members) < len(relevantContexts[j].Members)
			}
		}
		return relevantContexts[i].PeriodHeader[PureLabel] < relevantContexts[j].PeriodHeader[PureLabel]
	})
}

func periodString(context *hydratables.Context) LanguagePack {
	ret := LanguagePack{}
	if context.Period.Duration.EndDate != "" && context.Period.Duration.StartDate != "" {
		ret[PureLabel] = context.Period.Duration.StartDate + "/" + context.Period.Duration.EndDate
	} else if context.Period.Instant.CharData != "" {
		ret[PureLabel] = context.Period.Instant.CharData
	} else {
		ret[PureLabel] = ""
	}
	return ret
}

func periodMultilingualString(labels LabelPack, context *hydratables.Context) LanguagePack {
	ret := LanguagePack{}
	start := ""
	end := ""
	if context.Period.Duration.EndDate != "" && context.Period.Duration.StartDate != "" {
		start = context.Period.Duration.StartDate
		end = context.Period.Duration.EndDate
		ret[PureLabel] = start + "/" + end
		ret[BriefLabel] = ret[PureLabel]
	} else if context.Period.Instant.CharData != "" {
		start = context.Period.Instant.CharData
		ret[PureLabel] = start + "/" + end
		ret[BriefLabel] = ret[PureLabel]
	} else {
		ret[PureLabel] = ""
		return ret
	}
	ret = appendLanguagePackFromPeriod(ret, labels, start, end)
	return ret
}

func appendLanguagePackFromPeriod(langPack LanguagePack, labels LabelPack, start string, end string) LanguagePack {
	ret := langPack
	for _, item := range labels {
		for lang := range item {
			switch lang {
			case English:
				ret[English] = formatEnglishDates(start, end)
				break
			case Español:
				ret[Español] = formatSpanishDates(start, end)
				break
			default: //noop
			}
		}
	}
	return ret
}

func dedupEntities(h *hydratables.Hydratable) []Entity {
	entities := []Entity{}
	for _, instance := range h.Instances {
		for _, context := range instance.Contexts {
			entity := context.Entity
			scheme := entity.Identifier.Scheme
			entityid := entity.Identifier.CharData
			if len(scheme) > 0 && len(entityid) > 0 {
				entities = append(entities, Entity{
					Scheme:   scheme,
					CharData: entityid,
				})
			}
		}
	}

	uniques := func(arr []Entity) []Entity {
		occured := map[Entity]bool{}
		u := []Entity{}
		for e := range arr {
			if !occured[arr[e]] {
				occured[arr[e]] = true
				u = append(u, arr[e])
			}
		}
		return u
	}(entities)
	return uniques
}

func sortedEntities(h *hydratables.Hydratable) []Entity {
	schemedEntities := dedupEntities(h)
	sort.SliceStable(schemedEntities, func(i, j int) bool {
		if schemedEntities[i].Scheme == schemedEntities[j].Scheme {
			return schemedEntities[i].CharData < schemedEntities[j].CharData
		}
		return schemedEntities[i].Scheme < schemedEntities[j].Scheme
	})
	return schemedEntities
}

func stringify(e *Entity) string {
	if e == nil {
		return ""
	}
	return e.Scheme + "/" + e.CharData
}

func getRelevantContexts(schemedEntity string, h *hydratables.Hydratable,
	hrefs []string) ([]relevantContext, []myarcs.RArc, []myarcs.RArc, []LabelPack) {
	factuaHrefs := make([]string, 0, len(hrefs))
	for _, href := range hrefs {
		_, c, err := h.HashQuery(href)
		if err == nil && c != nil && !c.Abstract {
			factuaHrefs = append(factuaHrefs, href)
		}
	}
	if len(factuaHrefs) <= 0 {
		return []relevantContext{}, []myarcs.RArc{}, []myarcs.RArc{}, []LabelPack{}
	}
	segmentTypedDomainTrees := make([]myarcs.RArc, 0)
	scenarioTypedDomainTrees := make([]myarcs.RArc, 0)
	ret := make([]relevantContext, 0, len(hrefs)*4)
	labelPacks := make([]LabelPack, 0, len(hrefs)*4)
	for _, instance := range h.Instances {
		contextRefTaken := make(map[string]bool)
		for _, factualEdgeHref := range factuaHrefs {
			for _, fact := range instance.Facts {
				if _, taken := contextRefTaken[fact.ContextRef]; taken {
					continue
				}
				if factualEdgeHref == fact.Href {
					context := getContext(&instance, fact.ContextRef)
					entity := context.Entity
					contextualSchemedEntity := entity.Identifier.Scheme + "/" + entity.Identifier.CharData
					if contextualSchemedEntity == schemedEntity {
						contextualMembers, segmentTypedDomainTreesLocal, scenarioTypedDomainTreesLocal,
							contextualLabelPacks := getContextualMembers(context, h)
						labelPacks = append(labelPacks, contextualLabelPacks...)
						newItem := relevantContext{
							ContextRef:   context.ID,
							PeriodHeader: periodString(context),
							Members:      contextualMembers,
						}
						segmentTypedDomainTrees = append(segmentTypedDomainTrees, segmentTypedDomainTreesLocal...)
						scenarioTypedDomainTrees = append(scenarioTypedDomainTrees, scenarioTypedDomainTreesLocal...)
						ret = append(ret, newItem)
						contextRefTaken[fact.ContextRef] = true
					}
				}
			}
		}
	}
	sortContexts(ret)
	return ret, segmentTypedDomainTrees, scenarioTypedDomainTrees, labelPacks
}

func getContextualMembers(context *hydratables.Context,
	h *hydratables.Hydratable) ([]RelevantMember, []myarcs.RArc,
	[]myarcs.RArc, []LabelPack) {
	ret := make([]RelevantMember, 0)
	labelPacks := make([]LabelPack, 0)
	segmentTypedDomainTrees := make([]myarcs.RArc, 0)
	scenarioTypedDomainTrees := make([]myarcs.RArc, 0)
	if len(context.Entity.Segment.ExplicitMembers) > 0 {
		for _, explicitMember := range context.Entity.Segment.ExplicitMembers {
			member := explicitMember.Member.Href
			memberLabel := GetLabel(h, member)
			labelPacks = append(labelPacks, memberLabel)
			dimension := explicitMember.Dimension.Href
			dimensionLabel := GetLabel(h, dimension)
			labelPacks = append(labelPacks, dimensionLabel)
			ret = append(ret, RelevantMember{
				Dimension: Dimension{
					Href:  dimension,
					Label: GetLabel(h, dimension),
				},
				IsSegment: true,
				ExplicitMember: &ExplicitMember{
					Href:  member,
					Label: memberLabel,
				},
			})
		}
	}
	if len(context.Entity.Segment.TypedMembers) > 0 {
		for _, typedMember := range context.Entity.Segment.TypedMembers {
			dimension := typedMember.Dimension.Href
			dimensionLabel := GetLabel(h, dimension)
			labelPacks = append(labelPacks, dimensionLabel)
			typedDomainMembers, typedDomainArcs, typedDomainMemberLabels := getTypedMember(typedMember, Dimension{
				Href:  dimension,
				Label: dimensionLabel,
			}, true, h)
			segmentTypedDomainTrees = append(segmentTypedDomainTrees, *graph.Tree(typedDomainArcs, typedDomainArcRole))
			labelPacks = append(labelPacks, typedDomainMemberLabels...)
			ret = append(ret, typedDomainMembers...)
		}
	}
	if len(context.Scenario.ExplicitMembers) > 0 {
		for _, explicitMember := range context.Scenario.ExplicitMembers {
			member := explicitMember.Member.Href
			memberLabel := GetLabel(h, member)
			labelPacks = append(labelPacks, memberLabel)
			dimension := explicitMember.Dimension.Href
			dimensionLabel := GetLabel(h, dimension)
			labelPacks = append(labelPacks, dimensionLabel)
			ret = append(ret, RelevantMember{
				Dimension: Dimension{
					Href:  dimension,
					Label: GetLabel(h, dimension),
				},
				IsSegment: false,
				ExplicitMember: &ExplicitMember{
					Href:  member,
					Label: memberLabel,
				},
			})
		}
	}
	if len(context.Scenario.TypedMembers) > 0 {
		for _, typedMember := range context.Scenario.TypedMembers {
			dimension := typedMember.Dimension.Href
			dimensionLabel := GetLabel(h, dimension)
			labelPacks = append(labelPacks, dimensionLabel)
			typedDomainMembers, typedDomainArcs, typedDomainMemberLabels := getTypedMember(typedMember, Dimension{
				Href:  dimension,
				Label: dimensionLabel,
			}, false, h)
			scenarioTypedDomainTrees = append(scenarioTypedDomainTrees, *graph.Tree(typedDomainArcs, typedDomainArcRole))
			labelPacks = append(labelPacks, typedDomainMemberLabels...)
			ret = append(ret, typedDomainMembers...)
		}
	}
	return ret, segmentTypedDomainTrees, scenarioTypedDomainTrees, labelPacks
}
