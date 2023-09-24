package todo

import "errors"

// global variabel. digunakan untuk menyimpan object todo karena kita belum menggunakan database
var TodoItems []Todo

// buat sebuah struct repository
type TodoRepository struct {
	Data []Todo // property data, untuk menyinpan object karena kita belum menggunakan database
}

// buat function provider untuk membuat object baru dari struct TodoRepository (sepertu constructor)
func NewTodoRepository() *TodoRepository {
	return &TodoRepository{ // return pointer dari object TodoRepository yang dibuat
		Data: []Todo{}, // isi property Data dengan slice kosong (nanti diisi saat insert)
	}
}

// method get data Todo by id
func (t *TodoRepository) GetTodoById(inputId int) (*Todo, error) {
	for id, item := range t.Data { // looping semua data object Todo yang ada
		if id == inputId { // jika ada data yang IDnya sesuai dengan id yang kita pilih
			return &item, nil // returnkan data tersebut, error dibikin nil (tidak ada error)
		}
	}

	return nil, errors.New("record not found") // jika data tidak ketemu, returnkan error
}
