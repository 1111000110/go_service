package controller

import "github.com/gin-gonic/gin"

type data struct {
	Image_src     string `json:"image_src"`
	Open_type     string `json:"open_type"`
	Goods_id      int    `json:"goods_id"`
	Navigator_url string `json:"navigator_url"`
}

func GetSwiperList(ctx *gin.Context) {
	var list []data
	list = append(list, data{
		Image_src:     "https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/f96455c234fb063560c5acd314838de0.jpg?thumb=1&w=2452&h=920&f=webp&q=90",
		Open_type:     "navigate",
		Goods_id:      129,
		Navigator_url: "/pages/goods_datail/main?goods_id=129",
	})
	list = append(list, data{
		Image_src:     "https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/48027a1ab09d142a08ae0eb884680bf1.jpg?thumb=1&w=2452&h=920&f=webp&q=90",
		Open_type:     "navigate",
		Goods_id:      395,
		Navigator_url: "/pages/goods_datail/main?goods_id=395",
	})
	list = append(list, data{
		Image_src:     "https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/105e5ca001ebac00630dcbea2f1182bf.jpg?thumb=1&w=2452&h=920&f=webp&q=90",
		Open_type:     "navigate",
		Goods_id:      38,
		Navigator_url: "/pages/goods_datail/main?goods_id=38",
	})
	ctx.JSON(200, list)
}
