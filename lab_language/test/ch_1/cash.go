// https://www.jianshu.com/p/595eabe003c9
package ch_1

var Cash = make(map[string]string)

func Add(key, value string) {
	if _, ok := Cash[key]; !ok {
		Cash[key] = value
	}
}

func Delete(key string) {
	if _, ok := Cash[key]; ok {
		delete(Cash, key)
	}
}

func Update(key, value string) {
	Cash[key] = value
}

func Get(key string) string {
	if v, ok := Cash[key]; ok {
		return v
	}
	return ""
}

func Clean() {
	Cash = make(map[string]string)
}
