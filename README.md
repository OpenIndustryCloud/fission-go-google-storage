# Google Storage API

API for Google Cloud Storage, to upload objects to your bucket, download objects from your bucket, and remove objects from your bucket

## API reference

 - [Google Storage API](https://cloud.google.com/storage/docs/json_api/v1/)

## Authentication

Authentication is implemented using API token, you can either configure a [Secret in Kubernetes](https://kubernetes.io/docs/concepts/configuration/secret/) 

you can generate your API [Google Storage API](https://cloud.google.com/storage/docs/json_api/v1/how-tos/authorizing)


## Error hanlding

Empty Payload or malformed JSON would result in error reponse.

- Technical error :  `{"status":400,"message":"EOF"}`
- Malformed JSON `{"status":400,"message":"invalid character 'a' looking for beginning of object key string"}`

## Sample Input/Output

- Request payload

```
{
	"status": 200,
	"ticket_details": {
		"ticket": {
			"type": "incident",
			"subject": "Storm surge claim",
			"priority": "normal",
			"status": "new",
			"comment": {
				"html_body": "<p><b>If you have any other insurance or warranties covering your home, please advise us of the company name.</b> : sd</p><hr><p><b>In as much detail as possible, please use the text box below to describe the full extent of the damage to your home and how you discovered it.</b> : sdf</p><hr><p><b>Please describe the details of the condition of your home prior to discovering the damage</b> : sd</p><hr><p><b>If there has been any recent maintenance carried out on your home, please describe it</b> : sdf</p><hr><p><b>Would you like to upload more images?</b> : </p><hr><p><b>If you have been provided with a repair estimate from a contractor or tradesman, you can upload this by providing a clear photo of the document or upload an existing file.</b> <a href='https://admin.typeform.com/form/results/file/download/H8mm3s/wFpTHm7AZVNO/913579797420-22405371_888332421323094_6861338905885136899_n.jpg'>https://admin.typeform.com/form/results/file/download/H8mm3s/wFpTHm7AZVNO/913579797420-22405371_888332421323094_6861338905885136899_n.jpg</a></p><hr><p><b>Where did the incident happen? (City/town name)</b> : London</p><hr><p><b>Would you like to upload more images?</b> : </p><hr><p><b>Are you aware of anything else relevant to your claim that you would like to advise us of at this stage?</b> : sadsa</p><hr><p><b>If it is safe and possible to do so, please provide images of the damage to both the outside and the inside of your home.</b> <a href='https://admin.typeform.com/form/results/file/download/H8mm3s/63907299/30ecfd753d05-22405371_888332421323094_6861338905885136899_n.jpg'>https://admin.typeform.com/form/results/file/download/H8mm3s/63907299/30ecfd753d05-22405371_888332421323094_6861338905885136899_n.jpg</a></p><hr><p><b>Are you still have possession of the damage items (i.e. damaged guttering)?</b> : </p><hr><p><b>When did the incident happen?</b> : 2017-10-12</p><hr><p><b>If it is safe and possible to do so, please provide images of the damage to both the outside and the inside of your home.</b> <a href='https://admin.typeform.com/form/results/file/download/H8mm3s/j79cNctIvogK/ab33a3a7524b-22405371_888332421323094_6861338905885136899_n.jpg'>https://admin.typeform.com/form/results/file/download/H8mm3s/j79cNctIvogK/ab33a3a7524b-22405371_888332421323094_6861338905885136899_n.jpg</a></p><hr><p><b>If it is safe and possible to do so, please take a short video to include any areas of damage. Hold the camera in landscape orientation and include a voice narration if you can to help explain the situation.</b> <a href='https://admin.typeform.com/form/results/file/download/H8mm3s/63907004/a870839fc865-22405371_888332421323094_6861338905885136899_n.jpg'>https://admin.typeform.com/form/results/file/download/H8mm3s/63907004/a870839fc865-22405371_888332421323094_6861338905885136899_n.jpg</a></p><hr><p><b>We have made the following assumptions about your property, you and anyone living with you</b> : </p><hr>"
			},
			"requester": {
				"locale_id": 1,
				"name": "Amit",
				"email": "amitkumarvarman@gmail.com",
				"phone": "99999999999",
				"policy_number": "DUSSS2323232"
			}
		}
	},
	"weather_api_input": {
		"city": "London",
		"date": "20171012"
	},
	"tv_claim_data": {
		"tv_reciept_image_url": ""
	},
	"storm_claim_data": {
		"incident_place": "London",
		"incident_date": "2017-10-12",
		"damage_image_url_1": "https://admin.typeform.com/form/results/file/download/H8mm3s/63907299/30ecfd753d05-22405371_888332421323094_6861338905885136899_n.jpg",
		"damage_image_url_2": "https://admin.typeform.com/form/results/file/download/H8mm3s/j79cNctIvogK/ab33a3a7524b-22405371_888332421323094_6861338905885136899_n.jpg",
		"estimate_image_url": "https://admin.typeform.com/form/results/file/download/H8mm3s/wFpTHm7AZVNO/913579797420-22405371_888332421323094_6861338905885136899_n.jpg",
		"damage_video_url": "https://admin.typeform.com/form/results/file/download/H8mm3s/63907004/a870839fc865-22405371_888332421323094_6861338905885136899_n.jpg"
	},
	"weather-api-input": {
		"city": "birmingham",
		"country": "",
		"date": "20170101"
	}
}
```
- Response

```

{
    "status": 200,
    "media": [
        {
            "bucket": "artifacts-image",
            "name": "nyTbYLIlqA",
            "size": 384461,
            "media-link": "https://www.googleapis.com/download/storage/v1/b/artifacts-image/o/nyTbYLIlqA?generation=1513187644653616&alt=media",
            "original-link": "https://admin.typeform.com/form/results/file/download/H8mm3s/63907299/f550ae20714c-1111.jpg"
        },
        {
            "bucket": "artifacts-image",
            "name": "xuhPPyZTSy",
            "size": 384001,
            "media-link": "https://www.googleapis.com/download/storage/v1/b/artifacts-image/o/xuhPPyZTSy?generation=1513187647809516&alt=media",
            "original-link": "https://admin.typeform.com/form/results/file/download/H8mm3s/j79cNctIvogK/b05201c8626b-tree_falls_on_house.png"
        }
    ]
}

-----
on duplicate invocation 

{
    "status": 208,
    "message": "Image Exist in Storage"
}
----------

```


## Example Usage

## 1.  Deploy as Fission Functions

First, set up your fission deployment with the go environment.

```
fission env create --name go-env --image fission/go-env:1.8.1
```

To ensure that you build functions using the same version as the
runtime, fission provides a docker image and helper script for
building functions.


- Download the build helper script

```
$ curl https://raw.githubusercontent.com/fission/fission/master/environments/go/builder/go-function-build > go-function-build
$ chmod +x go-function-build
```

- Build the function as a plugin. Outputs result to 'function.so'

`$ go-function-build google-storage.go`

- Upload the function to fission

`$ fission function create --name google-storage --env go-env --package function.so`

- Map /google-storage to the google-storage function

`$ fission route create --method POST --url /google-storage --function google-storage`

- Run the function

```$ curl -d `sample request` -H "Content-Type: application/json" -X POST http://$FISSION_ROUTER/google-storage```

## 2. Deploy as AWS Lambda

> to be updated
