package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Определение флагов командной строки
	urlFlag := flag.String("url", "", "URL of the website to download")
	outputFlag := flag.String("output", "", "Output directory to save the website content")
	flag.Parse()

	// Проверка обязательного наличия URL
	if *urlFlag == "" {
		fmt.Println("Usage: go run main.go -url <URL> [-output <output directory>]")
		os.Exit(1)
	}

	// Определение выходной директории
	outputDir := *outputFlag
	if outputDir == "" {
		outputDir = "./"
	}

	// Проверка наличия слеша в конце URL
	if !strings.HasSuffix(*urlFlag, "/") {
		*urlFlag += "/"
	}

	// Парсинг URL
	baseURL, err := url.Parse(*urlFlag)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		os.Exit(1)
	}

	// Скачивание сайта
	err = downloadWebsite(baseURL, outputDir)
	if err != nil {
		fmt.Println("Error downloading website:", err)
		os.Exit(1)
	}

	fmt.Println("Website downloaded successfully!")
}

func downloadWebsite(baseURL *url.URL, outputDir string) error {
	// Выполнение GET-запроса к базовому URL
	resp, err := http.Get(baseURL.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Создание директории для сохранения файлов сайта
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return err
	}

	// Сохранение файла index.html
	htmlContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(outputDir, "index.html"), htmlContent, os.ModePerm)
	if err != nil {
		return err
	}

	// Парсинг ссылок на ресурсы на странице
	links := parseLinks(string(htmlContent))

	// Скачивание и сохранение всех ресурсов
	for _, link := range links {
		resourceURL, err := baseURL.Parse(link)
		if err != nil {
			fmt.Println("Error parsing resource URL:", err)
			continue
		}

		resp, err := http.Get(resourceURL.String())
		if err != nil {
			fmt.Println("Error downloading resource:", err)
			continue
		}
		defer resp.Body.Close()

		resourceContent, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading resource content:", err)
			continue
		}

		// Сохранение ресурса в файл
		resourcePath := filepath.Join(outputDir, filepath.Base(resourceURL.Path))
		err = ioutil.WriteFile(resourcePath, resourceContent, os.ModePerm)
		if err != nil {
			fmt.Println("Error saving resource to file:", err)
			continue
		}
	}

	return nil
}

func parseLinks(html string) []string {
	var links []string

	// Простой парсинг HTML на ссылки
	// В реальном проекте лучше использовать более надежные методы
	startIndex := 0
	for {
		hrefIndex := strings.Index(html[startIndex:], "href=\"")
		if hrefIndex == -1 {
			break
		}
		hrefIndex += startIndex

		linkStart := hrefIndex + 6
		linkEnd := strings.Index(html[linkStart:], "\"")
		if linkEnd == -1 {
			break
		}
		linkEnd += linkStart

		links = append(links, html[linkStart:linkEnd])

		startIndex = linkEnd
	}

	return links
}
