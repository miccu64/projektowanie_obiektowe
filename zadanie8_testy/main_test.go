package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"
	"testing"
	"time"
	"zadanie8_testy/utils"
)

func TestGetGoogle(t *testing.T) {
	test := func(driver selenium.WebDriver) {
		utils.TryGetSite(driver, "https://www.google.com")
		pageName, _ := driver.Title()

		assert.Contains(t, pageName, "Google")
	}
	utils.ProceedTest(test)
}

func TestSearchInBing(t *testing.T) {
	test := func(driver selenium.WebDriver) {
		utils.TryGetSite(driver, "https://www.bing.com")
		input, _ := driver.FindElement(selenium.ByID, "sb_form_q")
		_ = input.SendKeys("SzukanaFraza")
		_ = input.Submit()
		currentUrl, _ := driver.CurrentURL()

		assert.Contains(t, currentUrl, "SzukanaFraza")
	}
	utils.ProceedTest(test)
}

func TestGetWeatherForecast(t *testing.T) {
	test := func(driver selenium.WebDriver) {
		utils.TryGetSite(driver, "https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&current_weather=true&hourly=temperature_2m")
		html, _ := driver.PageSource()

		assert.Contains(t, html, "temperature")
	}
	utils.ProceedTest(test)
}

func TestGetOnetCheckAnyImagePresent(t *testing.T) {
	test := func(driver selenium.WebDriver) {
		utils.TryGetSite(driver, "https://onet.pl")
		elements, _ := driver.FindElements(selenium.ByTagName, "img")

		assert.Condition(t, func() (success bool) {
			return len(elements) > 0
		})
	}
	utils.ProceedTest(test)
}

func TestGetNonExistingPageThrowsError(t *testing.T) {
	test := func(driver selenium.WebDriver) {
		err := driver.Get("https://djasyd78asyhjkdhas79hddjas287h.pl")

		assert.Error(t, err)
	}
	utils.ProceedTest(test)
}

func TestUsosApiShouldDenyWithoutToken(t *testing.T) {
	test := func(driver selenium.WebDriver) {
		utils.TryGetSite(driver, "https://apps.usos.uj.edu.pl/services/tt/student")
		html, _ := driver.PageSource()

		assert.Contains(t, html, "method_forbidden")
	}
	utils.ProceedTest(test)
}

func TestGetNonExistingSiteFromWorkingServer(t *testing.T) {
	test := func(driver selenium.WebDriver) {
		utils.TryGetSite(driver, "https://www.mozilla.org/pl/dasdasdas")
		pageTitle, _ := driver.Title()

		assert.Contains(t, pageTitle, "404")
	}
	utils.ProceedTest(test)
}

func TestNavigateToRandomPageByClicking(t *testing.T) {
	test := func(driver selenium.WebDriver) {
		firstUrl := "https://stackoverflow.com/questions/6509628/how-to-get-http-response-code-using-selenium-webdriver"
		utils.TryGetSite(driver, firstUrl)
		elementToClick, _ := driver.FindElement(selenium.ByXPATH, "//a[contains(@href,'https')]")
		_ = elementToClick.Click()
		currentUrl, _ := driver.CurrentURL()

		assert.NotEqual(t, currentUrl, firstUrl)
	}
	utils.ProceedTest(test)
}

func TestGetOnlyImage(t *testing.T) {
	test := func(driver selenium.WebDriver) {
		imageUrl := "https://www.google.com/images/branding/googlelogo/2x/googlelogo_light_color_92x30dp.png"
		utils.TryGetSite(driver, imageUrl)
		html, _ := driver.PageSource()

		assert.Contains(t, html, imageUrl)
	}
	utils.ProceedTest(test)
}

func TestGetPageAndChangeIt(t *testing.T) {
	test := func(driver selenium.WebDriver) {
		firstUrl := "https://www.google.com"
		secondUrl := "https://www.onet.pl"
		thirdUrl := "https://stackoverflow.com"
		utils.TryGetSite(driver, firstUrl)
		utils.TryGetSite(driver, secondUrl)
		utils.TryGetSite(driver, thirdUrl)
		currentUrl, _ := driver.CurrentURL()

		assert.Contains(t, currentUrl, thirdUrl)
	}
	utils.ProceedTest(test)
}

func TestCheckStatusReadyOfDriverAfterChangeSite(t *testing.T) {
	test := func(driver selenium.WebDriver) {
		firstUrl := "https://www.google.com"
		secondUrl := "https://www.onet.pl"
		utils.TryGetSite(driver, firstUrl)
		utils.TryGetSite(driver, secondUrl)
		status, _ := driver.Status()

		assert.Equal(t, status.Ready, true)
	}
	utils.ProceedTest(test)
}

func TestReloadSiteFewTimesAndCheckHtmlNotChanged(t *testing.T) {
	test := func(driver selenium.WebDriver) {
		utils.TryGetSite(driver, "https://www.york.ac.uk/teaching/cws/wws/webpage1.html")
		initialHtml, _ := driver.PageSource()
		for i := 0; i < 10; i++ {
			_ = driver.Refresh()
		}
		finalHtml, _ := driver.PageSource()

		assert.Equal(t, initialHtml, finalHtml)
	}
	utils.ProceedTest(test)
}

func TestLoadWaitReloadCheckHtmlShouldChanged(t *testing.T) {
	test := func(driver selenium.WebDriver) {
		url := "https://www.zegary.rzeszow.pl/Zegar-atomowy-cabout-pol-289.html"
		utils.TryGetSite(driver, url)
		initialHtml, _ := driver.PageSource()
		_ = driver.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
			return false, nil
		}, time.Second*2)
		finalHtml, _ := driver.PageSource()

		assert.NotEqual(t, initialHtml, finalHtml)
	}
	utils.ProceedTest(test)
}

func TestShouldRedirect(t *testing.T) {
	test := func(driver selenium.WebDriver) {
		url := "https://outlook.office.com/mail/inbox/id/AAQkAGNWRkMzllYgAQAKFj5vBYxbNDph1QbKsKKGA"
		utils.TryGetSite(driver, url)
		currentUrl, _ := driver.CurrentURL()

		assert.NotEqual(t, url, currentUrl)
		assert.Contains(t, url, "https://outlook.office.com")
	}
	utils.ProceedTest(test)
}

func TestChangeSitesAndGoBack(t *testing.T) {
	test := func(driver selenium.WebDriver) {
		firstUrl := "https://www.google.com/"
		secondUrl := "https://www.onet.pl/"
		thirdUrl := "https://stackoverflow.com/"
		utils.TryGetSite(driver, firstUrl)
		utils.TryGetSite(driver, secondUrl)
		utils.TryGetSite(driver, thirdUrl)
		_ = driver.Back()
		_ = driver.Back()
		currentUrl, _ := driver.CurrentURL()

		assert.Equal(t, firstUrl, currentUrl)
	}
	utils.ProceedTest(test)
}

func TestChangeSitesGoBackAndGoForth(t *testing.T) {
	test := func(driver selenium.WebDriver) {
		firstUrl := "https://www.google.com/"
		secondUrl := "https://www.onet.pl/"
		thirdUrl := "https://stackoverflow.com/"
		utils.TryGetSite(driver, firstUrl)
		utils.TryGetSite(driver, secondUrl)
		utils.TryGetSite(driver, thirdUrl)
		_ = driver.Back()
		_ = driver.Back()
		_ = driver.Forward()
		_ = driver.Forward()
		_ = driver.Back()
		_ = driver.Forward()
		currentUrl, _ := driver.CurrentURL()

		assert.Equal(t, thirdUrl, currentUrl)
	}
	utils.ProceedTest(test)
}

func TestShouldLoadFasterThan10Seconds(t *testing.T) {
	test := func(driver selenium.WebDriver) {
		err := driver.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
			url := "https://www.google.com/"
			utils.TryGetSite(wd, url)
			return true, nil
		}, time.Second*10)

		if err != nil {
			panic(err)
		}

	}
	utils.ProceedTest(test)
}

func TestGetOnetCheckAnyATagWithQuestionHyperlinkClassPresent(t *testing.T) {
	test := func(driver selenium.WebDriver) {
		utils.TryGetSite(driver, "https://stackoverflow.com/questions/6509628/how-to-get-http-response-code-using-selenium-webdriver")
		elements, _ := driver.FindElements(selenium.ByXPATH, "//a[contains(@class,'question-hyperlink')]")

		assert.Condition(t, func() (success bool) {
			return len(elements) > 0
		})
	}
	utils.ProceedTest(test)
}

func TestFindElementByItsTextAndClickIt(t *testing.T) {
	test := func(driver selenium.WebDriver) {
		url := "https://pkg.go.dev/github.com/tebeka/selenium"
		utils.TryGetSite(driver, url)
		element, _ := driver.FindElement(selenium.ByLinkText, "Constants")
		_ = element.Click()
		currentUrl, _ := driver.CurrentURL()

		assert.NotEqual(t, url, currentUrl)
	}
	utils.ProceedTest(test)
}

func TestChangeTextInElementViaExecutingScript(t *testing.T) {
	test := func(driver selenium.WebDriver) {
		url := "https://stackoverflow.com/questions/6509628/how-to-get-http-response-code-using-selenium-webdriver"
		utils.TryGetSite(driver, url)

		text := "Blah blah blah"
		_, _ = driver.ExecuteScript("document.getElementById('question-header').innerHTML='"+text+"'", nil)
		element, _ := driver.FindElement(selenium.ByID, "question-header")
		currentText, _ := element.Text()

		assert.Equal(t, text, currentText)
	}
	utils.ProceedTest(test)
}
