package renderables

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func dedup(arr []string) []string {
	occured := map[string]bool{}
	ret := []string{}
	for e := range arr {
		if occured[arr[e]] != true {
			occured[arr[e]] = true
			ret = append(ret, arr[e])
		}
	}
	return ret
}
