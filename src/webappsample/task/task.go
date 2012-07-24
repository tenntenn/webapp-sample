package task

import (
	"time"
	"fmt"
	"strconv"
	"net/http"
	"encoding/json"	

	"appengine"
	"appengine/datastore"
)

// Initilize handlers
func Init() {
	http.HandleFunc("/task", handler)
}

// Handler about task
func handler(w http.ResponseWriter, r *http.Request) {
	switch (r.Method) {
	case "PUT":
		put(w, r)
	case "GET":
		get(w, r)
	}
}

// Nokey is represented by less than this value.
const NOKEY int64 = 0

// Model type of a task
type Task struct {

	Key int64 `datastore:"-" json: "key"`

	// Task name
	Name string `json:"name"`

	// Either the task is done or not.
	IsDone bool `json:"isDone"`

	// Deadline of the task
	Deadline time.Time `json: "deadline"`

	// Priority ot the task
	Priority int `json: "priority"`

}

// Get all tasks from datastore
func GetAll(c appengine.Context) (tasks []*Task, keys []*datastore.Key, err error) {
	
	// Query
	q := datastore.NewQuery("Task").Order("-Priority")
	
	if keys, err = q.GetAll(c, &tasks); err != nil {
		return nil, nil, err
	}

	for i, _ := range tasks {
		tasks[i].Key = keys[i].IntID()
	}

	return tasks, keys, err
}

// Get a task by key (int64)
func Get(c appengine.Context, intID int64) (task *Task, err error) {

	key := datastore.NewKey(c, "Task", "", intID, nil)

	err = datastore.Get(c, key, task)

	if err != nil {
		return nil, err
	}

	task.Key = key.IntID()

	return task, nil
}

// put a task to datastore
func put(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)

	if err := r.ParseForm(); err != nil {
		c.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Infof("%v", r.Form)

	var task Task
	err := json.Unmarshal([]byte(r.FormValue("task")), &task)

	if err != nil {
		c.Errorf("%v - %s", err, r.FormValue("task"))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var key *datastore.Key
	if task.Key < NOKEY {
		key = datastore.NewIncompleteKey(c, "Task", nil)
	} else {
		key = datastore.NewKey(c, "Task", "", task.Key, nil)
	}
	newKey, err := datastore.Put(c, key, &task)

	response := struct {
		Key int64 `json:"key"`
	}{
		newKey.IntID(),
	}

	responseText, err := json.Marshal(&response)
	if err != nil {
		c.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, responseText)
}

// get tasks from datastore
func get(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)

	if err := r.ParseForm(); err != nil {
		c.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	keyStr := r.FormValue("key")
	var tasks []*Task
	if keyStr != "" {
		key, err := strconv.ParseInt(keyStr, 10, 64)		
		if err != nil {
			c.Errorf("%v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// get a task by key
		task, err := Get(c, key)
		if err != nil {
			c.Errorf("%v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tasks = []*Task{task}
	} else {

		// get all tasks
		var err error
		tasks,_, err = GetAll(c)
		if err != nil {
			c.Errorf("%v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	responseText,err := json.Marshal(&tasks)
	if err != nil {
		c.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Infof("%s", responseText)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, responseText)
}