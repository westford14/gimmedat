variable "bucket" {
  type        = string
  description = "the name of the bucket"
}

variable "versioning" {
  type        = bool
  default     = true
  description = "(Optional, Default: true) Versioning state. Versioning is a means of keeping multiple variants of an object in the same bucket"
}

variable "block_public_acls" {
  type        = bool
  default     = true
  description = "(Optional) Whether Amazon S3 should block public ACLs for this bucket. Defaults to true."
}

variable "force_destroy" {
  type        = bool
  default     = false
  description = "(Optional) Whether or not to force destroy the bucket upon terraform destroy"
}