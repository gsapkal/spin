/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ConstraintState struct {

	DeliveryConfigName string `json:"deliveryConfigName,omitempty"`

	EnvironmentName string `json:"environmentName,omitempty"`

	Attributes *interface{} `json:"attributes,omitempty"`

	JudgedBy string `json:"judgedBy,omitempty"`

	Comment string `json:"comment,omitempty"`

	ArtifactVersion string `json:"artifactVersion,omitempty"`

	Type_ string `json:"type,omitempty"`

	JudgedAt string `json:"judgedAt,omitempty"`

	CreatedAt string `json:"createdAt,omitempty"`

	Status string `json:"status,omitempty"`
}
