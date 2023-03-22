package main

import (
	"fmt"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManyBuckets struct {
	pulumi.ResourceState
	BucketNames pulumi.StringArrayOutput
}

func NewManyBuckets(ctx *pulumi.Context, name string, count int, opts ...pulumi.ResourceOption) (*ManyBuckets, error) {
	component := &ManyBuckets{}
	err := ctx.RegisterComponentResource("rmadamanchi:hello:ManyBuckets", name, component, opts...)
	if err != nil {
		return nil, err
	}

	buckets := make([]*s3.Bucket, 0)
	for i := 0; i < count; i++ {
		bucket, err := s3.NewBucket(ctx, fmt.Sprintf("%s-%d", name, i),
			&s3.BucketArgs{}, pulumi.Parent(component))
		if err != nil {
			return nil, err
		}
		buckets = append(buckets, bucket)
	}

	bucketNames := make([]pulumi.StringOutput, 0)
	for _, bucket := range buckets {
		bucketNames = append(bucketNames, bucket.Bucket)
	}

	_ = ctx.RegisterResourceOutputs(component, pulumi.Map{
		"bucketNames": pulumi.ToStringArrayOutput(bucketNames),
	})

	component.BucketNames = pulumi.ToStringArrayOutput(bucketNames)

	return component, nil
}
