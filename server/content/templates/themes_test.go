package templates

import (
	"testing"

	"github.com/ketchuphq/ketchup/server/content/templates/dummystore"

	"github.com/ketchuphq/ketchup/server/content/templates/defaultstore"
	"github.com/ketchuphq/ketchup/server/content/templates/store"
	"github.com/stretchr/testify/assert"
)

func TestListThemes(t *testing.T) {
	m, stop := setup(false, testTheme)
	defer stop()

	noneTheme, err := (&defaultstore.DefaultStore{}).Get("none")
	assert.Nil(t, err)

	themes, err := m.ListThemes()
	assert.Equal(t, len(themes), 2)
	assert.Equal(t, noneTheme.Proto(), themes[0])
	assert.Equal(t, store.StripData(testTheme), themes[1])
	assert.Nil(t, err)
}

func TestGetTheme(t *testing.T) {
	m, stop := setup(false)
	defer stop()
	theme, ref, err := m.GetTheme("test-theme")
	assert.Nil(t, theme)
	assert.Empty(t, ref)
	assert.Nil(t, err)
}

func TestGetAsset(t *testing.T) {
	m, stop := setup(false)
	defer stop()
	asset, err := m.GetAsset("test-theme")
	assert.Nil(t, asset)
	assert.Nil(t, err)
}

func TestCheckThemeForUpdate(t *testing.T) {
	m, stop := setup(true, testTheme)
	defer stop()

	_, _, _, err := m.CheckThemeForUpdate("foo")
	assert.EqualError(t, err, "theme does not exist")

	hasUpdate, ref, _, err := m.CheckThemeForUpdate(testTheme.GetName())
	assert.Nil(t, err)
	assert.Empty(t, ref)
	assert.False(t, hasUpdate)

	store := m.themeStore.(*dummy.DummyTemplateStore)
	fakeHash := "565bb37cee6cc921caec9f35ca6ba1723bc401ee"
	store.Refs[testTheme.GetName()] = fakeHash

	getLatestRef = func(vcsURL string) (string, error) { return fakeHash, nil }
	hasUpdate, ref, newRef, err := m.CheckThemeForUpdate(testTheme.GetName())
	assert.Nil(t, err)
	assert.Equal(t, fakeHash, ref)
	assert.Equal(t, fakeHash, newRef)
	assert.False(t, hasUpdate)

	fakeNewHash := "d86965bbab6b50fc9312b0f74544bd8355b7cdfa"
	calledVCSURL := ""
	getLatestRef = func(vcsURL string) (string, error) {
		calledVCSURL = vcsURL
		return fakeNewHash, nil
	}

	hasUpdate, ref, newRef, err = m.CheckThemeForUpdate(testTheme.GetName())
	assert.Equal(t, testTheme.Package.GetVcsUrl(), calledVCSURL)
	assert.Nil(t, err)
	assert.Equal(t, fakeHash, ref)
	assert.Equal(t, fakeNewHash, newRef)
	assert.True(t, hasUpdate)
}
