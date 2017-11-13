// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Sample objects creates, list, deletes objects and runs
// other similar operations on them by using the Google Storage API.
// More documentation is available at
// https://cloud.google.com/storage/docs/json_api/v1/.
package main

import (
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"

	"cloud.google.com/go/storage"
)

var (
	bucket = "beta-claim-data" //default
)

func main() {
	println("staritng app..")
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8083", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {

	// bucketID := os.Getenv("GOOGLE_BUCKET_ID")
	// if bucketID == "" {
	// 	println("GOOGLE_BUCKET_ID environment variable must be set.\n")
	// }

	//Marhsal TYPE FORM DATA to TypeFormData struct
	var tranformedData = TranformedData{}
	err := json.NewDecoder(r.Body).Decode(&tranformedData)
	if err == io.EOF || err != nil {
		createErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//Storm Claim Data
	if tranformedData.StromClaimData.DamageImageURL1 != "" {
		if err := write(client, bucket,
			hex.EncodeToString([]byte(tranformedData.StromClaimData.DamageImageURL1)),
			tranformedData.StromClaimData.DamageImageURL1); err != nil {
			log.Fatalf("Cannot write object: %v", err)
		}
	}
	if tranformedData.StromClaimData.DamageImageURL2 != "" {
		if err := write(client, bucket,
			hex.EncodeToString([]byte(tranformedData.StromClaimData.DamageImageURL2)),
			tranformedData.StromClaimData.DamageImageURL2); err != nil {
			log.Fatalf("Cannot write object: %v", err)
		}
	}
	if tranformedData.StromClaimData.DamageVideoURL != "" {
		if err := write(client, bucket,
			hex.EncodeToString([]byte(tranformedData.StromClaimData.DamageVideoURL)),
			tranformedData.StromClaimData.DamageVideoURL); err != nil {
			log.Fatalf("Cannot write object: %v", err)
		}
	}
	//TV Claim Data
	if tranformedData.TVClaimData.DamageImageURL1 != "" {
		if err := write(client, bucket,
			hex.EncodeToString([]byte(tranformedData.TVClaimData.DamageImageURL1)),
			tranformedData.StromClaimData.DamageImageURL1); err != nil {
			log.Fatalf("Cannot write object: %v", err)
		}
	}
	if tranformedData.TVClaimData.DamageImageURL2 != "" {
		if err := write(client, bucket,
			hex.EncodeToString([]byte(tranformedData.TVClaimData.DamageImageURL2)),
			tranformedData.StromClaimData.DamageImageURL2); err != nil {
			log.Fatalf("Cannot write object: %v", err)
		}
	}
	if tranformedData.TVClaimData.TVReceiptImage != "" {
		if err := write(client, bucket,
			hex.EncodeToString([]byte(tranformedData.TVClaimData.TVReceiptImage)),
			tranformedData.TVClaimData.TVReceiptImage); err != nil {
			log.Fatalf("Cannot write object: %v", err)
		}
	}

	// switch operation {
	// case "write":
	// 	if err := write(client, bucket, object, url); err != nil {
	// 		log.Fatalf("Cannot write object: %v", err)
	// 	}
	// case "read":
	// 	data, err := read(client, bucket, object)
	// 	if err != nil {
	// 		log.Fatalf("Cannot read object: %v", err)
	// 	}
	// 	fmt.Printf("Object contents: %s\n", data)
	// case "metadata":
	// 	attrs, err := attrs(client, bucket, object)
	// 	if err != nil {
	// 		log.Fatalf("Cannot get object metadata: %v", err)
	// 	}
	// 	fmt.Printf("Object metadata: %v\n", attrs)
	// case "makepublic":
	// 	if err := makePublic(client, bucket, object); err != nil {
	// 		log.Fatalf("Cannot to make object public: %v", err)
	// 	}
	// case "delete":
	// 	if err := delete(client, bucket, object); err != nil {
	// 		log.Fatalf("Cannot to delete object: %v", err)
	// 	}
	// }

}

func write(client *storage.Client, bucket, object string, url string) error {
	ctx := context.Background()
	// [START upload_file]
	resp, err := http.Get(url)
	f, err := os.Open("README.md")
	if err != nil {
		return err
	}
	defer f.Close()

	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	if _, err = io.Copy(wc, resp.Body); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}
	// [END upload_file]
	return nil
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
