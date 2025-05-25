package handler

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"todo-list-cli/dto"
	"todo-list-cli/service"
)

// HandleCLI handles input from the terminal
func HandleCLI() {
	// CLI argument declaration
	add := flag.Bool("add", false, "Tambah tugas baru")
	list := flag.Bool("list", false, "Lihat semua tugas")
	done := flag.Int("done", 0, "Tandai tugas selesai berdasarkan ID")
	delete := flag.Int("delete", 0, "Hapus tugas berdasarkan ID")
	search := flag.String("search", "", "Cari tugas dengan keyword")
	title := flag.String("title", "", "Judul tugas")
	desc := flag.String("desc", "", "Deskripsi tugas")

	flag.Parse() // parsing the CLI argument

	switch {
	case *add:
		req := dto.CreateTaskRequest{
			Title:       *title,
			Description: *desc,
		}
		err := service.AddTask(req)
		if err != nil {
			fmt.Println("Gagal menambahkan tugas:", err)
		} else {
			fmt.Println("Tugas berhasil ditambahkan")
		}
	case *list:
		tasks, _ := service.GetAllTasks()
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0) // show table
		fmt.Fprintln(w, "ID\tJudul\tDeskripsi\tSelesai")
		for _, t := range tasks {
			fmt.Fprintf(w, "%d\t%s\t%s\t%v\n", t.ID, t.Title, t.Description, t.Done)
		}
		w.Flush()
	case *done != 0:
		err := service.MarkDone(*done)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Tugas ditandai selesai.")
		}
	case *delete != 0:
		err := service.DeleteTask(*delete)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Tugas berhasil dihapus.")
		}
	case *search != "":
		req := dto.SearchTaskRequest{Keyword: *search}
		tasks := service.SearchTask(req)
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tJudul\tDeskripsi\tSelesai")
		for _, t := range tasks {
			fmt.Fprintf(w, "%d\t%s\t%s\t%v\n", t.ID, t.Title, t.Description, t.Done)
		}
		w.Flush()
	default:
		fmt.Println("Gunakan --add, --list, --done, --delete, atau --search untuk menjalankan perintah.")
	}
}
