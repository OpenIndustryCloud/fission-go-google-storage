// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Sample objects creates, list, deletes objects and runs
// other similar operations on them by using the Google Storage API.
// More documentation is available at
// https://cloud.google.com/storage/docs/json_api/v1/.
package main

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	bucketId    = "artifacts-image" //default
)

func main() {
	println("staritng app..")
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8083", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {

	println("In Google Storage APP")
	// bucketID := os.Getenvgo("GOOGLE_BUCKET_ID")
	// if bucketID == "" {
	// 	println("GOOGLE_BUCKET_ID environment variable must be set.\n")
	// }

	//Marhsal TYPE FORM DATA to TypeFormData struct
	var tranformedData TranformedData
	err := json.NewDecoder(r.Body).Decode(&tranformedData)
	if err == io.EOF || err != nil {
		createErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	//create a client:
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	mediaBucket := MediaBucket{}
	//Storm Claim Data
	if (StromClaimData{}) != tranformedData.StromClaimData {
		if imgURL := tranformedData.StromClaimData.DamageImageURL1; imgURL != "" {
			if media, err := write(client, bucketId,
				RandStringRunes(10), imgURL); err == nil {
				mediaBucket.addAttachment(media)
			}
		}
		if imgURL := tranformedData.StromClaimData.DamageImageURL2; imgURL != "" {
			if media, err := write(client, bucketId,
				RandStringRunes(10), imgURL); err == nil {
				mediaBucket.addAttachment(media)
			}
		}
		if imgURL := tranformedData.StromClaimData.DamageVideoURL; imgURL != "" {
			if media, err := write(client, bucketId,
				RandStringRunes(10), imgURL); err == nil {
				mediaBucket.addAttachment(media)
			}
		}
	}
	//TV Claim Data
	if (TVClaimData{}) != tranformedData.TVClaimData {
		if imgURL := tranformedData.TVClaimData.DamageImageURL1; imgURL != "" {
			if media, err := write(client, bucketId,
				RandStringRunes(10), imgURL); err == nil {
				mediaBucket.addAttachment(media)
			}
		}
		if imgURL := tranformedData.TVClaimData.DamageImageURL2; imgURL != "" {
			if media, err := write(client, bucketId,
				RandStringRunes(10), imgURL); err == nil {
				mediaBucket.addAttachment(media)
			}
		}
		if imgURL := tranformedData.TVClaimData.TVReceiptImage; imgURL != "" {
			if media, err := write(client, bucketId,
				RandStringRunes(10), imgURL); err == nil {
				mediaBucket.addAttachment(media)
			}
		}
	}

	//add status 200 if no error
	if len(mediaBucket.Media) != 0 {
		mediaBucket.Status = 200
	} else {
		createErrorResponse(w, "No data Uploaded", http.StatusBadRequest)
	}

	//marshal to JSON
	mediaBucketJSON, err := json.Marshal(mediaBucket)
	if err != nil {
		createErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	println("Google Storage APP output : ", string(mediaBucketJSON))

	w.Header().Set("content-type", "application/json")
	w.Write([]byte(string(mediaBucketJSON)))
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func write(client *storage.Client, bucket, objectName string, url string) (Media, error) {
	println("writing media to cloud storage")
	ctx := context.Background()
	// [START upload_file]
	resp, err := http.Get(url)
	if err != nil {
		return Media{}, err
	}

	obj := client.Bucket(bucket).Object(objectName)
	wc := obj.NewWriter(ctx)
	if _, err = io.Copy(wc, resp.Body); err != nil {
		return Media{}, err
	}
	// Close, just like writing a file.
	if err := wc.Close(); err != nil {
		return Media{}, err
	}

	//Make Object Public
	acl := obj.ACL()
	if err := acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return Media{}, err
	}

	//read attribute
	objAttrs, err := obj.Attrs(ctx)
	if err != nil {
		return Media{}, err
	}
	println("object %s has size %d and can be read using %s\n",
		objAttrs.Name, objAttrs.Size, objAttrs.MediaLink)

	return Media{
		objAttrs.Bucket,
		objAttrs.Name,
		objAttrs.Size,
		objAttrs.MediaLink,
		url,
	}, err
}

type MediaBucket struct {
	Status int     `json:"status,omitempty"`
	Media  []Media `json:"media,omitempty"`
}

type Media struct {
	Bucket       string `json:"bucket,omitempty"`
	Name         string `json:"name,omitempty"`
	Size         int64  `json:"size,omitempty"`
	MediaLink    string `json:"media-link,omitempty"`
	OriginalLink string `json:"original-link,omitempty"`
}

func (mediaBucket *MediaBucket) addAttachment(media Media) {
	mediaBucket.Media = append(mediaBucket.Media, media)
}

func read(client *storage.Client, bucket, object string) ([]byte, error) {
	ctx := context.Background()
	// [START download_file]
	rc, err := client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return data, nil
	// [END download_file]
}

func attrs(client *storage.Client, bucket, object string) (*storage.ObjectAttrs, error) {
	ctx := context.Background()
	// [START get_metadata]
	o := client.Bucket(bucket).Object(object)
	attrs, err := o.Attrs(ctx)
	if err != nil {
		return nil, err
	}
	return attrs, nil
	// [END get_metadata]
}

func makePublic(client *storage.Client, bucket, object string) error {
	ctx := context.Background()
	// [START public]
	acl := client.Bucket(bucket).Object(object).ACL()
	if err := acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return err
	}
	// [END public]
	return nil
}

func delete(client *storage.Client, bucket, object string) error {
	ctx := context.Background()
	// [START delete_file]
	o := client.Bucket(bucket).Object(object)
	if err := o.Delete(ctx); err != nil {
		return err
	}
	// [END delete_file]
	return nil
}

func createErrorResponse(w http.ResponseWriter, message string, status int) {
	errorJSON, _ := json.Marshal(&Error{
		Status:  status,
		Message: message})
	//Send custom error message to caller
	w.WriteHeader(status)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(errorJSON))
}

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

//output data
type TranformedData struct {
	Status         int            `json:"status,omitempty"`
	TVClaimData    TVClaimData    `json:"tv_claim_data,omitempty"`
	StromClaimData StromClaimData `json:"storm_claim_data,omitempty"`
}

type TVClaimData struct {
	TVPrice         string `json:"tv_price,omitempty"`
	CrimeRef        string `json:"crime_ref,omitempty"`
	IncidentDate    string `json:"incident_date,omitempty"`
	TVModelNo       string `json:"tv_model_no,omitempty"`
	TVMake          string `json:"tv_make,omitempty"`
	TVSerialNo      string `json:"tv_serial_no,omitempty"`
	DamageImageURL1 string `json:"damage_image_url_1,omitempty"`
	DamageImageURL2 string `json:"damage_image_url_2,omitempty"`
	TVReceiptImage  string `json:"tv_reciept_image_url"`
	DamageVideoURL  string `json:"damage_video_url,omitempty"`
}

type StromClaimData struct {
	IncidentPlace       string `json:"incident_place,omitempty"`
	IncidentDate        string `json:"incident_date,omitempty"`
	DamageImageURL1     string `json:"damage_image_url_1,omitempty"`
	DamageImageURL2     string `json:"damage_image_url_2,omitempty"`
	RepairEstimateImage string `json:"estimate_image_url,omitempty"`
	DamageVideoURL      string `json:"damage_video_url,omitempty"`
}