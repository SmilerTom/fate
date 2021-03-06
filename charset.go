package fate

import (
	"log"
	"strconv"
	"strings"

	"github.com/godcong/excavator"
	"github.com/godcong/fate/debug"
	"github.com/godcong/fate/model"
	"github.com/satori/go.uuid"
)

type Charset []string

var (
	//木
	mu = Charset{
		"",
		"九",
		`干 工 弓 广 丌 及 孑 巾 久 口 亏 乞 彡 已 义`,
		`卞 丐 公 勾 讥 见 介 斤 今 仅 巨 开 亢 孔 木
	廿 牛 亓 气 欠 区 犬 劝 五 牙 艺 元 月 匀`,
		`本 东 尕 甘 功 古 瓜 归 宄 讧 卉 击 叽 饥 记
	加 甲 艽 叫 讦 纠 旧 句 卡 可 叩 邝 兰 卯 艿 巧 丘 囚 犰 去 术 外
	未 业 仪 仡 议 玉 驭 札`,

		`朳 芏 朵 纥 各 巩 共 关 观 犷 轨 过 玑 芨 机
	乩 肌 吉 级 伎 纪 夹 价 幵 奸 囝 件 讲 交 阶 阱 臼 讵 军 扛 伉 考
	扣 夸 匡 夼 纩 扩 芒 芃 朴 祁 芑 岂 企 扦 芊 乔 庆 曲 权 戎 阮 芍
	芄 伪 扤 吓 芗 芎 朽 吁 许 旭 厌 仰 尧 芋 聿 芝 朱 竹`,

		`芭 苄 材 苍 岑 杈 苌 床 苁 村 杜 扼 苊 呃 芳
	芬 芙 伽 尬 改 肝 杆 肛 纲 杠 告 更 攻 贡 估 谷 龟 妫 庋 呙 囯 旱
	何 吼 花 肓 鸡 技 忌 妓 歼 坚 间 角 疖 劫 芥 近 妗 劲 究 玖 局 拒
	苣 姖 君 佧 忾 闶 抗 囥 壳 克 抠 芤 库 快 狂 旷 困 苈 芦 杩 拟 你
	伲 芘 芪 杞 启 弃 佥 芡 羌 呛 芹 芩 穷 邱 虬 诎 驱 劬 却 芟 杉 杓
	苏 苇 我 芜 吾 吴 妩 芴 杌 匣 苋 县 肖 芯 杏 芽 严 言 杨 苡 吟 饫
	妪 芫 芸 纭 杖 芷 苎`,

		`枊 昂 茇 板 苞 杯 苯 苾 茌 杵 枞 妸 轭 枋 枫
	苻 斧 呷 该 陔 苷 秆 绀 杲 疙 肱 供 苟 狗 构 购 诟 咕 孤 姑 股 固
	呱 诖 乖 官 规 匦 诡 柜 国 果 杭 忽 昏 佶 亟 虮 季 剂 佳 迦 郏 驾
	肩 艰 拣 枧 建 降 郊 佼 杰 诘 届 卺 茎 京 经 肼 径 弪 迥 疚 拘 苴
	狙 居 驹 咀 具 卷 咔 咖 凯 苛 肯 空 苦 侉 侩 郐 诓 拦 枥 林 苓 茏
	茆 茅 茂 枚 杪 苠 茉 苜 呢 怩 拈 茑 杻 杷 枇 苤 苹 其 奇 歧 祈 肷
	枪 侨 茄 苘 顷 茕 穹 屈 诠 券 苒 枘 若 苫 枢 松 苔 苕 玩 枉 卧 析
	侠 狎 贤 枭 欣 厓 兖 侥 杳 宜 驿 茚 英 苑 岳 枕 枝 竺 杼 茁`,

		`柏 柲 荜 标 柄 草 茬 茶 尝 柽 柢 笃 俄 枹 尜
	柑 竿 缸 郜 诰 咯 革 虼 哏 郠 拱 枸 骨 牯 故 胍 挂 冠 咣 皈 鬼 癸
	訇 荭 虹 哄 逅 恢 荟 荤 咭 笈 急 挤 哜 既 珈 枷 荚 胛 叚 架 茧 俭
	荐 贱 牮 姜 绛 茭 娇 姣 骄 挢 狡 饺 绞 觉 皆 拮 结 界 疥 诫 衿 矜
	荩 胫 扃 赳 韭 柩 莒 矩 举 郡 胩 看 钪 拷 珂 柯 轲 科 咳 恪 客 眍
	绔 挎 哙 狯 哐 诳 贶 括 栏 荔 栎 柃 柳 栊 栌 荦 络 荬 面 茗 昵 逆
	柠 柈 枰 荠 契 恰 茜 荞 俏 诮 亲 轻 俅 酋 祛 朐 荃 畎 荛 饶 绕 荏
	茸 荣 茹 柿 树 竖 荪 莛 茼 统 柁 柝 柙 狭 相 枵 荇 荀 荨 研 俨 药
	荑 蚁 弈 奕 疫 羿 茵 荫 胤 郢 哟 柚 禺 竽 语 狱 柞 栅 栈 柘 祗 枳
	栉 荮 彦 茱 柱 斫`,

		`桉 笆 栢 梆 荸 笔 栟 柴 郴 莼 档 荻 屙 婀 莪
	饿 粉 莩 赅 酐 疳 赶 高 羔 哥 格 鬲 哿 根 耕 哽 绠 恭 蚣 躬 珙 鸪
	罟 顾 倌 莞 桄 逛 桂 桧 衮 郭 捍 悍 荷 核 桁 笏 桦 桓 获 唧 笄 屐
	姬 脊 觊 继 痂 家 恝 兼 捡 笕 健 舰 豇 桨 胶 轿 桀 蚧 紧 赆 痉 竞
	阄 桕 疽 桔 俱 倨 捐 娟 倦 桊 绢 捃 隽 莰 栲 疴 课 恳 倥 恐 哭 胯
	脍 宽 框 悝 捆 阃 悃 栝 莱 栳 莉 莅 栗 莲 莨 莽 莓 娩 难 匿 臬 耙
	莆 栖 桤 耆 颀 脐 起 悭 虔 悄 桥 窍 衾 倾 卿 逑 鸲 悛 辁 拳 缺 桡
	桑 莎 莘 倏 秫 栓 凇 颂 荽 笋 桃 梃 桐 砼 荼 桅 莴 悟 牺 奚 莶 校
	哮 栩 桠 唁 验 样 倚 挹 悒 谊 莺 莹 痈 邕 莜 莸 莠 娱 圄 峪 预 原
	笊 桢 茝 桎 株 桩 桌`,

		`菝 菢 笨 菶 萆 梐 笾 彬 菜 梣 菖 笞 菙 崔 萃
	笪 萏 菪 笛 兜 鄂 阏 谔 梵 菲 菔 桴 符 盖 敢 鸽 舸 梗 龚 笱 菰 菇
	蛄 蛊 梏 牿 鸹 馆 掼 绲 掴 帼 菡 萑 隍 谎 掎 偈 悸 寄 笳 袈 戛 袷
	假 菅 笺 检 趼 减 谏 矫 皎 脚 教 秸 婕 堇 菁 惊 颈 竟 婧 救 厩 掬
	菊 据 距 惧 鄄 眷 掘 桷 菌 皲 珺 康 氪 骒 啃 裉 控 寇 眶 盔 傀 匮
	梱 婪 笠 敛 啉 菱 棂 笼 萝 梅 萌 梦 萘 猊 旎 偶 笸 菩 萋 萁 骐 骑
	绮 掐 掮 乾 羟 筇 蚯 球 赇 蛆 躯 娶 圈 痊 绻 萨 啬 梢 笙 菽 梳 耜
	笥 菘 桫 梭 萄 梼 梯 笤 萜 桶 菟 菀 萎 隗 偓 梧 晤 菥 厢 萧 啸 偕
	谐 械 悻 悬 菸 阎 掩 厣 眼 谚 窑 翊 萤 营 萦 唷 萸 圉 庾 阈 菑 笮
	著 棁 梓 菹`,

		`棒 葆 筚 萹 策 搽 猹 蒇 趁 葱 蒂 椟 鹅 萼 遏
	愕 葑 搁 隔 葛 赓 辊 琯 猴 葫 鹄 慌 赍 犄 缉 葭 跏 颊 蛱 徦 犍 睑
	裥 楗 毽 腱 蒋 椒 蛟 搅 窖 揭 喈 街 颉 筋 腈 景 敬 窘 揪 啾 琚 椐
	犋 飓 鹃 喀 揩 蒈 慨 棵 颏 缂 筘 喾 裤 款 筐 揆 葵 蒉 喟 馈 愦 愧
	琨 蛞 榔 棱 联 椋 琳 蒌 椤 棉 募 葩 棚 椑 葡 期 欺 琪 琦 棋 蛴 祺
	葺 葜 谦 椠 嵌 强 愀 翘 趄 琴 琼 萩 遒 巯 趋 蛐 阒 筌 裙 森 厦 筛
	葚 覃 棠 葶 筒 椭 皖 葳 蒍 珷 稀 葸 筅 葙 缃 萱 雅 筵 雁 傜 椰 遗
	椅 硬 遇 御 鹆 寓 葬 棹 筝 植 筑 椎 孳 棕 最`,

		`嗄 蒡 蓓 蓖 蔀 槎 榇 筹 椽 槌 椿 戥 椴 蛾 腭
	蒽 摁 概 戤 感 幹 筻 搞 缟 塥 嗝 跟 遘 彀 媾 毂 鼓 痼 褂 瑰 跪 蒿
	嗬 鲎 槐 畸 蒺 楫 嫉 麂 蓟 筴 嫁 搛 蒹 缣 简 骞 谫 酱 跤 敫 睫 解
	骱 谨 靳 禁 缙 睛 舅 雎 裾 榉 龃 筠 楷 戡 嗑 稞 窠 窟 跨 蒯 筷 窥
	暌 魁 髡 廓 蓝 榄 楞 蓠 廉 楝 楼 楣 蓦 墓 幕 楠 睨 腻 筢 蓬 榀 蒲
	签 愆 跷 鄡 勤 楸 群 蓉 蓐 嗓 瑟 筲 椹 筮 蒴 蒜 蓑 蓊 皙 暇 献 筱
	楔 歇 蓄 楦 靴 罨 肄 缢 蓥 楹 颖 媵 榆 愚 瘐 蓣 瑗 楂 蓁 蒸 罪`,

		`蔼 榜 蔽 箅 槟 蔡 箣 蔟 榱 箪 凳 蔸 鹗 榧 嘎
	睾 膏 槁 歌 搿 膈 箇 觏 箍 鹘 管 蝈 蜾 赫 瘊 夥 箕 跽 鲚 暨 嘉 瘕
	戬 僭 僬 鲛 酵 截 鲒 竭 馑 廑 兢 儆 獍 僦 窭 槛 阚 慷 犒 颗 箜 蔻
	骷 酷 睽 蔹 蓼 蔺 榴 箩 蔓 甍 蔑 模 蔫 酿 箄 裴 嘁 綦 旗 搴 箝 歉
	蔷 敲 箧 箐 蕖 蜷 榷 榕 箬 槊 蔌 算 榫 榻 箢 蔚 斡 蓰 辖 箫 榭 榍
	蓿 酽 疑 蜴 瘗 龈 蝇 蜮 愿 箦 榨 寨 肇 蔗 榛 槠 箸 赚`,

		`箯 槽 樗 稻 噔 蕫 懂 额 颚 蕃 樊 噶 橄 稿 骼
	鲠 穀 鲧 横 骺 篌 糇 槲 蝗 篁 蕙 稽 畿 蕺 瘠 稷 鲫 鲣 鹣 翦 踺 箭
	僵 蕉 噍 羯 槿 觐 憬 樛 踞 屦 蕨 靠 瞌 蝌 聩 篑 醌 篓 樠 瞢 耦 篇
	蕲 槭 樯 鞒 憔 撬 擒 蝤 觑 蕤 蕊 蔬 樘 瞎 箱 橡 蝎 撷 蕈 颜 魇 谳
	毅 樱 龉 窳 蕴 樟 箴 橥 篆 蕞 醉`,

		`薜 篦 篰 橙 篪 篘 橱 薋 篡 噩 篚 擀 篙 糕 篝
	薅 嚆 薨 蕻 黉 圜 癀 墼 激 髻 冀 稼 彊 缰 耩 犟 徼 缴 噤 鲸 橘 踽
	遽 橛 麇 瞰 鲲 篮 蕾 篱 篥 橹 鲵 篷 器 褰 黔 橇 缲 樵 鞘 檎 螓 檠
	擎 磬 鼽 糗 磲 醛 穑 薯 薮 薇 蕹 樨 魈 薤 薪 薛 赝 邀 薏 嬴 瘿 橼
	樾 嘴 樽`,

		`藊 檦 檗 藏 簇 瞪 篼 簖 鳄 藁 臌 簋 馘 蟥 簧
	鳇 豁 羁 藉 謇 鹪 鹫 鞠 糠 髁 檑 檩 簏 懋 檬 藐 篾 蹊 襁 瞧 罄 鳅
	璩 龋 薷 篸 簌 薹 檀 魏 檄 罅 藓 薰 檐 鹬`,

		`藨 檫 簦 簟 瞽 鳏 颢 鞯 糨 襟 鞫 藜 藕 髂 瞿
	鬈 檮 藤 黠 嚣 鹰 簪`,

		`霭 簸 蹬 蘅 攉 藿 蠖 醮 警 鬏 髋 籁 麓 蘑 孽
	攀 麒 黢 蘧 蟹 籀 缵`,

		`蘩 鳜 籍 醵 蘖 黦 纂`,

		`赣 瓘 夔`,

		`鱇 氍 鬻 蘸`,

		`欑 蠲 癯 颧 鼹 躜 攥`,

		`衢`,

		`戆`,

		`蠼`,
	}
	//金
	jin = Charset{
		"",
		` 匕 厂 刀 儿 七 人 入 十 厶`,

		` 才 叉 亍 川 寸 千 刃 三 上 尸 士 夕 小`,

		` 仓 车 仇 戈 仁 认 仍 冗 少 什 升 氏 手 殳 书
		闩 双 兮 心 扎 专`,

		` 册 出 刍 处 乏 戋 节 刊 仟 阡 邛 仞 讱 扔 闪
		讪 申 生 失 石 史 矢 示 世 仕 市 甩 帅 司 丝 四 凸 仙 轧 乍 占 正
		卮 主`,

		` 伧 产 忏 伥 伡 臣 成 丞 冲 舛 创 此 次 汆 存
		忖 丢 而 钆 刚 划 吏 列 齐 迁 任 纫 礽 如 伞 扫 色 伤 舌 厍 设 师
		式 守 妁 死 寺 讼 夙 岁 孙 刎 问 西 吸 戏 先 纤 囟 邢 匈 旬 巡 驯
		页 曳 钇 伛 刖 再 在 早 则 吒 众 舟 州 伫 妆 庄 壮 字`,

		` 财 岔 刬 肠 怅 抄 吵 扯 忱 迟 赤 饬 沖 忡 初
		氚 串 怆 吹 纯 词 钉 兑 刭 陉 钌 判 刨 钋 抢 吣 忍 韧 轫 饪 妊 肜
		删 邵 劭 佘 社 伸 身 声 时 识 寿 抒 纾 束 吮 私 伺 祀 姒 忪 宋 诉
		忒 吻 饩 系 伭 辛 秀 序 玚 译 酉 皂 诈 钊 诏 针 吱 纸 忮 诌 肘 助
		抓 孜 邹 走 诅 佐 作 坐 阼`,

		` 参 厕 侧 侘 刹 诧 拆 钗 侪 昌 怊 弨 抻 衬 诚
		承 宠 怵 绌 钏 垂 佌 刺 钓 钒 庚 刮 怪 刿 刽 环 饯 金 净 剀 侃 刻
		刳 钔 狞 拧 孥 钕 凭 妻 钎 戗 戕 郄 怯 青 取 叁 丧 钐 衫 姗 陕 疝
		尚 绍 舍 呻 侁 诜 绅 审 肾 诗 实 使 始 驶 势 事 试 视 受 叔 述 刷
		咝 兕 驷 怂 肃 所 钍 兔 昔 穸 细 祆 限 线 详 些 性 姓 询 咋 甾 驵
		枣 责 迮 昃 闸 怔 织 侄 帜 终 周 妯 咒 宙 绉 帚 邾 侏 诛 拄 贮 拙
		宗 驺 卒 组 怍`,

		` 钡 残 恻 臿 查 差 姹 觇 钞 疢 宬 持 穿 舡 疮
		春 呲 祠 殂 促 钭 毒 独 度 钝 耏 贰 罚 阀 珐 钫 钙 钢 宫 钩 剐 郝
		剑 俓 钜 绝 钧 枯 闾 钠 陧 钮 钤 前 钦 侵 氢 秋 衽 狨 绒 柔 砂 珊
		舢 殇 珅 神 哂 矧 甚 胂 牲 省 胜 狮 施 拾 食 蚀 是 室 首 狩 姝 耍
		拴 顺 说 思 俟 娀 送 诵 叟 俗 虽 狲 钛 剃 恸 钨 郤 籼 险 庠 哓 骁
		信 星 咻 修 庥 须 胥 叙 宣 选 绚 削 逊 钥 俞 俣 哉 咱 昝 蚤 怎 眨
		咤 毡 战 帧 胗 狰 拯 挣 栀 胝 指 轵 咫 峙 陟 钟 昼 拽 咨 姿 兹 总
		俎 祖 昨 胙 祚`,

		`铋 钵 剥 钹 铂 蚕 舱 敇 豺 谄 倡 鬯 晁 眧 耖
		唓 宸 琤 称 乘 蚩 翅 帱 俶 陲 玼 瓷 脆 挫 厝 钿 凋 铎 罡 钴 壶 剞
		钾 借 剧 峻 骏 钶 铃 铆 钼 铌 倪 捏 聂 珮 铍 钷 剖 铅 钱 钳 倩 挈
		请 逡 辱 弱 扇 捎 哨 射 娠 谂 眚 逝 轼 殊 衰 谁 铄 鸶 耸 悚 素 速
		狻 谇 祟 损 隼 唆 索 唢 铊 钽 铁 捅 途 剜 紊 唏 息 席 蚬 猃 陷 祥
		逍 宵 绡 笑 胸 脩 袖 绣 訏 徐 痃 眩 铉 殉 剡 铀 钰 钺 悦 栽 载 宰
		赃 奘 唣 造 贼 哳 痄 斋 窄 债 真 畛 疹 钲 症 脂 挚 胵 酎 皱 珠 诸
		疰 谆 酌 诼 资 郰 诹 陬 租 钻 唑 座`,

		`铵 猜 彩 惭 惨 曹 掺 谗 婵 铲 猖 阊 娼 常 偿
		徜 惝 唱 巢 晨 谌 趻 偁 铖 脭 瓻 豉 舂 崇 铳 紬 惆 绸 偢 船 捶 啜
		疵 粗 猝 啐 悴 脞 措 铛 祷 得 掉 铞 铫 铤 铥 堵 铒 副 铬 铪 铧 秽
		祭 寂 铗 剪 铰 捷 旌 铠 勘 铐 馗 铑 率 铭 啮 聍 掊 跄 惬 圊 情 铨
		悫 雀 蚺 铷 铯 铩 啥 唼 商 绱 奢 赊 猞 蛇 赦 婶 绳 盛 匙 授 售 兽
		绶 鄃 孰 庶 唰 爽 偲 隋 随 琐 铴 铜 偷 推 欷 悉 惜 觋 袭 铣 徙 阋
		掀 衔 舷 馅 象 斜 偰 訢 衅 羞 鸺 宿 琇 酗 绪 续 谖 旋 铘 谒 铱 铟
		银 铕 谕 啧 帻 舴 铡 蚱 砦 粘 着 赈 睁 铮 趾 铢 蛀 综 族 做`,

		`锕 裁 曾 插 嗏 馇 馋 禅 孱 敞 超 朝 琛 惩 牚
		啻 畴 厨 锄 储 揣 遄 喘 窗 赐 琮 窜 毳 皴 搓 嵯 矬 痤 锉 貂 犊 牍
		锇 锋 锆 割 辜 锅 戟 缄 锏 锔 竣 锎 铿 铼 锒 铹 锂 链 靓 锍 锊 甯
		裒 铺 腔 禽 锓 揿 氰 蛩 阕 然 揉 鄏 锐 毵 散 搔 骚 嫂 痧 跚 善 觞
		赏 稍 畲 甥 剩 释 谥 舒 疏 暑 黍 属 税 舜 斯 蛳 缌 竦 搜 嗖 馊 酥
		粟 谡 飧 睃 锁 锑 替 童 晰 傒 舾 粞 犀 舄 痫 鹇 羡 缐 销 谢 锌 猩
		惺 锈 絮 婿 揎 喧 揄 逾 腴 愉 喻 裕 缊 凿 锃 揸 喳 掌 絷 殖 黹 蛛
		铸 装 紫 腙 尊 酢`,

		`锛 睬 粲 鼌 嗔 赪 絺 傺 搊 酬 稠 愁 裯 蜍 雏
		楚 搐 触 搋 歂 锤 辞 慈 跐 催 瘁 错 锝 锭 嘟 睹 鄜 辐 锢 锪 辑 鉴
		键 剿 锦 靖 锯 锩 锞 锟 锣 锚 锰 嗫 锘 锫 遣 慊 蜣 锖 嗪 寝 跫 鹊
		阙 缛 瑞 腮 塞 搡 裟 傻 歃 煞 骟 蛸 艄 慑 蜃 慎 摅 输 毹 署 蜀 鼠
		腧 数 睡 搠 肆 嗣 飕 稣 嗉 愫 睢 嗦 嗍 羧 锬 酮 骰 颓 腿 锡 媳 禊
		酰 跹 锨 嫌 跣 腺 想 像 新 歆 腥 貅 馐 嗅 煦 暄 瑜 觎 愈 锗 斟 甄
		缜 酯 瘃 锥 锱`,

		`锿 綵 嘈 锸 察 瘥 摻 蝉 僝 鋋 嫦 綝 酲 殠 摴
		僢 摐 雌 鹚 骢 摧 粹 翠 镀 锻 锷 镄 锾 劐 精 静 聚 镌 劂 锴 镂 銮
		镅 镁 綮 锵 嫱 锹 劁 谯 锲 蜻 銎 睿 赛 糁 缫 瘙 僧 酾 鄯 韶 誓 瘦
		摔 锶 锼 瞍 嗾 嗽 僳 觫 酸 隧 缩 僮 酴 蜥 僖 屣 鲜 需 舆 窬 臧 遭
		谮 蜘 摭 碡 翥 粽 僔`,

		`镑 艑 镔 镈 踩 憯 艖 廛 麨 瞋 漦 褫 瘛 憧 諔
		膗 噇 糍 璁 聪 醋 璀 撮 敷 镐 镉 镓 镏 镆 镎 镊 镍 劈 谴 噙 趣 髯
		蝾 镕 糅 褥 馓 磉 缮 奭 艏 熟 撕 嘶 螋 艘 鋈 嘻 膝 暹 儇 镒 劓 蝓
		鋆 糌 赜 憎 缯 谵 稹 镇 嘱 颛 撞 幢 踪 遵 撙 嘬`,

		`翱 镚 镖 餐 操 糙 幨 鲳 氅 鋹 踸 瘳 憷 踹 膪
		輴 錞 踳 蹉 鹾 镝 雕 镜 镘 穆 颞 缱 鲭 瘸 蹂 儒 噻 颡 懆 穇 擅 膳
		嬗 擞 镗 螅 羲 醒 髹 錡 錾 赞 噪 甑 蹅 瘵 整 麈 鲰 镞`,

		`擦 縩 璨 螬 艚 嚓 鲿 韔 懤 歜 黜 镩 镫 镦 徽
		蹇 镢 镧 镣 镥 縻 鍪 蹑 镤 镨 镪 嚅 孺 鳃 臊 膻 蟀 瞬 邃 镡 蟋 谿
		燨 蟓 擤 翼 糟 罾 鍼 骤 瞩`,

		`鏊 鎞 饆 襜 繟 儭 艟 躇 幮 蹙 镭 镰 鎏 謦 鞣蟮 雟 燹镱 镯 鬃`,

		`鏖 镲 蟾 谶 蹴 蹿 骥 襦 醯`,

		`犨 鰆 黩 鐍 鏻 黥 颥 蠕 繻 鳝 孀 璺 霰 馨 瓒
		躁 鐘`,

		`蠢 髓`,

		`镵 镶`,

		`鼷 鱚 趱`,

		`矗 鑫`,

		`鑶`,
	}
	//水
	shui = Charset{
		"",
		` 八 勹 卜 乜`,

		` 凡 飞 马 么 门 万 亡 习 下 乡 子`,

		` 巴 办 贝 比 币 不 歹 反 方 分 风 凤 夫 父 讣
		冈 互 戶 户 化 幻 计 毛 爿 匹 片 攴 仆 壬 卅 水 文 无 毋 勿 凶 云`,

		` 扒 叭 白 半 包 北 必 边 弁 卟 布 匆 发 犯 氾
		冯 弗 付 邗 汉 夯 号 禾 弘 乎 汇 劢 邙 矛 们 灭 民 皿 末 母 目 仫
		丕 皮 氕 平 叵 扑 讫 仨 疋 汀 务 写 兄 玄 穴 训 印 匝 汁`,

		` 百 阪 邦 毕 闭 邠 冰 并 汊 池 闯 凼 讹 伐 帆
		汎 邡 防 仿 访 妃 份 讽 缶 伏 凫 负 妇 亥 汗 行 好 合 红 后 冱 华
		欢 回 会 讳 汲 江 邟 妈 犸 吗 买 迈 忙 扪 米 糸 名 牟 仳 牝 乒 迄
		汔 汝 杀 汕 收 汜 汤 网 危 污 汐 向 协 邪 刑 兴 休 血 汛 杂`,

		` 夿 把 弝 吧 扳 伴 扮 报 陂 孛 邶 貝 狈 呗 伻
		皀 吡 佊 沘 妣 庇 邲 诐 抃 汳 汴 忭 别 兵 邴 伯 驳 补 步 吥 沧 沉
		呆 沌 返 饭 泛 妨 纺 吠 吩 纷 汾 沣 沨 佛 否 呋 扶 孚 抚 甫 汞 佝
		沟 汩 还 邯 含 罕 沆 诃 亨 宏 囫 护 沪 怀 奂 即 戒 鸠 况 泐 冷 沥
		沦 玛 祃 麦 尨 沒 没 每 闷 芈 汨 沔 妙 闵 亩 沐 尿 妞 扭 狃 忸 纽
		沤 彷 抛 沛 批 邳 伾 屁 评 沏 汽 沁 扰 沙 纱 沈 汰 汪 忘 沩 尾 纹
		汶 沃 希 闲 孝 忻 形 汹 妍 沂 妤 沅 沄 妘 沚 状`,

		` 岸 拔 爸 败 攽 版 拌 姅 绊 孢 饱 抱 卑 备 奔
		屄 彼 畀 贬 变 表 玢 秉 幷 並 拨 波 帛 瓝 泊 怖 法 范 贩 祊 肪 房
		放 非 肥 肺 狒 沸 氛 奋 忿 奉 肤 拂 服 怫 绂 绋 拊 府 咐 阜 驸 侅
		泔 岡 沽 卦 函 呵 和 劾 河 佫 轰 泓 郈 呼 狐 弧 虎 怙 诨 或 货 泾
		沮 炬 泪 泠 泷 泸 泺 码 卖 盲 氓 牦 泖 玫 妹 門 孟 弥 觅 泌 宓 黾
		苗 庙 玟 抿 泯 鸣 命 抹 殁 沫 陌 侔 拇 姆 牧 泥 泞 拍 泮 咆 狍 庖
		泡 呸 帔 佩 抨 怦 朋 披 拚 贫 郱 泼 迫 泣 浅 泅 沭 饲 泗 沲 沱 味
		武 物 郃 弦 冼 现 享 協 胁 泄 泻 绁 幸 泫 学 郇 沿 泱 泳 油 鱼 雨
		郁 咂 泽 沾 沼 治 注 狀`,

		` 疤 胈 拜 帮 绑 胞 保 鸨 背 贲 甭 泵 迸 秕 珌
		毖 哔 陛 窆 扁 昪 便 飑 骉 饼 玻 勃 钚 测 浐 泚 荡 点 洞 洱 畈 飛
		封 風 疯 罘 氟 俘 郛 祓 复 負 訃 阁 哈 孩 骇 顸 绗 阂 贺 很 狠 恨
		恒 洪 紅 侯 厚 轷 浒 祜 哗 骅 徊 洹 奐 宦 皇 挥 虺 洄 哕 浍 诲 绘
		浑 活 計 洎 济 挟 浃 茳 洚 浇 洁 津 荆 姥 洌 浏 洛 蚂 骂 脉 茫 昴
		冒 贸 眉 美 昧 虻 咪 迷 祢 洣 弭 眄 勉 眇 秒 咩 珉 闽 哞 某 浓 哌
		派 眅 盼 叛 胖 胚 毗 骈 拼 姘 品 俜 玶 洴 屏 珀 匍 柒 洽 泉 染 娆
		洳 洒 飒 洮 洼 洧 闻 郚 郗 洗 虾 咸 涎 宪 香 响 饷 项 巷 卸 洩 洫
		恤 泶 洵 浔 恂 徇 衍 洋 洇 荥 盈 拶 浈 洲 洙 浊`,

		`氨 捌 粑 罢 班 般 颁 舨 浜 蚌 趵 豹 倍 狽 悖
		被 唄 倴 畚 俾 舭 粃 毙 畢 狴 髟 俵 宾 病 砵 饽 袯 亳 浡 逋 庯 捕
		哺 部 涔 秤 臭 唇 涤 娥 舫 紡 匪 诽 紛 峰 逢 俸 蚨 浮 俯 釜 胲 海
		氦 害 浛 航 颃 蚝 耗 浩 盍 哼 珩 竑 候 華 换 唤 涣 浣 珲 晖 悔 恚
		贿 浹 贾 涧 浆 较 浸 涇 酒 涓 浚 涞 浪 涝 涟 流 馬 唛 脈 旄 浼 們
		勐 脒 敉 秘 眠 悯 冥 莫 秣 毪 畝 涅 紐 俳 畔 袢 旁 袍 疱 陪 配 旆
		倗 蚍 郫 疲 陴 胼 娉 瓶 哱 圃 浦 凄 訖 润 脎 涩 殺 紗 涉 涑 绥 娑
		泰 烫 涛 涕 涂 涒 涠 蚊 紋 涡 浯 務 浠 消 效 挾 脅 屑 訓 珢 涌 浴
		耘 涨 浙 浞`,
		`豝 敗 湴 絆 偝 绷 偪 閉 庳 敝 婢 匾 貶 彪 婊
		菠 舶 脖 啵 脣 淳 淙 淬 淡 淀 渎 訛 販 訪 婓 啡 绯 淝 悱 酚 偾 唪
		麸 趺 匐 涪 艴 紱 紼 辅 婦 袱 淦 髙 够 涫 蚶 琀 涵 毫 淏 菏 盒 涸
		痕 鸿 唿 惚 唬 瓠 扈 淮 患 逭 凰 彗 晦 阍 婚 馄 混 貨 祸 渐 蚵 淶
		淚 涼 猎 淋 淩 渌 淪 麻 麥 曼 猫 袤 捫 猛 眯 猕 谜 覓 密 绵 冕 渑
		喵 描 敏 眸 谋 淖 脲 排 徘 盘 脬 匏 烹 捧 啤 偏 谝 殍 票 貧 萍 颇
		婆 粕 脯 淇 淺 清 渠 渃 深 渖 渗 淑 涮 淞 挲 淌 淘 添 淟 晚 涴 望
		逶 偎 隈 阌 渦 淅 現 鄉 淆 涬 婞 訩 虚 雪 涯 淹 液 淫 渶 淤 雩 魚
		渔 鄅 淯 渊 渚 涿 淄 渍`,

		`鲃 跋 斑 綁 棓 傍 谤 報 悲 琲 辈 備 惫 賁 絣
		琫 逼 皕 赑 詖 愊 愎 弼 编 惼 揙 徧 遍 缏 猋 傧 斌 摒 博 葧 鹁 渤
		跛 補 鈈 瓿 測 滁 湻 渡 發 番 飯 鲂 扉 腓 斐 粪 愤 馮 渢 跗 稃 幅
		腑 赋 傅 富 溉 港 缑 蛤 酣 韩 寒 喊 絎 喝 訶 颌 訸 軤 猢 湖 琥 猾
		滑 換 喚 痪 渙 遑 徨 湟 惶 揮 辉 蛔 惠 喙 缋 渾 耠 惑 禍 湔 溅 渴
		溃 阔 落 傌 買 蛮 鄚 帽 貿 猸 湄 媒 渼 寐 媚 悶 幂 谧 湎 缅 淼 渺
		缈 缗 閔 湣 黽 鄍 谟 蛑 湳 琶 牌 蒎 湃 跑 赔 喷 湓 彭 琵 脾 痞 骗
		評 普 湫 惹 湿 溲 湯 湉 琠 湍 湾 萬 溈 湋 渭 雯 渥 無 喜 隙 閑 湘
		飨 項 亵 渫 雄 溆 渲 湮 欹 絏 颍 湧 游 渝 淵 雲 渣 湛 湞 滞 粥 煮
		滋`,
		`靶 鲅 搬 頒 斒 稖 龅 雹 飽 鲍 鹎 傰 鄙 閟 跸
		嗶 腷 痹 滗 裨 裱 滨 缤 摈 禀 搏 鲌 馎 滄 滀 鹑 滌 滇 蜂 缝 鳧 福
		滏 腹 鲋 缚 滚 傼 颔 頏 嗥 號 滈 阖 嗨 瑚 鄠 嘩 豢 滉 暉 毁 匯 賄
		會 琿 魂 溷 賈 較 粳 鳩 溘 滥 蒗 雷 漓 溧 漣 粱 溜 滦 媽 嗎 獁 满
		鄤 谩 漭 瑁 蒙 盟 腼 鹋 瞄 愍 溟 酩 摸 馍 嫫 貊 貉 漠 寞 溺 滂 辔
		搒 鹏 睥 辟 媲 犏 剽 频 嫔 聘 鲆 溥 溱 裘 溶 溽 滠 飼 溯 綏 溻 滩
		溏 滔 微 溦 雾 溪 携 溴 滟 溢 滢 源 滓`,

		`魃 蝂 蜯 褓 骳 褙 嘣 綳 鼻 碧 馝 弊 幣 稨 褊
		摽 幖 颮 滮 骠 賓 殡 膑 餅 駁 僰 箔 膊 漕 滻 漘 滴 蜚 緋 翡 僨 瘋
		鳳 孵 輔 腐 赙 複 澉 閣 嘏 寡 滾 撖 漢 豪 閡 褐 滹 鹕 滸 滬 瘓 漶
		潢 誨 霁 漸 漤 潋 漏 漉 漯 瑪 嘛 勱 嘜 馒 滿 幔 漫 慢 缦 髦 貌 瞀
		麽 麼 酶 鹛 魅 蜢 艋 嘧 蜜 綿 閩 暝 鳴 摹 膜 慕 暮 漚 膀 蜱 罴 漂
		缥 嫖 嘌 撇 鄱 魄 頗 僕 谱 漆 蜞 颯 滲 漱 霆 網 潍 瘟 聞 舞 熙 餉
		鲞 潇 熊 漩 踅 熏 窨 鲟 演 漾 漪 滎 潆 漁 漳 漲 潴 漬`,

		`澳 罷 瘢 魬 褒 鴇 緥 暴 輩 髲 駜 潷 蝙 編 緶
		膘 麃 憋 瘪 播 撥 餑 踣 餔 潺 潮 澈 澂 澄 蝽 醇 蕩 幡 範 魴 誹 憤
		麩 膚 蝠 幞 撫 駙 賦 蝮 蝜 緱 虢 骸 憨 頜 鹤 踝 鲩 璜 輝 麾 慧 澗
		漿 澆 潔 潰 澜 澇 犛 潦 凛 履 碼 罵 賣 邁 瞒 蝥 霉 緬 緲 廟 緡 憫
		瞑 摩 墨 瘼 潘 醅 賠 霈 噴 嘭 澎 翩 潑 撲 噗 潜 潛 嬈 潤 撒 鲨 潸
		澐 震 潲 澠 澍 澌 潭 漽 潼 潿 嬉 霄 勰 缬 寫 潯 潁`,

		`辦 薄 虣 鮑 憊 糒 甏 觱 鮅 避 嬖 辨 辩 飙 飚
		瘭 褾 儐 濒 馞 鮊 澶 醕 澹 澱 璠 霏 鲱 奮 縫 諷 縛 鲴 盥 還 駭 撼
		翰 頷 澣 憾 翮 醐 寰 缳 擐 噦 諱 澮 閽 餛 諢 霍 濑 澧 濂 霖 潞 螞
		颟 鞔 瞞 螨 蟒 醚 謎 螟 默 謀 霓 凝 濃 螃 耪 膨 駢 蹁 諞 瓢 瞟 瞥
		頻 鮃 璞 氆 憩 霎 潚 濉 燙 靦 閿 禧 憲 獬 邂 廨 懈 興 噱 學 澡 澤
		濁`,

		`癍 幫 謗 鲾 鞞 蹕 髀 斃 濞 臂 鳊 辫 豳 濱 擯
		襏 擘 點 繁 鼢 鲼 糞 黻 襆 黼 鳆 醢 鼾 韓 嚎 濠 鴻 觳 隳 濟 瀞 闊
		濫 蟎 縵 蟊 懑 朦 糜 彌 謐 邈 嬤 貘 嬷 濘 蹒 貔 螵 縹 嬪 皤 濮 濡
		澀 濕 霜 穗 濤 濰 霞 澩 獯 濯`,

		`鼥 奰 鞭 邊 瀌 蹩 殯 臏 鵓 餺 鵏 瀍 闖 瀆 藩
		翻 覆 馥 鹱 鯇 蟪 繢 瀐 濺 獵 瀏 濼 謾 懵 禰 饃 蹣 蟠 蟛 癖 鯆 瀑
		鳍 擾 瀋 霧 瀉 瀅 雜`,

		`瓣 鵯 鞴 襣 襞 臕 鳔 鳖 癟 瀕 髌 醭 簿 鶉 蹯
		鯡 羹 鞲 韝 鯝 瀚 懷 繯 繪 嚯 瀖 瀨 瀝 瀘 饅 鳗 艨 蠓 靡 鶓 鳘 髈
		鵬 羆 騙 蹼 譜 瀜 瀟 瀣 鳕 霪 瀛 瀠 藻`,

		`齙 鰏 躄 鯿 辮 穮 繽 鬓 襮 瀿 瀵 酆 鰒 灌 蠔
		鶘 護 驊 獾 鹹 瀾 瀲 瀧 鬕 顢 鶥 獼 蠛 魔 譬 蠙 瀼 攘 嚷 饗 響 瀷
		瀹 灂`,
		`黯 霸 贔 辯 驃 飆 飈 灃 灏 鶴 轟 歡 露 霹 鼙
		颦 禳 灄 攜 纈 醺`,

		`鷩 鰾 鱉 霽 霾 鰻 蘼 鰵 耱 轡 穰 瓤 灑 灘 鱈`,

		`變 髕 鱝 鬟 黴 鱘`,

		`灞 鬢 灝 襻 躞`,

		`蠻 灣 戅`,

		`灤`,
		"",
		"",
		`鬱`,

		`驫`,

		`灩`,
	}
	//火
	huo = Charset{

		"",
		` 丁 二 了 力 乃`,

		` 彳 大 孓 女 勺 巳 乇 幺 弋 丈 之`,

		` 尺 丑 从 丹 邓 仃 订 斗 队 丰 火 井 仂 历 六
		仑 内 日 太 天 厅 屯 乌 午 爻 长 仉 支 止 中`,

		` 丙 丛 打 代 旦 叨 忉 氐 电 叮 冬 对 尔 叻 乐
		礼 厉 立 辽 另 令 龙 卢 奶 尼 宁 奴 冉 让 他 它 台 叹 讨 田 仝 头
		讯 占 仗 召 只 左`,

		` 犴 场 尘 吃 弛 驰 虫 传 达 当 氘 导 灯 吊 玎
		动 多 夺 阨 耳 旮 亘 光 伙 尖 匠 尽 决 诀 旯 老 玏 耒 肋 利 尥 劣
		刘 甪 伦 论 吕 那 氖 囡 讷 年 农 乓 全 肉 她 饧 廷 同 团 氽 托 驮
		佤 纨 妄 寻 迅 扬 吆 宅 兆 贞 阵 争 执 旨 至 仲 纣 自`,

		` 灿 层 彻 陈 呈 辵 但 岛 低 狄 邸 诋 玓 弟 佃
		甸 阽 盯 町 疔 冻 抖 豆 妒 吨 旰 际 进 灸 抉 吭 来 劳 牢 里 丽 励
		呖 奁 连 良 两 疗 冽 邻 吝 伶 灵 陇 庐 卤 陆 卵 乱 抡 囵 纶 驴 呐
		纳 男 佞 弄 努 求 闰 邰 呔 忐 忑 体 条 听 佟 彤 投 抟 吞 囤 饨 佗
		陀 妥 肟 巫 旸 炀 妖 佁 灾 灶 张 帐 找 折 这 诊 证 志 豸 妐 住 灼
		姊 纵 足`,

		` 哎 佰 宝 采 畅 炒 坼 齿 侈 抶 抽 炊 佽 徂 怛
		妲 甙 岱 迨 绐 担 单 砀 宕 到 的 迪 籴 抵 底 典 店 耵 顶 定 咚 侗
		祋 炖 咄 剁 迩 佴 昉 拐 炅 昊 戽 诙 姐 咎 玦 炕 昆 拉 郎 佬 诔 例
		疠 戾 隶 怜 帘 练 拎 囹 呤 咙 拢 陋 炉 虏 坴 侓 录 轮 罗 侣 免 旻
		明 肭 奈 呶 闹 妮 念 咛 侬 驽 弩 疟 庞 炝 妾 炔 乳 软 朊 侍 沓 抬
		态 肽 贪 昙 帑 弢 屉 忝 佻 迢 帖 图 拖 驼 骀 拓 罔 炜 昕 炎 佯 疡
		易 绎 择 斩 账 胀 招 者 侦 征 郑 诤 知 肢 直 祉 郅 帙 制 质 炙 忠
		肿 驻 转 隹 卓`,

		` 炳 炽 俦 除 耷 哒 玳 带 殆 贷 待 怠 眈 胆 挡
		帝 玷 酊 胨 恫 陡 郖 段 怼 盹 盾 哆 哚 诶 饵 赴 曷 烀 恍 咴 钬 姞
		迹 柬 将 奖 炯 珏 俊 恺 剌 览 烂 类 厘 俚 郦 轹 俪 俐 疬 炼 俩 亮
		咧 临 玲 瓴 珑 胧 娄 轳 胪 孪 骆 律 哪 衲 娜 柰 耐 南 侽 挠 恼 哝
		怒 虐 挪 逄 炮 适 烁 挞 闼 胎 炱 炭 逃 畋 恬 殄 挑 贴 烃 亭 庭 挺
		突 彖 退 挖 歪 弯 挝 显 炫 紃 殃 徉 轺 咬 轶 荧 映 昱 怨 炸 昭 赵
		珍 轸 鸩 政 盅 种 重 轴 胄 炷 祝 追 耔 秭 籽 奏`,

		`龀 逞 骋 哧 鸱 耻 翀 畜 娖 耽 郸 疸 紞 党 捣
		倒 敌 递 娣 刁 爹 瓞 鸫 胴 都 蚪 逗 读 趸 顿 珥 烦 耿 烘 晃 晄 烩
		积 疾 晋 烬 狷 倔 烤 徕 狼 阆 朗 捞 唠 烙 狸 离 骊 逦 哩 娌 珕 猁
		恋 凉 谅 料 烈 赁 凌 陵 留 鸬 辂 赂 栾 挛 倮 珞 捋 旅 虑 耄 拿 孬
		脑 馁 能 娘 恁 脓 诺 恧 衄 哦 秦 逎 热 蚋 偌 晒 晌 烧 晟 恕 朔 趿
		郯 谈 袒 唐 倘 绦 陶 套 特 疼 剔 绨 倜 逖 悌 祧 调 珽 通 透 徒 涂
		鸵 娲 袜 顽 挽 倭 娭 玺 夏 晓 烜 珣 烟 鸯 秧 烊 珧 舀 窈 晔 烨 旃
		盏 展 站 哲 祯 振 朕 值 贽 致 晊 秩 舯 衷 冢 逐 烛 罜 捉 倬 赀 笫
		恣`,
		`欸 皑 捭 蛃 晡 阐 绰 蛏 眵 敕 逴 婥 凑 逮 袋
		聃 掸 啖 惮 弹 蛋 裆 盗 悼 羝 谛 掂 惦 谍 啶 断 惇 掇 敚 舵 鸸 烽
		晗 焓 焊 斛 焕 绩 接 琎 觖 龛 啦 翋 赉 啷 琅 廊 勒 累 梨 犁 理 蛎
		唳 粝 粒 琏 脸 殓 梁 辆 聊 捩 琌 聆 蛉 翎 羚 绫 领 琉 绺 聋 隆 偻
		颅 舻 掳 鹿 逯 鸾 掠 逻 脶 猡 绿 略 焖 捺 赧 捻 您 胬 喏 戚 婼 酞
		探 焘 掏 啕 惕 甜 掭 粜 停 屠 豚 脱 庹 唾 烷 脘 惋 绾 惘 焐 晞 烯
		谑 绽 章 辄 职 鸷 掷 痔 窒 啁 猪 舳 啭 骓 缀 啄 谘 缁 眦 偬`,

		`掰 焙 焯 掣 程 裎 欻 辍 搭 嗒 傣 殚 氮 谠 道
		登 觌 睇 缔 跌 揲 耋 喋 鼎 腚 董 痘 赌 短 缎 敦 遁 惰 焚 缓 践 焦
		嗟 晶 就 厥 焜 喇 腊 睐 阑 揽 缆 稂 痨 愣 鹂 喱 犂 雳 跞 詈 傈 痢
		裢 裣 量 晾 裂 硫 搂 喽 鲁 琭 禄 脔 缕 遖 喃 揇 婻 蛲 猱 辇 傩 晴
		赎 遂 塔 跆 毯 赕 傥 铽 提 啼 鹈 缇 腆 蜓 婷 艇 痛 酡 跎 腕 辋 喔
		窝 幄 循 巽 琰 焰 焱 蛘 谣 媛 暂 蛰 臸 智 痣 骘 彘 惴 缒 琢 辎 粢`,

		`摆 稗 煲 煏 煸 缠 塍 嗤 痴 媸 褚 辏 腠 瘅 亶
		嗲 殿 牒 叠 窦 督 煅 裰 躲 跺 觥 煳 煌 幌 跻 煎 赖 耢 酪 傫 蜊 缡
		粮 趔 零 龄 馏 旒 骝 遛 鲈 路 裸 煤 谬 睦 腩 瑙 鲇 暖 搦 稔 摄 遢
		鲐 摊 痰 搪 慆 腾 誊 裼 阗 龆 跳 嗵 蜕 煺 腽 畹 煨 蜗 瑄 煊 腰 摇
		徭 遥 虞 煜 詹 搌 鄣 障 照 罩 蜇 谪 置 雉 趑 訾 觜`,

		`叆 暧 熬 憏 踌 瞅 褡 靼 瘩 嘀 嫡 骶 端 鲕 裹
		摎 谲 蜡 瘌 辣 谰 罱 螂 嫪 嫘 缧 酹 嘞 璃 嫠 踉 僚 寥 撂 廖 粼 熘
		瘘 骡 摞 雒 缪 鼐 嫩 熔 煽 裳 谭 耥 瑭 韬 慝 舔 蜩 褪 箨 蜿 熄 鄩
		鞅 瑶 毓 摘 辗 獐 彰 嫜 幛 肈 遮 赘 禚 龇`,

		`僾 熛 噌 嘲 撤 撑 踟 憃 踔 龊 骴 撮 鞑 儋 德
		踮 蝶 齑 瑾 撅 噘 獗 褴 黎 鲡 鲤 鲢 撩 嘹 獠 寮 缭 遴 瘤 耧 蝼 撸
		噜 辘 戮 熳 蝻 撵 噢 僻 飘 熵 踏 瘫 羰 膛 躺 趟 滕 踢 题 髫 鲦 璇
		鹞 熠 熨 璋 赭 踬 膣 觯 撰 馔`,

		`薆 瞠 螭 鲷 蹀 憝 踱 燔 斓 懒 擂 罹 魉 燎 辚
		廪 懔 膦 鲮 鹨 癃 窿 氇 瘰 蟆 鲶 耨 燃 燊 燧 獭 螗 糖 醍 蹄 曈 暾
		橐 熹 燚 燠 燏 赠 瘴 辙 褶 鹧 臻 踵 髭 鲻`,

		`戴 黛 蹈 鲽 餲 爵 儡 臁 鹩 磷 瞵 璐 螺 麋 黏
		懦 赡 曙 騃 蹋 醣 螳 嚏 瞳 疃 臀 魍 龌 襄 燮 繇 赢 燥 蟑 螽 擢`,

		`雠 矁 戳 鞮 癜 邋 癞 髅 鹭 鳎 餮 曛 鳐 曜 瞻`,

		`爆 蹭 颤 魑 歠 蹲 蹶 鳓 羸 蠊 酃 蠃 曝 蹻 鼗
		瓆`,
		`嚼 矍 黧 醴 鳞 瓐 糯 鼍 曦 骧 耀 躅`,

		`癫 爝 鳢 躏 曩 鳣`,

		`躐 驎 囊 饕`,

		`攫 麟`,

		`纛 蠹`,

		`鬣 囔 馕 攮`,
	}
	//土
	tu = Charset{
		` 一 乙`,

		` 又`,

		` 个 己 山 土 丸 卫 兀 丫 也 亿 尢 于`,

		` 厄 切 瓦 王 韦 为 卬 夭 以 忆 尹 引 尤 友 与
		予 曰 允`,

		` 艾 凹 鸟 圣 戊 阢 央 叶 永 用 由 右 幼 邘 孕
		仔`,

		` 吖 安 充 地 圪 艮 圭 灰 圾 岌 圹 讴 圮 屺 似
		戍 吐 圩 伟 圬 邬 伍 仵 戌 压 伢 亚 讶 延 闫 羊 阳 爷 伊 衣 圯 夷
		屹 亦 异 因 阴 优 有 迂 纡 屿 宇 羽 约 圳`,

		` 阿 芺 岙 岜 坝 坂 坌 辰 肚 坊 坟 附 岗 坏 矶
		均 坎 坑 块 岚 坜 牡 呕 怄 妑 圻 岐 岍 岖 坍 坛 秃 完 违 围 帏 闱
		纬 位 呜 迕 庑 怃 忤 坞 氙 岘 轩 岈 迓 呀 冶 邺 医 诒 矣 抑 呓 邑
		佚 役 吲 饮 迎 应 佣 甬 攸 忧 邮 犹 卣 佑 玙 欤 余 园 员 远 狁 运
		址 坠`,

		` 肮 垇 坳 坢 衩 诞 坻 坫 岽 矾 废 坩 矸 岣 岵
		画 话 岬 坷 岢 矿 岿 坤 垃 岭 垅 垄 垆 峁 岷 坶 坭 拗 瓯 欧 殴 爬
		帕 怕 坯 坪 坡 坦 坨 宛 往 旺 玮 委 瓮 岫 盱 诩 押 岩 奄 怏 肴 耶
		夜 依 祎 饴 怡 迤 峄 佾 怿 诣 茔 拥 咏 呦 侑 盂 臾 育 鸢 玥 昀 郓
		征 肫`,

		` 哀 按 袄 砭 垞 昶 砗 城 垫 垤 垌 峒 砘 垛 垩
		垡 费 砜 垓 垢 闺 胡 砉 垲 砍 垦 垮 奎 垒 峦 娈 垴 鸥 趴 盆 砒 砌
		牵 峤 窃 埏 垧 哇 娃 威 畏 胃 诬 屋 侮 误 峡 型 勋 峋 鸦 哑 垭 砑
		娅 咽 恹 匽 砚 养 垚 姚 要 咿 咦 贻 姨 舣 音 姻 垠 俑 勇 幽 疣 羑
		囿 宥 诱 舁 禹 垣 爰 院 郧 陨 恽 峥 砖 窀`,

		`啊 埃 唉 挨 砹 爱 俺 胺 案 盎 敖 峬 埕 础 砥
		峨 恶 恩 砝 砩 埂 埚 监 砬 崃 崂 砺 砾 埒 砻 埋 袅 砰 破 埔 峭 容
		砷 砣 诿 娓 軎 翁 唔 捂 娴 顼 埙 鸭 琊 蚜 氩 胭 盐 艳 晏 宴 氧 恙
		眙 胰 酏 益 氤 殷 狺 蚓 狳 馀 谀 彧 眢 鸳 冤 袁 圆 阅 晕 砸 砟 砧
		轾 准`,

		`捱 庵 谙 埯 唵 崩 埠 埭 硐 堆 堕 硌 崮 惯 硅
		崞 黄 基 崨 崛 崆 堀 逵 壸 硭 埝 啪 盘 培 堋 埤 畦 崎 堑 硗 埽 硕
		眭 堂 窕 眺 堍 婉 唯 帷 惟 维 谓 尉 釫 牾 硒 硖 崤 硎 勖 崖 痖 焉
		崦 阉 偃 痒 掖 揶 野 猗 移 痍 埸 逸 翌 寅 隐 婴 庸 恿 悠 蚰 蚴 隅
		域 欲 跃 殒 酝 崭 埴`,

		`隘 啽 媕 揞 晻 媪 奡 傲 奥 堡 堛 嵖 堾 堤 奠
		堞 黑 堠 硷 堪 喹 塄 嵝 嵋 嵚 确 嵘 蛙 崴 琬 琟 嵬 猥 喂 猬 温 握
		硪 痦 婺 骛 翕 遐 翔 硝 揠 腌 蜒 堰 崵 崾 腋 揖 壹 堙 喑 愔 瑛 鱿
		釉 喁 嵛 琙 鼋 援 缘 越 粤 愠 崽 跖 蛭 嵫`,

		`矮 碍 嗳 嗌 嫒 鹌 腤 暗 嶅 遨 嗷 廒 骜 塝 碑
		碚 碜 碘 碉 碇 碓 痱 觟 嵴 跬 碌 硼 碰 嵊 嵩 塑 碎 塌 塘 填 碗 韪
		艉 痿 嗡 蜈 鹉 瑕 睚 衙 鄢 颐 裔 意 鄞 雍 蛹 猷 瘀 瑀 誉 塬 猿 氲
		韫 韵 稚`,

		`瑷 璈 獒 碥 碴 墋 磁 磋 碲 碟 鲑 碱 碣 境 墚
		墁 墙 墒 塾 墅 碳 鲔 稳 寤 鹜 墟 嘘 碹 嫣 旖 夤 撄 嘤 罂 缨 墉 慵
		踊 辕 翟 嶂`,
		`鞍 懊 磅 嶓 嶒 墀 嶝 墩 磙 嘿 蝴 糊 磕 蝰 磊
		嶙 碾 豌 慰 鞋 糈 噎 靥 璎 影 蝣 牖 墺 豫 增 磔`,

		`聱 螯 壁 衡 磺 磨 碛 融 歙 燕 噫 殪 瘾 鹦 壅
		螈`,
		`癌 闇 鮟 醠 謷 磴 礅 壕 壑 礁 嬲 嶷 翳 臆 膺
		臃 黝 龠`,
		`盫 鳌 蹦 璧 礓 礞 黟 彝 癔 鼬`,

		`爊 礤 巅 疆 嬿 繶 鳙`,

		`巉 壤 巍 鼯`,

		`礴 蠡`,

		`鹳 懿 饔`,

		`罐`,
	}
)

var scienceFixList = Charset{
	"",
	"一 乙",
	"二 乃 了 人 入 刀 力 卜 又 几 丁",
	"三 上 下 久 个 丸 乞 也 于 千 大 子 寸 小 山 川 工 己 匀 女 子",
	"四 中 丹 之 予 云 井 亢 介 仁 元 公 切 分 化 午 升 友 及 太 天 夫 少 引 心 户 支 文 斗 斤 方 日 月 木 火 水 比",
	"五 丘 且 世 丙 主 井 仕 仙 代 令 冬 出 加 功 包 北 半 占 卯 右 可 句 叶 古 司 只 召 外 本 巧 巨 市 布 平 弘 弗 必 戊 旦 正 民 永 玉 瓦 甘 生 用 田 由 甲 申 白 目 石 穴 立",
	"六 亘 交 仰 任 仲 伏 仔 光 先 兆 全 共 再 列 印 合 吉 向 后 同 名 宇 存 安 字 守 州 帆 年 旭 早 有 求 百 弛 竹 米 羊 羽 臣 自 至 舟 行 衣 西 回 如 充 成",
	"七 亨 吾 均 坐 壮 声 妙 孝 宏 局 希 序 志 戒 改 更 杏 材 村 位 佑 作 伯 伴 体 余 克 兑 兵 初 判 利 助 告 君 步 江 汗 汝 池 秀 究 良 见 言 彤 谷 豆 赤 车 辰",
	"八 并 事 亨 京 依 佳 侃 供 侍 使 佩 来 例 免 雨 其 具 典 冒 冽 函 刻 刷 刹 制 到 效 协 卓 卷 取 受 和 周 命 固 坤 垂 坦 坡 夜 奇 奈 奉 姑 始 妹 枚 板 林 欣 武 汲 决 沙 汰 冲 沛 沐 沃 汪 炎 炊 版 物 牧 玖 的 直 盲 知 社 空 究 舌 虎 采 金 长 昔 明 旺 服 朋 杭 果 枝 松 扭 东 门 青 季 孟 宜 官 宗 宙 定 尚 居 岳 岸 岩 岱 幸 庚 店 府 弦 征 彼 往 快 忽 忠 念 或 所 房 技 承 折 扶 政 放 齐 于 昂 昆 昌 升 昊 闰 兔",
	"九 亭 亮 系 侠 信 俊 保 便 侣 俞 冒 冠 克 前 则 劲 勉 勃 勇 南 厚 叙 咸 哄 品 垠 奎 奏 威 姻 姬 姜 妍 姿 客 宣 室 屋 巷 帝 幽 度 回 建 彦 待 律 思 性 易 招 拓 折 拜 抱 施 映 昨 是 春 星 昭 架 柯 查 柴 柔 柘 韦 柱 柏 柄 柳 段 油 泳 沿 河 况 注 泉 治 波 泊 法 冷 炳 帅 甚 界 皆 皇 盈 看 相 眉 祈 科 秋 秒 穿 突 竿 红 罕 美 耐 肖 衍 表 要 计 订 贞 军 重 门 面 革 音 风 飞 食 首 香 姣",
	"十 乘 倚 幸 仓 修 借 值 倍 仿 表 伦 党 兼 倡 刚 原 员 哥 唐 哲 城 夏 娥 宴 家 宫 宰 容 射 展 峡 岛 峰 师 席 库 航 般 芽 芹 花 芝 芳 娘 袁 衿 活 洪 洲 泰 洗 洞 派 洋 流 烘 烈 特 珂 珊 珍 玲 益 真 词 神 祝 组 祚 秦 秤 租 秘 并 竟 级 纱 纯 素 纳 纽 纺 纹 翁 者 耘 耿 育 股 皋 座 庭 径 徐 恩 恭 恢 恒 恤 息 恬 扇 拾 持 效 料 旁 旅 晏 晃 时 晋 书 朔 校 格 桂 根 栖 桃 桐 钊 殊 气 记 训 讨 豹 贡 财 起 马 支 高 娟 倩 娜 笔",
	"乾 伟 偕 健 偶 侧 停 侦 富 凰 剩 副 勘 动 务 区 卿 参 唯 启 商 唱 珠 般 产 皎 眼 婉 研 祥 移 竟 章 笛 伏 笙 弦 紫 绅 绍 绊 累 罩 习 翌 者 聊 胡 教 敏 斌 斜 旌 旋 族 晤 晨 晚 曹 望 朗 梗 梧 梓 梅 梨 梁 毫 球 海 浩 涉 涌 浴 烽 爽 崇 国 基 坚 执 堂 培 寅 寄 宿 寂 密 尉 寻 将 专 崔 巢 常 带 康 张 彩 彬 从 悦 悟 戚 挺 英 婕 若 苔 苗 茂 术 袖 许 责 赦 近 闭 雪 顷 顶 鹿 麦 扇 掘 卷 扫 舍 掌 迫 贰 量 开 闲 间 添 焰 无 为 犁 猛 球 麻 胜 佩 焕 盛",
	"仅 债 杰 催 惠 伤 舒 倦 传 勤 势 募 嗣 园 圆 块 干 廊 渡 湃 渺 照 煎 媒 炼 爷 琴 琢 琵 琶 棱 督 睦 碇 禁 禄 禽 稚 绢 义 圣 廉 巢 微 爱 意 惮 荣 雅 顺 翔 感 愚 想 愉 愈 斟 暑 会 极 楚 楠 枫 椰 榆 殿 景 钞 钦 发 农 退 乃 郊 电 佃 附 雌 雉 聘 肆 琛 唇 脱 台 获 莓 莫 蜀 街 词 衙 裟 装 裕 里 解 咏 夸 詹 资 迹 跳 路 载 浅 淘 淡 涯 涵 混 深 淑 清 饮 驯 驰 普 创 诏 雄 博 弼 智 贺 凯 等 筏 绞 给 绚 絮 绝 统 络 能 舜 砚 稀 稍 税 程 窗 竣 童 策 答 筑 筒 棠 栋 棒 棉 款 证 净 媛 登 授 捷 现 理 番 强 闵 雁 集 云 项 须 黑 备 接 注 评 象 贵 贴 贸 越 超 迪 茹 茫 众 传 割 劳 喜 乔 善 敢 散 敦 斑 斯 晶 晴 晰 最 替 棋 棍 栈 森 植 荒 草 曾 单 喻 围 堪 尧 场 堤 报 堡 媒 媚 寒 寓 寻 尊 岚 巽 帽 几 复 惟 情 荀 茜 茶",
	"团 图 寿 梦 奖 察 实 对 伪 侨 像 佟 煌 仆 僚 崭 廓 彰 愿 慈 慎 态 搬 业 温 港 渠 湖 湘 测 汤 滋 支 犒 猿 狮 瑟 鼓 监 尽 经 种 称 竭 端 个 算 精 紧 绰 绶 综 绯 绵 维 纶 绫 置 靖 晖 裳 暄 铃 群 郎 铅 阁 韶 领 饰 饱 仿 魁 魏 诗 试 询 诠 援 挥 扬 新 铁 莆 莎 莉 鸣 鼻 齐 瑛 碌 椿 琳 杨 虞 当 盟 酩 雷 颂 顿 预 鼎 鼓 雍 渊",
	"尽 俭 僻 剧 剑 劈 啸 豪 宽 番 寮 履 帜 广 熙 弹 影 微 彻 慰 慷 栗 醉 锐 锄 锋 铮 阅 院 阵 宵 霆 霈 落 蝶 冲 褓 复 课 谈 调 谅 敲 溪 源 溶 赐 质 赏 卖 趣 践 辈 轮 替 游 进 邮 部 醇 确 磁 磐 稼 稿 稻 期 朝 溢 沟 穷 箱 节 箭 范 篇 糊 纬 缘 缄 绪 线 缔 编 练 暑 义 铺 馆 苇 叶 葛 葵 管 著 虑 掴 摧 摩 数 暂 暴 暮 概 乐 槽 樟 枢 标 模 样 楼 欢 毅 演 汉 境 渐 涨 滞 漫 满 洋 熟 热 荧 瑶 玛 几 皓 盘 驾 驻 魄 鸦 华 惯 嘉 碧 樊 蒂 块 发 葆 渔 漆 纲 尝 彰 志 赫 辅 造 逍 速 逞 途 透 通 逢 连 静 萤 阴 银 铜 铭 齐 黄 旗 畅 槐 枪 歌 瑞 瑚 廖 硕 祯 福 翠 翡 台 与 舞 菊 果 菜 蜻 蜜 裙 诫 诰 诲 诚 誓 说 诞 认 貌 赈 宾 轻 赵 酸 宁 瑗 榕 诱 玮",
	"仪 仅 万 冀 剑 进 器 喷 论 坛 壁 奋 道 岭 憬 憧 抚 怜 战 撮 撤 撰 幢 播 扑 整 德 赓 晕 厉 机 陆 陵 鞘 头 余 默 桦 横 桥 橇 树 樽 橙 竖 洁 润 泄 贤 增 郭 赋 谷 莹 刘 烨 烧 燃 炖 磷 燎 芦 穆 窥 筛 筑 糖 县 罢 举 苍 蒸 席 辉 震 墨 娇 慧 萱 董 卫 衡 亲 谓 谒 诚 谏 谚 诸 豫 蹄 辑 办 运 遇 遂 道 达 都 醒 皓 锡 钱 总 橡 震 敬 磊 庆",
	"儒 优 赏 壕 壑 岳 应 撼 擒 擎 检 操 擅 择 擂 敛 檄 檀 褒 讲 谦 谢 豁 趋 龙 融 远 乡 键 针 钟 锻 阶 队 阳 馆 骏 鲜 点 齐 鸿 荫 襁 激 浓 颖 翰 致 雕 营 灿 烛 燧 微 响 独 瞰 瞬 禅 簇 纵 缝 声 聪 聊 临 艰 参 蔗 谅 蔬 钢 赛 璜 燃 兴 学 遥 晓 霖 澄 潮 潜 潭 蓉 蓄 茜 颖 璇 蓓 陶 陈 谘 璋 逸 霓 谋 颐 锦",
	"骏 澡 泽 黛 澧 澳 澹 激 逢 总 授 逊 斋 遣 忆 憾 懂 螺 恳 懈 应 懋 鲜 繁 韩 戴 搁 拟 擦 断 曜 曛 曙 涛 滨 爵 获 狞 猎 瞻 简 馈 箫 蔺 缮 翻 异 职 旧 荫 蕊 荡 蕃 蝉 声 讴 转 遭 适 鄙 医 锁 镇 锤 镰 聂 鸡 霜 离 雏 额 频 题 骑 鹃 灿 鞭 碧 霞 蒲 慕 鞠 隆",
	"劝 庐 扩 攀 旷 泻 溅 瀑 镜 关 雾 愿 类 鲸 鹊 鹏 麓 兽 猎 祷 稳 获 环 薄 绳 臆 臂 膺 薪 蔷 薇 襟 证 赞 赠 遵 郑 丰 礼 蕙 织 谨 济 归 濠 阔 湿",
	"还 释 钟 阐 露 飘 馨 璃 龄 宝 迈 怀 悬 胧 沥 献 琼 砾 籍 筹 篮 继 办 麒 辞 藏 萨 籍 薯 觉 触 译 议 警 赢 面 瀚 瀛 烁 蕾 选 辽 韵 绣",
	"翻 饶 驱 鸡 傈 属 巍 续 缠 腊 护 誉 贴 轰 辩 随 隐 霸 竞 耀 宝 识 篷 莲 罗",
	"顾 艺 俨 巅 摄 权 欢 灌 叠 穰 笼 听 澡 苏 览 赞 读 边 乡 餐 须 樱 铁 铎 栏 莺 鹤 锈",
	"岩 恋 织 藓 变 矿 显 驿 验 髓 体 乐 龚 懿 鉴 芦",
	"矗 罐 艳 燕 禳 酿 炉 陇 谒 灵 鹰 鑫 兰 麟",
	"听 篱 蛮 观 才 灵",
	"湾 瞩 赞 逻 爵 厌",
	"锣 銮 缆",
	"艳 欢 鹦",
	"麓",
	"鸾", //29
}

var nameList = Charset{
	``,
	``,
	`卜丁刀七力刁`,
	`于上山干士子千弓万`,
	`孔毛王文主尤牛尹元卡支巴仇戈公勾木水火井太`,
	`石央甘田白申包丘皮平令左右冉史世可由正名以丙玉布目仙市巨司召代弘`,
	`朱牟伊任伍米安羊全伏戎后百吉年向同匡有仲仰光自列老多羽守州印共危`,
	`李吴宋杜江何吕余佘辛谷巫车成利甫池岑系杞良求我伯言吾汝束里豆希贝冷而别步`,
	`岳宗沈卓狄屈杭牧居武幸宓尚明始长昌儿征析庚沙东汲沓帛虎知京念来委金孟季林易官扶和汪竺沃松艾於房祁`,
	`俞施柯段涂姚姜柴纪韦查侯柳风封秋咸竽柏羿禹南胥约勇河法革眉后计冠泰宦姬昭宣相红`,
	`花徐孙祖凌席班乌贡宫家祝桂唐真师宰起修留马恭轩容秘索铁桓仓桃展桐原肥洛袁秦柏奚倪时高夏洪翁益桑耿殷晁`,
	`张许梅章胡梁康范曹麦从崖那崔邢商寇苗尉英习鹿常崇国庸坚密涂假宿鱼符茅麻苟浦扈终巢`,
	`黄曾邵邱彭傅程阮项童贺乔富荆堵盛景荀闵喻云费焦舒理尧舜雄渊惠贵敦朝开冯单能强越嵇须邰茹钮`,
	`杨庄游雍贾雷莫虞楚温汤路裘衙督睦义新禄干郁嵩琴钳涂稠椿农经解湛`,
	`连廖熊赫郗管赵辈齐郎寿荣台逢造端辅通翟僮源闻韶凤慎郝郜臧`,
	`刘郭叶欧董葛万乐谈厉黎满鲁贤闾养樊墨陕颉谅广审挚`,
	`陈陶陆潘蒲记赖诸阎鲍骆钱龙运稽锡都衡颖桥噪声燕融穆`,
	`临励赛营襄优鸿翼辕邬蔡蒋蔚邹谢韩阳应龙钟`,
	`颜魏简阙聂丰储戴礼环浇济瞿庞隗`,
	`郑萧邓薛关谭祢薄`,
	`罗严钟释蓝党窦怀籍`,
	`巍顾珑饶铁`,
	`苏袭权蔺`,
	`兰体麟滩铄鳞栾严洒`,
}

func InitAll() {
	initCharFrom5156()
	updateStrokes()
	updateFivePhase()
	fixNameStrokes()
}

func initCharFrom5156() {
	list := excavator.Get("http://xh.5156edu.com")
	debug.Println("initCharFrom5156 start")
	log.Println(list)
	for k, v := range list {
		insertCharacter(v, k)
	}
	debug.Println("initCharFrom5156 end")
}

func insertCharacter(charset Charset, k string) {
	debug.Println("start: ", k)
	for _, v := range charset {
		s := strings.Split(v, "|")

		s0, _ := strconv.Atoi(s[0])
		if len(s) != 3 {
			debug.Println("wrong:", s)
			continue
		}
		c := model.Character{
			SimpleStrokes: s0,
			SimpleChar:    s[1],
			Pinyin:        s[2],
			Radical:       k,
		}
		i, e := c.Create()
		if e != nil {
			debug.Println(i, e)
		}
	}
	debug.Println("end", k)

}

func updateStrokes() {
	debug.Println("updateStrokes start")
	for k, v := range scienceFixList {
		ScienceUpdate(k, v)
	}
	debug.Println("updateStrokes end")
}

func ScienceUpdate(index int, s string) {
	debug.Println("science:", index)
	sa := strings.Split(s, " ")
	for _, v := range sa {
		if v == "" {
			continue
		}
		_ = v

		c := model.Character{SimpleChar: v}
		c = *c.Get()
		c.ScienceStrokes = index
		if uuid.FromStringOrNil(c.Id) == uuid.Nil {
			debug.Println("notfound: ", c, v)
			c.Create()
			continue
		}
		c.Update()
	}

}

func updateFivePhase() {
	log.Println("555")
}

func fixNameStrokes() {
	for k, v := range nameList {
		v := strings.Split(v, "")
		for _, value := range v {
			c := &model.Character{SimpleChar: value}
			c = c.Get()
			c.ScienceStrokes = k
			c.Update()
		}
	}
}
