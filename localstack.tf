provider "aws" {
  region = "eu-central-1"
  access_key = "test"
  secret_key = "test"
  skip_credentials_validation = true
  skip_requesting_account_id = true
  skip_metadata_api_check = true
  s3_force_path_style = true
  endpoints {
    s3 = "http://localhost:4566"
    sqs = "http://localhost:4566"
    sns = "http://localhost:4566"
  }
}

resource "aws_s3_bucket" "b" {
  bucket = "demo-bucket-terraform"
  acl    = "public-read"
}

resource "aws_sns_topic" "t"{
   name = "my-topic"
}

resource "aws_sqs_queue" "q" {
  name                      = "my-queue"
  delay_seconds             = 90
  max_message_size          = 2048
  message_retention_seconds = 86400
  receive_wait_time_seconds = 10

}
