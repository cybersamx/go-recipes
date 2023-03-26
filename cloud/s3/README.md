# List a S3 Bucket

An example that demonstrates the use of AWS Go SDK to list a public S3 bucket. There are some public S3 buckets, like those in [AWS Open Data Sponsorship Program](https://aws.amazon.com/opendata), that only allows access if you provide your access key id and access secret key when calling the AWS SDK.

These public S3 buckets are the ones I found that allow anonymous access.

* Snowflake Workshop Lab
* [Radiant MLHub](https://aws.amazon.com/marketplace/pp/prodview-yvrd3g43ui7ms)
* [Open Observatory of Network Interference (OONI)](https://aws.amazon.com/marketplace/pp/prodview-zabuwupohnpl4)

This program's output tries to emulate the output of `aws s3 ls` as closely as possible.

## Notes

While S3 bucket looks like a filesystem, it is rally a flat object storage system. Folders don't exist in S3. Each object, usually a blob object like a file, has a key, which you can namespace it by appending a prefix to the key name. So appending say `folder/subfolder/file` emulates a typical filesystem.

## Setup

1. Run the program.

   ```bash
   $ make run
   ```

1. We can also compare the result with the output of `aws s3 ls s3://snowflake-workshop-lab --no-sign-request` if awscli is installed in your system.

## Reference

* [AWS Open Data Sponsorship Program](https://aws.amazon.com/opendata)
* [Github: aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2)
* [Wikipedia: multi-byte unit](https://en.wikipedia.org/wiki/Byte#Multiple-byte_units)
