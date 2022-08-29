package main

//func test(){
//	// make a new bucket
//	opts := minio.MakeBucketOptions{
//		Region: "us-east-1",
//	}
//	bucketname, objectname := "enmotech-test", "testfile"
//	err := minioClient.MakeBucket(context.Background(),bucketname , opts)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	log.Println("Success")
//	// upload a new file from local
//	object, err := os.Open("C:\\Users\\dongdong\\Desktop\\apilog.txt")
//	if err != nil {
//		log.Fatalln(err)
//	}
//	defer object.Close()
//	objectStat, err := object.Stat()
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	n, err := minioClient.PutObject(context.Background(), bucketname, objectname , object, objectStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
//	if err != nil {
//		log.Fatalln(err)
//	}
//	log.Println("Uploaded", objectname, " of size: ", n, "Successfully.")
//
//	// list the new one
//	objopts := minio.ListObjectsOptions{
//		UseV1:     true,
//		Recursive: true,
//	}
//
//	// List all objects from a bucket-name with a matching prefix.
//	for object := range minioClient.ListObjects(context.Background(), bucketname, objopts) {
//		if object.Err != nil {
//			fmt.Println(object.Err)
//			return
//		}
//		fmt.Println(object)
//	}
//}
