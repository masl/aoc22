package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

func ErrorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Creates folder for the day and copy the template file
func CreateStructure(day uint) {
	if _, err := os.Stat(FormatDay(day)); os.IsNotExist(err) {
		os.Mkdir(FormatDay(day), 0777)
	}

	// Copy template file
	wd, _ := os.Getwd()
	srcPath := wd + "/utils/template"
	dstPath := wd + "/" + FormatDay(day) + "/main.go"

	// Open source file
	src, err := os.Open(srcPath)
	ErrorCheck(err)
	defer src.Close()

	// Create destination file
	dst, err := os.Create(dstPath)
	ErrorCheck(err)
	defer dst.Close()

	io.Copy(dst, src)

	DownloadInput(day)
	GetDescription(day)
}

// Formats the day to a string with leading zero
func FormatDay(day uint) string {
	return fmt.Sprintf("%02d", day)
}

// Reads the input file for given day
func ReadInput(day uint) []string {
	file, err := os.Open(FormatDay(day) + "/input.txt")
	ErrorCheck(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// Gets input from the aoc website using the session cookie and save it to a file
func DownloadInput(day uint) {
	var client http.Client

	cookie := &http.Cookie{
		Name:   "session",
		Value:  os.Getenv("AOC_SESSION"),
		MaxAge: 0,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2022/day/%d/input", day), nil)
	ErrorCheck(err)

	req.AddCookie(cookie)
	resp, err := client.Do(req)
	ErrorCheck(err)

	defer resp.Body.Close()
	fmt.Printf("StatusCode: %d\n", resp.StatusCode)

	// Read body
	body, err := io.ReadAll(resp.Body)

	// Write to file
	file, err := os.Create(FormatDay(day) + "/input.txt")
	ErrorCheck(err)

	_, err = file.WriteString(string(body))
	ErrorCheck(err)
	defer file.Close()
}

// Get the aoc description for the day
func GetDescription(day uint) {
	var client http.Client
	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2022/day/%d", day), nil)
	ErrorCheck(err)

	resp, err := client.Do(req)
	ErrorCheck(err)

	defer resp.Body.Close()
	fmt.Printf("StatusCode: %d\n", resp.StatusCode)

	// Read body and parse it
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	ErrorCheck(err)

	find := doc.Find("article.day-desc")
	html, err := find.Html()
	ErrorCheck(err)

	converter := md.NewConverter("adventofcode.com", true, nil)
	markdownContent, err := converter.ConvertString(html)
	ErrorCheck(err)

	// Write to file
	file, err := os.Create(FormatDay(day) + "/README.md")
	ErrorCheck(err)

	_, err = file.WriteString(markdownContent)
	ErrorCheck(err)
	defer file.Close()
}
