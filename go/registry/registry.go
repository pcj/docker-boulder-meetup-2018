package registry

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

type Registry struct {
	// Root path where we will serve/save files to
	Path, BlobsPath, ManifestsPath string
}

// Registry implementation that saves blobs and meta to the filesystem
// at the given root path.
func NewFilesystemRegistry(root string, router *mux.Router) (*Registry,error) {
	if root == "" {
		return nil, fmt.Errorf("Root argument must not be empty")
	}

	if !FileExists(root) {
		return nil, fmt.Errorf("Root dir does not exist: %s", root)
	}

	registry := &Registry{
		Path: root,
		BlobsPath: path.Join(root, "blobs"),
		ManifestsPath: path.Join(root, "manifests"),
	}

	if err := os.MkdirAll(registry.BlobsPath, os.ModePerm); err != nil {
		return nil, fmt.Errorf("Registry preparation failed: %v", err)
	}
	if err := os.MkdirAll(registry.ManifestsPath, os.ModePerm); err != nil {
		return nil, fmt.Errorf("Registry preparation failed: %v", err)
	}
	
	// Register all routes
	router.HandleFunc("/v2/{repository}/manifests/{reference}/", registry.handleGetManifest).Methods("GET")
	router.HandleFunc("/v2/{repository}/blobs/{reference}/", registry.handleHeadBlob).Methods("HEAD")
	router.HandleFunc("/v2/{repository}/blobs/uploads/", registry.handlePostBlobUpload).Methods("POST")
	router.HandleFunc("/v2/{repository}/blobs/uploads/{reference}", registry.handlePostBlobUpload).Methods("PUT")
	router.HandleFunc("/v2/", registry.handleApiVersionCheck)

	return registry, nil
}

func (reg *Registry) handleGetManifest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	repository := vars["repository"]
	reference := vars["reference"]
	log.Printf("Manifest request repository '%s', reference '%s'\n", repository, reference)
	if repository == "" || reference == "" {
		http.NotFound(w, r)
		return
	}

	manifestPath := path.Join(reg.ManifestsPath, repository, reference, "manifest.json")
	if !FileExists(manifestPath) {
		http.NotFound(w, r)
		return
	}
	
	parts := strings.SplitN(reference, ":", 2)
	digest := parts[0]
	hash := parts[1]
	log.Printf("Manifest request digest '%s', hash '%s' \n", digest, hash)

	if digest != "sha256" || len(hash) == 40 {
		log.Printf("Request not sha256 or bad sha256: '%s', hash '%s' \n", digest, hash)
		http.NotFound(w, r)
		return
	}

	http.NotFound(w, r)
}


func (reg *Registry) handleHeadBlob(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	repository := vars["repository"]
	reference := vars["reference"]

	log.Printf("Blob request repository '%s', reference '%s'\n", repository, reference)
	if repository == "" || reference == "" {
		http.NotFound(w, r)
		return
	}
	
	parts := strings.SplitN(reference, ":", 2)
	digest := parts[0]
	hash := parts[1]
	log.Printf("Blob request digest '%s', hash '%s' \n", digest, hash)

	if digest != "sha256" || len(hash) == 40 {
		log.Printf("Request not sha256 or bad sha256: '%s', hash '%s' \n", digest, hash)
		http.NotFound(w, r)
		return
	}

	blobPath := path.Join(reg.BlobsPath, hash)
	if !FileExists(blobPath) {
		log.Printf("Blob not found: %s\n", hash)
		http.NotFound(w, r)
	}

	// 200
}


func (reg *Registry) handlePostBlobUpload(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}


func (reg *Registry) handlePutBlobUpload(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)

	repository := vars["repository"]
	reference := vars["reference"]

	log.Printf("Blob request repository '%s', reference '%s'\n", repository, reference)
	if repository == "" || reference == "" {
		http.NotFound(w, r)
		return
	}
	
	parts := strings.SplitN(reference, ":", 2)
	digest := parts[0]
	hash := parts[1]
	log.Printf("Blob request digest '%s', hash '%s' \n", digest, hash)

	if digest != "sha256" || len(hash) == 40 {
		log.Printf("Request not sha256 or bad sha256: '%s', hash '%s' \n", digest, hash)
		http.NotFound(w, r)
		return
	}

	blobPath := path.Join(reg.BlobsPath, hash)
	if FileExists(blobPath) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	
	out, err := os.Create(blobPath)
	if err != nil {
		// panic?
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer out.Close()
	io.Copy(out, r.Body)

	log.Printf("Wrote blob %s\n", blobPath)
}

// Return 200 "we implement version v2"
func (reg *Registry) handleApiVersionCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Good to go!\n"))
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
