package config

type AirportInfoUint struct {
	Code             string // 机场代码
	AirPortName      string // 中文名称
	AirPortEnName    string // 英文名称
	AirPortCity      string //机场所在地区
	IntelligenceArea string // 情报区
}

var AirportInfo = struct {
	Info []AirportInfoUint
}{
	Info:[]AirportInfoUint{
		{
			Code: "ZBAA",
			AirPortName: "首都国际机场",
			AirPortEnName: "BEIJING/CAPITAL",
			AirPortCity: "北京/首都",
			IntelligenceArea: "北京情报区",
		},
		{
			Code: "ZBBB",
			AirPortName: "北京西郊机场",
			AirPortEnName: "BEIJING(CITY)",
			AirPortCity: "北京市",
			IntelligenceArea: "北京情报区",
		},
		{
			Code: "ZBNY",
			AirPortName: "北京南苑机场",
			AirPortEnName: "BEIJING/NANYUAN",
			AirPortCity: "北京/南苑",
			IntelligenceArea: "北京情报区",
		},
		{
			Code: "ZBCF",
			AirPortName: "土城子机场",
			AirPortEnName: "CHIFENG",
			AirPortCity: "赤峰",
			IntelligenceArea: "北京情报区",
		},
		{
			Code: "ZBCZ",
			AirPortName: "王村机场",
			AirPortEnName: "CHANGZHI/WANGCHUNG",
			AirPortCity: "长治",
			IntelligenceArea: "北京情报区",
		},
		{
			Code: "ZBDS",
			AirPortName: "鄂尔多斯机场",
			AirPortEnName: "DONGSHENG",
			AirPortCity: "东胜",
			IntelligenceArea: "北京情报区",
		},
		{
			Code: "ZBDT",
			AirPortName: "怀仁机场",
			AirPortEnName: "DATONG/HUAIREN",
			AirPortCity: "大同",
			IntelligenceArea: "北京情报区",
		},
		{
			Code: "ZBHH",
			AirPortName: "白塔国际机场",
			AirPortEnName: "HUHHOT",
			AirPortCity: "呼和浩特",
			IntelligenceArea: "北京情报区",
		},
		{
			Code: "ZBLA",
			AirPortName: "东山机场",
			AirPortEnName: "HAILAR/DONGSHAN",
			AirPortCity: "海拉尔",
			IntelligenceArea: "北京情报区",
		},
		{
			Code: "ZBOW",
			AirPortName: "二里半机场",
			AirPortEnName: "BAOTOU",
			AirPortCity: "包头",
			IntelligenceArea: "北京情报区",
		},
		{
			Code: "ZBSH",
			AirPortName: "山海关机场",
			AirPortEnName: "SHANHAIGUAN/QINHUA",
			AirPortCity: "秦皇岛",
			IntelligenceArea: "北京情报区",
		},
		{
			Code: "ZBSJ",
			AirPortName: "正定国际机场",
			AirPortEnName: "SHIJIAZHANG/ZHENGD",
			AirPortCity: "石家庄",
			IntelligenceArea: "北京情报区",
		},
		{
			Code: "ZBTJ",
			AirPortName: "滨海国际机场",
			AirPortEnName: "TIANJIN/ZHANGGUIZHUANG",
			AirPortCity: "天津/张贵庄（滨海）",
			IntelligenceArea: "北京情报区",
		},
		{
			Code: "ZBTL",
			AirPortName: "通辽机场",
			AirPortEnName: "TONGLIAO",
			AirPortCity: "通辽",
			IntelligenceArea: "北京情报区",
		},
		{
			Code: "ZBTX",
			AirPortName: "邢台机场",
			AirPortEnName: "XINGTAI",
			AirPortCity: "邢台",
			IntelligenceArea: "北京情报区",
		},
		{
			Code: "ZBUL",
			AirPortName: "乌兰浩特机场",
			AirPortEnName: "WULANHAOTE",
			AirPortCity: "乌兰浩特",
			IntelligenceArea: "北京情报区",
		},
		{
			Code: "ZBXH",
			AirPortName: "锡林浩特机场",
			AirPortEnName: "XILINHAOTE",
			AirPortCity: "锡林浩特",
			IntelligenceArea: "北京情报区",
		},
		{
			Code: "ZBYN",
			AirPortName: "武宿国际机场",
			AirPortEnName: "TAIYUAN/WUSU",
			AirPortCity: "太原",
			IntelligenceArea: "北京情报区",
		},
		{
			Code: "ZGBH",
			AirPortName: "福成机场",
			AirPortEnName: "BEIHAI/FUSHEN",
			AirPortCity: "北海",
			IntelligenceArea: "广州情报区",
		},
		{
			Code: "ZGHA",
			AirPortName: "黄花国际机场",
			AirPortEnName: "CHANGSHA/HUANGHUA",
			AirPortCity: "长沙",
			IntelligenceArea: "广州情报区",
		},
		{
			Code: "ZGDY",
			AirPortName: "荷花大庸机场",
			AirPortEnName: "ZHANGJIAJIE/DAYONG",
			AirPortCity: "张家界",
			IntelligenceArea: "广州情报区",
		},
		{
			Code: "ZGFS",
			AirPortName: "佛山沙堤机场",
			AirPortEnName: "FOSHAN",
			AirPortCity: "佛山",
			IntelligenceArea: "广州情报区",
		},
		{
			Code: "ZGGG",
			AirPortName: "白云国际机场",
			AirPortEnName: "GUANGZHOU/BAIYUN",
			AirPortCity: "广州/白云",
			IntelligenceArea: "广州情报区",
		},
		{
			Code: "ZGHA",
			AirPortName: "黄花国际机场",
			AirPortEnName: "CHANGSHA/HUANGHUA",
			AirPortCity: "长沙/黄花",
			IntelligenceArea: "广州情报区",
		},
		{
			Code: "ZJHK",
			AirPortName: "美兰国际机场",
			AirPortEnName: "HAIKOU/Meilan",
			AirPortCity: "海口/美兰",
			IntelligenceArea: "广州情报区",
		},
		{
			Code: "ZGHY",
			AirPortName: "衡阳机场",
			AirPortEnName: "HENGYANG",
			AirPortCity: "衡阳",
			IntelligenceArea: "广州情报区",
		},
		{
			Code: "ZGHZ",
			AirPortName: "惠州平潭机场",
			AirPortEnName: "HUIZHOU",
			AirPortCity: "惠州",
			IntelligenceArea: "广州情报区",
		},
		{
			Code: "ZGKL",
			AirPortName: "两江国际机场",
			AirPortEnName: "GUILIN",
			AirPortCity: "桂林",
			IntelligenceArea: "广州情报区",
		},
		{
			Code: "ZGMX",
			AirPortName: "梅县机场",
			AirPortEnName: "MEⅨIAN",
			AirPortCity: "梅县",
			IntelligenceArea: "广州情报区",
		},
		{
			Code: "ZGNN",
			AirPortName: "吴圩国际机场",
			AirPortEnName: "NANNING/WUXU",
			AirPortCity: "南宁",
			IntelligenceArea: "广州情报区",
		},
		{
			Code: "ZGOW",
			AirPortName: "外砂机场",
			AirPortEnName: "SHANTOU/WAISHA",
			AirPortCity: "汕头",
			IntelligenceArea: "广州情报区",
		},
		{
			Code: "ZGSD",
			AirPortName: "三灶国际机场",
			AirPortEnName: "CHANGDE",
			AirPortCity: "常德",
			IntelligenceArea: "广州情报区",
		},
		{
			Code: "ZGSY",
			AirPortName: "三亚凤凰机场",
			AirPortEnName: "SANYA/FENGHUANG",
			AirPortCity: "三亚",
			IntelligenceArea: "广州情报区",
		},
		{
			Code: "ZGSZ",
			AirPortName: "宝安国际机场",
			AirPortEnName: "SHENZHEN/HUANGTIAN",
			AirPortCity: "深圳",
			IntelligenceArea: "广州情报区",
		},
		{
			Code: "ZGWZ",
			AirPortName: "长洲岛机场",
			AirPortEnName: "WUZHOU/CHANGZHOUDA",
			AirPortCity: "梧州",
			IntelligenceArea: "广州情报区",
		},
		{
			Code: "ZGZH",
			AirPortName: "白莲机场",
			AirPortEnName: "LIUZHOU/BAILIAN",
			AirPortCity: "柳州",
			IntelligenceArea: "广州情报区",
		},
		{
			Code: "ZGZJ",
			AirPortName: "湛江机场",
			AirPortEnName: "ZHANJIANG",
			AirPortCity: "湛江",
			IntelligenceArea: "广州情报区",
		},
		{
			Code: "ZHCC",
			AirPortName: "新郑国际机场",
			AirPortEnName: "ZHENGZHOU",
			AirPortCity: "郑州",
			IntelligenceArea: "武汉情报区",
		},
		{
			Code: "ZHES",
			AirPortName: "许家坪机场",
			AirPortEnName: "ENSHI",
			AirPortCity: "恩施",
			IntelligenceArea: "武汉情报区",
		},
		{
			Code: "ZHHH",
			AirPortName: "天河国际机场",
			AirPortEnName: "WUHAN/TIANHE",
			AirPortCity: "武汉/天河",
			IntelligenceArea: "武汉情报区",
		},
		{
			Code: "ZHLY",
			AirPortName: "北郊机场",
			AirPortEnName: "LUOYANG",
			AirPortCity: "洛阳",
			IntelligenceArea: "武汉情报区",
		},
		{
			Code: "ZHNY",
			AirPortName: "姜营机场",
			AirPortEnName: "NANYANG",
			AirPortCity: "南阳",
			IntelligenceArea: "武汉情报区",
		},
		{
			Code: "ZHSS",
			AirPortName: "沙市机场",
			AirPortEnName: "JIONGSHA/SHASHI",
			AirPortCity: "沙市",
			IntelligenceArea: "武汉情报区",
		},
		{
			Code: "ZHXF",
			AirPortName: "刘集机场",
			AirPortEnName: "XIANGFAN/LIUJI",
			AirPortCity: "襄樊",
			IntelligenceArea: "武汉情报区",
		},
		{
			Code: "ZHYC",
			AirPortName: "三峡机场",
			AirPortEnName: "YICHANG/TUMENYA",
			AirPortCity: "宜昌",
			IntelligenceArea: "武汉情报区",
		},
		{
			Code: "ZLDH",
			AirPortName: "敦煌机场",
			AirPortEnName: "DUNHUAN",
			AirPortCity: "敦煌",
			IntelligenceArea: "兰州情报区",
		},
		{
			Code: "ZLHZ",
			AirPortName: "西关机场",
			AirPortEnName: "HANZHONG",
			AirPortCity: "汉中",
			IntelligenceArea: "兰州情报区",
		},
		{
			Code: "ZLIC",
			AirPortName: "河东机场",
			AirPortEnName: "YINCHUAN",
			AirPortCity: "银川",
			IntelligenceArea: "兰州情报区",
		},
		{
			Code: "ZLJQ",
			AirPortName: "嘉峪关机场",
			AirPortEnName: "JIAYUQUAN",
			AirPortCity: "嘉峪关",
			IntelligenceArea: "兰州情报区",
		},
		{
			Code: "ZLLL",
			AirPortName: "中川机场",
			AirPortEnName: "LANZHOU/ZHONGCHUAN",
			AirPortCity: "兰州",
			IntelligenceArea: "兰州情报区",
		},
		{
			Code: "ZLQY",
			AirPortName: "西峰镇机场",
			AirPortEnName: "QINGYANG/XIFENGXHE",
			AirPortCity: "庆阳",
			IntelligenceArea: "兰州情报区",
		},
		{
			Code: "ZLXN",
			AirPortName: "曹家堡机场",
			AirPortEnName: "XINING",
			AirPortCity: "西宁",
			IntelligenceArea: "兰州情报区",
		},
		{
			Code: "ZLXY",
			AirPortName: "咸阳国际机场",
			AirPortEnName: "XIAN/XIAN",
			AirPortCity: "西安",
			IntelligenceArea: "兰州情报区",
		},
		{
			Code: "ZLYA",
			AirPortName: "二十里堡机场",
			AirPortEnName: "YANAN",
			AirPortCity: "延安",
			IntelligenceArea: "兰州情报区",
		},
		{
			Code: "ZLYL",
			AirPortName: "西沙机场",
			AirPortEnName: "YULIN/XISHA",
			AirPortCity: "榆林",
			IntelligenceArea: "兰州情报区",
		},
		{
			Code: "ZPBS",
			AirPortName: "云端机场",
			AirPortEnName: "BAOSHAN",
			AirPortCity: "保山",
			IntelligenceArea: "昆明情报区",
		},
		{
			Code: "ZPDL",
			AirPortName: "大理机场",
			AirPortEnName: "DALI",
			AirPortCity: "大理",
			IntelligenceArea: "昆明情报区",
		},
		{
			Code: "ZPJH",
			AirPortName: "西双版纳机场",
			AirPortEnName: "JINGHONG/GASA",
			AirPortCity: "西双版纳",
			IntelligenceArea: "昆明情报区",
		},
		{
			Code: "ZPLJ",
			AirPortName: "三义机场",
			AirPortEnName: "LIJIANG",
			AirPortCity: "丽江",
			IntelligenceArea: "昆明情报区",
		},
		{
			Code: "ZPMS",
			AirPortName: "面上机场",
			AirPortEnName: "MANGSHI/LUXI",
			AirPortCity: "芒市",
			IntelligenceArea: "昆明情报区",
		},
		{
			Code: "ZPPP",
			AirPortName: "巫家坝国际机场",
			AirPortEnName: "KUNMING/WUJIABA",
			AirPortCity: "昆明",
			IntelligenceArea: "昆明情报区",
		},
		{
			Code: "ZPSM",
			AirPortName: "思茅机场",
			AirPortEnName: "SIMAO",
			AirPortCity: "思茅",
			IntelligenceArea: "昆明情报区",
		},
		{
			Code: "ZPZT",
			AirPortName: "昭通机场",
			AirPortEnName: "ZHAOTONG",
			AirPortCity: "昭通",
			IntelligenceArea: "昆明情报区",
		},
		{
			Code: "ZSAM",
			AirPortName: "高崎国际机场",
			AirPortEnName: "XIAMEN/GAOQI",
			AirPortCity: "厦门",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSAQ",
			AirPortName: "天柱山机场",
			AirPortEnName: "ANQING/DALONGSHAN",
			AirPortCity: "安庆",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSCG",
			AirPortName: "奔牛机场",
			AirPortEnName: "CHANGZHOU/BENNEU",
			AirPortCity: "常州",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSCN",
			AirPortName: "昌北国际机场",
			AirPortEnName: "NANCHANG",
			AirPortCity: "南昌",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSFY",
			AirPortName: "西关机场",
			AirPortEnName: "FUYANG",
			AirPortCity: "阜阳",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSFZ",
			AirPortName: "长乐国际机场",
			AirPortEnName: "FUZHOU",
			AirPortCity: "福州",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSGZ",
			AirPortName: "黄金机场",
			AirPortEnName: "GANZHOU/HUANGJIN",
			AirPortCity: "赣州",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSHC",
			AirPortName: "萧山国际机场",
			AirPortEnName: "HANGZHOU/JIANQIAO",
			AirPortCity: "杭州",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSJD",
			AirPortName: "罗家机场",
			AirPortEnName: "JINGDEZHEN",
			AirPortCity: "景德镇",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSJG",
			AirPortName: "济宁机场",
			AirPortEnName: "JINING",
			AirPortCity: "济宁",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSJJ",
			AirPortName: "庐山机场",
			AirPortEnName: "JIUJIANG",
			AirPortCity: "九江",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSJN",
			AirPortName: "遥墙国际机场",
			AirPortEnName: "JINAN/YAOQIANG",
			AirPortCity: "济南",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSJU",
			AirPortName: "衢州机场",
			AirPortEnName: "QUZHOU",
			AirPortCity: "衢州",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSLG",
			AirPortName: "白塔埠机场",
			AirPortEnName: "LIANYUNGAN/BAITAFU",
			AirPortCity: "连云港",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSLQ",
			AirPortName: "黄岩路桥机场",
			AirPortEnName: "HUANGYAN/LUQIAO",
			AirPortCity: "黄岩",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSLY",
			AirPortName: "临沂机场",
			AirPortEnName: "LINYI",
			AirPortCity: "临沂",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSNB",
			AirPortName: "栎社国际机场",
			AirPortEnName: "NINGBO/LISHE",
			AirPortCity: "宁波",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSNJ",
			AirPortName: "禄口国际机场",
			AirPortEnName: "NANJING",
			AirPortCity: "南京",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSNT",
			AirPortName: "兴东机场",
			AirPortEnName: "NANTONG/XINDONG",
			AirPortCity: "南通",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSOF",
			AirPortName: "骆岗机场",
			AirPortEnName: "HEHUI/LUOGANG",
			AirPortCity: "合肥",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSPD",
			AirPortName: "浦东机场",
			AirPortEnName: "SHANGHAI/PUDONG",
			AirPortCity: "上海/浦东",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSQD",
			AirPortName: "流亭国际机场",
			AirPortEnName: "QINGDAO",
			AirPortCity: "青岛",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSQZ",
			AirPortName: "晋江机场",
			AirPortEnName: "JINJIANG",
			AirPortCity: "泉州",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSSS",
			AirPortName: "虹桥机场",
			AirPortEnName: "SHANGHAI/HONGQIAO",
			AirPortCity: "上海/虹桥",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSTX",
			AirPortName: "屯溪机场",
			AirPortEnName: "HUANGSHAN/TUNXI",
			AirPortCity: "黄山",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSWF",
			AirPortName: "潍坊机场",
			AirPortEnName: "WEIFANG",
			AirPortCity: "潍纺",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSWH",
			AirPortName: "大水泊国际机场",
			AirPortEnName: "WEIHAI",
			AirPortCity: "威海",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSWX",
			AirPortName: "硕放机场",
			AirPortEnName: "WUXI",
			AirPortCity: "无锡",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSWY",
			AirPortName: "武夷山机场",
			AirPortEnName: "WUYISHAN",
			AirPortCity: "武夷山",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSWZ",
			AirPortName: "永强机场",
			AirPortEnName: "WENZHOU/YONGQIANG",
			AirPortCity: "温州",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSXZ",
			AirPortName: "观音机场",
			AirPortEnName: "XUZHOU/JIULISHAN",
			AirPortCity: "徐州",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSYN",
			AirPortName: "南洋机场",
			AirPortEnName: "YANCHENG",
			AirPortCity: "盐城",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSYT",
			AirPortName: "莱山机场",
			AirPortEnName: "YANTAI/PENGLAI",
			AirPortCity: "烟台",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZSYW",
			AirPortName: "义乌机场",
			AirPortEnName: "YIWU",
			AirPortCity: "义乌",
			IntelligenceArea: "上海情报区",
		},
		{
			Code: "ZUBD",
			AirPortName: "邦达机场",
			AirPortEnName: "CHANGDU/BANGDA",
			AirPortCity: "昌都",
			IntelligenceArea: "成都情报区",
		},
		{
			Code: "ZUCK",
			AirPortName: "江北国际机场",
			AirPortEnName: "CHONGQING/JIANGBEI",
			AirPortCity: "重庆",
			IntelligenceArea: "成都情报区",
		},
		{
			Code: "ZUDX",
			AirPortName: "河市机场",
			AirPortEnName: "DAXIAN/HESHIBA",
			AirPortCity: "达县",
			IntelligenceArea: "成都情报区",
		},
		{
			Code: "ZUGH",
			AirPortName: "广汉机场",
			AirPortEnName: "GUANGHAN",
			AirPortCity: "广汉",
			IntelligenceArea: "成都情报区",
		},
		{
			Code: "ZUGY",
			AirPortName: "龙洞堡国际机场",
			AirPortEnName: "GUIYANG/LEIZHUAN",
			AirPortCity: "贵阳",
			IntelligenceArea: "成都情报区",
		},
		{
			Code: "ZULS",
			AirPortName: "贡嘎机场",
			AirPortEnName: "LHASA",
			AirPortCity: "拉萨",
			IntelligenceArea: "成都情报区",
		},
		{
			Code: "ZULZ",
			AirPortName: "蓝田机场",
			AirPortEnName: "LUZHOU/XUANTIAN",
			AirPortCity: "泸州",
			IntelligenceArea: "成都情报区",
		},
		{
			Code: "ZUNC",
			AirPortName: "高坪机场",
			AirPortEnName: "NANCHONG",
			AirPortCity: "南充",
			IntelligenceArea: "成都情报区",
		},
		{
			Code: "ZUTR",
			AirPortName: "大兴机场",
			AirPortEnName: "TONGREN",
			AirPortCity: "铜仁",
			IntelligenceArea: "成都情报区",
		},
		{
			Code: "ZUUU",
			AirPortName: "双流国际机场",
			AirPortEnName: "CHENGDU/SHUANGLIU",
			AirPortCity: "成都",
			IntelligenceArea: "成都情报区",
		},
		{
			Code: "ZUXC",
			AirPortName: "青山机场",
			AirPortEnName: "XICHANG/QISHAN",
			AirPortCity: "西昌",
			IntelligenceArea: "成都情报区",
		},
		{
			Code: "ZUYB",
			AirPortName: "菜坝机场",
			AirPortEnName: "YIBIN",
			AirPortCity: "宜宾",
			IntelligenceArea: "成都情报区",
		},
		{
			Code: "ZWAK",
			AirPortName: "阿克苏机场",
			AirPortEnName: "AKESU/WENSU",
			AirPortCity: "阿克苏",
			IntelligenceArea: "乌鲁木齐情报区",
		},
		{
			Code: "ZWAT",
			AirPortName: "阿勒泰机场",
			AirPortEnName: "ALETAI",
			AirPortCity: "阿尔泰",
			IntelligenceArea: "乌鲁木齐情报区",
		},
		{
			Code: "ZWFY",
			AirPortName: "可托托海机场",
			AirPortEnName: "FUYUN/KEKETOTOHAI",
			AirPortCity: "富蕴",
			IntelligenceArea: "乌鲁木齐情报区",
		},
		{
			Code: "ZWHM",
			AirPortName: "哈密机场",
			AirPortEnName: "HAMI",
			AirPortCity: "哈密",
			IntelligenceArea: "乌鲁木齐情报区",
		},
		{
			Code: "ZWKC",
			AirPortName: "库车机场",
			AirPortEnName: "KUQA",
			AirPortCity: "库车",
			IntelligenceArea: "乌鲁木齐情报区",
		},
		{
			Code: "ZWKL",
			AirPortName: "库尔勒机场",
			AirPortEnName: "KUERLE",
			AirPortCity: "库尔勒",
			IntelligenceArea: "乌鲁木齐情报区",
		},
		{
			Code: "ZWKM",
			AirPortName: "克拉玛依机场",
			AirPortEnName: "KELAMAYI",
			AirPortCity: "克拉玛依",
			IntelligenceArea: "乌鲁木齐情报区",
		},
		{
			Code: "ZWSH",
			AirPortName: "喀什机场",
			AirPortEnName: "KASHI",
			AirPortCity: "喀什",
			IntelligenceArea: "乌鲁木齐情报区",
		},
		{
			Code: "ZWTC",
			AirPortName: "塔城机场",
			AirPortEnName: "TACHENG",
			AirPortCity: "塔城",
			IntelligenceArea: "乌鲁木齐情报区",
		},
		{
			Code: "ZWYN",
			AirPortName: "伊宁机场",
			AirPortEnName: "YINING",
			AirPortCity: "伊宁",
			IntelligenceArea: "乌鲁木齐情报区",
		},
		{
			Code: "ZYCC",
			AirPortName: "龙嘉国际机场",
			AirPortEnName: "CHANGCHUN/DAFANGSH",
			AirPortCity: "长春",
			IntelligenceArea: "沈阳情报区",
		},
		{
			Code: "ZYCY",
			AirPortName: "朝阳机场",
			AirPortEnName: "CHAOYANG",
			AirPortCity: "朝阳",
			IntelligenceArea: "沈阳情报区",
		},
		{
			Code: "ZYDD",
			AirPortName: "浪头机场",
			AirPortEnName: "DANDONG/LANGTOU",
			AirPortCity: "丹东",
			IntelligenceArea: "沈阳情报区",
		},
		{
			Code: "ZYHB",
			AirPortName: "太平国际机场",
			AirPortEnName: "HARBIN/YANJIAGANG",
			AirPortCity: "哈尔滨",
			IntelligenceArea: "沈阳情报区",
		},
		{
			Code: "ZYJL",
			AirPortName: "二台子机场",
			AirPortEnName: "JILIN/ERTAIZI",
			AirPortCity: "吉林",
			IntelligenceArea: "沈阳情报区",
		},
		{
			Code: "ZYJM",
			AirPortName: "东郊机场",
			AirPortEnName: "JIAMUSI",
			AirPortCity: "佳木斯",
			IntelligenceArea: "沈阳情报区",
		},
		{
			Code: "ZYJZ",
			AirPortName: "小岭子机场",
			AirPortEnName: "JINZHOU",
			AirPortCity: "锦州",
			IntelligenceArea: "沈阳情报区",
		},
		{
			Code: "ZYMD",
			AirPortName: "江海浪机场",
			AirPortEnName: "MUDANJIANG",
			AirPortCity: "牡丹江",
			IntelligenceArea: "沈阳情报区",
		},
		{
			Code: "ZYQQ",
			AirPortName: "南苑机场",
			AirPortEnName: "QIQIHAER/SHANGJIAZI",
			AirPortCity: "齐齐哈尔",
			IntelligenceArea: "沈阳情报区",
		},
		{
			Code: "ZYTL",
			AirPortName: "周水子国际机场",
			AirPortEnName: "DALIAN",
			AirPortCity: "大连",
			IntelligenceArea: "沈阳情报区",
		},
		{
			Code: "ZYTX",
			AirPortName: "桃仙国际机场",
			AirPortEnName: "SHENYANG/TAOXIAN",
			AirPortCity: "沈阳/桃仙",
			IntelligenceArea: "沈阳情报区",
		},
		{
			Code: "ZYYJ",
			AirPortName: "朝阳川国际机场",
			AirPortEnName: "YANJI/CHAOYANCHUAN",
			AirPortCity: "延吉",
			IntelligenceArea: "沈阳情报区",
		},
		{
			Code: "RCBS",
			AirPortName: "尚义机场",
			AirPortEnName: "JINMEN",
			AirPortCity: "金门/尚义",
			IntelligenceArea: "中国台湾",
		},
		{
			Code: "RCKH",
			AirPortName: "小港国际机场",
			AirPortEnName: "GAOⅪONG",
			AirPortCity: "高雄",
			IntelligenceArea: "中国台湾",
		},
		{
			Code: "RCKU",
			AirPortName: "嘉义机场",
			AirPortEnName: "JIAYI",
			AirPortCity: "嘉义",
			IntelligenceArea: "中国台湾",
		},
		{
			Code: "RCNN",
			AirPortName: "台南机场",
			AirPortEnName: "TAINAN",
			AirPortCity: "台南",
			IntelligenceArea: "中国台湾",
		},
		{
			Code: "RCQC",
			AirPortName: "太武机场",
			AirPortEnName: "MAGONG",
			AirPortCity: "马公",
			IntelligenceArea: "中国台湾",
		},
		{
			Code: "RCSS",
			AirPortName: "松山机场",
			AirPortEnName: "TAIBEI/SONGSHAN",
			AirPortCity: "台北/松山",
			IntelligenceArea: "中国台湾",
		},
		{
			Code: "RCTP",
			AirPortName: "桃园国际机场",
			AirPortEnName: "TAIBEI",
			AirPortCity: "台北国际机场",
			IntelligenceArea: "中国台湾",
		},
		{
			Code: "VHHH",
			AirPortName: "赤腊角机场",
			AirPortEnName: "HONGKONG/INTL",
			AirPortCity: "香港/国际机场",
			IntelligenceArea: "中国香港",
		},
		{
			Code: "VTBD",
			AirPortName: "曼谷廊曼国际机场",
			AirPortEnName: "BANGKOK",
			AirPortCity: "曼谷",
			IntelligenceArea: "国外机场",
		},
	},
}