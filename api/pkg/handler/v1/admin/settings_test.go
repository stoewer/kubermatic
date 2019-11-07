package admin_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	apiv1 "github.com/kubermatic/kubermatic/api/pkg/api/v1"
	kubermaticv1 "github.com/kubermatic/kubermatic/api/pkg/crd/kubermatic/v1"
	"github.com/kubermatic/kubermatic/api/pkg/handler/test"
	"github.com/kubermatic/kubermatic/api/pkg/handler/test/hack"
	"k8s.io/apimachinery/pkg/runtime"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCreateRoleBinding(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name                   string
		expectedResponse       string
		httpStatus             int
		existingAPIUser        *apiv1.User
		existingKubermaticObjs []runtime.Object
	}{
		// scenario 1
		{
			name:                   "scenario 1: unauthorized user gets settings",
			expectedResponse:       `{"error":{"code":403,"message":"forbidden: \"bob@acme.com\" doesn't have admin rights"}}`,
			httpStatus:             http.StatusForbidden,
			existingKubermaticObjs: test.GenDefaultKubermaticObjects(),
			existingAPIUser:        test.GenDefaultAPIUser(),
		},
		// scenario 2
		{
			name:                   "scenario 2: authorized user gets settings first time",
			expectedResponse:       `{"id":"globalSettings","name":"globalSettings","creationTimestamp":"0001-01-01T00:00:00Z","globalSettings":{"customLinks":[],"cleanupOptions":{"Enabled":false,"Enforced":false},"defaultNodeCount":10,"clusterTypeOptions":10,"displayDemoInfo":false,"displayAPIDocs":false,"displayTermsOfService":false}}`,
			httpStatus:             http.StatusOK,
			existingKubermaticObjs: []runtime.Object{genUser("Bob", "bob@acme.com", true)},
			existingAPIUser:        test.GenDefaultAPIUser(),
		},
		// scenario 3
		{
			name:             "scenario 3: authorized user gets existing global settings",
			expectedResponse: `{"id":"globalSettings","name":"globalSettings","creationTimestamp":"0001-01-01T00:00:00Z","globalSettings":{"customLinks":[{"label":"label","url":"url:label","icon":"icon","location":"EU"}],"cleanupOptions":{"Enabled":true,"Enforced":true},"defaultNodeCount":5,"clusterTypeOptions":5,"displayDemoInfo":true,"displayAPIDocs":true,"displayTermsOfService":true}}`,
			httpStatus:       http.StatusOK,
			existingKubermaticObjs: []runtime.Object{genUser("Bob", "bob@acme.com", true),
				genDefaultGlobalSettings()},
			existingAPIUser: test.GenDefaultAPIUser(),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			kubernetesObj := []runtime.Object{}
			kubeObj := []runtime.Object{}
			req := httptest.NewRequest("GET", "/api/v1/admin/settings", strings.NewReader(""))
			res := httptest.NewRecorder()
			kubermaticObj := []runtime.Object{}
			kubermaticObj = append(kubermaticObj, tc.existingKubermaticObjs...)
			ep, _, err := test.CreateTestEndpointAndGetClients(*tc.existingAPIUser, nil, kubeObj, kubernetesObj, kubermaticObj, nil, nil, hack.NewTestRouting)
			if err != nil {
				t.Fatalf("failed to create test endpoint due to %v", err)
			}

			ep.ServeHTTP(res, req)

			if res.Code != tc.httpStatus {
				t.Fatalf("Expected HTTP status code %d, got %d: %s", tc.httpStatus, res.Code, res.Body.String())
			}

			test.CompareWithResult(t, res, tc.expectedResponse)
		})
	}
}

func genUser(name, email string, isAdmin bool) *kubermaticv1.User {
	user := test.GenUser("", name, email)
	user.Spec.IsAdmin = isAdmin
	return user
}

func genDefaultGlobalSettings() *kubermaticv1.KubermaticSetting {
	return &kubermaticv1.KubermaticSetting{
		ObjectMeta: v1.ObjectMeta{
			Name: kubermaticv1.GlobalSettingsName,
		},
		Spec: kubermaticv1.SettingSpec{
			CustomLinks: []kubermaticv1.CustomLink{
				{
					Label:    "label",
					URL:      "url:label",
					Icon:     "icon",
					Location: "EU",
				},
			},
			CleanupOptions: kubermaticv1.CleanupOptions{
				Enabled:  true,
				Enforced: true,
			},
			DefaultNodeCount:      5,
			ClusterTypeOptions:    5,
			DisplayDemoInfo:       true,
			DisplayAPIDocs:        true,
			DisplayTermsOfService: true,
		},
	}
}