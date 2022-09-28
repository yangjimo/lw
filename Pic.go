package lwapi

//
//  CmpPic
//  @Description: 在绑定的窗口的指定点上（x,y）截取一张与提供的待对比图片等尺寸的图像，然后与提供的待对比图片做相似度对比，返回两张图片的相似度，返回值是一个0到100之间的整数，100表示完全一样，0表示完全不一样。
//   x 截取区域的左上X坐标
//   y 截取区域的左上Y坐标
//   pic 待对比图片,只能是单张图片，比如"test.bmp"，也可以是缓区中的图片，如“10000”
//   int  返回两张图片的相似度，返回值是一个0到100之间的整数，100表示完全一样，0表示完全不一样。
//
func (com *LwSoft) CmpPic(x int, y int, pic string) int {
	var ret, _ = com.lw.CallMethod("CmpPic", x, y, pic)
	return int(ret.Val)
}
