package endpoint

import (
	"encoding/json"
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"sort"
)




type StoreHandler struct {

	Store map[string]string
	
}


func NewstoreHandlers() *StoreHandler {

	k := StoreHandler{Store: map[string]string{}}
	k.Store["abc-1"] = "this"
	k.Store["abc-2"] = "is"
	k.Store["xyz-1"] = "a"
	k.Store["xyz-2"] = "test"
	return &k

}




func (k *StoreHandler) Getkeys(w http.ResponseWriter, r *http.Request) {

	key := strings.Split(r.URL.Path, "/")[2]
	fmt.Println(key)
	jsonBytes, err := json.Marshal(k.Store[key])
	 if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	 }
	 w.Write(jsonBytes)
}

func (k *StoreHandler) Setkeys(w http.ResponseWriter, r *http.Request) {

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("need content-type 'application/json', but got '%s'", ct)))
		return
	}

	keystore := map[string]string{}

	err = json.Unmarshal(bodyBytes, &keystore)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(keystore)

	for key, _ := range keystore{
		k.Store[key] = keystore[key]

	}
	

	jsonBytes, err := json.Marshal(k.Store)
	if err != nil {

	   w.WriteHeader(http.StatusInternalServerError)
	   w.Write([]byte(err.Error()))
	}
	w.Write(jsonBytes)


}





func (k *StoreHandler) Searchkeys(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
    query_map := r.URL.Query()

	var result[]string
	
	if value, ok := query_map["prefix"]; ok {

		prefix_len := len(value[0])
		for key, _ := range k.Store {
			if key[0:prefix_len] == value[0] {
				result = append(result,key)

			}

		}	
	}

	if value, ok := query_map["suffix"]; ok {

		sufix_len := len(value[0])
		for key, _ := range k.Store {
			if key[len(key)-sufix_len:] == value[0] {
				result = append(result,key)

			}

		}	
	}
	
	sort.Strings(result)


	fmt.Println(result)
	jsonBytes, err := json.Marshal(result)
	if err != nil {

	   w.WriteHeader(http.StatusInternalServerError)
	   w.Write([]byte(err.Error()))
	}
	w.Write(jsonBytes)


 

}