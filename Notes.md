# Updates

## Remove Code

Remove code related to Certificates, Keys, Access Policies, Hardware Security Modules, and Managed Storage

### Deleted files
- internal/services/keyvault/key_vault_access_policy_resource.go
- internal/services/keyvault/key_vault_access_policy_resource_test.go
- internal/services/keyvault/key_vault_certificate_contacts_resource.go
- internal/services/keyvault/key_vault_certificate_contacts_resource_test.go
- internal/services/keyvault/key_vault_certificate_issuer_resource.go
- internal/services/keyvault/key_vault_certificate_resource.go
- internal/services/keyvault/key_vault_key_resource.go
- internal/services/keyvault/key_vault_managed_hardware_security_module_resource.go
- internal/services/keyvault/key_vault_managed_storage_account.go
- internal/services/keyvault/key_vault_managed_storage_account_test.go
- internal/services/keyvault/key_vault_managed_storage_account_sas_token_definition_resource.go
- internal/services/keyvault/key_vault_managed_storage_account_sas_token_definition_resource_test.go

### Modified files
- internal/provider/features.go
   - Removed the following settings:
      - purge_soft_deleted_certificates_on_destroy
      - purge_soft_deleted_hardware_security_modules_on_destroy
      - purge_soft_deleted_keys_on_destroy
      - recover_soft_deleted_keys
      - recover_soft_deleted_certificates

- internal/services/keyvault/registration.go
   - Removed the following resources:
      - azurerm_key_vault_access_policy
      - azurerm_key_vault_certificate
      - azurerm_key_vault_certificate_issuer
      - azurerm_key_vault_key
      - azurerm_key_vault_managed_hardware_security_module
      - azurerm_key_vault_managed_storage_account
      - azurerm_key_vault_managed_storage_account_sas_token_definition
      - KeyVaultCertificateContactsResource

### Other files

These resource files were kept because they contain types and functions that are used by the corresponding data-source files.

- internal/services/keyvault/key_vault_certificate_issuer_resource_test.go
- internal/services/keyvault/key_vault_certificate_resource_test.go
- internal/services/keyvault/key_vault_key_resource_test.go
- internal/services/keyvault/key_vault_managed_hardware_security_module_resource_test.go

## Prevent State Updates

Add logic to prevent changes to the following attributes from triggering state updates: content_type, effective_date, not_before_date

### DiffSuppressFunc to suppress state changes

Created file: internal/tf/suppress/always.go

File contains function to always return true when schema is compared. This function is used with the `content_type`, `effective_date`, and `not_before_date` fields.

```go
// Always is a SchemaDiffSuppressFunc that always returns true.
// It can be used for the DiffSuppressFunc field in schema values.
// When used, Terraform will ignore any changes to that schema value when a plan is created.
func Always(_, _, _ string, _ *schema.ResourceData) bool {
	return true
}
```

Also, used `DiffSuppressOnRefresh: true` to ignore changes during state refreshes (in addition to plans).

### Attribute changes
- content_type
   - `key_vault_certificate`: Skipped (this was skipped because certificates are being removed altogether)
   - `key_vault_secret`: Added `DiffSuppressFunc` and `DiffSuppressOnRefresh`
- effective_date
   - This field is not used in any Key Vault resources
- not_before_date
   - `key_vault_key`: Skipped (this was skipped because keys are being removed altogether)
   - `key_vault_secret`: Added `DiffSuppressFunc` and `DiffSuppressOnRefresh`

# TODO
- Update tests
- Update examples


# Questions/Comments
- I don't see any Key Vault resources that use `effective_date`... only the logz_monitor_resource uses it.
   - Was it supposed to be `expiration_date` (which does indeed appear in `key_vault_secret`)?
- I only modified resources so far... not data sources, since data sources shouldn't affect state. Should I also remove data sources?
- Clarify "Update Secrets code to allow for no or ‘dummy’ value and eliminate value from persistent state storage":
   - Does that mean allow for options of real secret, dummy secret, and empty/no secret?
   - Does it also say to delete the secret from the state file? (If so, then we probably don't need "real" secrets)
   - Does Azure allow an empty secret? If not, should I ignore the error that is returned from the API?
- Clarify "Where certain internal functions may have secrets dependencies, those functions will be made private"
   - Can you give hypothetical example?
   - Does "secrets" mean just secrets... or also certificates, keys, etc.?
- Just to confirm... I should delete code, not just hide/disable it? Disabling might make it easier to pull and merge updates from upstream repo.
- Just to confirm... we're keeping secrets and removing certificates, keys, access policies, hardware security modules, and managed storage, right?