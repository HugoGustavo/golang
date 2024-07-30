variable "name" {
  description = "The name to identify the Group Membership"
  type        = string
}

variable "user" {
  description = "A user name to associate with the Group"
  type        = string
}

variable "group" {
  description = "The IAM Group name to attach the list of users to"
  type        = string
}