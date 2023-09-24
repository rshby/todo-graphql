package todo

import "errors"

// global variabel. digunakan untuk menyimpan object todo karena kita belum menggunakan database
var TodoItems []Todo

var count = 1 // global variabel untuk incrmeent property Id

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

// method to get all data todos
func (t *TodoRepository) GetAllTodos() ([]Todo, error) {
	if len(t.Data) == 0 { // cek jika panjang data yang ada di slice 0, atau slicenya masih kosong
		return nil, errors.New("record not found") // returnkan error bahwa data tidak ditemukan
	}

	return t.Data, nil // jika ada datanya, returnkan semua data
}

// method add/insert Todo
func (t *TodoRepository) Add(inputTitle string) (*Todo, error) {
	newData := Todo{ // buat object baru dari struct Todo
		Id:        count,      // isi value property Id dengan nilai count
		Title:     inputTitle, // isi value property Title dengan inputan kita
		Completed: false,      // isi value property completed dengan false
	}

	t.Data = append(t.Data, newData) // append atau tambahkan object baru tersebut ke slice of Todo (t.Data)
	count++                          // naikkan increment, supaya add data berikutnya countnya bertambah
	return &newData, nil             // success add data, returnkan object baru kita tadi, dan errornya nil
}

// method untuk update
func (t *TodoRepository) Update(inputId int, inputTitle string, inputCompleted bool) (*Todo, error) {
	var updatedData *Todo // buat variabel kosong untuk menampung hasil data yang sudah diupdate, tipe datanya pointer struct Todo

	for idx, item := range t.Data { // looping semua value slice yang ada, mencari id yang sesuai
		if idx == inputId { // jika object dengan value id sesuai dengan yang ingin kita update
			item.Title = inputTitle         // ganti value property Title menjadi inputan kita
			item.Completed = inputCompleted // ganti value property completed meenjadi inputan kita

			updatedData = &item // assign variabel kosong tadi dengan object dengan value yang sudah diupdate
			break               // break untuk keluar looping karena sudah ketemu datanya, tidak perlu looping sisa object
		}
	}

	return updatedData, nil // returnkan object baru dan errornya nil
}
