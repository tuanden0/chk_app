package v1

import (
	"chk/internal/chk_apis/v1/handlers"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/render"
)

type handler struct {
	repo *repo
}

func NewHandler(repo *repo) (h *handler) {
	return &handler{repo: repo}
}

func (h *handler) uploadCSVFile(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	defer r.Body.Close()
	r.ParseForm()

	fCSV, fHeader, err := r.FormFile("file")
	if err != nil {
		render.JSON(w, r, failedRequest(http.StatusBadRequest, "read_file_error"))
		return
	}
	defer fCSV.Close()

	fPath, uploadErr := uploadFiles(fHeader.Filename, fCSV)
	if uploadErr != nil {
		render.JSON(w, r, failedRequest(http.StatusInternalServerError, "write_file_error"))
		return
	}

	data, err := handlers.HandleCSVFile(fPath)
	if err != nil {
		render.JSON(w, r, failedRequest(http.StatusInternalServerError, err.Error()))
		return
	}

	out, err := h.repo.Create(ctx, convertRowsToCSV(data))
	if err != nil {
		log.Println(err)
		render.JSON(w, r, failedRequest(http.StatusInternalServerError, "create_data_error"))
		return
	}

	render.JSON(w, r, successResponse(out))
}

func (h *handler) listCSVData(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	in := &handlers.ListRequest{}
	if err := json.NewDecoder(r.Body).Decode(in); err != nil && err.Error() != "EOF" {
		log.Println(err)
		render.JSON(w, r, invalidRequest())
		return
	}

	d, _ := json.Marshal(in)
	log.Println(string(d))

	out, err := h.repo.List(ctx, in)
	if err != nil {
		log.Println(err)
		render.JSON(w, r, failedRequest(http.StatusInternalServerError, "fetch_data_error"))
		return
	}

	render.JSON(w, r, out)
}
