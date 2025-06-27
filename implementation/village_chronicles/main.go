package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("*** GO Village Chronicles ***")
	// farmers_inventory()
	// notice_board()
	// route_planner()
	village_messenger()
}

func farmers_inventory() {

	// to-do: make it like a menu driven app while asking prompts from the farmers itself
	fruits := []string{"apple", "mango", "grapes"}

	// morning and evening reports
	morningCount := map[string]int{
		"apple":  3,
		"mango":  5,
		"grapes": 10,
	}

	eveningCount := map[string]int{
		"apple":  6,
		"mango":  2,
		"grapes": 19,
	}

	getCount := func(name, time string) int {
		if time == "morning" {
			return morningCount[name]
		}
		if time == "evening" {
			return eveningCount[name]
		}

		return -1
	}

	// showing difference between morning evening count
	// getting count for possible theft
	for _, fruit := range fruits {
		morning := getCount(fruit, "morning")
		evening := getCount(fruit, "evening")
		diff := evening - morning

		fmt.Printf("%s: Morning=%d, Evening=%d, Difference: %d", fruit, morning, evening, diff)
		if diff < 0 {
			fmt.Println("")
			fmt.Printf("⚠️ Alert: Possible Theft Detected in %s!", fruit)
		}
		fmt.Println("")
	}

	addNewFruit := func(fruitName string, count int) {
		fruits = append(fruits, fruitName)
		morningCount[fruitName] += count
		eveningCount[fruitName] += count
	}

	addNewFruit("banana", 10)
	addNewFruit("mango", 3)
	fmt.Println("\nAfter adding new fruit")
	for _, v := range fruits {
		morning := getCount(v, "morning")
		evening := getCount(v, "evening")

		fmt.Println("")
		fmt.Printf("%s: Morning=%d, Evening=%d", v, morning, evening)

	}

}

func notice_board() {

	// to-do: add some loops to make the notice board continuos.....
	generalAnnouncement := []string{
		"Welcome to the Village Chronicles!",
		"Reminder: Market day is every Saturday.",
		"Please report any missing items to the village chief.",
		"Harvest festival preparations begin next week.",
	}

	privateAnnouncement := []string{
		"Community meeting scheduled for Friday evening at the town hall.",
		"Planning Budget right now",
	}

	eventReader := func(access string) func() string {
		h, p := 0, 0

		return func() (evt string) {
			event := ""

			switch access {
			case "head":
				event = privateAnnouncement[h]
				h += 1
			case "public":
				event = generalAnnouncement[p]
				p += 1
			}

			if h > len(privateAnnouncement)-1 || p > len(generalAnnouncement)-1 {
				h, p = 0, 0
			}
			evt = event
			return
		}
	}

	updateEvent := func(messageFor string) func(...string) {
		if messageFor == "head" {
			return func(args ...string) {
				privateAnnouncement = append(privateAnnouncement, args...)
			}
		}
		return func(args ...string) {
			generalAnnouncement = append(generalAnnouncement, args...)
		}
	}

	event := eventReader("public")
	fmt.Println(event())
	update := updateEvent("public")
	update("Event 1", "event 2")
	fmt.Println(event())
	fmt.Println(event())
	fmt.Println(event())
	fmt.Println(event())
	fmt.Println(event())
	fmt.Println(event())
	fmt.Println(event())
	fmt.Println(event())
	fmt.Println(event())
	fmt.Println(event())
	fmt.Println(event())
	fmt.Println(event())
	fmt.Println(event())
	fmt.Println(event())
}

// chapter 3: uses struct that i haven't used currently

func route_planner() {

	mails := map[string]map[int]bool{
		"sector1": map[int]bool{
			101: true,
			102: true,
			103: false,
		},
		"sector2": map[int]bool{
			201: true,
			202: true,
			203: false,
			204: true,
		},
		"sector3": map[int]bool{
			301: false,
		},
	}

	markDelivered := func(houseNum int) map[string]map[int]bool {

		defer func() {
			errMessage := recover()
			if errMessage != nil && len(errMessage.(string)) > 0 {
				fmt.Println(errMessage)
			}
		}()

		if houseNum < 100 {
			panic("Invalid House Num")
		}

		for _, mail := range mails {
			if mail[houseNum] == true {
				mail[houseNum] = false
			}
		}
		return mails
	}

	housesWithMail := func(mails map[string]map[int]bool) map[int]bool {
		result := make(map[int]bool)
		for _, sector := range mails {
			for house, hasMail := range sector {
				if hasMail {
					result[house] = hasMail
				}
			}
		}
		return result
	}

	fmt.Println(markDelivered(120))
	fmt.Println(housesWithMail(mails))

}

// chapter 6: village messenger app
func village_messenger() {

	msg_queue := make([]string, 0, 5)

	msg_queue = append(msg_queue, "rohan", "chris", "baby", "bobyy", "boobo")
	send_messages(&msg_queue)

}

// ch-6 helper functions

func process_messages(message string) error {
	fmt.Println("processing message....")
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println(r)
	// 	}
	// }()

	if message == "panic" {
		// panic("Panic Explosion")
		return fmt.Errorf("Panic explosion while processing %s", message)
	}
	fmt.Println("Processed: ", message)
	return nil
}

func send_messages(msg_queue *[]string) {

	msg := ""
	fmt.Print("Enter a message: ")
	fmt.Scanln(&msg)

	if len(*msg_queue) < cap(*msg_queue) {
		*msg_queue = append(*msg_queue, msg)
	} else {
		*msg_queue = slices.Delete(*msg_queue, 0, 1)
		*msg_queue = append(*msg_queue, msg)
	}
	err := retry(func() error {
		return process_messages(msg)
	}, 3)

	if err != nil {
		fmt.Println(err)
	}

}

func retry(action func() error, maxRetries int) (err error) {

	for attempt := 1; attempt <= maxRetries; attempt++ {
		fmt.Printf("Attempting %d....\n", attempt)

		err = action()

		if err == nil {
			return nil
		}
		fmt.Println("Error: ", err)

	}

	return fmt.Errorf("All %d retries failed. Last error: %w", maxRetries, err)

}
