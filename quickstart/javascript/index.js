"use strict";
const pulumi = require("@pulumi/pulumi");
const aws = require("@pulumi/aws");
const awsx = require("@pulumi/awsx");

const bucketNames =  []
for (let i = 0; i < 10; i++) {
    const bucket = new aws.s3.Bucket("my-bucket-" + i);
    bucketNames.push(bucket.bucket)
}

exports.bucketNames = bucketNames;
