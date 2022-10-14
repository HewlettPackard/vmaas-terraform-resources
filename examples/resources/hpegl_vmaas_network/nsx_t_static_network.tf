# (C) Copyright 2021 Hewlett Packard Enterprise Development LP

resource "hpegl_vmaas_network" "test_net" {
  static_network {
    name         = "tf_nsx_t_static_network"
    description  = "Static Network create using tf"
    group_id     = data.hpegl_vmaas_group.default_group.id
    scope_id     = data.hpegl_vmaas_transport_zone.tf_zone.provider_id
    cidr         = "168.72.10.1/18"
    gateway      = "168.72.10.1"
    primary_dns  = "8.8.8.8"
    scan_network = false
    pool_id = 7
    active       = true
    config {
      connected_gateway = data.hpegl_vmaas_router.tier1_router.provider_id
    }
    resource_permissions {
      all = true
    }
  } 
}