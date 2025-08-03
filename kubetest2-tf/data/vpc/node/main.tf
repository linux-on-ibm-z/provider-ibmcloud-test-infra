resource "ibm_is_instance" "node" {
  name              = var.node_name
  instance_template = var.node_instance_template_id
  primary_network_interface {
    subnet          = var.subnet_id
    security_groups = [var.security_group_id]
  }
}

resource "ibm_is_floating_ip" "node" {
  name           = "${var.node_name}-ip"
  target         = ibm_is_instance.node.primary_network_interface[0].id
  resource_group = var.resource_group
}
