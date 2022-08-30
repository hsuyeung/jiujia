package parseutil

import "strings"

// 解析文本数据工具类

// ExtractCookieAndTkFromHeader 从 header 中提取出 cookie 和 tk
func ExtractCookieAndTkFromHeader(header string) (cookie map[string]string, tk string) {
	if len(header) == 0 {
		return nil, ""
	}
	header = strings.ReplaceAll(header, "cookie: ", "Cookie: ")

	// tk
	startIdx := strings.Index(header, "tk: ")
	header = header[startIdx:]
	endIdx := strings.Index(header, "\n")
	tk = strings.TrimSpace(header[len("tk: "):endIdx])
	header = header[endIdx:]

	// cookie
	startIdx = strings.Index(header, "Cookie: ")
	header = header[startIdx:]
	endIdx = strings.Index(header, "\n")
	cookieStr := strings.TrimSpace(header[len("Cookie: "):endIdx])
	//header = header[endIdx:]
	// 将多个 cookie 拼接而成的字符串解析为单个的 cookie 放入 map 中
	cookies := strings.Split(cookieStr, ";")
	if len(cookies) == 0 {
		return nil, ""
	}
	cookie = make(map[string]string, len(cookies))
	for _, c := range cookies {
		kv := strings.Split(strings.TrimSpace(c), "=")
		if len(kv) != 2 {
			return nil, ""
		}
		cookie[kv[0]] = kv[1]
	}

	return cookie, tk
}
