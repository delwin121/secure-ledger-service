output "bucket_name" {
  value = aws_s3_bucket.ledger_bucket.id
}

output "policy_arn" {
  value = aws_iam_policy.bucket_read_only.arn
}
