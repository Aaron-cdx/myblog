/**
 * @Author: caoduanxi
 * @Date: 2021/3/17 16:46
 * @Motto: Keep thinking, keep coding!
 * 关于我的模块about me
 */

package controllers

type AboutMeController struct {
	BaseController
}

func (c *AboutMeController) Get() {
	c.Data["wechat"] = "微信: 15779236476"
	c.Data["qq"] = "QQ: 1315020626"
	c.Data["tel"] = "Tel: 15779236476"
	c.TplName = "aboutme.html"
}
