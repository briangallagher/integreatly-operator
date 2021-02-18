// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	oauthv1 "github.com/openshift/api/oauth/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeUserOAuthAccessTokens implements UserOAuthAccessTokenInterface
type FakeUserOAuthAccessTokens struct {
	Fake *FakeOauthV1
}

var useroauthaccesstokensResource = schema.GroupVersionResource{Group: "oauth.openshift.io", Version: "v1", Resource: "useroauthaccesstokens"}

var useroauthaccesstokensKind = schema.GroupVersionKind{Group: "oauth.openshift.io", Version: "v1", Kind: "UserOAuthAccessToken"}

// Get takes name of the userOAuthAccessToken, and returns the corresponding userOAuthAccessToken object, and an error if there is any.
func (c *FakeUserOAuthAccessTokens) Get(ctx context.Context, name string, options v1.GetOptions) (result *oauthv1.UserOAuthAccessToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(useroauthaccesstokensResource, name), &oauthv1.UserOAuthAccessToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*oauthv1.UserOAuthAccessToken), err
}

// List takes label and field selectors, and returns the list of UserOAuthAccessTokens that match those selectors.
func (c *FakeUserOAuthAccessTokens) List(ctx context.Context, opts v1.ListOptions) (result *oauthv1.UserOAuthAccessTokenList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(useroauthaccesstokensResource, useroauthaccesstokensKind, opts), &oauthv1.UserOAuthAccessTokenList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &oauthv1.UserOAuthAccessTokenList{ListMeta: obj.(*oauthv1.UserOAuthAccessTokenList).ListMeta}
	for _, item := range obj.(*oauthv1.UserOAuthAccessTokenList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested userOAuthAccessTokens.
func (c *FakeUserOAuthAccessTokens) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(useroauthaccesstokensResource, opts))
}

// Create takes the representation of a userOAuthAccessToken and creates it.  Returns the server's representation of the userOAuthAccessToken, and an error, if there is any.
func (c *FakeUserOAuthAccessTokens) Create(ctx context.Context, userOAuthAccessToken *oauthv1.UserOAuthAccessToken, opts v1.CreateOptions) (result *oauthv1.UserOAuthAccessToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(useroauthaccesstokensResource, userOAuthAccessToken), &oauthv1.UserOAuthAccessToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*oauthv1.UserOAuthAccessToken), err
}

// Update takes the representation of a userOAuthAccessToken and updates it. Returns the server's representation of the userOAuthAccessToken, and an error, if there is any.
func (c *FakeUserOAuthAccessTokens) Update(ctx context.Context, userOAuthAccessToken *oauthv1.UserOAuthAccessToken, opts v1.UpdateOptions) (result *oauthv1.UserOAuthAccessToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(useroauthaccesstokensResource, userOAuthAccessToken), &oauthv1.UserOAuthAccessToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*oauthv1.UserOAuthAccessToken), err
}

// Delete takes name of the userOAuthAccessToken and deletes it. Returns an error if one occurs.
func (c *FakeUserOAuthAccessTokens) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(useroauthaccesstokensResource, name), &oauthv1.UserOAuthAccessToken{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUserOAuthAccessTokens) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(useroauthaccesstokensResource, listOpts)

	_, err := c.Fake.Invokes(action, &oauthv1.UserOAuthAccessTokenList{})
	return err
}

// Patch applies the patch and returns the patched userOAuthAccessToken.
func (c *FakeUserOAuthAccessTokens) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *oauthv1.UserOAuthAccessToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(useroauthaccesstokensResource, name, pt, data, subresources...), &oauthv1.UserOAuthAccessToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*oauthv1.UserOAuthAccessToken), err
}
