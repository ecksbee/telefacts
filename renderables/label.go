package renderables

import (
	"ecksbee.com/telefacts/attr"
	"ecksbee.com/telefacts/hydratables"
)

func GetLabel(h *hydratables.Hydratable, href string) LabelPack {
	ret := LabelPack{}
	ret[Default] = LanguagePack{}
	ret[Default][PureLabel] = href
	ret = appendLabelModifiersFromHref(ret, h, href)
	return ret
}

func appendLabelModifiersFromHref(labelPack LabelPack, h *hydratables.Hydratable, href string) LabelPack {
	for _, labels := range h.LabelLinkbases {
		for _, labelLink := range labels.LabelLink {
			for _, loc := range labelLink.Locs {
				if loc.Href == href {
					labelArcs := labelLink.LabelArcs
					for _, labelArc := range labelArcs {
						if labelArc.From == loc.Label && labelArc.Arcrole == attr.LabelArcrole {
							for _, labelLinkLabel := range labelLink.Labels {
								if labelLinkLabel.Label == labelArc.To {
									charData := labelLinkLabel.CharData
									switch labelLinkLabel.Role {
									case attr.Label:
										labelPack = appendLanguage(&labelLinkLabel, Default, charData, labelPack)
										break
									case attr.TerseLabel:
										labelPack = appendLanguage(&labelLinkLabel, Terse, charData, labelPack)
										break
									case attr.VerboseLabel:
										labelPack = appendLanguage(&labelLinkLabel, Verbose, charData, labelPack)
										break
									default: //noop
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return labelPack
}

func appendLanguage(labelLinkLabel *hydratables.LabelLinkLabel, labelRole LabelRole, charData string, labelPack LabelPack) LabelPack {
	switch labelLinkLabel.Lang {
	case "en-US":
		labelPack[labelRole][English] = charData
		break
	default: //noop
	}
	return labelPack
}
