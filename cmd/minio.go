package cmd
import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"minio/api"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	Path string
	Prefix string
)


var minioClient *minio.Client
var MinioCmd = &cobra.Command{
	Use:   "minio",
	Short: "Operations of objects with minio.",
	Long: `MINIO allows you to make operations of objects with minio. For example:
    mgo minio bucket
    mgo minio object`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("minio called")
		if len(args) == 0 {
			fmt.Println(`Error: You must specify the type of resource to operate.  Valid resource types include:
    * bucket
    * object`)
		} else {
			switch args[0] {
			case "bucket", "object":
				break
			default:
				fmt.Println(`Error: You must specify the type of resource to operate.  Valid resource types include:
    * bucket
    * object`)
			}
		}
	},
}
var bucketCmd = &cobra.Command{
	Use:   "bucket",
	Short: "Operations of bucket with minio.",
	Long: `bucket allows you to make operations of buckets with minio. For example:
    mgo minio bucket ls
    mgo minio bucket make
	mgo minio bucket rm`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("bucket called")
		if len(args) == 0 {
			fmt.Println(`Error: You must specify the type of operation for bucket.  Valid operation types include:
    * ls
    * make
	* rm`)
		} else {
			switch args[0] {
			case "ls", "make", "rm":
				break
			default:
				fmt.Println(`Error: You must specify the type of operation for bucket.  Valid operation types include:
    * ls
    * make
	* rm`)
			}
		}
	},
}
var objectCmd = &cobra.Command{
	Use:   "object",
	Short: "Operations of object with minio.",
	Long: `bucket allows you to make operations of objects with minio. For example:
    mgo minio object ls
    mgo minio object dl
	mgo minio object ul`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("bucket called")
		if len(args) == 0 {
			fmt.Println(`Error: You must specify the type of operation for bucket.  Valid operation types include:
    * ls
    * make
	* rm`)
		} else {
			switch args[0] {
			case "ls", "make", "rm":
				break
			default:
				fmt.Println(`Error: You must specify the type of operation for bucket.  Valid operation types include:
    * ls
    * make
	* rm`)
			}
		}
	},
}

var listBucketCmd = &cobra.Command{
	Use:   "ls",
	Short: "List Minio buckets",
	Long: `List all buckets of a minio. For example:

    mgo minio bucket ls`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("bucket ls called")

		if len(args) > 0 {
			fmt.Println(`Error: No parameter is required for this command.`)
			os.Exit(1)
		}

		api.GetBucketList(minioClient)
	},
}

var makeBucketCmd = &cobra.Command{
	Use:   "make",
	Short: "Make Minio buckets",
	Long: `Make a new bucket in minioClient path. For example:

    mgo minio bucket make bucket-name`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("bucket make called")

		if len(args) != 1 {
			fmt.Println(`Error: A single bucket name is required for this command.`)
			os.Exit(1)
		}

		api.MakeBucket(minioClient, args[0])
	},
}

var removeBucketCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove Minio buckets",
	Long: `Remove a bucket in minioClient path. For example:

    mgo minio bucket rm bucket-name`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("bucket rm called")

		if len(args) != 1 {
			fmt.Println(`Error: A single bucket name is required for this command.`)
			os.Exit(1)
		}

		api.RemoveBucket(minioClient, args[0])
	},
}


var listObjectsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List Minio objects",
	Long: `List all buckets of a bucket of a minio with prefix. For example:

    mgo minio object ls bucket-name --prefix prefix-name`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("object ls called")

		if len(args) < 1 {
			fmt.Println(`Error: At least a bucket name is required for this command.`)
			os.Exit(1)
		}
		api.ListObjects(minioClient, args[0], Prefix)
	},
}

var getObjectFileCmd = &cobra.Command{
	Use:   "getf",
	Short: "Get Minio object as a file.",
	Long: `Get object file with a bucket-name of a minio as a file(filepath could be omit). For example:

    mgo minio object getf bucket-name object-name -p filepath`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("object getf called")
		if len(args) < 2 {
			fmt.Println(`Error: At least bucket-name and object-name is required for this command.`)
			os.Exit(1)
		}
		if Path == "" {
			Path = args[1]
		}
		api.GetObjectFile(minioClient, args[0], args[1], Path)
	},
}
var putObjectFileCmd = &cobra.Command{
	Use:   "putf",
	Short: "Upload a file as a Minio object.",
	Long: `Upload a file as a Minio object with a bucket-name. For example:

    mgo minio object putf bucket-name object-name -p filepath`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("object putf called")
		if len(args) < 2 {
			fmt.Println(`Error: At least bucket-name and object-name is required for this command.`)
			os.Exit(1)
		}
		api.PutObjectFile(minioClient, args[0], args[1], Path)
	},
}
var statObjectCmd = &cobra.Command{
	Use:   "stat",
	Short: "Get status of a Minio object .",
	Long: `Get status of a object with a bucket-name of a minio. For example:

    mgo minio object stat bucket-name object-name`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("object stat called")
		if len(args) < 2 {
			fmt.Println(`Error: At least bucket-name and object-name is required for this command.`)
			os.Exit(1)
		}
		api.StatObject(minioClient, args[0], args[1])
	},
}
var removeObjectCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a Minio object .",
	Long: `Remove a object with a bucket-name of a minio. For example:

    mgo minio object rm bucket-name object-name`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("object stat called")
		if len(args) < 2 {
			fmt.Println(`Error: At least bucket-name and object-name is required for this command.`)
			os.Exit(1)
		}
		api.RemoveObject(minioClient, args[0], args[1])
	},
}
var removeObjectsCmd = &cobra.Command{
	Use:   "rmall",
	Short: "Removeall Minio object .",
	Long: `Remove all objects with a bucket-name of a minio. For example:

    mgo minio object rm bucket-name`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("object stat called")
		if len(args) < 1 {
			fmt.Println(`Error: At least bucket-name is required for this command.`)
			os.Exit(1)
		}
		api.RemoveObjects(minioClient, args[0], Prefix)
	},
}


func init() {
	minioClient = api.GetClient()
	RootCmd.AddCommand(MinioCmd)

	MinioCmd.AddCommand(bucketCmd)
	bucketCmd.AddCommand(listBucketCmd)
	bucketCmd.AddCommand(makeBucketCmd)
	bucketCmd.AddCommand(removeBucketCmd)



	MinioCmd.AddCommand(objectCmd)
	objectCmd.AddCommand(listObjectsCmd)
	objectCmd.AddCommand(getObjectFileCmd)
	objectCmd.AddCommand(putObjectFileCmd)
	objectCmd.AddCommand(statObjectCmd)
	objectCmd.AddCommand(removeObjectCmd)
	objectCmd.AddCommand(removeObjectsCmd)

	listObjectsCmd.Flags().StringVarP(&Prefix, "prefix", "", "", "The prefix of objects. If specified, overrides the value default.")
	getObjectFileCmd.Flags().StringVarP(&Path, "path", "p", "", "The path to store object(like /home/xxx/xxx.sh). If specified, overrides the value default.")
	putObjectFileCmd.Flags().StringVarP(&Path, "path", "p", "", "The path of file to upload (like /home/xxx/xxx.sh). If specified, overrides the value default.")
	removeObjectsCmd.Flags().StringVarP(&Prefix, "prefix", "", "", "The prefix of objects. If specified, overrides the value default.")
}