output "BUCKET_URL" {
  value = "https://${aws_s3_bucket.bucket.bucket_regional_domain_name}"
}
