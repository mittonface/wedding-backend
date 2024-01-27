variable "app_name" {
    type        = string
    description = "Application Name"
}

variable "app_environment" {
    type = string
    description = "The environment"
}

variable "acm_cert_arn" {
  type = string
  default = "arn:aws:acm:us-east-1:009960124252:certificate/d4969b0d-2cc9-48c1-a996-51519eb9752f"
}

variable "aws_region" {
    type = string
    description = "Region to host the application in"
}

variable "cidr" {
  description = "The CIDR block for the VPC."
  default     = "10.0.0.0/16"
}

variable "public_subnets" {
  description = "List of public subnets"
}

variable "private_subnets" {
  description = "List of private subnets"
}

variable "availability_zones" {
  description = "List of availability zones"
}

variable "supabase_url" {
    description = "The URL that we make requests to our database through"
}

variable "supabase_secret" {
    description = "The api secret for our database"
variable "USE_SSL" {
    description = "Use ssl?"
    default = "true"
}

variable "commit_sha" {
 description = "Commit SHA from CI/CD pipeline"
 type       = string
}
