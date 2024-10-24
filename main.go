package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 3 {
	} else {
		fmt.Println("[!] Error: [Provide 2 arguments.]")
		return
	}
	input, output := os.Args[1], os.Args[2]
	if os.Args[1][len(os.Args[1])-4:len(os.Args[1])] != ".txt" || os.Args[2][len(os.Args[2])-4:len(os.Args[2])] != ".txt" {
		fmt.Println("[!] Error: [Incorrect file name (0-9a-zA-Z_) or extension (.txt).]")
		fmt.Println("Program exited.")
		log.Fatalln()
	}
	text_byte, err := os.ReadFile(input)
	if err != nil {
		fmt.Printf("[!] Error: [%s]\n\n", err.Error())
		fmt.Println("Program exited.")
		log.Fatalln()
	} else {
		fmt.Println("[+] Data read successfully.")
	}
	formattedText := string(text_byte)
	for range "123" {
		formattedText = FormatPunctuation(FormatText(FormatPunctuation(formattedText)))
	}
	formattedText = regexp.MustCompile(`(\()\s*(\w+)`).ReplaceAllString(formattedText, "$1$2")
	formattedText = strings.ReplaceAll(formattedText, "\r", "")
	for range "123" {
		formattedText = regexp.MustCompile(`(\S+)(\n*)(\s*)([:;.!?])(\w*)`).ReplaceAllString(formattedText, "$1$4$5")
	}
	formattedText = article(formattedText)
	formatedApostrophe := regexp.MustCompile(`'\s*(.*?)\s*'`).ReplaceAllString(formattedText, "'$1'")
	formattedText = regexp.MustCompile(`(.*?) +' +(.*?)`).ReplaceAllString(formatedApostrophe, "$1' $2")
	formattedText = regexp.MustCompile(`"\s*(.*?)\s*"`).ReplaceAllString(formattedText, `"$1"`)
	err = os.WriteFile(output, []byte(formattedText), 0o777)
	if err != nil {
		fmt.Printf("[!] Error: [%s]\n\n", err.Error())
		fmt.Println("Program exited.")
		log.Fatalln()
	} else {
		fmt.Println("[+] File saved successfully.")
	}
}
func article(text string) string {
	an := regexp.MustCompile(`((^a)|(\sa))(\s+[aeiouhAEIOUH])`)
	index := an.FindStringIndex(text)
	for len(an.FindStringIndex(text)) != 0 {
		if index[0] == 0 {
			text = "an" + text[1:]
			index = an.FindStringIndex(text)
		} else {
			text = text[:index[0]+1] + "an" + text[index[0]+2:]
			index = an.FindStringIndex(text)
		}
	}
	an = regexp.MustCompile(`((^A)|(\sA))(\s+[aeiouhAEIOUH])`)
	index = an.FindStringIndex(text)
	for len(an.FindStringIndex(text)) != 0 {
		if index[0] == 0 {
			text = "An" + text[1:]
			index = an.FindStringIndex(text)
		} else {
			text = text[:index[0]+1] + "An" + text[index[0]+2:]
			index = an.FindStringIndex(text)
		}
	}
	an = regexp.MustCompile(`((^an)|(\san))(\s+[^aeiouhAEIOUH\s])`)
	index = an.FindStringIndex(text)
	if index != nil && index[0] == 0 {
		text = "a" + text[2:]
		index = an.FindStringIndex(text)
	}
	for len(an.FindStringIndex(text)) != 0 {
		if index[0] == 0 {
			text = "a" + text[1:]
			index = an.FindStringIndex(text)
		} else {
			text = text[:index[0]+1] + "a" + text[index[0]+3:]
			index = an.FindStringIndex(text)
		}
	}
	an = regexp.MustCompile(`((^An) |(^AN)|(\sAn)| (\sAN))(\s+[^aeiouhAEIOUH\s])`)
	index = an.FindStringIndex(text)
	if index != nil && index[0] == 0 {
		text = "A" + text[2:]
		index = an.FindStringIndex(text)
	}
	for len(an.FindStringIndex(text)) != 0 {
		if index[0] == 0 {
			text = "A" + text[1:]
			index = an.FindStringIndex(text)
		} else {
			text = text[:index[0]+1] + "A" + text[index[0]+3:]
			index = an.FindStringIndex(text)
		}
	}
	return text
}
func FormatPunctuation(text string) string {
	remove_spaces := regexp.MustCompile(` {1,}`).ReplaceAllString(text, " ")
	bracketWithoutSpace := regexp.MustCompile(`\(\s*(.*?)\s*\)`).ReplaceAllString(remove_spaces, "($1)")
	formated := regexp.MustCompile(` *([.,!?:;]) *`).ReplaceAllString(bracketWithoutSpace, "$1")
	formated = regexp.MustCompile(`(\))(\()`).ReplaceAllString(formated, "$1 $2")
	formated = regexp.MustCompile(`(\))(\()`).ReplaceAllString(formated, "$1 $2")
	formated = regexp.MustCompile(`(\()([0-9a-fA-F]+)`).ReplaceAllString(formated, "$1 $2")
	formated = regexp.MustCompile(`(\()(\s)([b-c][ia])`).ReplaceAllString(formated, "$1$3")
	formated = regexp.MustCompile(`(\w)(\()`).ReplaceAllString(formated, "$1 $2")
	formated = regexp.MustCompile(`(\')(\()`).ReplaceAllString(formated, "$1 $2")
	formated = regexp.MustCompile(`([.,!?:;]+)([^ ]|$)`).ReplaceAllString(formated, "$1 $2")
	formatedApostrophe := regexp.MustCompile(`'\s*(.*?)\s*'`).ReplaceAllString(formated, "'$1'")
	formated = regexp.MustCompile(`(.*?) +' +(.*?)`).ReplaceAllString(formatedApostrophe, "$1' $2")
	formatedQuotes := regexp.MustCompile(`"\s*(.*?)\s*"`).ReplaceAllString(formated, `"$1"`)
	spaceStartLine := regexp.MustCompile(`^\s*`).ReplaceAllString(formatedQuotes, "")
	txt := regexp.MustCompile(`\(lower case\)`).ReplaceAllString(spaceStartLine, "(low)")
	fmt.Println(txt)
	return txt
}
func FormatText(inputText string) string {
	reForConvert := regexp.MustCompile(`\((cap|up|low|hex|bin)((,\s*\w+\[\S])|(,\s*\w+))?\)`) //`\((cap|up|low|hex|bin)((,\s*\S+)|(,\s*\S+\W+))?\)
	matches := reForConvert.FindAllString(inputText, -1)
	for _, match := range matches {
		reMatchSlash := regexp.MustCompile(`([()])`).ReplaceAllString(match, `\$1`)
		reAllPrevWords := regexp.MustCompile(`(?s).*?` + reMatchSlash)
		count := 0
		for _, char := range inputText {
			if char == '\'' {
				count++
			}
		}
		allPrevWords := reAllPrevWords.FindString(inputText)
		allPrevWords_sub := strings.Split(allPrevWords, match)
		allPrevWords = allPrevWords_sub[0] + match
		converted := convertMatch(allPrevWords, match, reMatchSlash)
		inputText = strings.Replace(inputText, allPrevWords, converted, 1)
	}
	return inputText
}
func convertMatch(allPrevWords, match, reMatchSlash string) string {
	isNewLine := regexp.MustCompile(`\n` + reMatchSlash)
	for i := 0; i < 5; i++ {
		if isNewLine.MatchString(allPrevWords) {
			allPrevWords = strings.TrimSuffix(allPrevWords, match)
		}
	}
	wordCase := regexp.MustCompile(`\b(cap|up|low|hex|bin)\b`).FindString(match)
	wordNum, err := strconv.Atoi(regexp.MustCompile(`\d+`).FindString(match))
	if err != nil {
		wordNum = 1
	}
	onlyPrevWords := strings.TrimSuffix(allPrevWords, match)
	onlyPrevWords = strings.Replace(onlyPrevWords, "\n", " \n", -1)
	onlyPrevWords = regexp.MustCompile(` {1,}`).ReplaceAllString(onlyPrevWords, " ")
	words := regexp.MustCompile(` `).Split(onlyPrevWords, -1)
	for i, elem := range words {
		fmt.Print(i)
		fmt.Print(": ")
		fmt.Printf("%q\n", elem)
	}
	wordCount := 0
	for index := len(words) - 1; index >= 0; index-- {
		wordTemp := words[index]
		if wordNum == wordCount {
			break
		}
		if len(words[index]) != 0 && strings.ContainsAny(words[index], "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890") && words[index] != "\n''\r" && words[index] != "()" && words[index] != "\n\r" && words[index] != "\r" && words[index] != "\n" && words[index] != "'\r"  && words[index] != "\n:" {
			wordCount++
		}
		switch wordCase {
		case "bin":
			wordTemp1 := ""
			add := ""
			for _, char := range wordTemp {
				if char != '\n' && char != '\r' {
					wordTemp1 += string(char)
				} else if char == '\n' {
					add = "\n"
				}
			}
			decimal, _ := strconv.ParseInt(wordTemp1, 2, 64)
			convertedNum := strconv.Itoa(int(decimal))
			if convertedNum == "0" {
				words[index] = wordTemp + " "
			} else {
				if add == "\n" {
					words[index] = add + convertedNum + " "
				} else {
					words[index] = convertedNum + " "
				}
			}
		case "hex":
			wordTemp1 := ""
			add := ""
			for _, char := range wordTemp {
				if char != '\n' && char != '\r' {
					wordTemp1 += string(char)
				} else if char == '\n' {
					add = "\n"
				}
			}
			decimal, _ := strconv.ParseInt(wordTemp1, 16, 64)
			convertedNum := strconv.Itoa(int(decimal))
			if convertedNum == "0" {
				words[index] = wordTemp + " "
			} else {
				if add == "\n" {
					words[index] = add + convertedNum + " "
				} else {
					words[index] = convertedNum + " "
				}

			}
		case "up":
			words[index] = strings.ToUpper(wordTemp)
		case "low":
			words[index] = strings.ToLower(wordTemp)
		case "cap":
			words[index] = capitalizeWord(strings.ToLower(wordTemp))
		}
	}
	return strings.Join(words, " ")
}
func capitalizeWord(s string) string {
	runes := []rune(s)
	capitalize := true
	count := 0
	for i, r := range runes {
		if isAlpha := (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z'); isAlpha && count < 2 {
			if capitalize && r >= 'a' && r <= 'z' {
				runes[i] -= 32
			} else if !capitalize && r >= 'A' && r <= 'Z' {
				runes[i] += 32
			}
			capitalize = false
		} else if i == 0 && r == '\'' || r == '(' || r == '/' || r == '\\' { //r == '\'' ||
			count++
			capitalize = true
		} else if r != '\n' {
			capitalize = false
		}
	}
	return string(runes)
}
