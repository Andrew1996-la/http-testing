package main

import (
	"io"
	"net/http"
	"strconv"
	"strings"
)

var cafeList = map[string][]string{
	"moscow": []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент", "Ложка и вилка"},
	"tula":   []string{"Пир и мир", "Красиво есть не запретишь", "Поздний завтрак"},
}

func mainHandle(w http.ResponseWriter, r *http.Request) {
	var err error

	// получение колличества кафе из запроса. 25 по умолчанию
	count := 25
	countStr := r.FormValue("count")
	if countStr != "" {
		count, err = strconv.Atoi(countStr)
		if err != nil {
			http.Error(w, "incorrect count", http.StatusBadRequest)
			return
		}
	}

	city := r.FormValue("city")
	cafe, ok := cafeList[city]
	if !ok {
		http.Error(w, "unknown city", http.StatusBadRequest)
		return
	}

	if search := r.FormValue("search"); search != "" {
		var found []string

		for _, v := range cafe {
			if strings.Contains(strings.ToLower(v), strings.ToLower(search)) {
				found = append(found, v)
			}
		}

		cafe = found
	}

	count = min(count, len(cafe))

	answer := strings.Join(cafe[:count], ",")
	io.WriteString(w, answer)
}

func main() {
	http.HandleFunc("/cafe", mainHandle)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
