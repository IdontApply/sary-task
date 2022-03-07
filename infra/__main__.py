"""An AWS Python Pulumi program"""

import pulumi
from pulumi_aws import s3, ecr
import pulumi_eks as eks

# Create an AWS resource (S3 Bucket)
bucket = s3.Bucket('my-bucket')

# Export the name of the bucket
pulumi.export('bucket_name', bucket.id)

# Create an EKS cluster with the default configuration.
cluster = eks.Cluster('my-cluster',
    eks.ClusterArgs(
        desired_capacity = 3,
        min_size = 2,
        max_size = 3,
    )
)

repo = ecr.Repository('my-repo')
image_name = repo.repository_url
registry_info = None

# Export the cluster's kubeconfig.
pulumi.export('repo', repo.repository_url)
