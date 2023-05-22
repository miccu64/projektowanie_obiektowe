package utils

import (
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func ProceedTest(callback func(driver selenium.WebDriver)) {
	service, err := selenium.NewChromeDriverService("./chromedriver", 4444)
	if err != nil {
		panic(err)
	}
	defer func(service *selenium.Service) {
		if service.Stop() != nil {
			panic(err)
		}
	}(service)

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"window-size=1920x1080",
		"--no-sandbox",
		"--disable-dev-shm-usage",
		"disable-gpu",
		"--headless",
	}})

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		panic(err)
	}

	callback(driver)
}

func TryGetSite(driver selenium.WebDriver, url string) {
	err := driver.Get(url)
	if err != nil {
		panic(err)
	}
}
