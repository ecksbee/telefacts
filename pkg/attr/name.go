package attr

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/beevik/etree"
)

const xmlns = "xmlns"

type space string

type prefix string

type NameProvider struct {
	originPrefixes map[prefix]space
	targetPrefixes map[prefix]prefix
}

func NewNameProvider(attrs []xmlquery.Attr) (*NameProvider, error) {
	ret := NameProvider{}
	originPrefixes, targetPrefixes := processPrefixes(attrs)
	ret.originPrefixes = originPrefixes
	ret.targetPrefixes = targetPrefixes
	return &ret, nil
}

func (np *NameProvider) NsAttrs() []etree.Attr {
	ret := make([]etree.Attr, 0)
	for originPrefix, targetPrefix := range np.targetPrefixes {
		a := etree.Attr{}
		if targetPrefix == "" {
			a.Space = ""
			a.Key = xmlns
		} else {
			a.Space = xmlns
			a.Key = string(targetPrefix)
		}
		targetNs := np.originPrefixes[originPrefix]
		a.Value = string(targetNs)
		ret = append(ret, a)
	}
	return ret
}

func (np *NameProvider) ProvideConceptName(prefixed string) string {
	i := strings.IndexRune(prefixed, ':')
	if i < 0 {
		return prefixed
	}
	prefix := prefix(prefixed[:i])
	if _, hit1 := np.originPrefixes[prefix]; hit1 {
		if targetPrefix, hit2 := np.targetPrefixes[prefix]; hit2 {
			return string(targetPrefix) + prefixed[i:]
		}
	}
	panic("prefix, " + prefix + ", does not match a namespace")
}

func (np *NameProvider) ProvideXmlName(prefixed string) (*xml.Name, error) {
	i := strings.IndexRune(prefixed, ':')
	if i < 0 {
		return &xml.Name{
			Space: "",
			Local: prefixed[i:],
		}, nil
	}
	prefix := prefix(prefixed[:i])
	if space, hit1 := np.originPrefixes[prefix]; hit1 {
		return &xml.Name{
			Space: string(space),
			Local: prefixed[i+1:],
		}, nil
	}
	return nil, fmt.Errorf("prefix, %s, does not match a namespace", prefix)
}

func (np *NameProvider) ProvideName(ns string, local string) string {
	var resultPrefix prefix
	found := false
	for originPrefix, iSpace := range np.originPrefixes {
		if iSpace == space(ns) {
			resultPrefix = np.targetPrefixes[originPrefix]
			found = true
			break
		}
	}
	if resultPrefix != "" && found {
		return string(resultPrefix) + ":" + local
	} else if resultPrefix == "" {
		return local
	}
	panic("not a valid name space, " + ns)
}

func (np *NameProvider) ProvideOutputXml(node *xmlquery.Node, self bool) string {
	var recur func(recurringNode *xmlquery.Node, nosiblings bool) string
	{
	}
	recur = func(recurringNode *xmlquery.Node, nosiblings bool) string {
		if recurringNode.FirstChild == nil {
			return recurringNode.InnerText()
		} else {
			acc := ""
			temp := etree.NewDocument()
			targetName := np.ProvideName(recurringNode.NamespaceURI, recurringNode.Data)
			attrs := " "
			for _, myAttr := range recurringNode.Attr {
				targetNameAttr := np.ProvideName(myAttr.NamespaceURI, myAttr.Name.Local)
				attrs = attrs + " " + targetNameAttr + "=" + "\"" + strings.ReplaceAll(myAttr.Value, "\"", "'") + "\""
			}
			err := temp.ReadFromString("<" + targetName + attrs + ">" + recur(recurringNode.FirstChild, false) + "</" + targetName + ">")
			if err != nil {
				panic(err)
			}
			str, err := temp.WriteToString()
			if err != nil {
				panic(err)
			}
			acc += str
			curr := recurringNode.NextSibling
			if !nosiblings {
				for {
					if curr == nil {
						break
					}
					acc += recur(curr, false)
					curr = curr.NextSibling
				}
			}
			return acc
		}
	}
	if self {
		return recur(node, true)
	}
	if node.FirstChild == nil {
		return node.InnerText()
	}
	return recur(node.FirstChild, false)
}

func processPrefixes(attrs []xmlquery.Attr) (map[prefix]space, map[prefix]prefix) {
	nameMap := map[space]prefix{
		space(XBRLI):   prefix(""),
		space(XBRLDI):  prefix("xbrldi"),
		space(XLINK):   prefix("xlink"),
		space(XSI):     prefix("xsi"),
		space(ISO4217): prefix("iso4217"),
		space(LINK):    prefix("link"),
	}
	origin := make(map[prefix]space)
	prefixes := map[prefix]space{
		nameMap[space(XBRLI)]:   space(XBRLI),
		nameMap[space(XBRLDI)]:  space(XBRLDI),
		nameMap[space(XLINK)]:   space(XLINK),
		nameMap[space(XSI)]:     space(XSI),
		nameMap[space(ISO4217)]: space(ISO4217),
		nameMap[space(LINK)]:    space(LINK),
	}
	target := make(map[prefix]prefix)
	for _, attr := range attrs {
		if attr.NamespaceURI != xmlns {
			continue
		}
		o := prefix(attr.Name.Local)
		t, currSpace := targetPrefix(attr, nameMap, prefixes)
		nameMap[space(attr.Value)] = t
		prefixes[o] = currSpace
		origin[o] = currSpace
		target[o] = t
	}
	return origin, target
}

func targetPrefix(curr xmlquery.Attr, acc map[space]prefix, prefixes map[prefix]space) (prefix, space) {
	origin := prefix(curr.Name.Local)
	currSpace := space(curr.Value)
	prev, used := acc[currSpace]
	if !used {
		return newPrefix(origin, prefixes), currSpace
	}
	return prev, currSpace
}

func newPrefix(origin prefix, prefixes map[prefix]space) prefix {
	curr := origin
	acc := 0
	for {
		hit := false
		for p := range prefixes {
			if curr == p {
				hit = true
				break
			}
		}
		if !hit {
			return curr
		}
		curr = prefix(string(origin) + "-" + toChar(acc))
		acc++
		hit = false
	}
}

var arr = [...]string{"a", "b", "c", "d", "e",
	"f", "g", "h", "i", "j", "k", "l", "m", "n",
	"o", "p", "q", "r", "s", "t", "u", "v", "w",
	"x", "y", "z"}

func toChar(i int) string {
	if i < len(arr) {
		return arr[i]
	}
	return arr[len(arr)-1] + toChar(i-len(arr)-1)
}
