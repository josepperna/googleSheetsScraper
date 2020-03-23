package main

import (
    "fmt"
    "log"
    //"time"

    //"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    
    "encoding/json"
    "io/ioutil"
    "net/http"
    "os"
    "strconv"
    "strings"

    "golang.org/x/net/context"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    "google.golang.org/api/sheets/v4"
)

// You will be using this Trainer type later in the program
type Answer struct {
    Date string
    Regio string
    Parc string
    Cinta int
    Vestit1 int
    Vestit2 int
    Caputxa int
    Guants1 int
    Guants2 int
    Mascares1 int
    Mascares2 int
    Lleixiu bool
    Sabo bool
    Guardia bool
}

type Trainer struct {
        Name string
        Age  int
        City string
    }

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
    // The file token.json stores the user's access and refresh tokens, and is
    // created automatically when the authorization flow completes for the first
    // time.
    tokFile := "token.json"
    tok, err := tokenFromFile(tokFile)
    if err != nil {
            tok = getTokenFromWeb(config)
            saveToken(tokFile, tok)
    }
    return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
    authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
    fmt.Printf("Go to the following link in your browser then type the "+
            "authorization code: \n%v\n", authURL)

    var authCode string
    if _, err := fmt.Scan(&authCode); err != nil {
            log.Fatalf("Unable to read authorization code: %v", err)
    }

    tok, err := config.Exchange(context.TODO(), authCode)
    if err != nil {
            log.Fatalf("Unable to retrieve token from web: %v", err)
    }
    return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
    f, err := os.Open(file)
    if err != nil {
            return nil, err
    }
    defer f.Close()
    tok := &oauth2.Token{}
    err = json.NewDecoder(f).Decode(tok)
    return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
    fmt.Printf("Saving credential file to: %s\n", path)
    f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
    if err != nil {
            log.Fatalf("Unable to cache oauth token: %v", err)
    }
    defer f.Close()
    json.NewEncoder(f).Encode(token)
}

func getData() [] interface{} {
    ansList := make([] interface{}, 1)
    b, err := ioutil.ReadFile("credentials.json")
    if err != nil {
            log.Fatalf("Unable to read client secret file: %v", err)
    }

    // If modifying these scopes, delete your previously saved token.json.
    config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
    if err != nil {
            log.Fatalf("Unable to parse client secret file to config: %v", err)
    }
    client := getClient(config)

    srv, err := sheets.New(client)
    if err != nil {
            log.Fatalf("Unable to retrieve Sheets client: %v", err)
    }

    // Prints the names and majors of students in a sample spreadsheet:
    // https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
    spreadsheetId := "1rbOKKUHX32RyH7DfPJFLVwmEvBDdb3gOM608UGvluRs"
    readRange := "Respostes al formulari 1!A2:N"
    resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
    if err != nil {
            log.Fatalf("Unable to retrieve data from sheet: %v", err)
    }

    if len(resp.Values) == 0 {
            fmt.Println("No data found.")
    } else {
            //fmt.Println(resp.Values)
            for _, row := range resp.Values {
                if len(row) == 14{
                        var tempList []int
                        for _, item := range row[3:11]{
                                temp, err := strconv.Atoi(fmt.Sprintf("%v",item))
                                if err == nil {
                                        tempList = append(tempList, temp)
                                }else{
                                        if s, err := strconv.ParseFloat(fmt.Sprintf("%v",item), 32); err == nil {
                                                tempList = append(tempList, int(s))
                                            }else{
                                                    tempList = append(tempList, 0)
                                            }
                                        
                                }
                        }
                        tempEntry := Answer{fmt.Sprintf("%v",row[0])[0:10], fmt.Sprintf("%v",row[1]), fmt.Sprintf("%v", removeSpacesAndMakeLower(row[2])), tempList[0],
                        tempList[1], tempList[2], tempList[3], tempList[4], tempList[5], tempList[6], tempList[7], convertToBool(row[10]), convertToBool(row[12]), convertToBool(row[13])}
                        ansList = append(ansList, tempEntry)
                        //fmt.Println(tempEntry)
                }
            }
    }
    return ansList
}
func convertToBool(item interface{}) bool{
        if(len(fmt.Sprintf("%v",item)) == 0){
                return false
        }
        if(strings.ToLower(fmt.Sprintf("%v",item))[0] == 's'){
                return true
        }
        return false
}
func removeSpacesAndMakeLower(item interface{}) string{
        newItem := strings.ToLower(fmt.Sprintf("%v",item))
        newItem = strings.TrimSpace(newItem)
        return newItem
}

func main() {
         // Rest of the code will go here
        ans := getData()
        clientOptions := options.Client().ApplyURI("mongodb+srv://helixcatofen:M-72lawLATW@cluster0-wniia.mongodb.net/test?retryWrites=true&w=majority")
        var client, err = mongo.Connect(context.TODO(), clientOptions)
        if err != nil{
                log.Fatal(err)
        }
        collection := client.Database("test").Collection("answers")
        opts := options.InsertMany().SetOrdered(false)
        updateResult, err := collection.InsertMany(context.TODO(), ans[2:], opts)
        if err != nil {
                log.Fatal(err)
                
        }
        fmt.Println(updateResult)


}

