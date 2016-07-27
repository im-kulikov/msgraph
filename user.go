package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/guregu/null"
)

func init() {
	resources["UserV1"] = Resource{"User", APIVersionV1, "users"}
}

type AssignedLicense struct {
	// A collection of the unique identifiers for plans that have
	// been disabled.
	DisabledPlans []string `json:"disabledPlans"`

	// The unique identifier for the SKU.
	SKU string `json:"skuId"`
}

type AssignedPlan struct {
	// The date and tiem at which the plan was assigned, in UTC.
	AssignedDateTime time.Time `json:"assignedDateTime"`

	// Whether the plan is enabled. For example, "Enabled".
	CapabilityStatus string `json:"capabilityStatus"`

	// The name of the service.
	Service string `json:"service"`

	// A GUID that identifies the service plan.
	ServicePlanId string `json:"servicePlanID"`
}

type PasswordProfile struct {
	// Indicates whether the user must change their password on the
	// next login.
	ForceChangePasswordOnNextSignIn bool `json:"forceChangePasswordNextSignIn"`

	// The password for the user. This property is required when a
	// user is created. It can be updated, but the user will be
	// required to change the password on the next login. The
	// password must satisfy minimum requirements as specified by
	// the user's PasswordPolicies property. By default, a strong
	// password is required.
	Password string `json:"password"`
}

type ProvisionedPlan struct {
	// Whether or not this plan is enabled. For example, "Enabled".
	CapabilityStatus string `json:"capabilityStatus"`

	// The provisioning status of this plan. For example, "Success".
	ProvisioningStatus string `json:provisioningStatus"`

	// The name of the service; for example, "AccessControlS2S".
	Service string `json:"service"`
}

type ProxyAddress struct {
	// The address' protocol.
	Protocol string

	// The address.
	Address string
}

type User struct {
	// A freeform text entry field for the user to describe
	// themselves.
	AboutMe string `json:"aboutMe,omitempty"`

	// Whether or not the account is enabled. This property is
	// required when a user is created.
	AccountEnabled null.Bool `json:"accountEnabled,omitempty"`

	// The licenses that are assigned to the user. Not nullable.
	AssignedLicenses []*AssignedLicense `json:"assignedLicenses,omitempty"`

	// The plans that are assigned to the user. Read-only. Not
	// nullable.
	AssignedPlans []*AssignedPlan `json:"assignedPlans,omitempty"`

	// The birthday of the user, in UTC.
	Birthday null.Time `json:"birthday,omitempty"`

	// The city in which the user is located.
	City string `json:"city,omitempty"`

	// The country or regiion in which the user is located.
	Country string `json:"country,omitempty"`

	// The name for the department in which the user works.
	Department string `json:"department,omitempty"`

	// The name displayed in the address book for the user. This is
	// usually the combination of the user's first name, middle
	// initial and last name. This property is required when a user
	// is created and it cannot be cleared during updates.
	DisplayName string `json:"displayName,omitempty"`

	// The given name (first name) of the user.
	GivenName string `json:"givenName,omitempty"`

	// The hire date of the user, in UTC.
	HireDate null.Time `json:"hireDate,omitempty"`

	// The unique identifier for the user. Not nullable. Read-only.
	ID string `json:"id,omitempty"`

	// A list for the user to describe their interests.
	Interests []string `json:"interests,omitempty"`

	// The user's job title.
	JobTitle string `json:"jobTitle,omitempty"`

	// The SMTP address for the user. Read-only.
	EmailAddress string `json:"mail,omitempty"`

	// The mail alias for the user. This property must be specified
	// when a user is created.
	MailNickname string `json:"mailNickname,omitempty"`

	// The primary cellular telephone number for the user.
	MobilePhone string `json:"mobilePhone,omitempty"`

	// The URL for the user's personal site.
	MySite string `json:"mySite,omitempty"`

	// The office location in the user's place of business.
	OfficeLocation string `json:"officeLocation,omitempty"`

	// This property is used to associate an on-premises Active
	// Directory user account to their Azure AD user object. This
	// property must be specified when createing a new user account
	// in the Graph if you are a federated domain for the user's
	// userPrincipalName (UPN) property. Important: the $ and _
	// characters cannot be used when specifying this property.
	OnPremisesImmutableID string `json:"onPremisesImmutableId,omitempty"`

	// Indicates the last time at which the object was synced with
	// the on-premises directory, in UTC. Read-only.
	OnPremisesLastSyncDateTime null.Time `json:"onPremisesLastSyncDateTime,omitempty"`

	// Contains the on-premises security identifier (SID) for the
	// user that was synchronized from on-premises to the cloud.
	// Read-only.
	OnPremisesSecurityIdentifier string `json:"onPremisesSecurityIdentifier,omitempty"`

	// Indicates whether or not this object is synced from an
	// on-premises directory. "True" indicates that this object is
	// synced from an on-premises directory; "False" indicates that
	// this object was originally synced from an on-premises
	// directory but is no longer synced; "Null" indicates that this
	// object has never been synced from an on-premises directory
	// (default). Read-only.
	OnPremisesSyncEnabled *bool `json:"onPremisesSyncEnabled,omitempty"`

	// Specifies password policies for this user. This value is an
	// enumeration with one possible value being
	// "DisableStrongPassword", which allows weaker passwords than
	// the default policy to be specified.
	// "DisablePasswordExpiration" can also be specified. The two
	// may be specified together; for example:
	// "DisablePasswordExpiration,DisableStrongPassword".
	PasswordPolicies string `json:"passwordPolicies,omitempty"`

	// Specifies the password profile for the user. The proile
	// contains the user's password. This property is required when
	// a user is created. The password in the profile must satisfy
	// minimum requirements as specified by the PasswordPolicies
	// property. By default, a strong password is required.
	PasswordProfile *PasswordProfile `json:"passwordProfile,omitempty"`

	// A list for the user to enumerate their past projects.
	PastProjects []string `json:"pastProjects,omitempty"`

	// The postal code for the user's postal address. The postal
	// code is specific to the user's country/region. In the United
	// States of America, this attribute contains the ZIP code.
	PostalCode string `json:"postalCode,omitempty"`

	// The preferred language for the user. Should follow ISO 639-1
	// Code; for example, "en-US".
	PreferredLanguage string `json:"preferredLanguage,omitempty"`

	// The preferred name for the user.
	PreferredName string `json:"preferredName,omitempty"`

	// The plans that are provisioned for the user. Read-only. Not
	// nullable.
	ProvisionedPlans []*ProvisionedPlan `json:"provisionedPlans,omitempty"`

	// An array of email addresses associated with the user.
	// TODO: change this to []ProxyAddress.
	ProxyAddresses []string `json:"proxyAddresses,omitempty"`

	// A list for the user to enumerate their responsibilities.
	Responsibilities []string `json:"responsibilities,omitempty"`

	// A list for the user to enumerate the schools they have
	// attended.
	Schools []string `json:"schools,omitempty"`

	// A list for the user to enumerate their skills.
	Skills []string `json:"skills,omitempty"`

	// The state or province in the user's address.
	State string `json:"state,omitempty"`

	// The street address of the user's place of business.
	StreetAddress string `json:"streetAddress,omitempty"`

	// The user's surname (family name or last name).
	Surname string `json:"surname,omitempty"`

	// A two-letter country code (ISO standard 3166). Required for
	// users that will be assigned licenses due to legal requirement
	// to change for availability of services in countries. Not
	// nullable.
	UsageLocation string `json:"usageLocation,omitempty"`

	// The user principal name (UPN) of the user. The UPN is an
	// Internet-style login name for the user based on the Internet
	// standard RFC 822. By convention, this should map to the
	// user's email name. The general format is "alias@domain",
	// where "domain" must be present in the tenant's collection of
	// verified domains. This property is required when a user is
	// created. The verified domains for the tenant can be accessed
	// from the VerifiedDomains property of the Organization.
	UserPrincipalName string `json:"userPrincipalName,omitempty"`

	// A string value that can be used to classify user types in
	// your directory, such as "Member" and "Guest".
	UserType string `json:"userType,omitempty"`
}

// GetUser retrieves the properties and relationships of a user object.
// The id parameter can be either a user ID or user principal name.
func (api *GraphAPI) GetUser(id string, properties []string) (user User, err error) {
	log.WithFields(log.Fields{
		"user": id,
	}).Info("Getting user from Graph API")

	endpoint := fmt.Sprintf("%s/%s", api.GetResourceEndpoint(resources["UserV1"]), id)
	if len(properties) > 0 {
		endpoint = fmt.Sprintf("%s?$select=%s", endpoint, strings.Join(properties, ","))
	}

	client, err := api.Client()
	if err != nil {
		return
	}

	resp, err := client.Get(endpoint)
	log.Debugf("Response: %v", resp)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if err = json.Unmarshal(body, &user); err != nil {
		return
	}

	return
}
