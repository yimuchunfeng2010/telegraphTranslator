package config

type MessageRule struct {
	MessageType string
	Rule        []int // 编组编码
}

var Config = struct {
	MessageRuleList []MessageRule
}{
	MessageRuleList: []MessageRule{
		{
			MessageType: "ARR",
			Rule: []int{
				7, 13, 16, 17,
			},
		},
		{
			MessageType: "DEP",
			Rule: []int{
				7, 13, 16, 18,
			},
		},
		{
			MessageType: "FPL",
			Rule: []int{
				7, 8, 9, 10, 13, 15, 16, 18,
			},
		},
		{
			MessageType: "CHG",
			Rule: []int{
				7, 13, 16, 18, 22,
			},
		},
		{
			MessageType: "CNL",
			Rule: []int{
				7, 13, 16, 18,
			},
		},
		{
			MessageType: "DLA",
			Rule: []int{
				7, 13, 16, 18,
			},
		},
		{
			MessageType: "EST",
			Rule: []int{
				7, 13, 14, 16,
			},
		},
		{
			MessageType: "CPL",
			Rule: []int{
				7, 8, 9, 10, 13, 14, 15, 16, 18,
			},
		},
		{
			MessageType: "CDN",
			Rule: []int{
				7, 13, 16, 22,
			},
		},
		{
			MessageType: "ACP",
			Rule: []int{
				7, 13, 16,
			},
		},
		{
			MessageType: "RQP",
			Rule: []int{
				7, 13, 16, 18,
			},
		},
		{
			MessageType: "RQS",
			Rule: []int{
				7, 13, 16, 18,
			},
		},
		{
			MessageType: "SPL",
			Rule: []int{
				7, 13, 16, 18, 19,
			},
		},
		{
			MessageType: "ALR",
			Rule: []int{
				5, 7, 8, 9, 10, 13, 15, 16, 18, 19, 20,
			},
		},
		{
			MessageType: "RCF",
			Rule: []int{
				7, 21,
			},
		},
	},
}

type Group3_unit struct {
	MessageType string
	Desc string
}
var Group3_info = struct {
	Info []Group3_unit
}{
	Info:[]Group3_unit{
		{
			"ARR",
			"落地报",
		},
		{
			"DEP",
			"起飞报",
		},
		{
			"FPL",
			"領航计划报",
		},
		{
			"CHG",
			"修改領航报",
		},
		{
			"CNL",
			"取消领航报",
		},
		{
			"DLA",
			"延误报",
		},
		{
			"EST",
			"预计飞越报",
		},
		{
			"CPL",
			"飞行计划变更报",
		},
		{
			"CDN",
			"管制协调报",
		},
		{
			"ACP",
			"协调接受报",
		},
		{
			"LAM",
			"逻辑确认报",
		},
		{
			"RQP",
			"请求飞行计划报",
		},
		{
			"RQS",
			"请求领航计划补充信息报",
		},
		{
			"SPL",
			"领航计划补充信息报",
		},
		{
			"ALR",
			"告警报",
		},
		{
			"RCF",
			"无线电通信失效报",
		},
	},
}

