package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Meetings []Meeting

// STRUCTURE FOR DEFINING MEETING AS GIVEN IN PDF
type Meeting struct {
	IdMeeting          string `json:"IdMeeting"`
	Title              string `json:"Title"`
	Participants       string `json:"Participants"`
	Start_Time         string `json:"Start_Time"`
	End_Time           string `json:"End_Time"`
	Creation_Timestamp string `json:"Creation_Timestamp"`
}

// STRUCTURE FOR DEFINING PARTICIPANT AS GIVEN IN PDF
type Participant struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
	RSVP  string `json:"RSVP"`
}

// END POINT FOR "GET A MEETING USING ID" AS GIVEN IN PDF : : COVERING END POINTS
func returnMeetingOfId(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint For: Listing Meeting of id X")
	id := strings.SplitN(r.URL.Path, "/", 3)[2]

	for i := 0; i < len(Meetings); i++ {
		if id == Meetings[i].IdMeeting 
		{
			json.NewEncoder(w).Encode(Meetings[i])
		}
	}
}

func MeetingOperations(w http.ResponseWriter, r *http.Request) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://kaustubh:abc123def@cluster0.kfjyr.mongodb.net/appointly?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	fmt.Println("CONNECTION TO MONGO DB IS SUCCESSFULL")

	var presented_query = r.URL.Query()
	if len(presented_query["participant"]) != 0 {
		w.Header().Set("Content-Type", "application/json")
		for i := 0; i < len(Meetings); i++ {
			fmt.Println(Meetings[i].Participants)
			fmt.Println(presented_query["participant"][0])
			if presented_query["participant"][0] == "\""+Meetings[i].Participants+"\"" 
			{
				json.NewEncoder(w).Encode(Meetings[i])
			}
		}
		fmt.Println("THE ROUTE FOR PARTICIPANT CORRECT DONE")
	}
	if len(presented_query["start"]) != 0 && len(presented_query["end"]) != 0 
	{
		for i := 0; i < len(Meetings); i++ {
			fmt.Println(presented_query["start"][0])
			fmt.Println(presented_query["end"][0])
			fmt.Println(Meetings)
			if presented_query["start"][0] == "\""+Meetings[i].Start_Time+"\"" && presented_query["end"][0] == "\""+Meetings[i].End_Time+"\"" {
				json.NewEncoder(w).Encode(Meetings[i])
			}

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
		log.Println(inMeet)
		appointlyDB := client.Database("appointlyDB")
		meetingsCollection := appointlyDB.Collection("meetings")
		meetingResult, err := meetingsCollection.InsertOne(ctx, bson.D{
			{Key: "Id", Value: inMeet.IdMeeting},
			{Key: "Title", Value: inMeet.Title},
			{Key: "Participants", Value: inMeet.Participants},
			{Key: "Start_Time", Value: inMeet.Start_Time},
			{Key: "End_Time", Value: inMeet.End_Time},
			{Key: "Creation_Timestamp", Value: inMeet.Creation_Timestamp}})

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(meetingResult.InsertedID)
		just_now := time.Now()
		var newMeeting Meeting = Meeting{IdMeeting: inMeet.IdMeeting, Title: inMeet.Title, Participants: inMeet.Participants, Start_Time: inMeet.Start_Time, End_Time: inMeet.End_Time, Creation_Timestamp: just_now.String()}
		Meetings = append(Meetings, newMeeting)
		json.NewEncoder(w).Encode(inMeet)
		fmt.Println("SCHEDULE MEETING ROUTE")
	}

}

func handleRequests() {
	http.HandleFunc("/meeting/", returnMeetingOfId)
	http.HandleFunc("/meetings", MeetingOperations)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {

	Meetings = []Meeting{}
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://kaustubh:abc123def@cluster0.kfjyr.mongodb.net/appointly?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	fmt.Println("CONNECTION TO MONGODB IS SUCCESSFULL")

	appointlyDB := client.Database("appointlyDB")
	meetingsCollection := appointlyDB.Collection("meetings")

	cursor, err := meetingsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var allMeetingsDB []bson.M
	if err = cursor.All(ctx, &allMeetingsDB); err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(allMeetingsDB); i++ {
		var newMeeting Meeting = Meeting{IdMeeting: fmt.Sprint(allMeetingsDB[i]["Id"]), Title: fmt.Sprint(allMeetingsDB[i]["Title"]), Participants: fmt.Sprint(allMeetingsDB[i]["Participants"]), Start_Time: fmt.Sprint(allMeetingsDB[i]["Start_Time"]), End_Time: fmt.Sprint(allMeetingsDB[i]["End_Time"]), Creation_Timestamp: fmt.Sprint(allMeetingsDB[i]["Creation_Timestamp"])}
		Meetings = append(Meetings, newMeeting)
	}

	handleRequests()
}
