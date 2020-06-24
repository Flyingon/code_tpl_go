package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func parseHtml(htmlBody string) error {
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return err
	}
	fmt.Println(doc.Attr)
	return nil
}


func parseContentWithNovel(htmlBody string) (isXiaoShuo bool, xiaoshuoData map[string]interface{}, err error) {
	var xiaoDataBase64 string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "proteus" {
			for _, a := range n.Attr {
				if a.Key == "type" && a.Val == "bookcard" {  // 小说标签: <proteus type="bookcard" </proteus>
					isXiaoShuo = true
				}
				if a.Key == "data" {
					xiaoDataBase64 = a.Val
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return
	}
	f(doc)

	if xiaoDataBase64 == "" {
		return
	}
	xiaoDataJson, err := base64.StdEncoding.DecodeString(xiaoDataBase64)
	if err != nil {
		return
	}
	err = json.Unmarshal(xiaoDataJson, &xiaoshuoData)
	if err != nil {
		return
	}
	return
}


func main() {
	//htmlBody := "<p>宋祁和他哥宋庠，早年家境不好，吃了很多苦。还曾发生过没有酒肉，无法过年的事。最后，两兄弟只能把祖传宝剑剑鞘上的裹银拿去当了，换成钱，买了酒肉回家过年。</p><p>于是，宋祁和哥哥宋庠寒窗苦读，终于一跃龙门，同榜考中了进士。宋祁比哥哥更高一筹，考中状元。但皇太后觉得这样不妥，她说：“自古以来，长幼有序，哪有弟弟排到哥哥前面的道理？”</p><p>结果，宋庠被定为状元，宋祁成了第十名，不过世人还是称他们兄弟俩为<strong>“双状元</strong>”。</p><p>入了仕途，苦尽甘来，一切都好了起来。弟弟宋祁好客，经常在家里大宴宾客，夜夜笙歌。主人和宾客相对饮酒，观看歌舞，忘记了时间。等到揭开窗幕，才发现到了第二天早上。所以宋祁的府邸得了一个<strong>“不晓天</strong>”的名号。</p><p>而哥哥宋庠呢？即使在最热闹的上元节，他也不去看花灯，而是夜里跑到冷寂的书院里，一个人静静地读《周易》。他听说弟弟点上华灯，被歌妓簇拥着，喝酒通宵达旦时，很是痛心，第二天便写了书信去斥责他：“闻昨夜烧灯夜宴，穷极奢侈，<strong>不知记得某年上元同在某州州学内吃薤煮饭时否？</strong>”</p><p>薤是一种野菜。两兄弟当年负笈求学困窘之时,没少吃这种粗糙食物。宋庠是想告诉弟弟，不要忘记过去吃糠咽菜的苦日子，要惜福，不要过分铺张奢华。</p><p>宋祁收到信后却笑了。他回信给哥哥说：“不知某年同某处吃薤煮饭是为甚底？”弟弟的意思很明白，当年忍饥挨冻，孜孜苦读，不正是为了今天的享乐吗？</p><p>宋祁的这句话，很容易让人产生不好的联想。曾几何时，享乐就一直和罪恶、堕落联系在一起。官员的腐败，和贪图享乐有关；一支队伍人心涣散，和贪图享乐有关；甚至许多王朝的倾覆，也被认为是贪图享乐所致。因此，享乐在公开的意识形态里仿佛是洪水猛兽，有人避之犹恐不及，有人虽“心向往之”，但至少也会做足表面文章。</p><p>但宋祁却大张旗鼓向世人宣告自己的喜好，自然也招致一些人的不满。</p><p>有一次，宋仁宗想外放宋祁到四川成都做一把手，宰相陈执中表示反对：“成都是一个喜欢享乐的城市，宋祁是一个喜欢吃喝玩乐的人。他去管理成都，恐怕会忘乎所以，耽误公务，不太合适。”但皇帝仍然批准他上任了。</p><p>到了成都，宋祁果然不负皇帝期望，如鱼得水，积极倡导享乐，也大大提振了当地的饮食娱乐业。</p><p>宋祁不仅自己带头吃喝玩乐，给官员百姓起到榜样作用，而且还自创很多新玩法、新项目，让大家学习。他在蜀中大倡游宴，用实际行动推动当地经济文化的发展。成都作为享乐之都的名号，愈发响亮了。宋祁离开成都后，他的影响力依然巨大，后来历届成都太守也都效仿他，主持并带头游宴。</p><p>在享乐之余，宋祁遍访民间，实地考察，写下一本极具史料价值的<strong>《益部方物略记》</strong>。这是一部优秀的古代生物学典籍，详细描述了成都及附近地区的鸟兽、草木等，共记载动物15种，植物50多种，许多是前人尚未记录过的。除此之外，此书还以专业角度介绍了四川人做菜的烹饪原料、技艺等。有专家指出，川菜能够成为“八大菜系之一”，《益部方物略记》有非常之功。</p><p>不享乐，毋宁死，这恐怕是宋祁一生追求的。而他也明白，享乐不能逾越自己的底线。他喜欢享乐不假，但并没有贪污腐败，与民争利。</p><p>在成都任上，宋祁也不是一味玩乐，不理政务。他给成都人民做了许多大好事、大实事。以至于去职时，老百姓不舍得他走。他去世时，成都百姓哭于其祠者，有数千人之多。</p><p>宋祁也没有因为贪图享乐使学术荒疏。他年少成名，中年更是精进向学。在大儒巨贤辈出的北宋，他被皇帝钦点去编撰<strong>《新唐书》</strong>，这很不容易。《新唐书》另一位主要编撰者，是我们都很熟悉的、鼎鼎大名的欧阳修。</p><p>修《新唐书》10多年，宋祁平时出入，总是携带纸笔，随时记录。在成都，他吃喝不辍，修史也没有耽误过。他每晚开门垂帘工作到深夜，在两柱巨大的灯烛下，侍女丫鬟环绕身边，帮他研墨伸纸，远近都知道是他在编修《唐书》，<strong>看上去像神仙一般</strong>。</p><p>宋祁工作也不忘记享乐。他有大才，但他的享乐作风还是影响到了自己的仕途发展。朝廷曾想提拔他做中央最高财政长官，包拯表示反对，说他生活奢靡，不能担此重任。朝廷最后采纳了包拯的意见。</p><proteus class=\"addons-place\"  type=\"bookcard\" data=\"eyJhcnRfaWQiOjQ4OTIwMSwiYXJ0X3RpdGxlIjoi56WB5bidIiwiYXJ0X2F1dGhvciI6IuWQkeW3pueci+acneWPs+i1sCIsImFydF9waWMiOiJodHRwczovL2Njc3RhdGljLTEyNTIzMTc4MjIuZmlsZS5teXFjbG91ZC5jb20vYm9va2NvdmVyaW1nL2NvdmVyLzIwMTItMDktMDEvNjg4ZTI5YzRkY2Q0ZjAwNTM2MTU0NzA3NjYwYTQyZDEuanBnIiwiYXJ0X2RldGFpbCI6IiIsImtkX3RvcGljaWQiOiIiLCJjYXRlZ29yeV9pZCI6MjAwNTAsImNhdGVnb3J5X25hbWUiOiLmuLjmiI8iLCJ1bnNoZWxmIjowLCJzb3VyY2UiOjAsInB1Ymxpc2hfdGltZSI6IjIwMTktMDctMTkgMDY6MTM6MzYiLCJ1cGRhdGVfdGltZSI6IjIwMjAtMDMtMjMgMjA6MDQ6NDMiLCJjYmlkIjozNjQyODk4NzAzMzM5MzAyLCJjc2JpZCI6NDg5MjAxLCJhZHZmcmVlc3RhdHVzIjoxfQ==\"></proteus>"
	//isXiaoShuo, xiaoshuoData, err := parseContentWithNovel(htmlBody)
	//fmt.Println(isXiaoShuo, xiaoshuoData, err )
	parseHtml("书籍UIC<em>测</em><em>试</em>")
}
