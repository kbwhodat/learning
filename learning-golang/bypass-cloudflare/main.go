package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/utils"
	// "github.com/go-rod/rod/lib/proto"
	"fmt"
)

func init() {
	launcher.NewBrowser().MustGet()
}

func main() {
	// This is the tricky part, it's the same as this https://ttt.musc.edu/knowledge-base/how-to-enable-and-disable-the-chrome-pdf-viewer
	// If we don't set this, the pdf will be opened by the pdf viewer not the downloader.
	// The pdf viewer is not visible for the cdp protocol.
	// utils.OutputFile("tmp/test/Default/Preferences", `{
	// 	"plugins": { "always_open_pdf_externally": true }
	// }`)

	// u := launcher.New().
	// 	Set("disable-blink-features", "AutomationControlled").
	// 	Set("disable-web-security", "").
	// 	Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36").
	// 	Set("window-size", "0,0").
	// 	UserDataDir("tmp/test").
	// 	Headless(true).MustLaunch()
	u := launcher.NewUserMode().
		Leakless(true).
		UserDataDir("tmp/t").
		Set("disable-default-apps").
		Set("no-first-run").
		Headless(true).
		MustLaunch()

	// browser := rod.New().ControlURL(u).MustConnect()
	target := "https://en.rarbg-official.com/episodes/true-detective-2014-season-4-episode-1"

	page := rod.New().ControlURL(u).MustConnect().NoDefaultDevice().MustPage(target)


	// wait := browser.MustWaitDownload()



	// page := browser.NoDefaultDevice().MustPage(target)
	// page.MustWaitLoad()
	utils.Sleep(10)

	// f := page.MustElement("iframe").MustFrame()
	// p := page.Browser().MustPageFromTargetID(proto.TargetTargetID(f.FrameID))
	// p.MustWaitStable()
	// p.MustElement("input[type=checkbox]").MustClick()

	// utils.Sleep(2)
	html, _ := page.HTML()
	fmt.Println(html)
	// utils.OutputFile("test.pdf", wait())
}
