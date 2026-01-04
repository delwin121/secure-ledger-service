provider "aws" {
  region = "us-east-1"
}

resource "aws_s3_bucket" "ledger_bucket" {
  bucket_prefix = "ledger-service-mvp-"
  force_destroy = true
}

resource "aws_s3_bucket_public_access_block" "ledger_bucket_block" {
  bucket = aws_s3_bucket.ledger_bucket.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_s3_bucket_versioning" "ledger_bucket_versioning" {
  bucket = aws_s3_bucket.ledger_bucket.id
  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket_server_side_encryption_configuration" "ledger_bucket_encrypt" {
  bucket = aws_s3_bucket.ledger_bucket.id

  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

resource "aws_iam_policy" "bucket_read_only" {
  name        = "LedgerBucketReadOnly"
  description = "Read-only access to the ledger bucket"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "s3:ListBucket",
          "s3:GetObject"
        ]
        Resource = [
          aws_s3_bucket.ledger_bucket.arn,
          "${aws_s3_bucket.ledger_bucket.arn}/*"
        ]
      }
    ]
  })
}
