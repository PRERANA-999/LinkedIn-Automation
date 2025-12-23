package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

// Human-like delay
func HumanDelay(min, max int) {
	time.Sleep(time.Duration(rand.Intn(max-min)+min) * time.Millisecond)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// -------------------------------
	// Launch browser (VISIBLE)
	// -------------------------------
	u := launcher.New().
		Headless(false).
		NoSandbox(true).
		MustLaunch()

	browser := rod.New().
		ControlURL(u).
		MustConnect()
	defer browser.MustClose()

	// -------------------------------
	// Open LinkedIn
	// -------------------------------
	page := browser.MustPage("https://www.linkedin.com")
	page.MustWaitLoad()

	fmt.Println("🔐 Please login manually.")
	fmt.Println("✅ After login + feed loads, press ENTER here.")
	fmt.Scanln()

	// Ensure feed is loaded
	page.MustWaitLoad()
	page.MustWaitIdle()
	HumanDelay(2000, 4000)

	fmt.Println("📍 Current URL:", page.MustInfo().URL)

	// -------------------------------
	// CLICK MESSAGING ICON
	// -------------------------------
	fmt.Println("🔎 Locating Messaging icon...")

	msgIcon := page.MustElementX(
		"//a[contains(@href,'/messaging')]",
	)

	msgIcon.MustScrollIntoView()
	HumanDelay(800, 1500)
	msgIcon.MustHover()
	HumanDelay(500, 1000)
	msgIcon.MustClick()

	fmt.Println("💬 Clicked Messaging")

	// Wait for messaging page
	page.MustWaitLoad()
	page.MustWaitIdle()
	HumanDelay(3000, 4000)

	// -------------------------------
	// WAIT FOR CONVERSATION LIST
	// -------------------------------
	fmt.Println("⏳ Waiting for conversation list...")

	page.MustWaitElementsMoreThan(
		"li.msg-conversation-listitem",
		0,
	)

	HumanDelay(1500, 2500)

	// -------------------------------
	// OPEN FIRST CONVERSATION
	// -------------------------------
	fmt.Println("🗂 Opening first conversation...")

	chats := page.MustElements("li.msg-conversation-listitem")
	firstChat := chats[0]

	firstChat.MustScrollIntoView()
	HumanDelay(700, 1200)
	firstChat.MustHover()
	HumanDelay(400, 800)
	firstChat.MustClick()

	page.MustWaitIdle()
	HumanDelay(2000, 3000)

	fmt.Println("✅ Conversation opened")

	// -------------------------------
	// READ LAST MESSAGE
	// -------------------------------
	fmt.Println("📩 Reading last message...")

// Wait for message container
page.MustElement("div.msg-spinmail-thread-presenter__message")

HumanDelay(1200, 1800)

// Grab all message lines
lines := page.MustElements("p.msg-spinmail-thread-presenter__message-body p")

if len(lines) == 0 {
	fmt.Println("⚠ No message text found")
	return
}

// Collect text safely
var messages []string
for _, l := range lines {
	txt := strings.TrimSpace(l.MustText())
	if txt != "" {
		messages = append(messages, txt)
	}
}

if len(messages) == 0 {
	fmt.Println("⚠ Message container empty")
	return
}

// Print last logical message
lastMessage := messages[len(messages)-1]
fmt.Println("💬 Last message:", lastMessage)

	// -------------------------------
	// SCROLL MESSAGE LIST
	// -------------------------------
	fmt.Println("📜 Scrolling messages...")

	scrollBox := page.MustElement("div.msg-s-message-list-container")

	for i := 0; i < 5; i++ {
		scrollBox.MustEval(`el => { el.scrollTop += 300 }`)
		HumanDelay(800, 1200)
	}

	fmt.Println("✅ Automation demo completed successfully")

	// Keep browser open for demo
	select {}
}
