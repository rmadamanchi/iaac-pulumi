package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an AWS resource (S3 Bucket)

		buckets := make([]*s3.Bucket, 0)
		for i := 0; i < 20; i++ {
			bucket, err := s3.NewBucket(ctx, "my-bucket-"+strconv.Itoa(i), nil)
			if err != nil {
				return err
			}
			buckets = append(buckets, bucket)
		}

		// Export the name of the bucket
		bucketIds := make([]pulumi.StringOutput, 0)
		for _, bucket := range buckets {
			bucketIds = append(bucketIds, bucket.Arn)
		}

		ctx.Export("Buckets", pulumi.ToStringArrayOutput(bucketIds))
		return nil
	})
}
