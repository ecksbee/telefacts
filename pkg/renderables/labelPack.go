package renderables

type LabelRole string
type Lang string

const Default = LabelRole("Default")
const Terse = LabelRole("Terse")
const Verbose = LabelRole("Verbose")

type LabelPack map[LabelRole]LanguagePack

const PureLabel = Lang("Unlabelled")
const BriefLabel = Lang("Truncated")
const English = Lang("en - english")
const Deutsch = Lang("de - deutsch")
const Français = Lang("fr - français")
const Español = Lang("es - español")
const Hindi = Lang("hi - हिन्दी")

type LanguagePack map[Lang]string

func getPureLabel(labelPack LabelPack) string {
	return getLabel(labelPack, Default, PureLabel)
}

func getLabel(labelPack LabelPack, labelRole LabelRole, lang Lang) string {
	return labelPack[labelRole][lang]
}

func reduce(labelPacks []LabelPack) *LabelPack {
	if len(labelPacks) <= 0 {
		return nil
	}
	if len(labelPacks) == 1 {
		return &labelPacks[0]
	}
	ret := make(LabelPack)
	for i := 0; i < len(labelPacks); i++ {
		item := labelPacks[i]
		for labelRole, langPack := range item {
			ret[labelRole] = make(LanguagePack)
			for lang, chardata := range langPack {
				ret[labelRole][lang] = chardata
			}
		}
	}
	return &ret
}

func destruct(labelPack LabelPack) ([]LabelRole, []Lang) {
	labelRoles := make([]LabelRole, 0, 20)
	langs := make([]Lang, 0, 8)
	for labelRole, langPack := range labelPack {
		labelRoles = append(labelRoles, labelRole)
		for lang := range langPack {
			langs = append(langs, lang)
		}
	}
	return labelRoles, langs
}

func dedupLang(langs []Lang) []Lang {
	arr := make([]string, 0, len(langs))
	for _, lang := range langs {
		arr = append(arr, string(lang))
	}
	uniqs := dedup(arr)
	ret := make([]Lang, 0, len(uniqs))
	for _, uniq := range uniqs {
		ret = append(ret, Lang(uniq))
	}
	return ret
}

func dedupLabelRole(labelRoles []LabelRole) []LabelRole {
	arr := make([]string, 0, len(labelRoles))
	for _, labelRole := range labelRoles {
		arr = append(arr, string(labelRole))
	}
	uniqs := dedup(arr)
	ret := make([]LabelRole, 0, len(uniqs))
	for _, uniq := range uniqs {
		ret = append(ret, LabelRole(uniq))
	}
	return ret
}
