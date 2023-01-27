import sys
from awsglue.transforms import *
from awsglue.utils import getResolvedOptions
from pyspark.context import SparkContext
from awsglue.context import GlueContext
from awsglue.job import Job

args = getResolvedOptions(sys.argv, ["JOB_NAME"])
sc = SparkContext()
glueContext = GlueContext(sc)
spark = glueContext.spark_session
job = Job(glueContext)
job.init(args["JOB_NAME"], args)

# # Script generated for node S3 bucket
# S3bucket_node1 = glueContext.create_dynamic_frame.from_options(
#     format_options={},
#     connection_type="s3",
#     format="parquet",
#     connection_options={"paths": ["s3://prtstoryprod-main-snapshot"], "recurse": True},
#     transformation_ctx="S3bucket_node1",
# )

# # Script generated for node ApplyMapping
# ApplyMapping_node2 = ApplyMapping.apply(
#     frame=S3bucket_node1,
#     mappings=[
#         ("exportTaskIdentifier", "string", "exportTaskIdentifier", "string"),
#         ("sourceArn", "string", "sourceArn", "string"),
#         ("exportOnly", "array", "exportOnly", "array"),
#         ("snapshotTime", "string", "snapshotTime", "string"),
#         ("taskStartTime", "string", "taskStartTime", "string"),
#         ("taskEndTime", "string", "taskEndTime", "string"),
#         ("s3Bucket", "string", "s3Bucket", "string"),
#         ("s3Prefix", "string", "s3Prefix", "string"),
#         ("exportedFilesPath", "string", "exportedFilesPath", "string"),
#         ("iamRoleArn", "string", "iamRoleArn", "string"),
#         ("kmsKeyId", "string", "kmsKeyId", "string"),
#         ("status", "string", "status", "string"),
#         ("percentProgress", "int", "percentProgress", "int"),
#         ("totalExportedDataInGB", "double", "totalExportedDataInGB", "double"),
#     ],
#     transformation_ctx="ApplyMapping_node2",
# )

# # Script generated for node S3 bucket
# S3bucket_node3 = glueContext.write_dynamic_frame.from_options(
#     frame=ApplyMapping_node2,
#     connection_type="s3",
#     format="json",
#     connection_options={"path": "s3://story-test-bigquery", "partitionKeys": []},
#     transformation_ctx="S3bucket_node3",
# )

print('sample-job! Update Success!')

job.commit()
