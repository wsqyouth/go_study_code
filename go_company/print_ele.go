package main

import (
	"fmt"
	"strings"
)

func PrintElement(src []string) {
	for _, ele := range src {
		templateStr := "{\n\"value\": \"" + ele +
			"\",\n\"desc\": \"" + ele + "\",\n\"status\": \"normal\"\n},"
		fmt.Println(templateStr)
	}
}

func PrintEnum(src []string) {
	for _, ele := range src {
		templateStr := "<element name=\"" + ele + "\" type=\"" + ele + "\" require=\"no\" allow_zero=\"yes\" />"
		fmt.Println(templateStr)
	}
}

func PrintComplexType() {
	complexSpecMap := map[string]string{
		"end_cover_spec":          "结束页",
		"mask_cover_spec":         "氛围图",
		"scan_spec":               "扫一扫",
		"double_button_page_spec": "双按钮",
	}

	for key, val := range complexSpecMap {
		templateStr := "<complexType name=\"" + key + "\" extends=\"struct\">\n" +
			"	<attribute name=\"description\" type=\"string\" value=\"" + val + "能力\" />\n" +
			"	<attribute name=\"restraint\" type=\"string\" value=\"" + val + "能力\" />\n" +
			"	<attribute name=\"errormsg\" type=\"string\" value=\"" + val + "能力错误\" />\n</complexType>"
		fmt.Println(templateStr)
	}
}

// simpleType bool
func PrintSimpleTypeBool() {
	specMap := map[string]string{
		"end_cover_switch":     "结束页",
		"mask_cover_switch":    "氛围图",
		"scan_switch":          "扫一扫",
		"double_button_switch": "双按钮",
	}
	for key, val := range specMap {
		templateStr := "<simpleType name=\"" + key + "\" extends=\"boolean\">\n" +
			"	<attribute name=\"description\" type=\"string\" value=\"" + val + "开关\" />\n" +
			"	<attribute name=\"restraint\" type=\"string\" value=\"" + val + "开关\" />\n" +
			"	<attribute name=\"errormsg\" type=\"string\" value=\"" + val + "开关状态不正确\" />\n</simpleType>"
		fmt.Println(templateStr)
	}
}

// simpleType enum
func PrintSimpleTypeEnum() {
	specMap := map[string]string{
		"end_cover_type":  "结束页",
		"mask_cover_type": "氛围图",
	}

	for key, val := range specMap {
		keyNew := "Api" + camelString(key)
		templateStr := "<simpleType name=\"" + key + "\" extends=\"string\">\n" +
			"	<attribute name=\"description\" type=\"string\" value=\"" + val + "\" />\n" +
			"	<attribute name=\"restraint\" type=\"string\" value=\"" + val + "区分类型\" />\n" +
			"	<attribute name=\"errormsg\" type=\"string\" value=\"" + val + "不正确\" />\n" +
			"	<restriction>\n" +
			"	<validate type=\"enum\" source=\"" + keyNew + "\"></validate>" +
			"\n	</restriction>\n</simpleType>"
		fmt.Println(templateStr) //ApiTitlePosition
	}
}

//simpleType string
func PrintSimpleTypeString() {
	specMap := map[string]string{
		"end_cover_image_url": "结束页图片url",
		"end_cover_desc":      "结束页描述",
		"end_cover_title":     "结束页标题",

		"mask_cover_ambient_video_url": "视频氛围图url",
		"mask_cover_ambient_end_url":   "结束页氛围图url",
		"scan_bg_image":                "扫一扫背景图url",
		"scan_desc":                    "扫一扫描述",
		"scan_desc_icon":               "扫一扫特定icon",
		"scan_detect_succ_icon":        "扫一扫成功显示icon",
		"left_button_name":             "左按钮文案",
		"right_button_name":            "右按钮文案",
		"left_button_page_id":          "左按钮page_id",
		"right_button_page_id":         "右按钮page_id",
	}

	for key, val := range specMap {
		templateStr := "<simpleType name=\"" + key + "\" extends=\"string\">\n" +
			"	<attribute name=\"max_length\" type=\"integer\" value=\"200\" />\n" +
			"	<attribute name=\"min_length\" type=\"integer\" value=\"1\" />\n" +
			"	<attribute name=\"description\" type=\"string\" value=\"" + val + "内容\" />\n" +
			"	<attribute name=\"restraint\" type=\"string\" value=\"" + val + "内容\" />\n" +
			"	<attribute name=\"errormsg\" type=\"string\" value=\"" + val + "内容不正确\" />\n" +
			"	<restriction>\n" +
			"	<validate type=\"string\">\n" + "		<maxLength value=\"@max_length\" />\n" +
			"		<minLength value=\"@min_length\" />\n" +
			"	</validate>" +
			"\n	</restriction>\n</simpleType>"
		fmt.Println(templateStr)
	}
}

func main() {
	srcStr := []string{
		"end_cover_switch",
		"end_cover_type",
		"end_cover_image_url",
		"end_cover_desc",
		"end_cover_title",
		"end_cover_action_title",
		"mask_cover_switch",
		"mask_cover_type",
		"mask_cover_ambient_video_url",
		"mask_cover_ambient_end_url",
		"scan_switch",
		"scan_bg_image",
		"scan_desc",
		"scan_desc_icon",
		"scan_detect_succ_icon",
		"double_button_switch",
		"left_button_name",
		"right_button_name",
		"left_button_page_id",
		"right_button_page_id",
		"end_cover_spec",
		"mask_cover_spec",
		"scan_spec",
		"double_button_page_spec",
	}

	//PrintElement(srcStr)
	PrintEnum(srcStr)
	PrintComplexType()
	PrintSimpleTypeBool()
	PrintSimpleTypeEnum()
	PrintSimpleTypeString()
}

/**
 * 驼峰转蛇形 snake string
 * @description XxYy to xx_yy , XxYY to xx_y_y
 * @date 2020/7/30
 * @param s 需要转换的字符串
 * @return string
 **/
func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

/**
 * 蛇形转驼峰
 * @description xx_yy to XxYx  xx_y_y to XxYY
 * @date 2020/7/30
 * @param s要转换的字符串
 * @return string
 **/
func camelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}
