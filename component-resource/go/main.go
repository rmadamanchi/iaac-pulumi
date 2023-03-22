package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an AWS resource (S3 Bucket)

		manyBuckets, err := NewManyBuckets(ctx, "my-manybuckts", 10)
		if err != nil {
			return err
		}

		ctx.Export("Buckets", manyBuckets.BucketNames)
		return nil
	})
}
