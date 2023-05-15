import pytest
from moto import mock_s3
import boto3
from unittest.mock import MagicMock, patch

def create_bucket(bucket_name):
    s3 = boto3.client('s3')
    s3.create_bucket(Bucket=bucket_name)
    print("Bucket name created.")

def upload_file_to_s3(file_path, bucket_name, object_key):
    s3 = boto3.client('s3')
    s3.upload_file(file_path, bucket_name, object_key)

def test_create_bucket():
    with mock_s3():
        bucket_name = 'my-test-bucket'
        s3 = boto3.client('s3')
        s3.create_bucket(Bucket=bucket_name)
        
        buckets = s3.list_buckets()
        assert bucket_name in [bucket['Name'] for bucket in buckets['Buckets']]

@patch('boto3.client')
def test_upload_file_to_s3(mock_boto3_client):
    mock_s3_client = MagicMock()
    mock_boto3_client.return_value = mock_s3_client

    upload_file_to_s3('test.txt', 'mybucket', 'test.txt')

    mock_boto3_client.assert_called_once_with('s3')
    mock_s3_client.upload_file.assert_called_once_with('test.txt', 'mybucket', 'test.txt')