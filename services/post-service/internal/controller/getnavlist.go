package controller

import "github.com/gin-gonic/gin"

type getNavListHandlerdata struct {
	Image_src     string `json:"image_src"`
	Open_type     string `json:"open_type"`
	Name          string `json:"name"`
	Navigator_url string `json:"navigator_url"`
}

func getNavList(ctx *gin.Context) {
	var list []getNavListHandlerdata
	list = append(list, getNavListHandlerdata{
		Image_src:     "https://s1.aigei.com/src/img/png/53/53b8e9cae4f34ba58c8933b7f0df562e.png?e=1735488000&token=P7S2Xpzfz11vAkASLTkfHN7Fw-oOZBecqeJaxypL:EbBJtQGfe9qGhwRkCUp40CTD46M=",
		Open_type:     "swtichTab",
		Name:          "分类",
		Navigator_url: "/pages/category/main",
	})
	list = append(list, getNavListHandlerdata{
		Name:      "秒杀拍",
		Image_src: "https://s1.aigei.com/src/img/png/aa/aa60e74b4b584b5c8fcad445b98ce77a.png?imageMogr2/auto-orient/thumbnail/!282x282r/gravity/Center/crop/282x282/quality/85/%7CimageView2/2/w/282&e=1735488000&token=P7S2Xpzfz11vAkASLTkfHN7Fw-oOZBecqeJaxypL:FT_STsLhn_LHJgnQ7TubAbLABS0=",
	})
	list = append(list, getNavListHandlerdata{
		Name:      "超市购",
		Image_src: "https://s1.aigei.com/src/img/png/d4/d41719edad3647439c1342e9a1db3c1c.png?e=1735488000&token=P7S2Xpzfz11vAkASLTkfHN7Fw-oOZBecqeJaxypL:szNuUciLpEdB5gwpMHKjSH2du7g=",
	})
	list = append(list, getNavListHandlerdata{
		Name:      "母婴品",
		Image_src: "https://s1.aigei.com/src/img/png/fb/fb3d1f4b16cb435f9ea722c41416cde0.png?imageMogr2/auto-orient/thumbnail/!282x282r/gravity/Center/crop/282x282/quality/85/%7CimageView2/2/w/282&e=1735488000&token=P7S2Xpzfz11vAkASLTkfHN7Fw-oOZBecqeJaxypL:AH2NEgiYBiHZao1zj5YBq84meVo=",
	})
	ctx.JSON(200, list)
}
