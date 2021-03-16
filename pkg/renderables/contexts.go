package renderables

import (
	"sort"

	"ecksbee.com/telefacts/pkg/hydratables"
)

func getContext(instance *hydratables.Instance, contextRef string) *hydratables.Context {
	for _, context := range instance.Contexts {
		if context.ID == contextRef {
			return &context
		}
	}
	return nil
}

type RelevantContext struct {
	ContextRef   string
	PeriodHeader LanguagePack
	Dimensions   []ContextualDimension
}

type ContextualDimension struct {
	Element    string
	IsExplicit bool
	Dimension  ContextConcept
	Member     ContextConcept
}

type ContextConcept struct {
	Href  string
	Label LabelPack
}

func sortContexts(relevantContexts []RelevantContext) {
	sort.SliceStable(relevantContexts, func(i int, j int) bool {
		if relevantContexts[i].PeriodHeader[PureLabel] == relevantContexts[j].PeriodHeader[PureLabel] {
			if len(relevantContexts[i].Dimensions) == len(relevantContexts[j].Dimensions) {
				for c := 0; c < len(relevantContexts[i].Dimensions); c++ {
					a := relevantContexts[i].Dimensions[c]
					b := relevantContexts[j].Dimensions[c]
					if a.Element == b.Element {
						if a.IsExplicit == b.IsExplicit {
							return a.Member.Href < b.Member.Href
						}
						return !a.IsExplicit
					}
					return a.Element < b.Element
				}
			} else {
				return len(relevantContexts[i].Dimensions) < len(relevantContexts[j].Dimensions)
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
			if occured[arr[e]] != true {
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
	hrefs []string) ([]RelevantContext, int, []LabelPack) {
	factuaHrefs := make([]string, 0, len(hrefs))
	for _, href := range hrefs {
		_, c, err := h.HashQuery(href)
		if err == nil && !c.Abstract {
			factuaHrefs = append(factuaHrefs, href)
		}
	}
	if len(factuaHrefs) <= 0 {
		return []RelevantContext{}, 0, []LabelPack{}
	}
	maxDepth := 0
	ret := make([]RelevantContext, 0, len(hrefs)*4)
	labelPacks := make([]LabelPack, 0, len(hrefs)*4)
	for _, instance := range h.Instances {
		contextRefTaken := make(map[string]bool)
		for _, factualEdgeHref := range factuaHrefs {
			for _, fact := range instance.Facts {
				if _, taken := contextRefTaken[fact.ContextRef]; taken {
					continue
				}
				if factualEdgeHref == fact.Href {
					var context *hydratables.Context
					context = getContext(&instance, fact.ContextRef)
					entity := context.Entity
					contextualSchemedEntity := entity.Identifier.Scheme + "/" + entity.Identifier.CharData
					if contextualSchemedEntity == schemedEntity {
						contextualDimensions, contextualLabelPacks := getContextualDimensions(context, h)
						labelPacks = append(labelPacks, contextualLabelPacks...)
						newItem := RelevantContext{
							ContextRef:   context.ID,
							PeriodHeader: periodString(context),
							Dimensions:   contextualDimensions,
						}
						if len(newItem.Dimensions) > maxDepth {
							maxDepth = len(newItem.Dimensions)
						}
						ret = append(ret, newItem)
						contextRefTaken[fact.ContextRef] = true
					}
				}
			}
		}
	}
	sortContexts(ret)
	return ret, maxDepth, labelPacks
}

func getContextualDimensions(context *hydratables.Context, h *hydratables.Hydratable) ([]ContextualDimension, []LabelPack) {
	ret := make([]ContextualDimension, 0)
	labelPacks := make([]LabelPack, 0)
	if len(context.Entity.Segment.ExplicitMembers) > 0 {
		for _, explicitMember := range context.Entity.Segment.ExplicitMembers {
			member := explicitMember.Member.Href
			memberLabel := GetLabel(h, member)
			labelPacks = append(labelPacks, memberLabel)
			dimension := explicitMember.Dimension.Href
			dimensionLabel := GetLabel(h, dimension)
			labelPacks = append(labelPacks, dimensionLabel)
			ret = append(ret, ContextualDimension{
				Element:    "segment",
				IsExplicit: true,
				Dimension: ContextConcept{
					Href:  dimension,
					Label: dimensionLabel,
				},
				Member: ContextConcept{
					Href:  member,
					Label: memberLabel,
				},
			})
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
			ret = append(ret, ContextualDimension{
				Element:    "scenario",
				IsExplicit: true,
				Dimension: ContextConcept{
					Href:  dimension,
					Label: dimensionLabel,
				},
				Member: ContextConcept{
					Href:  member,
					Label: memberLabel,
				},
			})
		}
	}
	//todo typedMembers
	return ret, labelPacks
}
