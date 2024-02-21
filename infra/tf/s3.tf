resource "random_string" "random" {
  length           = 6
  special          = false
  numeric = false
  upper = false
}

resource "aws_s3_bucket" "bucket" {
  bucket = "bucket-${random_string.random.result}"
}
