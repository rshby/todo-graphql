package todo

// create object todo
type Todo struct {
	Id        int    `json:"id"`        // property id dengan tipe data int
	Title     string `json:"title"`     // property untuk menyinpan informasi Title dengan tipe data string
	Completed bool   `json:"completed"` // property untuk menyimpan informasi sudah complete/belum, dengan tipe data boolen. nantinya berisi true/false
}
