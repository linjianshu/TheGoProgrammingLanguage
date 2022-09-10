package main

func main() {
	m := make(map[string]struct{})
	if _, ok := m["ljs"]; !ok {
		//记录首次出现
		m["ljs"] = struct{}{}
	}
}
