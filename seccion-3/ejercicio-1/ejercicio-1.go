// API, concurrencia y Errors.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ItemDetails struct {
	Item
	Details string `json:"details"`
}

func raiz(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Estás en la ruta raiz")
}

func getItems(w http.ResponseWriter, r *http.Request) {
	// Establecemos el encabezado "Content-Type" de la respuesta HTTP como "application/json"
	w.Header().Set("Content-Type", "application/json")

	pageUser := r.URL.Query().Get("page")
	itemsUser := r.URL.Query().Get("itemsPerPage")

	// Convertir los parámetros a enteros
	page, err := strconv.Atoi(pageUser)
	if err != nil {
		page = 1
	}
	itemsPerPage, err := strconv.Atoi(itemsUser)
	if err != nil {
		itemsPerPage = 10
	}

	inicio := (page - 1) * itemsPerPage

	// Obtener los elementos del slice que corresponden a la página solicitada
	var resultado []Item
	if inicio >= 0 && inicio < len(items) {
		final := inicio + itemsPerPage
		if final > len(items) {
			final = len(items)
		}
		resultado = items[inicio:final]
	}

	// Función para obtener todos los elementos
	json.NewEncoder(w).Encode(resultado)
}

func getItemId(w http.ResponseWriter, r *http.Request) {
	// Establecemos el encabezado "Content-Type" de la respuesta HTTP como "application/json"
	w.Header().Set("Content-Type", "application/json")

	// Función para obtener un elemento específico
	parametros := mux.Vars(r)

	for _, item := range items {
		if item.ID == parametros["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	fmt.Fprintf(w, "Usuario no encontrado")
	json.NewEncoder(w).Encode(&Item{})

}

func getItemName(w http.ResponseWriter, r *http.Request) {
	// Establecemos el encabezado "Content-Type" de la respuesta HTTP como "application/json"
	w.Header().Set("Content-Type", "application/json")

	// Función para obtener un elemento específico
	parametros := mux.Vars(r)
	var itemsIguales []Item
	count := 0
	for _, item := range items {
		if strings.EqualFold(item.Name, parametros["name"]) {
			itemsIguales = append(itemsIguales, item)
			count++
		}
	}

	if count > 1 {
		json.NewEncoder(w).Encode(itemsIguales)
		return
	}

	fmt.Fprintf(w, "Usuario no encontrado")
	json.NewEncoder(w).Encode(&Item{})

}

func getItemDetails(id string) ItemDetails {
	// Simula la obtención de detalles desde una fuente externa con un time.Sleep
	time.Sleep(100 * time.Millisecond)
	var foundItem Item
	for _, item := range items {
		if item.ID == string(id) {
			foundItem = item
			break
		}
	}
	//Obviamente, aquí iria un SELECT si es SQL o un llamado a un servicio externo
	//pero esta busqueda del item junto con Details, la hacemos a mano.
	return ItemDetails{
		Item:    foundItem,
		Details: fmt.Sprintf("Detalles para el item %d", id),
	}
}

func getDetails(w http.ResponseWriter, r *http.Request) {
	// Establecemos el encabezado "Content-Type" de la respuesta HTTP como "application/json"
	w.Header().Set("Content-Type", "application/json")

	wg := &sync.WaitGroup{}
	detailsChannel := make(chan ItemDetails, len(items))
	var detailedItems []ItemDetails

	for _, item := range items {
		wg.Add(1) // Creamos el escucha, sin aun crearse la gorutina
		go func(id string) {
			defer wg.Done() //Completamos el trabajo del escucha, al final de esta ejecución
			detailsChannel <- getItemDetails(id)
		}(item.ID)
	}

	go func() {
		wg.Wait()
		close(detailsChannel)
	}()

	for details := range detailsChannel {
		detailedItems = append(detailedItems, details)
	}

	fmt.Println(detailedItems)
	json.NewEncoder(w).Encode(detailedItems)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	// Función para crear un nuevo elemento
	var newItem Item
	rqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte un item válido")
		return
	}

	json.Unmarshal(rqBody, &newItem)

	id := len(items) + 1
	// Agregar el nuevo elemento al slice
	newItem.ID = "item" + strconv.Itoa(id)
	items = append(items, newItem)

	// Enviar una respuesta exitosa al cliente
	// Establecemos el encabezado "Content-Type" de la respuesta HTTP como "application/json"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
	fmt.Println("¡Item creadoexitosamente!")
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	// Función para actualizar un elemento existente
	var UpdateItem Item
	parametros := mux.Vars(r)
	rqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte un item válido")
		return
	}
	json.Unmarshal(rqBody, &UpdateItem)

	for i, item := range items {
		if item.ID == parametros["id"] {
			items = append(items[:i], items[i+1:]...)
			UpdateItem.ID = parametros["id"]
			items = append(items, UpdateItem)
		}
	}

	// Enviar una respuesta exitosa al cliente
	// Establecemos el encabezado "Content-Type" de la respuesta HTTP como "application/json"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(UpdateItem)
	fmt.Println("¡Item actualizado!")
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	// Establecemos el encabezado "Content-Type" de la respuesta HTTP como "application/json"
	w.Header().Set("Content-Type", "application/json")

	// Función para eliminar un elemento
	parametros := mux.Vars(r)

	for i, item := range items {
		if item.ID == parametros["id"] {
			items = append(items[:i], items[i+1:]...)
			fmt.Fprintf(w, "La tarea con el ID %d fue eliminada", item.ID)
			return
		}
	}

	fmt.Fprintf(w, "Usuario no encontrado")
}

var items []Item

func main() {

	for i := 1; i <= 10; i++ {
		items = append(items, Item{ID: fmt.Sprintf("item%d", i), Name: fmt.Sprintf("Item %d", i)})
	}

	router := mux.NewRouter()

	router.HandleFunc("/", raiz).Methods("GET")
	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/id/{id}", getItemId).Methods("GET")
	router.HandleFunc("/items/name/{name}", getItemName).Methods("GET")
	router.HandleFunc("/items/details", getDetails).Methods("GET")
	router.HandleFunc("/items", createItem).Methods("POST")
	router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	var portNumber int = 9999
	fmt.Println("Listen in port ", portNumber)
	err := http.ListenAndServe(":"+strconv.Itoa(portNumber), router)
	if err != nil {
		fmt.Println(err)
	}

}
