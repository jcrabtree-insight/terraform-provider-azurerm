# Updates

## Prevent state update triggers (lifecycle ignore changes)

### content_type
- `api_management_api_schema`: skipped (not part of key vault)
- `api_management`: skipped (not part of key vault)
- `app_configuration_key`: skipped (this can be of type `vault`, but content_type is only used for type `kv`... see [link](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/app_configuration_key#argument-reference))
- `key migration` (v0_to_v1 and v1_to_v2): skipped (see app_configuration_key explanation)
- `key_vault_certificate`: skipped (this was skipped because certificates are being removed altogether)
- `key_vault_secret`: **completed**
- `media_content_key_policy`: skipped (not part of key vault)
- `content_key_policy` (v0 to v1): skipped (not part of key vault... part of media_content_key_policy)
- `service_bus_subscription_rule`: skipped (not part of key vault)
- `storage_blob`: skipped (not part of key vault)
- `storage_share_file`: skipped (not part of key vault)
- `blob migration` (v0 to v1): skipped (not part of key vault)
- `content_types_to_compress` (in cdn resources): skipped (not part of key vault)

### effective_date
- `logz_monitor`: skipped (not part of key vault)

### not_before_date
- `key_vault_key`: skipped (this was skipped because keys are being removed altogether)
- `key_vault_secret`: **completed**

# TODO
- "It's possible to define Key Vault Certificate Contacts both within the azurerm_key_vault resource via the contact block and by using the azurerm_key_vault_certificate_contacts resource. However it's not possible to use both methods to manage Certificate Contacts within a KeyVault, since there'll be conflicts."

# Questions/Comments
- I don't see any Key Vault resources that use `effective_date`... only logz_monitor_resource uses it. Was it supposed to be `expiration_date` (which does appear in `key_vault_secret`)?
- I only modified resources... not data sources, since data sources shouldn't affect state.