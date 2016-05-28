package main


func filehelper(response http.ResponseWriter, request *http.Request) {
var user User
var session Session
//get the session
ctx := appengine.NewContext(request)
item, session_id, err := getSession(request)
json.Unmarshal(item.Value, &user)
session.User = user
session.Session_id = session_id

if err != nil{
//no session found
log.Errorf(ctx, "*** Error Debug: In files, user is impossible to find: %v ***", err)
logout(response, request)
}

//create a new object handler
client, err := storage.NewClient(ctx)
if err != nil {
log.Errorf(ctx, "*** Error Debug: In files, storage.NewClient: %s", err)
session.Message = "Oooops! Something went wrong try again"
tpl.ExecuteTemplate(response, "files.html", session)
return
}
defer client.Close()

//query the gcs so we could get all the files of the LOGGED IN user
//query delimiter is user.Id/
query := &storage.Query{ Prefix:  strconv.Itoa(int(user.Id)) + "/" }
objs, _ := client.Bucket(gcsBucket).List(ctx, query) //return a list of file objects

//GET method renders all the files to the browser
if request.Method == "GET" {
var files_list []File
for _, obj := range objs.Results {
fileName := strings.TrimPrefix(obj.Name, strconv.Itoa(int(user.Id)) + "/")
file := File{
Name: fileName,
//build the file link manually. Don't know if good practice, but heck it works.
Source_Link: "https://storage.googleapis.com/" + gcsBucket + "/" + strconv.Itoa(int(user.Id)) + "/" + fileName,
Download_Link: obj.MediaLink,
}
files_list = append(files_list, file)
}

//send the files_list to files.js
err = json.NewEncoder(response).Encode(files_list)
}

//delete a file
if request.Method == "DELETE" {
filename := request.FormValue("filename") //filename is passed and retrieve here from the url
//again build the filename in the form user.Id/filename and delete it
err = client.Bucket(gcsBucket).Object( strconv.Itoa(int(user.Id)) + "/" + filename).Delete(ctx)
if err != nil {
log.Errorf(ctx, "*** Error Debug: In filehelper, Delete: %s", err)
}
//any string here
io.WriteString(response, "done")
}
}