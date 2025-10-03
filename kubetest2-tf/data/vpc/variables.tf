variable "vpc_api_key" {
  sensitive = true
}

variable "vpc_resource_group" {
  default = "default"
}

variable "vpc_ssh_key" {}

variable "vpc_name" {
  type        = string
  description = "Specify VPC name. If none is provided, it will create a new VPC named {cluster_name}-vpc"
  default     = ""
}

variable "node_image" {
  default = "ibm-ubuntu-22-04-2-minimal-s390x-1"
}

variable "node_profile" {
  default = "bz2-2x8"
}

variable "vpc_region" {
  description = "Denotes which IBM Cloud zone to connect to - .i.e: eu-de-1 eu-de-2  us-south etc."
}

variable "vpc_zone" {
  description = "Denotes which IBM Cloud zone to connect to - .i.e: eu-de-1 eu-de-2  us-south etc."
}
