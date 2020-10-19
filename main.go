package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

var Meetings []Meeting

// STRUCTURE FOR DEFINING MEETING AS GIVEN IN THE PDF
type Meeting struct {
	Id                 string    `json:"Id"`
	Title              string    `json:"Title"`
	Participants       string    `json:"Participants"`
	Start_Time         time.Time `json:"Start_Time"`
	End_Time           time.Time `json:"End_Time"`
	Creation_Timestamp time.Time `json:"Creation_Timestamp"`
}

// STRUCTURE FOR DEFINING PARTICIPANT AS GIVEN IN THE PDF
type Participant struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
	RSVP  string `json:"RSVP"`
}

// END POINT FOR "GET A MEETING USING ID"
func returnMeetingOfId(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint For: Listing Meeting of id X")
	id := strings.SplitN(r.URL.Path, "/", 3)[2]

	for i := 0; i < len(Meetings); i++ {
		if id == Meetings[i].Id {
			json.NewEncoder(w).Encode(Meetings[i])
		}
	}
}

func MeetingOperations(w http.ResponseWriter, r *http.Request) {
	var presented_query = r.URL.Query()
	if len(presented_query["participant"]) != 0 {
		w.Header().Set("Content-Type", "application/json")

		for i := 0; i < len(Meetings); i++ {
			fmt.Println(Meetings[i].Participants)
			fmt.Println(presented_query["participant"][0])
			if presented_query["participant"][0] == "\""+Meetings[i].Participants+"\"" {
				fmt.Println("sad")
				json.NewEncoder(w).Encode(Meetings[i])
			}
		}
		fmt.Println("THE ROUTE FOR PARTICIPANT CORRECT DONE")
	}
	if len(presented_query["start"]) != 0 && len(presented_query["end"]) != 0 {
		for i := 0; i < len(Meetings); i++ {
			start_meeting_time, err := time.Parse(time.ANSIC, presented_query["start"][0])
			end_meeting_time, err := time.Parse(time.ANSIC, presented_query["end"][0])
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(start_meeting_time)
			fmt.Println(end_meeting_time)
		}
		fmt.Println("ROUTE FOR RANGE MEETING CORRECTLY DONE")
	}
	if len(presented_query["participant"]) == 0 && len(presented_query["start"]) == 0 && len(presented_query["end"]) == 0 {
		decoder := json.NewDecoder(r.Body)
		var inMeet Meeting
		w.Header().Set("Content-Type", "application/json")
		err := decoder.Decode(&inMeet)
		if err != nil {
			panic(err)
		}
	
		just_now := time.Now()
		var newMeeting Meeting = Meeting{Id: inMeet.Id, Title: inMeet.Title, Participants: inMeet.Participants, Start_Time: inMeet.Start_Time, End_Time: inMeet.End_Time, Creation_Timestamp: just_now}
		Meetings = append(Meetings, newMeeting)
		json.NewEncoder(w).Encode(inMeet)
		fmt.Println("SCHEDULE MEETING ROUTE")
	}

}

func handleRequests() {
	http.HandleFunc("/meeting/", returnMeetingOfId)
	http.HandleFunc("/meetings", MeetingOperations)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {

	Meetings = []Meeting{
		Meeting{Id: "1", Title: "Appointy General meeting", Participants: "general@appointy.com", Start_Time: time.Now().UTC(), End_Time: time.Now().UTC(), Creation_Timestamp: time.Now().UTC()},
		Meeting{Id: "2", Title: "Appointy Executive Meeting", Participants: "executive@appointy.com", Start_Time: time.Now().UTC(), End_Time: time.Now().UTC(), Creation_Timestamp: time.Now().UTC()},
		Meeting{Id: "3", Title: "Appointy and Kaustubh's Meeting", Participants: "kaustubh@appointy.com", Start_Time: time.Now().UTC(), End_Time: time.Now().UTC(), Creation_Timestamp: time.Now().UTC()},
	}
	handleRequests()
}
